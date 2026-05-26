// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Dot is an operator to perform the dot product over two matrices.
// y = x1 dot x2
type Dot[O mat.Tensor] struct {
	x1 O
	x2 O
}

// NewDot returns a new Dot Function.
func NewDot[O mat.Tensor](x1 O, x2 O) *Dot[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Dot[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *Dot[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Dot[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
