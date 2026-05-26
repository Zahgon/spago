// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Reshape is a Function which reshapes an operand into a new matrix of given
// rows × columns size.
type Reshape[O mat.Tensor] struct {
	x    O
	rows int
	cols int
}

// NewReshape returns a new Reshape Function.
func NewReshape[O mat.Tensor](x O, r, c int) *Reshape[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *Reshape[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the node.
func (r *Reshape[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *Reshape[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
