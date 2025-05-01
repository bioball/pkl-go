// Code generated from Pkl module `org.foo.BugHolder`. DO NOT EDIT.
package bugholder

import "github.com/apple/pkl-go/pkl"

type Person interface {
	Being

	GetBike() *Bike

	GetFirstName() pkl.Option[uint16]

	GetLastName() map[string]pkl.Option[uint32]

	GetThings() map[int]struct{}
}

var _ Person = (*PersonImpl)(nil)

// A Person!
type PersonImpl struct {
	IsAlive bool `pkl:"isAlive"`

	Bike *Bike `pkl:"bike"`

	// The person's first name
	FirstName pkl.Option[uint16] `pkl:"firstName"`

	// The person's last name
	LastName map[string]pkl.Option[uint32] `pkl:"lastName"`

	Things map[int]struct{} `pkl:"things"`
}

func (rcv *PersonImpl) GetIsAlive() bool {
	return rcv.IsAlive
}

func (rcv *PersonImpl) GetBike() *Bike {
	return rcv.Bike
}

// The person's first name
func (rcv *PersonImpl) GetFirstName() pkl.Option[uint16] {
	return rcv.FirstName
}

// The person's last name
func (rcv *PersonImpl) GetLastName() map[string]pkl.Option[uint32] {
	return rcv.LastName
}

func (rcv *PersonImpl) GetThings() map[int]struct{} {
	return rcv.Things
}
