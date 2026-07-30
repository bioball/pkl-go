//go:build !libpkl

package pkl

import (
	"context"
	"os"
	"testing"

	"github.com/apple/pkl-go/pkl/internal"
	"github.com/stretchr/testify/assert"
)

func TestLoadProjectWithExternalReaders(t *testing.T) {
	t.Skip("native: panic: runtime error: invalid memory address or nil pointer dereference [recovered]")

	manager := NewEvaluatorManager()
	version, err := manager.(*evaluatorManager).getVersion()
	if err != nil {
		t.Fatal(err)
	}
	if internal.PklVersion0_27.IsGreaterThan(version) {
		t.SkipNow()
	}

	tempDir := t.TempDir()
	_ = os.Mkdir(tempDir+"/pigeons", 0o777)
	writeFile(t, tempDir+"/pigeons/PklProject", project4Contents)

	project, err := LoadProject(context.Background(), tempDir+"/pigeons/PklProject")
	if assert.NoError(t, err) {
		t.Run("evaluatorSettings", func(t *testing.T) {
			expectedSettings := ProjectEvaluatorSettings{
				ExternalModuleReaders: map[string]ProjectEvaluatorSettingExternalReader{
					"scheme1": {Executable: "reader1"},
					"scheme2": {Executable: "reader2", Arguments: []string{"with", "args"}},
				},
				ExternalResourceReaders: map[string]ProjectEvaluatorSettingExternalReader{
					"scheme3": {Executable: "reader3"},
					"scheme4": {Executable: "reader4", Arguments: []string{"with", "args"}},
				},
			}
			assert.Equal(t, expectedSettings, project.EvaluatorSettings)
		})
	}
}
