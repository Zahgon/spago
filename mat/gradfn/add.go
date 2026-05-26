// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Add is an operator to perform element-wise sum over two values.
// y = x1 + x2
type Add[O mat.Tensor] struct {
	x1 O
	x2 O
}

// NewAdd returns a new Add Function.
func NewAdd[O mat.Tensor](x1, x2 O) *Add[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Add[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *Add[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Add[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
