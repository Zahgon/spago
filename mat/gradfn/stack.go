// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Stack is a Function which stacks together all given operand matrices,
// producing a single bigger matrix as result.
type Stack[O mat.Tensor] struct {
	xs []O
}

// NewStack returns a new Stack Function.
func NewStack[O mat.Tensor](xs []O) *Stack[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Stack[O]) Operands() []O {
	_ = "STUB: not implemented"

	// Forward computes the output of the function.
	return nil
}

func (r *Stack[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Stack[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
