// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

const cacheLineSize = 64

func initOptions() { _ = "STUB: not implemented"; return }

func archInit() { _ = "STUB: not implemented"; return }

// Most platforms don't seem to allow reading these registers.
//
// OpenBSD:
// See https://golang.org/issue/31746

// setMinimalFeatures fakes the minimal ARM64 features expected by
// TestARM64minimalFeatures.
func setMinimalFeatures() { _ = "STUB: not implemented"; return }

func readARM64Registers() { _ = "STUB: not implemented"; return }

func parseARM64SystemRegisters(isar0, isar1, pfr0 uint64) {
	_ = "STUB: not implemented"
	// ID_AA64ISAR0_EL1
	return
}

// ID_AA64ISAR1_EL1

// ID_AA64PFR0_EL1

func extractBits(data uint64, start, end uint) uint { _ = "STUB: not implemented"; return 0 }
