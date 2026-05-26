// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package multiheadattention

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/mat/rand"
	"github.com/nlpodyssey/spago/nn"
	"github.com/nlpodyssey/spago/nn/attention/selfattention"
	"github.com/nlpodyssey/spago/nn/linear"
)

var _ nn.Model = &Model{}

// Model contains the serializable parameters.
type Model struct {
	nn.Module
	Heads       []*selfattention.Model
	OutputMerge *linear.Model
}

func init() {
	gob.Register(&Model{})
}

// New returns a new model with parameters initialized to zeros.
func New[T float.DType](size, numOfHeads int, useCausalMask, isCrossAttention bool) *Model {
	_ = "STUB: not implemented"
	return nil
}

// Init initializes the self-attention heads and the merge layer with uniform Xavier random distribution.
func (m *Model) Init(rng *rand.LockedRand) { _ = "STUB: not implemented"; return }

func makeAttentionHeads[T float.DType](dm, n int, useCausalMask, isCrossAttention bool) []*selfattention.Model {
	_ = "STUB: not implemented"
	return nil
}

// Cache contains the self-attention cache for each head.
type Cache []selfattention.Cache

func (r Cache) At(i int) selfattention.Cache {
	_ = "STUB: not implemented"
	return *new(selfattention.Cache)
}

// Forward performs the forward step for each input node and returns the result.
func (m *Model) Forward(cache Cache, q, x []mat.Tensor) ([]mat.Tensor, [][]mat.Tensor, Cache) {
	_ = "STUB: not implemented"
	return nil, nil, *new(Cache)
}

func (m *Model) project(heads [][]mat.Tensor, seqLen int) []mat.Tensor {
	_ = "STUB: not implemented"
	return nil
}

// shares the same backing array with buf
