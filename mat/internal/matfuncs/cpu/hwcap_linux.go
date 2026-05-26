// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

const (
	_AT_HWCAP  = 16
	_AT_HWCAP2 = 26

	procAuxv = "/proc/self/auxv"

	uintSize = int(32 << (^uint(0) >> 63))
)

// For those platforms don't have a 'cpuid' equivalent we use HWCAP/HWCAP2
// These are initialized in cpu_$GOARCH.go
// and should not be changed after they are initialized.
var hwCap uint
var hwCap2 uint

func readHWCAP() error { _ = "STUB: not implemented"; return nil }

// e.g. on android /proc/self/auxv is not accessible, so silently
// ignore the error and leave Initialized = false. On some
// architectures (e.g. arm64) doinit() implements a fallback
// readout and will set Initialized = true again.
