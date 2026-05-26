// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// At is an operator to obtain the i,j-th value of a matrix.
type At[O mat.Tensor] struct {
	x       O
	indices []int
}

// NewAt returns a new At Function.
func NewAt[O mat.Tensor](x O, indices ...int) *At[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *At[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *At[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *At[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
