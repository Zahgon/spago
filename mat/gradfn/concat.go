// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Concat is an operator to perform vector concatenation.
type Concat[O mat.Tensor] struct {
	xs    []O
	ySize int
}

// NewConcat returns a new Concat Function.
func NewConcat[O mat.Tensor](xs []O) *Concat[O] { _ = "STUB: not implemented"; return nil }

// assigned during the Forward()

// Operands returns the list of operands.
func (r *Concat[O]) Operands() []O {
	_ = "STUB: not implemented"

	// Forward computes the output of the function.
	return nil
}

func (r *Concat[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// reset output size

// Backward computes the backward pass.
func (r *Concat[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }
