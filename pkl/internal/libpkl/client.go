//===----------------------------------------------------------------------===//
// Copyright © 2025 Apple Inc. and the Pkl project authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//===----------------------------------------------------------------------===//

//go:build libpkl

package libpkl

/*
#cgo pkg-config: libpkl
#include <stdlib.h>
#include <pkl.h>

// Bridge function to handle Go callbacks from C
// This function will be called by the C library and will forward to Go
void go_pkl_message_handler_bridge(unsigned int length, char *message, void *userData);

// Static C function that acts as the bridge to Go
static void c_pkl_message_handler_bridge(unsigned int length, char *message, void *userData) {
   go_pkl_message_handler_bridge(length, message, userData);
}

// Helper function to get the bridge function pointer
static pkl_message_response_handler get_bridge_handler() {
   return c_pkl_message_handler_bridge;
}
*/
import "C"

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	"github.com/google/uuid"
)

var handlerMap sync.Map

// MessageHandler is the Go equivalent of PklMessageResponseHandler
// The userData parameter will be the unsafe.Pointer passed to pkl_init
type MessageHandler func(message []byte, userData unsafe.Pointer)

//export go_pkl_message_handler_bridge
//goland:noinspection GoSnakeCaseUsage
func go_pkl_message_handler_bridge(length C.uint, message *C.char, userData unsafe.Pointer) {
	handler, exists := handlerMap.Load(userData)
	if !exists {
		return
	}

	// Convert C data to Go data
	messageBytes := C.GoBytes(unsafe.Pointer(message), C.int(length))

	// Call the Go handler with the original userData provided by the user
	handler.(MessageHandler)(messageBytes, userData)
}

type PklClient struct {
	handler MessageHandler
	pexec   *C.pkl_exec_t

	id        uuid.UUID
	idPointer unsafe.Pointer

	closed bool
	mu     sync.Mutex
	jobs   chan *job
	stop   chan struct{}
}

type job struct {
	fn   func()
	done chan struct{}
}

// New initializes the Pkl executor with a Go callback
func New(handler MessageHandler) (*PklClient, error) {
	id := uuid.New()

	client := &PklClient{
		handler:   handler,
		id:        id,
		idPointer: unsafe.Pointer(&id),
		jobs:      make(chan *job),
		stop:      make(chan struct{}),
	}

	go client.run()
	var err error
	j := &job{
		fn: func() {
			var cerr C.pkl_error_t
			var pexec *C.pkl_exec_t

			// Call the C function with our bridge handler
			if ret := C.pkl_init(C.get_bridge_handler(), client.idPointer, &pexec, &cerr); ret != 0 {
				err = fmt.Errorf("pkl_init failed: %s", pklErrorMessage(&cerr))
			}
			client.pexec = pexec
		},
		done: make(chan struct{}),
	}
	client.jobs <- j
	<-j.done
	if err != nil {
		return nil, err
	}
	handlerMap.Store(client.idPointer, client.handler)

	return client, nil
}

func (c *PklClient) run() {
	// Don't need to unlock; once goroutine terminates, the Go runtime destroys the underlying OS thread.
	// This is defensive; prevents any thread state mutations on the C/Pkl side from being side-effecting.
	runtime.LockOSThread()
	for {
		select {
		case <-c.stop:
			return
		case j := <-c.jobs:
			j.fn()
			j.done <- struct{}{}
		}
	}
}

// SendMessage sends a message to Pkl
func (c *PklClient) SendMessage(message []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return fmt.Errorf("pkl client is closed")
	}

	if len(message) == 0 {
		return fmt.Errorf("message cannot be empty")
	}

	// Convert Go slice to C data
	cMessage := C.CBytes(message)
	defer C.free(cMessage)

	var err error
	j := &job{
		fn: func() {
			var cerr C.pkl_error_t
			result := C.pkl_send_message(c.pexec, C.uint(len(message)), (*C.char)(cMessage), &cerr)
			if result != 0 {
				err = fmt.Errorf("pkl_send_message failed: %s", pklErrorMessage(&cerr))
			}
		},
		done: make(chan struct{}),
	}
	c.jobs <- j
	<-j.done
	return err
}

// Close cleans up resources
func (c *PklClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil
	}
	var err error
	j := &job{
		fn: func() {
			var cerr C.pkl_error_t
			if ret := C.pkl_close(c.pexec, &cerr); ret != 0 {
				err = fmt.Errorf("pkl_close failed: %s", pklErrorMessage(&cerr))
			}

		},
		done: make(chan struct{}),
	}
	c.jobs <- j
	<-j.done
	if err != nil {
		return err
	}
	c.stop <- struct{}{}
	c.closed = true

	handlerMap.Delete(c.id)

	return nil
}

// pklErrorMessage returns the message contained in a pkl_error_t, or a
// generic fallback if the C library didn't populate one.
func pklErrorMessage(err *C.pkl_error_t) string {
	if err.message == nil {
		return "unknown error"
	}
	return C.GoString(err.message)
}

func Version() string {
	version := C.GoString(C.pkl_version())
	return strings.Clone(version)
}
