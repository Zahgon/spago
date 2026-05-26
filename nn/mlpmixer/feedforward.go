// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mlpmixer

import (
	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
	"github.com/nlpodyssey/spago/nn/activation"
)

// FeedForward is the model for feed-forward operations of a MixerBlock.
type FeedForward struct {
	nn.Module
	Layers nn.ModuleList[nn.StandardModel]
}

func newFeedForward[T float.DType](dim, hiddenDim int, act activation.Activation, dropout T) *FeedForward {
	_ = "STUB: not implemented"
	return nil
}

// dropout.New(dropout),

// dropout.New(dropout),

func (m *FeedForward) Forward(xs ...mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }
