// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Swish is an operator to perform element-wise swish function: y = x * sigmoid(x).
type Swish[O mat.Tensor] struct {
	x O
}

// NewSwish returns a new Swish Function.
func NewSwish[O mat.Tensor](x O) *Swish[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (l *Swish[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (l *Swish[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (l *Swish[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }

func swishDeriv(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }
