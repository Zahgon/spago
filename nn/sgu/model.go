// Copyright 2021 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sgu implements the Spatial Gating Unit (SGU).
// Reference: `Pay Attention to MLPs` by Liu et al, 2021 (https://arxiv.org/pdf/2105.08050.pdf)
package sgu

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
	"github.com/nlpodyssey/spago/nn/activation"
	"github.com/nlpodyssey/spago/nn/convolution/conv1x1"
	"github.com/nlpodyssey/spago/nn/normalization/layernorm"
)

// Model contains the serializable parameters.
type Model struct {
	nn.Module
	Config Config
	Norm   *layernorm.Model
	Proj   *conv1x1.Model
	Act    *activation.Model
}

var _ nn.Model = &Model{}

// Config provides configuration parameters for Model.
type Config struct {
	Dim        int
	DimSeq     int
	InitEps    float64
	Activation activation.Activation
}

func init() {
	gob.Register(&Model{})
}

// New returns a new Model initialized to zeros.
func New[T float.DType](config Config) *Model { _ = "STUB: not implemented"; return nil }

// Initialize set the projection weights as near-zero values and the biases as ones to improve training stability.
func (m *Model) Initialize(seed uint64) { _ = "STUB: not implemented"; return }

// Forward performs the forward step for each input node and returns the result.
func (m *Model) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }
