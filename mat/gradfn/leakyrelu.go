// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// LeakyReLU is an operator to perform the LeakyReLU activation function.
// LeakyReLU(x) = max(0,x) + slope ° min(0,x)
type LeakyReLU[O mat.Tensor] struct {
	x     O
	alpha O // scalar
}

// NewLeakyReLU returns a new LeakyReLU Function.
func NewLeakyReLU[O mat.Tensor](x, alpha O) *LeakyReLU[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *LeakyReLU[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *LeakyReLU[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *LeakyReLU[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
