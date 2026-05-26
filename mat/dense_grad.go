// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

// Value returns the value of the Matrix itself.
func (d *Dense[T]) Value() Tensor {
	_ = "STUB: not implemented"

	// Grad returns the gradients accumulated during the backward pass.
	return *new(Tensor)
}

func (d *Dense[T]) Grad() Tensor { _ = "STUB: not implemented"; return *new(Tensor) }

// AccGrad accumulates the gradients.
// It accumulates the gradients even if the requiresGrad flag is false.
func (d *Dense[T]) AccGrad(grad Tensor) { _ = "STUB: not implemented"; return }

// HasGrad reports whether there are accumulated gradients.
func (d *Dense[T]) HasGrad() bool { _ = "STUB: not implemented"; return false }

// RequiresGrad reports whether the Variable requires gradients.
func (d *Dense[T]) RequiresGrad() bool { _ = "STUB: not implemented"; return false }

// SetRequiresGrad sets the requiresGrad flag.
func (d *Dense[T]) SetRequiresGrad(v bool) {
	_ = "STUB: not implemented"

	// ZeroGrad zeroes the gradients, setting the value of Grad to nil.
	return
}

func (d *Dense[T]) ZeroGrad() { _ = "STUB: not implemented"; return }
