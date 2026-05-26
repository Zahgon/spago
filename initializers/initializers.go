// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package initializers

import (
	"math"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/rand"
	"github.com/nlpodyssey/spago/nn/activation"
)

var sqrt2 = math.Sqrt(2.0)

// Gain returns a coefficient that help to initialize the params in a way to keep gradients stable.
// Use it to find the gain value for Xavier initializations.
func Gain(f activation.Activation) float64 { _ = "STUB: not implemented"; return 0 }

// Uniform fills the input matrix m with a uniform distribution where a is the lower bound and b is the upper bound.
//
// The matrix is returned for convenience.
func Uniform(m mat.Matrix, min, max float64, generator *rand.LockedRand) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

// Normal fills the input matrix with random samples from a normal (Gaussian)
// distribution.
//
// The matrix is returned for convenience.
func Normal(m mat.Matrix, mean, std float64, generator *rand.LockedRand) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

// Constant fills the input matrix with the value n.
//
// The matrix is returned for convenience.
func Constant(m mat.Matrix, n float64) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

// Ones fills the input matrix with the scalar value `1`.
//
// The matrix is returned for convenience.
func Ones(m mat.Matrix) mat.Matrix {
	_ = "STUB: not implemented"
	return *

	// Zeros fills the input matrix with the scalar value `0`.
	//
	// The matrix is returned for convenience.
	new(mat.Matrix)
}

func Zeros(m mat.Matrix) mat.Matrix {
	_ = "STUB: not implemented"
	return *

	// XavierUniform fills the input `m` with values according to the method described in `Understanding the difficulty of training deep
	// feedforward  neural networks` - Glorot, X. & Bengio, Y. (2010), using a uniform distribution.
	//
	// The matrix is returned for convenience.
	new(mat.Matrix)
}

func XavierUniform(m mat.Matrix, gain float64, generator *rand.LockedRand) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

// XavierNormal fills the input matrix with values according to the method
// described in "Understanding the difficulty of training deep feedforward
// neural networks" - Glorot, X. & Bengio, Y. (2010), using a normal
// distribution.
//
// The matrix is returned for convenience.
func XavierNormal(m mat.Matrix, gain float64, generator *rand.LockedRand) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

// Achlioptas fills the input matrix with values according to the mthod
// described on "Database-friendly random projections: Johnson-Lindenstrauss
// with binary coins", by Dimitris Achlioptas 2001
// (https://core.ac.uk/download/pdf/82724427.pdf)
//
// The matrix is returned for convenience.
func Achlioptas(m mat.Matrix, generator *rand.LockedRand) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}
