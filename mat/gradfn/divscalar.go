// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// DivScalar is an operator to perform element-wise division with a scalar value.
type DivScalar[O mat.Tensor] struct {
	x1 O
	x2 O // scalar
}

// NewDivScalar returns a new DivScalar Function.
func NewDivScalar[O mat.Tensor](x1 O, x2 O) *DivScalar[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *DivScalar[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *DivScalar[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *DivScalar[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
