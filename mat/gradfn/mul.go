// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Mul is an operator to perform matrix-vector multiplication.
type Mul[O mat.Tensor] struct {
	x1 O // matrix
	x2 O // vector
}

// NewMul returns a new Mul Function.
func NewMul[O mat.Tensor](x1 O, x2 O) *Mul[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Mul[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *Mul[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Mul[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }

//r.x2.AccGrad(gy.T().Mul(r.x1).T()) // alternative method
