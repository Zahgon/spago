// Copyright 2021 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gmlp

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
	"github.com/nlpodyssey/spago/nn/activation"
)

var _ nn.Model = &Block{}

// Block is the core model of the gMLP.
type Block struct {
	nn.Module
	Layers nn.ModuleList[nn.StandardModel]
}

// BlockConfig provides configuration parameters for a single Block of the gMLP Model.
type BlockConfig struct {
	Dim        int
	DimFF      int
	SeqLen     int
	Activation activation.Activation
}

func init() {
	gob.Register(&Block{})
}

// NewBlock returns a new Block.
func NewBlock[T float.DType](config BlockConfig) *Block { _ = "STUB: not implemented"; return nil }

func (m *Block) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }
