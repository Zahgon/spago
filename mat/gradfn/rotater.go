// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import "github.com/nlpodyssey/spago/mat"

// RotateR is a function to perform a right circular shift of a vector.
type RotateR[O mat.Tensor] struct {
	x O
	i int
}

// NewRotateR returns a new RotateR Function. `i` is the number of places by
// which the elements are shifted.
func NewRotateR[O mat.Tensor](x O, i int) *RotateR[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *RotateR[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *RotateR[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *RotateR[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }

func rotate(m mat.Matrix, i int) mat.Matrix { _ = "STUB: not implemented"; return *new(mat.Matrix) }
