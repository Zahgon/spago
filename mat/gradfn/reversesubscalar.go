// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// ReverseSubScalar is the element-wise subtraction function over two values.
type ReverseSubScalar[O mat.Tensor] struct {
	x1 O
	x2 O // scalar
}

// NewReverseSubScalar returns a new ReverseSubScalar Function.
func NewReverseSubScalar[O mat.Tensor](x1 O, x2 O) *ReverseSubScalar[O] {
	_ = "STUB: not implemented"
	return nil
}

// Operands returns the list of operands.
func (r *ReverseSubScalar[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *ReverseSubScalar[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *ReverseSubScalar[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
