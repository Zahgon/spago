// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package optimizers

import (
	"github.com/nlpodyssey/spago/nn"
)

// OptimizationStrategy is the interface implemented by AdaGrad, Adam, etc.
type OptimizationStrategy interface {
	OptimizeParams(*nn.Param) error
}

// Optimizer is an optimizer that can optimize a set of parameters.
type Optimizer struct {
	// parameters is a function that returns a channel of parameters to optimize.
	parameters nn.ParamChannelFunc
	// strategy is the optimization strategy to use.
	strategy OptimizationStrategy
}

// New returns a new optimizer.
func New(parameters nn.ParamChannelFunc, strategy OptimizationStrategy) *Optimizer {
	_ = "STUB: not implemented"
	return nil
}

// Optimize performs the optimization of the parameters.
func (o *Optimizer) Optimize() error { _ = "STUB: not implemented"; return nil }

// As soon as an error occurs, stop the iteration over parameters
// Wait for running goroutines to finish
