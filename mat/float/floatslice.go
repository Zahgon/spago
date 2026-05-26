// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package float

// Slice is implemented by any value that can be resolved
// to a slice of a type constrained by DType.
type Slice interface {
	F32() []float32
	F64() []float64
	BitSize() int
	Len() int
	Equals(other Slice) bool
	InDelta(other Slice, delta float64) bool
}

// Make converts a concrete slice of DType values to an internal
// representation compatible with Slice.
func Make[T DType](v ...T) Slice {
	_ = "STUB: not implemented"
	return *

	// SliceValueOf converts a Slice value to a concrete slice
	// of DType values.
	new(Slice)
}

func SliceValueOf[T DType](v Slice) []T { _ = "STUB: not implemented"; return nil }

// floatSlice is the built-in implementation of a Slice.
type floatSlice[T DType] []T

// F32 returns the value as []float32, converting it if necessary.
func (fs floatSlice[T]) F32() []float32 { _ = "STUB: not implemented"; return nil }

// F64 returns the value as []float64, converting it if necessary.
func (fs floatSlice[T]) F64() []float64 { _ = "STUB: not implemented"; return nil }

// BitSize returns the size in bits of the internal float value type.
func (fs floatSlice[T]) BitSize() int { _ = "STUB: not implemented"; return 0 }

// Len returns the length of the slice.
func (fs floatSlice[_]) Len() int {
	_ = "STUB: not implemented"

	// Equals reports whether the content of the receiver is equal to the
	// content of the other slice.
	// The data type of the other slice is converted to the same type of
	// the receiver, if necessary.
	return 0
}

func (fs floatSlice[T]) Equals(other Slice) bool { _ = "STUB: not implemented"; return false }

// InDelta reports whether the receiver and the other slice have the same
// length and all their values at the same positions are within delta.
// The data type of the other slice is converted to the same type of
// the receiver, if necessary.
func (fs floatSlice[T]) InDelta(other Slice, delta float64) bool {
	_ = "STUB: not implemented"
	return false
}

// convertFloatSlice converts the given source slice to the destination type
// if necessary. If the source is nil, it returns nil. If the source type
// is identical to the destination type, the source is returned directly.
func convertFloatSlice[S, D DType](source []S) []D { _ = "STUB: not implemented"; return nil }
