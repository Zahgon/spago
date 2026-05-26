// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package adanorm implements the Adaptive Normalization (AdaNorm) method.
//
// Reference: "Understanding and Improving Layer Normalization" by Jingjing Xu, Xu Sun,
// Zhiyuan Zhang,Guangxiang Zhao, Junyang Lin (2019).
// (https://papers.nips.cc/paper/8689-understanding-and-improving-layer-normalization.pdf)
package adanorm

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
)

var _ nn.Model = &Model{}

// Model contains the scaling factor.
type Model struct {
	nn.Module
	Scale *nn.Buffer
}

func init() {
	gob.Register(&Model{})
}

// New returns a new model.
func New[T float.DType](scale float64) *Model { _ = "STUB: not implemented"; return nil }

// Forward performs the forward step for each input node and returns the result.
func (m *Model) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

// detach the gradient of fi and only treat it as a changeable constant in implementation

// Mean computes the mean of the input.
func (m *Model) Mean(xs []mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

// StdDev computes the standard deviation of the input.
func (m *Model) StdDev(meanVectors []mat.Tensor, xs []mat.Tensor) []mat.Tensor {
	_ = "STUB: not implemented"
	return nil
}
