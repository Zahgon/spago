// Copyright 2020 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package activation

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/nn"
)

var _ nn.Model = &Model{}

type Model struct {
	nn.Module
	Activation Activation
	Params     []*nn.Param
}

func init() {
	gob.Register(&Model{})
}

func New(activation Activation, params ...*nn.Param) *Model { _ = "STUB: not implemented"; return nil }

func (m *Model) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

func (m *Model) activationFunc() (func(x mat.Tensor) mat.Tensor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
