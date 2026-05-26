// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || amd64 || amd64p32
// +build 386 amd64 amd64p32

package cpu

const cacheLineSize = 64

func initOptions() { _ = "STUB: not implemented"; return }

// These capabilities should always be enabled on amd64:

func archInit() { _ = "STUB: not implemented"; return }

// For XGETBV, OSXSAVE bit is required and sufficient.

// Check if XMM and YMM registers have OS support.

// Darwin doesn't save/restore AVX-512 mask registers correctly across signal handlers.
// Since users can't rely on mask register contents, let's not advertise AVX-512 support.
// See issue 49233.

// Check if OPMASK and ZMM registers have OS support.

// Because avx-512 foundation is the core required extension

func isSet(bitpos uint, value uint32) bool { _ = "STUB: not implemented"; return false }
