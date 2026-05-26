// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// ReduceMean is an operator to perform reduce-mean function.
type ReduceMean[O mat.Tensor] struct {
	x O
}

// NewReduceMean returns a new ReduceMean Function.
func NewReduceMean[O mat.Tensor](x O) *ReduceMean[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *ReduceMean[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of this node.
func (r *ReduceMean[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *ReduceMean[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
