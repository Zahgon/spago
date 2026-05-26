// Copyright 2022 The NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !amd64 || !gc || purego

package matfuncs

// Log32 computes the natural logarithm of each element of x, storing the result in y (32 bits).
func Log32(x, y []float32) {
	_ = "STUB: not implemented"

	// Log64 computes the natural logarithm of each element of x, storing the result in y (64 bits).
	return
}

func Log64(x, y []float64) { _ = "STUB: not implemented"; return }
