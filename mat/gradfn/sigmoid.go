// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Sigmoid is an operator to perform element-wise sigmoid function.
type Sigmoid[O mat.Tensor] struct {
	x O
}

// NewSigmoid returns a new Log Function.
func NewSigmoid[O mat.Tensor](x O) *Sigmoid[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (l *Sigmoid[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (l *Sigmoid[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	// TODO: cache the sigmoid value in the forward pass for the backward pass?
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (l *Sigmoid[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }

// derivative of the sigmoid function
