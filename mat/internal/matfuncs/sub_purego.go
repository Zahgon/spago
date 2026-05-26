// Copyright 2022 The NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !amd64 || !gc || purego

package matfuncs

// Sub32 subtracts x2 from x1, element-wise, storing the result in y (32 bits).
func Sub32(x1, x2, y []float32) {
	_ = "STUB: not implemented"

	// Sub64 subtracts x2 from x1, element-wise, storing the result in y (64 bits).
	return
}

func Sub64(x1, x2, y []float64) { _ = "STUB: not implemented"; return }

func sub[F float32 | float64](x1, x2, y []F) { _ = "STUB: not implemented"; return }
