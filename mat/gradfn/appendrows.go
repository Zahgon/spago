// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import "github.com/nlpodyssey/spago/mat"

// AppendRows is a Function which appends new tail rows to a matrix.
type AppendRows[O mat.Tensor] struct {
	x  O
	vs []O
}

// NewAppendRows returns a new AppendRows Function.
func NewAppendRows[O mat.Tensor](x O, vs ...O) *AppendRows[O] {
	_ = "STUB: not implemented"
	return nil
}

// Operands returns the list of operands.
func (a *AppendRows[O]) Operands() []O { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (a *AppendRows[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (a *AppendRows[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
