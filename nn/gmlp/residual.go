// Copyright 2021 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gmlp

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/nn"
)

var _ nn.Model = &Residual{}

// Residual is a helper model to perform residual connections.
type Residual struct {
	nn.Module
	PreNorm *PreNorm
}

func init() {
	gob.Register(&Residual{})
}

// NewResidual returns a new Residual.
func NewResidual(preNorm *PreNorm) *Residual { _ = "STUB: not implemented"; return nil }

// Forward performs the forward step.
func (m *Residual) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }
