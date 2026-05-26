// Copyright 2021 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lamb

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
)

// Config provides configuration settings for Lamb optimizer.
type Config struct {
	StepSize float64
	Beta1    float64
	Beta2    float64
	Epsilon  float64
	Lambda   float64
}

// NewConfig returns a new Lamb Config.
func NewConfig(stepSize, beta1, beta2, epsilon, lambda float64) Config {
	_ = "STUB: not implemented"
	return *new(Config)
}

// NewDefaultConfig returns a new Config with generically reasonable default values.
func NewDefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// Lamb implements the Lamb gradient descent optimization method.
type Lamb[T float.DType] struct {
	Config
	Alpha    float64
	TimeStep int
}

// New returns a new Lamb optimizer, initialized according to the given configuration.
func New[T float.DType](c Config) *Lamb[T] { _ = "STUB: not implemented"; return nil }

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

func (o *Lamb[T]) newState(shape ...int) *State { _ = "STUB: not implemented"; return nil }

// IncExample beats the occurrence of a new example.
func (o *Lamb[_]) IncExample() { _ = "STUB: not implemented"; return }

func (o *Lamb[T]) updateAlpha() { _ = "STUB: not implemented"; return }

// CalcDelta returns the difference between the current params and where the method wants it to be.
func (o *Lamb[T]) CalcDelta(state *State, cur mat.Matrix, grads mat.Matrix) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

// v = v*beta1 + grads*(1.0-beta1)
// m = m*beta2 + (grads*grads)*(1.0-beta2)
// weights = ||params|| / || (v / (sqrt(m) + eps)) + (lambda * weights)
// d = (v / (sqrt(m) + eps)) + (lambda * weights) * alpha
func (o *Lamb[T]) calculateParamUpdate(grads mat.Matrix, state *State, weights mat.Matrix) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

// v = v*beta1 + grads*(1.0-beta1)
func updateV(grads mat.Matrix, state *State, beta1 float64) { _ = "STUB: not implemented"; return }

// m = m*beta2 + (grads*grads)*(1.0-beta2)
func updateM(grads mat.Matrix, state *State, beta2 float64) { _ = "STUB: not implemented"; return }

func norm(grads mat.Matrix) float64 { _ = "STUB: not implemented"; return 0 }

func (o *Lamb[T]) OptimizeParams(param *nn.Param) error { _ = "STUB: not implemented"; return nil }
