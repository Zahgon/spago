// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ag

import (
	"github.com/nlpodyssey/spago/mat"
)

// Backward initiates back-propagation from the input tensors.
//
// The function operates according to the following mutually exclusive rules:
//   - If the tensors already has gradients (likely assigned externally via node.AccGrads()), those gradients are used.
//   - If the tensors does not have gradients assigned and is a scalar, the output gradients are automatically assigned
//     by finding the derivative of the tensors with respect to itself (dy/dy = 1).
//   - If the tensors does not have gradients assigned and is not a scalar, it returns an error.
//
// During the back-propagation process, the gradients of all tensors, except for the given tensors, are summed to the existing gradients.
// Unless you intend to do so, ensure that all tensors have zero gradients.
func Backward(xs ...mat.Tensor) error { _ = "STUB: not implemented"; return nil }

// The three for loops below are intentionally executed in sequence.
// These steps must occur in this order, so the loops cannot be combined due to their sequential dependencies.

// 1. Prepare the backward pass for each operator.

// 2. Assign the output gradients for each operator.

// 3. Process the backward pass for each operator in parallel using wait groups.

// filterOperators returns a list of operators from a list of tensors.
func filterOperators(nodes []mat.Tensor) []*Operator { _ = "STUB: not implemented"; return nil }
