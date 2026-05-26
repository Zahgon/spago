// Copyright 2020 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// SparseMaxLoss function implementation, based on https://github.com/gokceneraslan/SparseMax.torch
type SparseMaxLoss[O mat.Tensor] struct {
	x   O
	tau float64    // computed during the forward pass
	y   mat.Matrix // computed during forward pass
}

// NewSparseMaxLoss returns a new SparseMaxLoss Function.
func NewSparseMaxLoss[O mat.Tensor](x O) *SparseMaxLoss[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *SparseMaxLoss[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *SparseMaxLoss[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *SparseMaxLoss[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
