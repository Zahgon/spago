// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Copy is an operator to perform copy function.
// y = x
type Copy[O mat.Tensor] struct {
	x O
}

// NewCopy returns a new Copy Function.
func NewCopy[O mat.Tensor](x O) *Copy[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Copy[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *Copy[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Copy[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
