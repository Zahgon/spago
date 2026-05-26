// Copyright 2020 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Min is an operator to perform element-wise min.
// y = min(x1, x2)
type Min[O mat.Tensor] struct {
	x1 O
	x2 O
}

// NewMin returns a new Min Function.
func NewMin[O mat.Tensor](x1 O, x2 O) *Min[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Min[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *Min[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Min[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }

// FIXME: avoid casting to specific type
