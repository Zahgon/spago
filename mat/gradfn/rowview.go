// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// RowView is a function to extract the i-th row from the input matrix.
type RowView[O mat.Tensor] struct {
	x O
	i int
}

// NewRowView returns a new RowView Function.
func NewRowView[O mat.Tensor](x O, i int) *RowView[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *RowView[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *RowView[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *RowView[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
