// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// ColView is an operator to extract the i-th column from a matrix.
type ColView[O mat.Tensor] struct {
	x O
	i int
}

// NewColView extracts the i-th column from the input matrix.
func NewColView[O mat.Tensor](x O, i int) *ColView[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *ColView[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *ColView[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *ColView[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
