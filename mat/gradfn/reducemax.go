// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// ReduceMax is an operator to perform reduce-max function.
// It gets the maximum element of the Operand x
type ReduceMax[O mat.Tensor] struct {
	x      O
	argmax int
}

// NewReduceMax returns a new ReduceMax Function.
func NewReduceMax[O mat.Tensor](x O) *ReduceMax[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *ReduceMax[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of this function.
func (r *ReduceMax[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *ReduceMax[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
