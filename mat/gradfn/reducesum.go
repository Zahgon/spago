// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// ReduceSum is an operator to perform reduce-sum function.
type ReduceSum[O mat.Tensor] struct {
	x O
}

// NewReduceSum returns a new ReduceSum Function.
func NewReduceSum[O mat.Tensor](x O) *ReduceSum[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *ReduceSum[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of this function.
func (r *ReduceSum[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *ReduceSum[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
