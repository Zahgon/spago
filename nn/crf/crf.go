// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crf

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
	Size             int
	TransitionScores *nn.Param
}

func init() {
	gob.Register(&Model{})
}

// New returns a new convolution Model, initialized according to the given configuration.
func New[T float.DType](size int) *Model { _ = "STUB: not implemented"; return nil }

// +1 for start and end transitions

// Decode performs viterbi decoding.
func (m *Model) Decode(emissionScores []mat.Tensor) []int { _ = "STUB: not implemented"; return nil }

// NegativeLogLoss computes the negative log loss with respect to the targets.
func (m *Model) NegativeLogLoss(emissionScores []mat.Tensor, target []int) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

func (m *Model) goldScore(emissionScores []mat.Tensor, target []int) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// start transition

// end transition

func (m *Model) totalScore(predicted []mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

func (m *Model) totalScoreStart(stepVec mat.Tensor) []mat.Tensor {
	_ = "STUB: not implemented"
	return nil
}

func (m *Model) totalScoreEnd(stepVec []mat.Tensor) []mat.Tensor {
	_ = "STUB: not implemented"
	return nil
}

func (m *Model) totalScoreStep(totalVec []mat.Tensor, stepVec []mat.Tensor) []mat.Tensor {
	_ = "STUB: not implemented"
	return nil
}
