// Copyright 2021 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gmlp

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
	"github.com/nlpodyssey/spago/nn/normalization/layernorm"
)

var _ nn.Model = &PreNorm{}

// PreNorm is a helper model to perform pre-normalization.
type PreNorm struct {
	nn.Module
	Block *Block
	Norm  *layernorm.Model
}

func init() {
	gob.Register(&PreNorm{})
}

// NewPreNorm returns a new PreNorm.
func NewPreNorm[T float.DType](dim int, block *Block) *PreNorm {
	_ = "STUB: not implemented"
	return nil
}

// Forward performs the forward step.
func (m *PreNorm) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }
