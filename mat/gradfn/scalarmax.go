// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// ScalarMax is an operator to perform reduce-max function on a list of scalars.
// It gets the maximum element of the Operand x
type ScalarMax[O mat.Tensor] struct {
	xs     []O
	argmax int
}

// NewScalarMax returns a new ScalarMax Function.
func NewScalarMax[O mat.Tensor](xs []O) *ScalarMax[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *ScalarMax[O]) Operands() []O {
	_ = "STUB: not implemented"

	// Forward computes the output of this function.
	return nil
}

func (r *ScalarMax[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// FIXME: avoid casting to specific type

// Backward computes the backward pass.
func (r *ScalarMax[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
