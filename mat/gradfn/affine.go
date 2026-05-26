// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Affine is an operator to apply the affine function y = b + W1x1 + W2x2 + ... + WnXn.
type Affine[O mat.Tensor] struct {
	b       O
	w1      O
	x1      O
	wxPairs []O
}

// NewAffine returns a new Affine Function.
//
// The affine transformation is in the form y = b + W1x1 + W2x2 + ... + WnXn.
// The function accepts the bias "b", the first mandatory pair of "w1" and "x1",
// and then an optional arbitrary list of (w, x) pairs, that is
// "w2", "x2", "w3", "x3", ..., "wN", "xN".
//
// For each additional pair (w, x), x is allowed to be nil. In this case, the
// pair is completely ignored.
// For example, given the arguments (b, w1, x1, w2, x2, w3, x3), if only x2
// is nil, the actual operation performed will be y = b + w1x1 + w3x3.
//
// If no additional (w, x) pair is given, or all x values of wxPairs are nil,
// the actual function is just y = b + W1x1.
//
// It panics if the length of wxs is not even.
func NewAffine[O mat.Tensor](b, w1, x1 O, wxPairs ...O) *Affine[O] {
	_ = "STUB: not implemented"
	return nil
}

// Operands returns the list of operands.
func (a *Affine[O]) Operands() []O { _ = "STUB: not implemented"; return nil }

// Forward computes the output of the function.
func (a *Affine[O]) Forward() (mat.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(mat.Tensor), nil
}

// Backward computes the backward pass.
func (a *Affine[O]) Backward(gy mat.Tensor) error { _ = "STUB: not implemented"; return nil }

func isNil[O mat.Tensor](o O) bool { _ = "STUB: not implemented"; return false }
