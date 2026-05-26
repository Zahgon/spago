// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Sqrt is an operator to perform element-wise square root function.
type Sqrt[O mat.Tensor] struct {
	x O
}

// NewSqrt returns a new Sqrt Function.
func NewSqrt[O mat.Tensor](x O) *Sqrt[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Sqrt[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *Sqrt[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Sqrt[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
