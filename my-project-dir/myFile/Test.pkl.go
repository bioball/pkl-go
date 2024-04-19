// Code generated from Pkl module `test`. DO NOT EDIT.
package myfile

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Test struct {
	Reg string `pkl:"reg"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Test
func LoadFromPath(ctx context.Context, path string) (ret *Test, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Test
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Test, error) {
	var ret Test
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
