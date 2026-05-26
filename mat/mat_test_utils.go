// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mattest provides utilities for testing code involving spaGO matrices.
package mat

// T requires a subset of methods from testing.TB.
// This interface is primarily useful to simplify the testing of the
// package itself.
type T interface {
	Helper()
	Error(args ...any)
	Errorf(format string, args ...any)
	FailNow()
}

// AssertMatrixEquals tests whether expected is equal to actual; if not,
// T.Error or T.Errorf are called, providing useful information.
//
// It returns whether the comparison succeeded.
//
// The expected matrix is not allowed to be nil, otherwise the function always
// produces an error.
func AssertMatrixEquals(t T, expected, actual Tensor, args ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// RequireMatrixEquals tests whether expected is equal to actual; if not,
// T.Error or T.Errorf are called, providing useful information, followed by
// T.FailNow.
//
// The expected matrix is not allowed to be nil, otherwise the function always
// produces an error and fails.
func RequireMatrixEquals(t T, expected, actual Matrix, args ...any) {
	_ = "STUB: not implemented"
	return
}

// AssertMatrixInDelta tests whether expected and actual have the same shape
// and all elements at the same positions are within delta; if not, T.Error
// or T.Errorf are called, providing useful information.
//
// It returns whether the comparison succeeded.
//
// The expected matrix is not allowed to be nil, otherwise the function always
// produces an error.
func AssertMatrixInDelta(t T, expected, actual Matrix, delta float64, args ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// RequireMatrixInDelta tests whether expected and actual have the same shape
// and all elements at the same positions are within delta; if not, T.Error
// or T.Errorf are called, providing useful information, followed by T.FailNow.
//
// The expected matrix is not allowed to be nil, otherwise the function always
// produces an error and fails.
func RequireMatrixInDelta(t T, expected, actual Matrix, delta float64, args ...any) {
	_ = "STUB: not implemented"
	return
}
