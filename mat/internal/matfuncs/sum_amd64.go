// Copyright 2022 The NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 && gc && !purego

package matfuncs

var (
	sum32 = SumSSE32
	sum64 = SumSSE64
)

func init() {
	if hasAVX {
		sum32 = SumAVX32
		sum64 = SumAVX64
	}
}

// Sum32 returns the sum of all values of x (32 bits).
func Sum32(x []float32) float32 {
	_ = "STUB: not implemented"

	// Sum64 returns the sum of all values of x (64 bits).
	return 0
}

func Sum64(x []float64) float64 { _ = "STUB: not implemented"; return 0 }
