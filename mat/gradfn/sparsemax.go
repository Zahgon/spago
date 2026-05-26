// Copyright 2020 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// SparseMax function implementation, based on https://github.com/gokceneraslan/SparseMax.torch
type SparseMax[O mat.Tensor] struct {
	x O
	y mat.Matrix // initialized during the forward pass, required by the backward pass
}

// NewSparseMax returns a new SparseMax Function.
func NewSparseMax[O mat.Tensor](x O) *SparseMax[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *SparseMax[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *SparseMax[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// translate the input by max for numerical stability

// Backward computes the backward pass.
func (r *SparseMax[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }

func sparseMaxCommon(v mat.Matrix) (zs, cumSumInput mat.Matrix, bounds []float64, tau float64) {
	_ = "STUB: not implemented"
	// FIXME: avoid casting to specific type
	return *new(mat.Matrix), *new(mat.Matrix), nil, 0
}

// Sort zs in descending order.
