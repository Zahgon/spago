// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// SwishB function: f(x) = x * sigmoid.
//
// Reference: "Searching for Activation Functions" by Ramachandran et al, 2017.
// (https://arxiv.org/pdf/1710.05941.pdf)
type SwishB[O mat.Tensor] struct {
	x    O
	beta O // scalar
}

// NewSwishB returns a new SwishB Function.
func NewSwishB[O mat.Tensor](x O, beta O) *SwishB[O] { _ = "STUB: not implemented"; return nil }

// Operands returns the list of operands.
func (r *SwishB[O]) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (r *SwishB[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (r *SwishB[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }

// FIXME: avoid casting to specific type
