// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// MaxPooling is an operator to perform max pooling.
type MaxPooling[O mat.Tensor] struct {
	x    O
	rows int
	cols int
	// initialized during the forward pass
	y       mat.Matrix
	argmaxI [][]int
	argmaxJ [][]int
}

// NewMaxPooling returns a new MaxPooling Function.
func NewMaxPooling[O mat.Tensor](x O, r, c int) *MaxPooling[O] {
	_ = "STUB: not implemented"
	return nil
}

// Operands returns the list of operands.
func (r *MaxPooling[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *MaxPooling[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// output argmax row index
// output argmax column index

// FIXME: avoid casting to specific type

// makeIntMatrix returns a new 2-dimensional slice of int.
func makeIntMatrix(indices []int) [][]int { _ = "STUB: not implemented"; return nil }

// Backward computes the backward pass.
func (r *MaxPooling[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
