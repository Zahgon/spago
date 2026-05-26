// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (386 || amd64 || amd64p32) && gccgo
// +build 386 amd64 amd64p32
// +build gccgo

package cpu

//extern gccgoGetCpuidCount
func gccgoGetCpuidCount(eaxArg, ecxArg uint32, eax, ebx, ecx, edx *uint32)

func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

//extern gccgoXgetbv
func gccgoXgetbv(eax, edx *uint32)

func xgetbv() (eax, edx uint32) { _ = "STUB: not implemented"; return 0, 0 }

// gccgo doesn't build on Darwin, per:
// https://github.com/Homebrew/homebrew-core/blob/HEAD/Formula/gcc.rb#L76
func darwinSupportsAVX512() bool { _ = "STUB: not implemented"; return false }
