// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crf

import (
	"github.com/nlpodyssey/spago/mat"
)

// FIXME: ViterbiStructure currently works with float64 only

// ViterbiStructure implements Viterbi decoding.
type ViterbiStructure struct {
	scores       mat.Matrix
	backpointers []int
}

// NewViterbiStructure returns a new ViterbiStructure ready to use.
func NewViterbiStructure(size int) *ViterbiStructure { _ = "STUB: not implemented"; return nil }

// Viterbi decodes the xs sequence according to the transitionMatrix.
func Viterbi(transitionMatrix mat.Matrix, xs []mat.Tensor) []int {
	_ = "STUB: not implemented"
	return nil
}

func viterbiStepStart(transitionMatrix, maxVec mat.Matrix) *ViterbiStructure {
	_ = "STUB: not implemented"
	return nil
}

func viterbiStepEnd(transitionMatrix, maxVec mat.Matrix) *ViterbiStructure {
	_ = "STUB: not implemented"
	return nil
}

func viterbiStep(transitionMatrix, maxVec, stepVec mat.Matrix) *ViterbiStructure {
	_ = "STUB: not implemented"
	return nil
}
