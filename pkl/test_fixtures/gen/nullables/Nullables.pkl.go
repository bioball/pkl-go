// Code generated from Pkl module `nullables`. DO NOT EDIT.
package nullables

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Nullables struct {
	Res0 pkl.Option[string] `pkl:"res0"`

	Res1 pkl.Option[string] `pkl:"res1"`

	Res2 pkl.Option[int] `pkl:"res2"`

	Res3 pkl.Option[int] `pkl:"res3"`

	Res4 pkl.Option[int8] `pkl:"res4"`

	Res5 pkl.Option[int8] `pkl:"res5"`

	Res6 pkl.Option[int16] `pkl:"res6"`

	Res7 pkl.Option[int16] `pkl:"res7"`

	Res8 pkl.Option[int32] `pkl:"res8"`

	Res9 pkl.Option[int32] `pkl:"res9"`

	Res10 pkl.Option[uint] `pkl:"res10"`

	Res11 pkl.Option[uint] `pkl:"res11"`

	Res12 pkl.Option[uint8] `pkl:"res12"`

	Res13 pkl.Option[uint8] `pkl:"res13"`

	Res14 pkl.Option[uint16] `pkl:"res14"`

	Res15 pkl.Option[uint16] `pkl:"res15"`

	Res16 pkl.Option[uint32] `pkl:"res16"`

	Res17 pkl.Option[uint32] `pkl:"res17"`

	Res18 pkl.Option[float64] `pkl:"res18"`

	Res19 pkl.Option[float64] `pkl:"res19"`

	Res20 pkl.Option[bool] `pkl:"res20"`

	Res21 pkl.Option[bool] `pkl:"res21"`

	Res22 pkl.Option[map[string]string] `pkl:"res22"`

	Res23 pkl.Option[map[string]string] `pkl:"res23"`

	Res25 pkl.Option[map[pkl.Option[string]]pkl.Option[string]] `pkl:"res25"`

	Res26 pkl.Option[[]pkl.Option[int]] `pkl:"res26"`

	Res27 pkl.Option[[]pkl.Option[int]] `pkl:"res27"`

	Res28 pkl.Option[*MyClass] `pkl:"res28"`

	Res29 pkl.Option[*MyClass] `pkl:"res29"`

	Res30 pkl.Option[*MyClass] `pkl:"res30"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Nullables
func LoadFromPath(ctx context.Context, path string) (ret *Nullables, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Nullables
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Nullables, error) {
	var ret Nullables
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
