// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"github.com/nlpodyssey/spago/mat/float"
)

// SmallestNonzero returns the smallest positive, non-zero value representable by the type.
func SmallestNonzero[T float.DType]() T { _ = "STUB: not implemented"; return *new(T) }

// Pi mathematical constant.
func Pi[T float.DType]() T {
	_ = "STUB: not implemented"

	// Pow returns x**y, the base-x exponential of y.
	return *new(T)
}

func Pow[T float.DType](x, y T) T { _ = "STUB: not implemented"; return *new(T) }

// Cos returns the cosine of the radian argument x.
func Cos[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }

// Sin returns the sine of the radian argument x.
func Sin[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }

// Cosh returns the hyperbolic cosine of x.
func Cosh[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }

// Sinh returns the hyperbolic sine of x.
func Sinh[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }

// Exp returns e**x, the base-e exponential of x.
func Exp[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }

// Abs returns the absolute value of x.
func Abs[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }

// Sqrt returns the square root of x.
func Sqrt[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }

// Log returns the natural logarithm of x.
func Log[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }

// Tan returns the tangent of the radian argument x.
func Tan[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }

// Tanh returns the hyperbolic tangent of x.
func Tanh[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }

// Max returns the larger of x or y.
func Max[T float.DType](x, y T) T { _ = "STUB: not implemented"; return *new(T) }

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf[T float.DType](sign int) T {
	_ = "STUB: not implemented"
	return *

	// IsInf reports whether f is an infinity, according to sign.
	new(T)
}

func IsInf[T float.DType](f T, sign int) bool { _ = "STUB: not implemented"; return false }

// NaN returns an IEEE 754 “not-a-number” value.
func NaN[T float.DType]() T {
	_ = "STUB: not implemented"
	return *

	// Ceil returns the least integer value greater than or equal to x.
	new(T)
}

func Ceil[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }

// Floor returns the greatest integer value less than or equal to x.
func Floor[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }

// Round returns the nearest integer, rounding half away from zero.
func Round[T float.DType](x T) T { _ = "STUB: not implemented"; return *new(T) }
