// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// ProdScalar is an operator to perform element-wise product with a scalar value.
type ProdScalar[O mat.Tensor] struct {
	x1 O
	x2 O // scalar
}

// NewProdScalar returns a new ProdScalar Function.
func NewProdScalar[O mat.Tensor](x1 O, x2 O) *ProdScalar[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *ProdScalar[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the node.
func (r *ProdScalar[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *ProdScalar[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
