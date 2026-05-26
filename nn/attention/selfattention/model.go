// Copyright 2020 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package selfattention

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/mat/rand"
	"github.com/nlpodyssey/spago/nn"
	"github.com/nlpodyssey/spago/nn/linear"
)

var _ nn.Model = &Model{}

// Cache contains the projected keys and values at index 0, 1 respectively.
type Cache [2]mat.Tensor

// HasValues reports whether both values in Cache are not nil.
func (c Cache) HasValues() bool { _ = "STUB: not implemented"; return false }

// Model contains the serializable parameters.
type Model struct {
	nn.Module
	Config
	Query       *linear.Model
	Key         *linear.Model
	Value       *linear.Model
	ScaleFactor *nn.Buffer
}

// Config provides configuration settings for a Self-Attention Model.
type Config struct {
	InputSize        int
	QuerySize        int
	KeySize          int
	ValueSize        int
	ScaleFactor      float64
	UseCausalMask    bool
	IsCrossAttention bool
}

func init() {
	gob.Register(&Model{})
}

// New returns a new model with parameters initialized to zeros.
func New[T float.DType](config Config) *Model { _ = "STUB: not implemented"; return nil }

// Init initializes the query, key and value linear layers with uniform Xavier random distribution.
func (m *Model) Init(rng *rand.LockedRand) { _ = "STUB: not implemented"; return }

// Forward performs the forward step for each input node and returns the result.
func (m *Model) Forward(cache Cache, q, x []mat.Tensor) ([]mat.Tensor, []mat.Tensor, Cache) {
	_ = "STUB: not implemented"
	return nil, nil, *new(Cache)
}
