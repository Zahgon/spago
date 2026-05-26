// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// MulT is an operator to perform matrix-vector multiplication.
type MulT[O mat.Tensor] struct {
	x1 O // matrix
	x2 O // vector
}

// NewMulT returns a new MulT Function.
func NewMulT[O mat.Tensor](x1 O, x2 O) *MulT[O] { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *MulT[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Operands returns the list of operands.
func (r *MulT[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Backward computes the backward pass.
func (r *MulT[O]) Backward(gy mat.Tensor) error {
	_ = "STUB: not implemented"
	//	if !(r.x1.Value().Shape()[0] == gy.Shape()[0] && r.x2.Value().Shape()[1] == gy.Shape()[1]) {
	//		panic("fn: matrices with not compatible size")
	//	}
	return nil
}

//r.x2.AccGrad(gy.T().MulT(r.x1).T()) // alternative method
