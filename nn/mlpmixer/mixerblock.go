// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mlpmixer implements the MLP-Mixer (Tolstikhin et al., 2021).
package mlpmixer

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
	"github.com/nlpodyssey/spago/nn/activation"
	"github.com/nlpodyssey/spago/nn/normalization/layernorm"
)

var _ nn.Model = &MixerBlock{}

// MixerBlock contains the serializable parameters.
type MixerBlock struct {
	nn.Module
	Config
	TokenLayerNorm   *layernorm.Model
	TokenMixerFF     *FeedForward
	ChannelLayerNorm *layernorm.Model
	ChannelMixerFF   *FeedForward
}

// Config provides configuration settings for a MixerBlock.
type Config struct {
	InputSize               int
	HiddenSizeTokenMixer    int
	HiddenSizeChannelMixer  int
	Channels                int
	ActFunctionTokenMixer   activation.Activation
	ActFunctionChannelMixer activation.Activation
	Eps                     float64
}

func init() {
	gob.Register(&MixerBlock{})
}

// New returns a new model with parameters initialized to zeros.
func New[T float.DType](config Config) *MixerBlock { _ = "STUB: not implemented"; return nil }

// Forward performs the forward step for each input node and returns the result.
func (m *MixerBlock) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

func (m *MixerBlock) tokenMix(xs []mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

func (m *MixerBlock) channelMix(xs []mat.Tensor) []mat.Tensor {
	_ = "STUB: not implemented"
	return nil
}

func (m *MixerBlock) residual(xs []mat.Tensor, residual []mat.Tensor) []mat.Tensor {
	_ = "STUB: not implemented"
	return nil
}
