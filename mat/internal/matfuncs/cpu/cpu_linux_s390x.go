// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

const (
	// bit mask values from /usr/include/bits/hwcap.h
	hwcap_ZARCH  = 2
	hwcap_STFLE  = 4
	hwcap_MSA    = 8
	hwcap_LDISP  = 16
	hwcap_EIMM   = 32
	hwcap_DFP    = 64
	hwcap_ETF3EH = 256
	hwcap_VX     = 2048
	hwcap_VXE    = 8192
)

func initS390Xbase() {
	_ = "STUB: not implemented"
	// test HWCAP bit vector
	return
}

// mandatory

// optional
