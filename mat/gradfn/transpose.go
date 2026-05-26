// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Transpose is a Function to calculate the transpose of the matrix-operand.
type Transpose[O mat.Tensor] struct {
	x O
}

// NewTranspose returns a new Transpose Function.
func NewTranspose[O mat.Tensor](x O) *Transpose[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Transpose[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the node.
func (r *Transpose[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Transpose[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
