// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// SoftPlus function: f(x) = 1 / β ∗ log(1 + exp(β ∗ x))
type SoftPlus[O mat.Tensor] struct {
	x         O
	beta      O
	threshold O
}

// NewSoftPlus returns a new SoftPlus Function.
func NewSoftPlus[O mat.Tensor](x O, beta, threshold O) *SoftPlus[O] {
	_ = "STUB: not implemented"
	return nil
}

// Operands returns the list of operands.
func (r *SoftPlus[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *SoftPlus[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *SoftPlus[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
