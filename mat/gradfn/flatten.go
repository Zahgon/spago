// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Flatten is a Function to reshape a matrix-operand into a "flattened" row vector.
type Flatten[O mat.Tensor] struct {
	x O
}

// NewFlatten returns a new Flatten Function.
func NewFlatten[O mat.Tensor](x O) *Flatten[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Flatten[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the node.
func (r *Flatten[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Flatten[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
