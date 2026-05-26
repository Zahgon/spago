// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package attention

import (
	"github.com/nlpodyssey/spago/mat"
)

// ScaledDotProductAttention is a self-attention mechanism relating different positions of a single
// sequence to compute a representation of the same sequence.
// This method requires that the query, the key and the value vectors have already been obtained
// from the input sequence. The scaled factor is the square root of the dimension of the key vectors.
func ScaledDotProductAttention(q []mat.Tensor, k, v, scaleFactor mat.Tensor, useCausalMask bool) ([]mat.Tensor, []mat.Tensor) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: use external cache for causal mask?

// makeCausalMask returns a slice of size seqLength filled with zeros until curIndex, and the rest with -inf.
// FIXME: avoid specific float64 type, later passed to NewVec
func makeCausalMask(curIndex, seqLength int) []float64 { _ = "STUB: not implemented"; return nil }
