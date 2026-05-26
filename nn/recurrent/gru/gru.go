// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gru

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
)

var _ nn.Model = &Model{}

// Model contains the serializable parameters.
type Model struct {
	nn.Module
	WPart    *nn.Param
	WPartRec *nn.Param
	BPart    *nn.Param
	WRes     *nn.Param
	WResRec  *nn.Param
	BRes     *nn.Param
	WCand    *nn.Param
	WCandRec *nn.Param
	BCand    *nn.Param
}

// State represent a state of the GRU recurrent network.
type State struct {
	R mat.Tensor
	P mat.Tensor
	C mat.Tensor
	Y mat.Tensor
}

func init() {
	gob.Register(&Model{})
}

// New returns a new model with parameters initialized to zeros.
func New[T float.DType](in, out int) *Model { _ = "STUB: not implemented"; return nil }

func newGateParams[T float.DType](in, out int) (w, wRec, b *nn.Param) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Forward performs the forward step for each input node and returns the result.
func (m *Model) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Next performs a single forward step, producing a new state.
//
// r = sigmoid(wr (dot) x + br + wrRec (dot) yPrev)
// p = sigmoid(wp (dot) x + bp + wpRec (dot) yPrev)
// c = f(wc (dot) x + bc + wcRec (dot) (yPrev * r))
// y = p * c + (1 - p) * yPrev
func (m *Model) Next(state *State, x mat.Tensor) (s *State) { _ = "STUB: not implemented"; return nil }

// tryProd returns the product if 'a' il not nil, otherwise nil
func tryProd(a, b mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }
