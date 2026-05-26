// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Softmax is a single-input softmax function.
type Softmax[O mat.Tensor] struct {
	x O
	y mat.Matrix // initialized during the forward pass (required by the backward pass)
}

// NewSoftmax returns a new Softmax Function.
func NewSoftmax[O mat.Tensor](x O) *Softmax[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Softmax[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of this function.
func (r *Softmax[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Softmax[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
