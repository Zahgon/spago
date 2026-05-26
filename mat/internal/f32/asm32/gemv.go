// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm32

// GemvN computes
//
//	y = alpha * A * x + beta * y
//
// where A is an m×n dense matrix, x and y are vectors, and alpha and beta are scalars.
func GemvN(m, n uintptr, alpha float32, a []float32, lda uintptr, x []float32, incX uintptr, beta float32, y []float32, incY uintptr) {
	_ = "STUB: not implemented"
	return
}

// GemvT computes
//
//	y = alpha * Aᵀ * x + beta * y
//
// where A is an m×n dense matrix, x and y are vectors, and alpha and beta are scalars.
func GemvT(m, n uintptr, alpha float32, a []float32, lda uintptr, x []float32, incX uintptr, beta float32, y []float32, incY uintptr) {
	_ = "STUB: not implemented"
	return
}

// beta == 0 is special-cased to memclear
