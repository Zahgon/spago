// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Pow is an operator to perform element-wise pow function.
type Pow[O mat.Tensor] struct {
	x     O
	power float64
}

// NewPow returns a new Pow Function.
func NewPow[O mat.Tensor](x O, power float64) *Pow[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Pow[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *Pow[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Pow[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
