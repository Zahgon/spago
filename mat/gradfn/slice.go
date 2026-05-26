// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Slice is a function to extract a portion of a matrix.
type Slice[O mat.Tensor] struct {
	x       O
	fromRow int
	fromCol int
	toRow   int
	toCol   int
}

// NewSlice returns a new Slice Function.
func NewSlice[O mat.Tensor](x O, fromRow, fromCol, toRow, toCol int) *Slice[O] {
	_ = "STUB: not implemented"
	return nil
}

// Operands returns the list of operands.
func (s *Slice[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (s *Slice[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (s *Slice[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
