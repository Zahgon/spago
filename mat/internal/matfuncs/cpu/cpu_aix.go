// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix
// +build aix

package cpu

const (
	// getsystemcfg constants
	_SC_IMPL     = 2
	_IMPL_POWER8 = 0x10000
	_IMPL_POWER9 = 0x20000
)

func archInit() { _ = "STUB: not implemented"; return }

func getsystemcfg(label int) (n uint64) { _ = "STUB: not implemented"; return 0 }
