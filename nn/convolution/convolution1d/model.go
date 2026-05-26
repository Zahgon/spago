// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package convolution1d

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
	"github.com/nlpodyssey/spago/nn/activation"
)

var _ nn.Model = &Model{}

// Config provides configuration settings for a convolution Model.
type Config struct {
	KernelSizeX    int
	KernelSizeY    int
	YStride        int
	InputChannels  int
	OutputChannels int
	Mask           []int
	DepthWise      bool // Special case od depth-wise convolution, where output channels == input channels
	Activation     activation.Activation
}

// Model contains the serializable parameters for a convolutional neural network model.
type Model struct {
	nn.Module
	Config Config
	K      []*nn.Param
	B      []*nn.Param
}

func init() {
	gob.Register(&Model{})
}

// New returns a new convolution Model, initialized according to the given configuration.
func New[T float.DType](config Config) *Model { _ = "STUB: not implemented"; return nil }

// Forward performs the forward step for each input node and returns the result.
func (m *Model) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

func (m *Model) forward(xs []mat.Tensor, outputChannel int) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// TODO: refactor for performance
