// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// SoftShrink function: f(x) = x − λ if x > λ; x + λ if x < −λ; 0 otherwise.
type SoftShrink[O mat.Tensor] struct {
	x      O
	lambda O // scalar
}

// NewSoftShrink returns a new SoftShrink Function.
func NewSoftShrink[O mat.Tensor](x O, lambda O) *SoftShrink[O] {
	_ = "STUB: not implemented"
	return nil
}

// Operands returns the list of operands.
func (r *SoftShrink[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *SoftShrink[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *SoftShrink[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
