// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adagrad

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
)

// Config provides configuration settings for an AdaGrad optimizer.
type Config struct {
	LR      float64
	Epsilon float64
}

// NewConfig returns a new AdaGrad Config.
func NewConfig(lr, epsilon float64) Config { _ = "STUB: not implemented"; return *new(Config) }

type State struct {
	M mat.Matrix // sum of squares of historical gradients
}

func init() {
	gob.Register(&State{})
}

// NewDefaultConfig returns a new Config with generically reasonable default values.
func NewDefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// AdaGrad assigns a different learning rate to each parameter using the sum of squares of its all historical gradients.
// References
//
//	Adaptive Subgradient Methods for Online Learning and Stochastic Optimization
//	http://www.jmlr.org/papers/volume12/duchi11a/duchi11a.pdf
type AdaGrad[T float.DType] struct {
	Config
}

// New returns a new AdaGrad optimizer, initialized according to the given configuration.
func New[T float.DType](c Config) *AdaGrad[T] { _ = "STUB: not implemented"; return nil }

func (o *AdaGrad[T]) newState(shape ...int) *State { _ = "STUB: not implemented"; return nil }

// m = m + grads*grads
// delta = (grads / (sqrt(m) + eps)) * lr
func (o *AdaGrad[T]) calculateParamUpdate(grads mat.Matrix, state *State) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

func (o *AdaGrad[T]) OptimizeParams(param *nn.Param) error { _ = "STUB: not implemented"; return nil }
