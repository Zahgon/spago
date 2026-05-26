// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradclipper

import (
	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/nn"
)

// GradClipper performs gradient clipping on a set of parameters.
type GradClipper interface {
	// ClipGrads clips the gradients in place.
	ClipGrads(parameters nn.ParamChannelFunc)
}

// ValueClipper is a GradClipper which clips the values of a matrix between -Value and +Value.
type ValueClipper struct {
	Value float64
}

// ClipGrads clips the gradients in place between -Value and +Value.
func (c *ValueClipper) ClipGrads(parameters nn.ParamChannelFunc) { _ = "STUB: not implemented"; return }

// NormClipper is a GradClipper which clips the values of a matrix according to the NormType.
type NormClipper struct {
	MaxNorm  float64
	NormType float64
}

// validateNormType ensures that the NormType is greater than 1.
func (c *NormClipper) validateNormType() { _ = "STUB: not implemented"; return }

// calculateTotalNorm calculates the total norm based on NormType and matrices gs.
func (c *NormClipper) calculateTotalNorm(gs []mat.Tensor) float64 {
	_ = "STUB: not implemented"
	return 0
}

// ClipGradients clips the gradients, multiplying each parameter by the MaxNorm, divided by n-norm of the overall gradients.
// NormType is the n-norm. Can be “Double.POSITIVE_INFINITY“ for infinity norm (default 2.0)
func (c *NormClipper) ClipGradients(parameters nn.ParamChannelFunc) {
	_ = "STUB: not implemented"
	return
}

// collectGradients collects all the gradients from the parameters channel and returns them as a slice.
func collectGradients(parameters nn.ParamChannelFunc) []mat.Tensor {
	_ = "STUB: not implemented"
	return nil
}
