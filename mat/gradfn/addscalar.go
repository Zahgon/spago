// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// AddScalar is an operator to perform element-wise addition over two values.
type AddScalar[O mat.Tensor] struct {
	x1 O
	x2 O // scalar
}

// NewAddScalar returns a new AddScalar Function.
func NewAddScalar[O mat.Tensor](x1, x2 O) *AddScalar[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *AddScalar[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
// It doesn't backward on the scalar value x2.
func (r *AddScalar[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *AddScalar[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
