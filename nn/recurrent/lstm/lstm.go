// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lstm

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/mat/rand"
	"github.com/nlpodyssey/spago/nn"
)

var _ nn.Model = &Model{}

// Model contains the serializable parameters.
type Model struct {
	nn.Module
	UseRefinedGates bool

	// Input gate
	WIn    *nn.Param
	WInRec *nn.Param
	BIn    *nn.Param

	// Output gate
	WOut    *nn.Param
	WOutRec *nn.Param
	BOut    *nn.Param

	// Forget gate
	WFor    *nn.Param
	WForRec *nn.Param
	BFor    *nn.Param

	// Candiate gate
	WCand    *nn.Param
	WCandRec *nn.Param
	BCand    *nn.Param
}

// State represent a state of the LSTM recurrent network.
type State struct {
	InG  mat.Tensor
	OutG mat.Tensor
	ForG mat.Tensor
	Cand mat.Tensor
	Cell mat.Tensor
	Y    mat.Tensor
}

// Option allows to configure a new Model with your specific needs.
type Option func(*Model)

func init() {
	gob.Register(&Model{})
}

// New returns a new model with parameters initialized to zeros.
func New[T float.DType](in, out int) *Model { _ = "STUB: not implemented"; return nil }

// Input gate

// Output gate

// Forget gate

// Candiate gate

// WithRefinedGates sets whether to use refined gates.
// Refined Gate: A Simple and Effective Gating Mechanism for Recurrent Units
// (https://arxiv.org/pdf/2002.11338.pdf)
//
// Refined gates setting requires input size and output size be the same.
func (m *Model) WithRefinedGates(value bool) *Model { _ = "STUB: not implemented"; return nil }

// Init initializes the parameters using Xavier uniform randomization.
// It follows the LSTM bias hack setting the Forget gate to 1 (http://proceedings.mlr.press/v37/jozefowicz15.pdf).
func (m *Model) Init(rndGen *rand.LockedRand) *Model { _ = "STUB: not implemented"; return nil }

// Forward performs the forward step for each input node and returns the result.
func (m *Model) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Next performs a single forward step, producing a new state.
//
// It computes the results with the following equations:
// inG = sigmoid(wIn (dot) x + bIn + wInRec (dot) yPrev)
// outG = sigmoid(wOut (dot) x + bOut + wOutRec (dot) yPrev)
// forG = sigmoid(wFor (dot) x + bFor + wForRec (dot) yPrev)
// cand = f(wCand (dot) x + bC + wCandRec (dot) yPrev)
// cell = inG * cand + forG * cellPrev
// y = outG * f(cell)
func (m *Model) Next(state *State, x mat.Tensor) (s *State) { _ = "STUB: not implemented"; return nil }
