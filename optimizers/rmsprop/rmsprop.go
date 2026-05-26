// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rmsprop

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
)

// Config provides configuration settings for an RMSProp optimizer.
type Config struct {
	LR      float64
	Epsilon float64
	Decay   float64
}

// NewConfig returns a new RMSProp Config.
func NewConfig(lr, epsilon, decay float64) Config { _ = "STUB: not implemented"; return *new(Config) }

// NewDefaultConfig returns a new Config with generically reasonable default values.
func NewDefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

//var _ optimizers.Strategy = &RMSProp[float32]{}

// The RMSProp method is a variant of AdaGrad where the squared sum of previous gradients is replaced with a moving average.
// References:
//
//	RMSProp: Divide the gradient by a running average of its recent magnitude
//	http://www.cs.toronto.edu/~tijmen/csc321/slides/lecture_slides_lec6.pdf
type RMSProp[T float.DType] struct {
	Config
}

// New returns a new RMSProp optimizer, initialized according to the given configuration.
func New[T float.DType](c Config) *RMSProp[T] { _ = "STUB: not implemented"; return nil }

type State struct {
	V mat.Matrix // first moment vector
}

func init() {
	gob.Register(&State{})
}

func (o *RMSProp[T]) newState(shape ...int) *State { _ = "STUB: not implemented"; return nil }

func (o *RMSProp[T]) calculateParamUpdate(grads mat.Matrix, state *State) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

func (o *RMSProp[T]) OptimizeParams(param *nn.Param) error { _ = "STUB: not implemented"; return nil }
