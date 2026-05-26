// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package float

// Float is implemented by any value that can be resolved
// to the types constrained by DType.
type Float interface {
	// F32 returns the value as float32, converting it if necessary.
	F32() float32
	// F64 returns the value as float64, converting it if necessary.
	F64() float64
	// BitSize returns the size in bits of the internal float value type.
	BitSize() int
}

// Interface converts a concrete DType value to an internal representation
// compatible with Float.
func Interface[T DType](v T) Float {
	_ = "STUB: not implemented"
	return *

	// ValueOf converts a Float value to a concrete DType.
	new(Float)
}

func ValueOf[T DType](i Float) T { _ = "STUB: not implemented"; return *new(T) }

// float is the built-in implementation of a Float.
type float[T DType] struct {
	v T
}

// F32 returns the value as float32, converting it if necessary.
func (f float[_]) F32() float32 { _ = "STUB: not implemented"; return 0 }

// F64 returns the value as float64, converting it if necessary.
func (f float[_]) F64() float64 { _ = "STUB: not implemented"; return 0 }

// BitSize returns the size in bits of the internal float value type.
func (f float[T]) BitSize() int { _ = "STUB: not implemented"; return 0 }

// String returns the value as a string.
func (f float[_]) String() string { _ = "STUB: not implemented"; return "" }
