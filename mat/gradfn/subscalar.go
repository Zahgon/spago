// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// SubScalar is an element-wise subtraction function with a scalar value.
type SubScalar[O mat.Tensor] struct {
	x1 O
	x2 O // scalar
}

// NewSubScalar returns a new SubScalar Function.
func NewSubScalar[O mat.Tensor](x1 O, x2 O) *SubScalar[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *SubScalar[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the node.
func (r *SubScalar[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *SubScalar[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }

// equals to gy.ProdScalar(1.0)
