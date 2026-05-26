// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package radam

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
)

// Config provides configuration settings for a RAdam optimizer.
type Config struct {
	StepSize float64
	Beta1    float64
	Beta2    float64
	Epsilon  float64
}

// NewConfig returns a new RAdam Config.
// It panics if beta1 or beta2 are not in the range [0.0, 1.0).
func NewConfig(stepSize, beta1, beta2, epsilon float64) Config {
	_ = "STUB: not implemented"
	return *new(Config)
}

// NewDefaultConfig returns a new Config with generically reasonable default values.
func NewDefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// RAdam implements the RAdam gradient descent optimization method.
type RAdam[T float.DType] struct {
	Config
	RoMax    float64 // The maximum length of the approximated SMA.
	TimeStep int
}

// New returns a new RAdam optimizer, initialized according to the given configuration.
func New[T float.DType](c Config) *RAdam[T] { _ = "STUB: not implemented"; return nil }

type State struct {
	M    mat.Matrix // first moment vector
	V    mat.Matrix // second moment vector
	Buf1 mat.Matrix // buffer for first moment vector
	Buf2 mat.Matrix // buffer for second moment vector
	Buf3 mat.Matrix // buffer for second moment vector
}

func init() {
	gob.Register(&State{})
}

// newState returns a new state.
func (o *RAdam[T]) newState(shape ...int) *State { _ = "STUB: not implemented"; return nil }

// IncBatch beats the occurrence of a new batch.
func (o *RAdam[_]) IncBatch() { _ = "STUB: not implemented"; return }

func (o *RAdam[T]) calculateParamUpdate(grads mat.Matrix, state *State) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

// m = m*beta1 + grads*(1.0-beta1)
func updateM(grads mat.Matrix, state *State, beta1 float64) { _ = "STUB: not implemented"; return }

// v = v*beta2 + (grads*grads)*(1.0-beta2)
func updateV(grads mat.Matrix, state *State, beta2 float64) { _ = "STUB: not implemented"; return }

func (o *RAdam[T]) calcAlpha() float64 { _ = "STUB: not implemented"; return 0 }

// i.e. if the variance is tractable

func (o *RAdam[T]) OptimizeParams(param *nn.Param) error { _ = "STUB: not implemented"; return nil }
