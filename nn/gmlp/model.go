// Copyright 2021 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gmlp implements a model composed by basic MLP layers with gating mechanism.
// Reference: `Pay Attention to MLPs` by Liu et al, 2021 (https://arxiv.org/pdf/2105.08050.pdf)
package gmlp

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
	"github.com/nlpodyssey/spago/nn/activation"
)

var _ nn.Model = &Model{}

// Model contains the serializable parameters.
type Model struct {
	nn.Module
	Config Config
	Layers nn.ModuleList[nn.StandardModel]
}

// Config provides configuration parameters for a the gMLP Model.
type Config struct {
	Dim        int
	Depth      int
	SeqLen     int
	FFMult     int
	Activation activation.Activation
	// TODO: ProbSurvival T
}

func init() {
	gob.Register(&Model{})
}

// New returns a new Model.
func New[T float.DType](config Config) *Model { _ = "STUB: not implemented"; return nil }

// Forward performs the forward step. It adds pads if necessary.
func (m *Model) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }
