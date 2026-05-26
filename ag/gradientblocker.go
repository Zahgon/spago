// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ag

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
)

// GradientBlocker embeds any tensors implementation disabling gradients handling and
// blocking gradients accumulation.
type GradientBlocker struct {
	mat.Tensor
}

func init() {
	gob.Register(&GradientBlocker{})
}

// StopGrad creates a new GradientBlocker that stops the accumulated gradients from
// flowing through the wrapped Node.
func StopGrad(t mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Grad always returns nil on a GradientBlocker Node.
func (r *GradientBlocker) Grad() mat.Tensor {
	_ = "STUB: not implemented"

	// AccGrad has no effects on a GradientBlocker Node.
	return *new(mat.Tensor)
}

func (r *GradientBlocker) AccGrad(_ mat.Tensor) {
	_ = "STUB: not implemented"

	// HasGrad always returns false on a GradientBlocker Node.
	return
}

func (r *GradientBlocker) HasGrad() bool {
	_ = "STUB: not implemented"

	// RequiresGrad always returns false on a GradientBlocker Node.
	return false
}

func (r *GradientBlocker) RequiresGrad() bool {
	_ = "STUB: not implemented"

	// ZeroGrad has no effects on a GradientBlocker Node.
	return false
}

func (r *GradientBlocker) ZeroGrad() { _ = "STUB: not implemented"; return }
