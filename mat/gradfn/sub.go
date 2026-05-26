// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Sub is an element-wise subtraction function over two values.
type Sub[O mat.Tensor] struct {
	x1 O
	x2 O
}

// NewSub returns a new Sub Function.
func NewSub[O mat.Tensor](x1 O, x2 O) *Sub[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Sub[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the node.
func (r *Sub[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Sub[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
