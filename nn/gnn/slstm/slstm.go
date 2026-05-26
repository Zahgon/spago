// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package slstm implements a Sentence-State LSTM graph neural network.
//
// Reference: "Sentence-State LSTM for Text Representation" by Zhang et al, 2018.
// (https://arxiv.org/pdf/1805.02474.pdf)
package slstm

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
)

var _ nn.Model = &Model{}

// TODO(1): code refactoring using a structure to maintain states.
// TODO(2): use a gradient policy (i.e. reinforcement learning) to increase the context with dynamic skip connections.

// Model contains the serializable parameters.
type Model struct {
	nn.Module
	Config                 Config
	InputGate              *HyperLinear4
	LeftCellGate           *HyperLinear4
	RightCellGate          *HyperLinear4
	CellGate               *HyperLinear4
	SentCellGate           *HyperLinear4
	OutputGate             *HyperLinear4
	InputActivation        *HyperLinear4
	NonLocalSentCellGate   *HyperLinear3
	NonLocalInputGate      *HyperLinear3
	NonLocalSentOutputGate *HyperLinear3
	StartH                 *nn.Param
	EndH                   *nn.Param
	InitValue              *nn.Param
}

// Config provides configuration settings for a Sentence-State LSTM Model.
type Config struct {
	InputSize  int
	OutputSize int
	Steps      int
}

const windowSize = 3 // the window is fixed in this implementation

// HyperLinear4 groups multiple params for an affine transformation.
type HyperLinear4 struct {
	nn.Module
	W *nn.Param
	U *nn.Param
	V *nn.Param
	B *nn.Param
}

// HyperLinear3 groups multiple params for an affine transformation.
type HyperLinear3 struct {
	nn.Module
	W *nn.Param
	U *nn.Param
	B *nn.Param
}

// State contains nodes used during the forward step.
type State struct {
	xUi []mat.Tensor
	xUl []mat.Tensor
	xUr []mat.Tensor
	xUf []mat.Tensor
	xUs []mat.Tensor
	xUo []mat.Tensor
	xUu []mat.Tensor

	ViPrevG mat.Tensor
	VlPrevG mat.Tensor
	VrPrevG mat.Tensor
	VfPrevG mat.Tensor
	VsPrevG mat.Tensor
	VoPrevG mat.Tensor
	VuPrevG mat.Tensor
}

func init() {
	gob.Register(&Model{})
}

// New returns a new model with parameters initialized to zeros.
func New[T float.DType](config Config) *Model { _ = "STUB: not implemented"; return nil }

func newGate4[T float.DType](in, out int) *HyperLinear4 { _ = "STUB: not implemented"; return nil }

func newGate3[T float.DType](size int) *HyperLinear3 { _ = "STUB: not implemented"; return nil }

// Forward performs the forward step for each input node and returns the result.
func (m *Model) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

// the result is shared among all steps

// the result is shared among all nodes of the same step

func (m *Model) computeUx(s *State, xs []mat.Tensor) { _ = "STUB: not implemented"; return }

func (m *Model) computeVg(s *State, prevG mat.Tensor) { _ = "STUB: not implemented"; return }

func (m *Model) updateHiddenNodes(s *State, prevH []mat.Tensor, prevC []mat.Tensor, prevG mat.Tensor) ([]mat.Tensor, []mat.Tensor) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Model) updateSentenceState(prevH []mat.Tensor, prevC []mat.Tensor, prevG mat.Tensor) (mat.Tensor, mat.Tensor) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), *new(mat.Tensor)
}

func (m *Model) processNode(s *State, i int, prevH []mat.Tensor, prevC []mat.Tensor, prevG mat.Tensor) (h mat.Tensor, c mat.Tensor) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), *new(mat.Tensor)
}
