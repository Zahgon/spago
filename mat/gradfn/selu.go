// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// SELU function: f(x) = scale ∗ (max(0,x) + min(0, α ∗ (exp(x) − 1)))
type SELU[O mat.Tensor] struct {
	x     O
	alpha O // scalar
	scale O // scalar
}

// NewSELU returns a new SELU Function.
func NewSELU[O mat.Tensor](x O, alpha, scale O) *SELU[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *SELU[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *SELU[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *SELU[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
