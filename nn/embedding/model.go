// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package embedding

import (
	"encoding/gob"
	"sync"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
)

var _ nn.ParamsTraverser = &Model{}

// Model implements a simple lookup table that stores fixed-size embeddings
// for a predefined dictionary. It is commonly used to store and retrieve word
// embeddings using their corresponding indices.
type Model struct {
	nn.Module
	Size         int
	Dim          int
	Weights      []*nn.Param
	embedGradIdx map[int]struct{}
	mu           sync.Mutex
}

func init() {
	gob.Register(&Model{})
}

// New returns a new embeddings Model.
func New[T float.DType](size int, dim int) *Model { _ = "STUB: not implemented"; return nil }

// TraverseParams allows embeddings with gradients to be traversed for optimization.
func (m *Model) TraverseParams(callback func(param *nn.Param)) { _ = "STUB: not implemented"; return }

func (m *Model) Embedding(idx int) (*Embedding, error) { _ = "STUB: not implemented"; return nil, nil }

// Encode returns the embedding values associated with the input indices.
// It returns an error if one of the input elements is out of range.
func (m *Model) Encode(input []int) ([]mat.Tensor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustEncode returns the embedding values associated with the input indices.
func (m *Model) MustEncode(input []int) []mat.Tensor { _ = "STUB: not implemented"; return nil }

// checkInput returns an error if one of the input elements is out of range.
func (m *Model) checkInput(input []int) error { _ = "STUB: not implemented"; return nil }

func (m *Model) CountEmbedWithGrad() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) ZeroGrad() { _ = "STUB: not implemented"; return }
