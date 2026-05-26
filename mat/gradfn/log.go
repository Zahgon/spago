// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Log is an operator to perform element-wise natural logarithm function.
type Log[O mat.Tensor] struct {
	x O
}

// NewLog returns a new Log Function.
func NewLog[O mat.Tensor](x O) *Log[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (l *Log[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (l *Log[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (l *Log[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }

func safeLogDeriv(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }
