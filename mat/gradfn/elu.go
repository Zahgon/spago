// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// ELU is an operator to perform the ELU activation function.
// ELU(x) = max(0,x) + min(0,α ∗ (exp(x) − 1))
type ELU[O mat.Tensor] struct {
	x     O
	alpha O // scalar
}

// NewELU returns a new ELU Function.
func NewELU[O mat.Tensor](x O, alpha O) *ELU[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *ELU[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *ELU[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *ELU[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
