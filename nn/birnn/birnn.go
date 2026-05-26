// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package birnn

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/nn"
)

// MergeType is the enumeration-like type used for the set of merging methods
// which a BiRNN model Processor can perform.
type MergeType int

const (
	// Concat merging method: the outputs are concatenated together (the default)
	Concat MergeType = iota
	// Sum merging method: the outputs are added together
	Sum
	// Prod merging method: the outputs multiplied element-wise together
	Prod
	// Avg merging method: the average of the outputs is taken
	Avg
)

var _ nn.Model = &Model{}

// Model contains the serializable parameters.
type Model struct {
	nn.Module
	Positive  nn.StandardModel // positive time direction a.k.a. left-to-right
	Negative  nn.StandardModel // negative time direction a.k.a. right-to-left
	MergeMode MergeType
}

func init() {
	gob.Register(&Model{})
	gob.Register(&Model{})
}

// New returns a new model with parameters initialized to zeros.
func New(positive, negative nn.StandardModel, merge MergeType) *Model {
	_ = "STUB: not implemented"
	return nil
}

// Forward performs the forward step for each input node and returns the result.
func (m *Model) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

func reversed(ns []mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

func (m *Model) merge(a, b mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}
