// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adam

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/nn"
)

// Config provides configuration settings for an Adam optimizer.
type Config struct {
	StepSize float64
	Beta1    float64
	Beta2    float64
	Epsilon  float64
	Lambda   float64 // AdamW
}

// NewConfig returns a new Adam Config.
func NewConfig(stepSize, beta1, beta2, epsilon float64) Config {
	_ = "STUB: not implemented"
	return *new(Config)
}

// NewAdamWConfig returns a new Adam Config.
func NewAdamWConfig(stepSize, beta1, beta2, epsilon, lambda float64) Config {
	_ = "STUB: not implemented"
	return *new(Config)
}

// NewDefaultConfig returns a new Config with generically reasonable default values.
func NewDefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

//var _ optimizers.Strategy = &Adam[float32]{}

// Adam implements the Adam gradient descent optimization method.
type Adam struct {
	Config
	Alpha    float64
	TimeStep int
	adamw    bool
}

// New returns a new Adam optimizer, initialized according to the given configuration.
func New(c Config) *Adam { _ = "STUB: not implemented"; return nil }

// initialize 'alpha' coefficient

type State struct {
	V    mat.Matrix // first moment vector
	M    mat.Matrix // second raw moment vector
	Buf1 mat.Matrix // contains 'grads.ProdScalar(1.0 - beta1)'
	Buf2 mat.Matrix // contains 'grads.Prod(grads).ProdScalar(1.0 - beta2)'
	Buf3 mat.Matrix
}

func init() {
	gob.Register(&State{})
}

func (o *Adam) newStateFor(param mat.Matrix) *State { _ = "STUB: not implemented"; return nil }

// IncExample beats the occurrence of a new example.
func (o *Adam) IncExample() { _ = "STUB: not implemented"; return }

func (o *Adam) updateAlpha() { _ = "STUB: not implemented"; return }

// v = v*beta1 + grads*(1.0-beta1)
// m = m*beta2 + (grads*grads)*(1.0-beta2)
// d = (v / (sqrt(m) + eps)) * alpha
func (o *Adam) calculateParamUpdate(grads mat.Matrix, state *State) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

// v = v*beta1 + grads*(1.0-beta1)
// m = m*beta2 + (grads*grads)*(1.0-beta2)
// d = (v / (sqrt(m) + eps))  + (lambda * weights) + alpha
func (o *Adam) calculateParamUpdateW(grads mat.Matrix, state *State, weights mat.Matrix) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

// v = v*beta1 + grads*(1.0-beta1)
func updateV(grads mat.Matrix, state *State, beta1 float64) { _ = "STUB: not implemented"; return }

// m = m*beta2 + (grads*grads)*(1.0-beta2)
func updateM(grads mat.Matrix, state *State, beta2 float64) { _ = "STUB: not implemented"; return }

func (o *Adam) OptimizeParams(param *nn.Param) error { _ = "STUB: not implemented"; return nil }
