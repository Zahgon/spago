// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gccgo
// +build gccgo

package cpu

// haveAsmFunctions reports whether the other functions in this file can
// be safely called.
func haveAsmFunctions() bool {
	_ = "STUB: not implemented"

	// TODO(mundaym): the following feature detection functions are currently
	// stubs. See https://golang.org/cl/162887 for how to fix this.
	// They are likely to be expensive to call so the results should be cached.
	return false
}

func stfle() facilityList     { _ = "STUB: not implemented"; return *new(facilityList) }
func kmQuery() queryResult    { _ = "STUB: not implemented"; return *new(queryResult) }
func kmcQuery() queryResult   { _ = "STUB: not implemented"; return *new(queryResult) }
func kmctrQuery() queryResult { _ = "STUB: not implemented"; return *new(queryResult) }
func kmaQuery() queryResult   { _ = "STUB: not implemented"; return *new(queryResult) }
func kimdQuery() queryResult  { _ = "STUB: not implemented"; return *new(queryResult) }
func klmdQuery() queryResult  { _ = "STUB: not implemented"; return *new(queryResult) }
