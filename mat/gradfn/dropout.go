// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/rand"
)

// Dropout is an operator to perform elements dropout with a probability.
type Dropout[O mat.Tensor] struct {
	x       O
	prob    float64
	q       float64 // 1 - p
	randGen *rand.LockedRand
	mask    mat.Matrix // filled during the forward
}

// NewDropout returns a new Dropout Function.
func NewDropout[O mat.Tensor](x O, p float64, randGen *rand.LockedRand) *Dropout[O] {
	_ = "STUB: not implemented"
	return nil
}

// Operands returns the list of operands.
func (r *Dropout[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *Dropout[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// FIXME: avoid casting to specific type

// Backward computes the backward pass.
func (r *Dropout[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
