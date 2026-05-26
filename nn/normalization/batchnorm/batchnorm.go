// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package batchnorm

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
	W        *nn.Param
	B        *nn.Param
	Mean     *nn.Buffer
	StdDev   *nn.Buffer
	Momentum *nn.Buffer
}

const epsilon = 1e-5
const defaultMomentum = 0.9

func init() {
	gob.Register(&Model{})
}

// NewWithMomentum returns a new model with supplied size and momentum.
func NewWithMomentum[T float.DType](size int, momentum T) *Model {
	_ = "STUB: not implemented"
	return nil
}

// New returns a new model with the supplied size and default momentum
func New[T float.DType](size int) *Model { _ = "STUB: not implemented"; return nil }

// Forward performs the forward step for each input node and returns the result.
func (m *Model) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

// ForwardT performs the forward step for each input node and returns the result.
func (m *Model) ForwardT(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

func (m *Model) process(xs []mat.Tensor, devVector mat.Tensor, meanVector mat.Tensor) []mat.Tensor {
	_ = "STUB: not implemented"
	return nil
}

func (m *Model) updateBatchNormParameters(meanVector, devVector mat.Matrix) {
	_ = "STUB: not implemented"
	return
}

// Mean computes the mean of the input.
func (m *Model) mean(xs []mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// StdDev computes the standard deviation of the input.
func (m *Model) stdDev(meanVector mat.Tensor, xs []mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}
