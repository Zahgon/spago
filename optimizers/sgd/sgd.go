// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sgd

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
)

// Config provides configuration settings for an SGD optimizer.
type Config struct {
	LR       float64
	Mu       float64
	Nesterov bool
}

// NewConfig returns a new SGD Config.
func NewConfig(lr, momentum float64, nesterov bool) Config {
	_ = "STUB: not implemented"
	return *new(Config)
}

//var _ optimizers.Strategy = &SGD[float32]{}

// SGD implements the SGD gradient descent optimization method.
type SGD[T float.DType] struct {
	Config
	Alpha float64
}

// New returns a new SGD optimizer, initialized according to the given configuration.
func New[T float.DType](c Config) *SGD[T] { _ = "STUB: not implemented"; return nil }

type State struct {
	V     mat.Matrix // velocity
	Buf   mat.Matrix // buffer
	VPrev mat.Matrix // previous velocity
	VTmp  mat.Matrix // temporary velocity
}

func init() {
	gob.Register(&State{})
}

func (o *SGD[T]) newStateFor(param mat.Matrix) *State { _ = "STUB: not implemented"; return nil }

func (o *SGD[T]) calculateParamUpdate(grads mat.Matrix, state *State) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

func (o *SGD[T]) calculateParamUpdateSGD(grads mat.Matrix, state *State) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

func (o *SGD[T]) calculateParamUpdateMomentum(grads mat.Matrix, state *State) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

func (o *SGD[T]) calculateParamUpdateNesterovMomentum(grads mat.Matrix, state *State) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

// += grad * alpha

func (o *SGD[T]) OptimizeParams(param *nn.Param) error { _ = "STUB: not implemented"; return nil }
