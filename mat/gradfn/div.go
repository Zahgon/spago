// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Div is an operator to perform element-wise division over two values.
type Div[O mat.Tensor] struct {
	x1 O
	x2 O
}

// NewDiv returns a new Div Function.
func NewDiv[O mat.Tensor](x1 O, x2 O) *Div[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Div[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *Div[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Div[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
