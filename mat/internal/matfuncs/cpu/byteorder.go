// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

// byteOrder is a subset of encoding/binary.ByteOrder.
type byteOrder interface {
	Uint32([]byte) uint32
	Uint64([]byte) uint64
}

type littleEndian struct{}
type bigEndian struct{}

func (littleEndian) Uint32(b []byte) uint32 {
	_ = "STUB: not implemented"
	// bounds check hint to compiler; see golang.org/issue/14808
	return 0
}

func (littleEndian) Uint64(b []byte) uint64 {
	_ = "STUB: not implemented"
	// bounds check hint to compiler; see golang.org/issue/14808
	return 0
}

func (bigEndian) Uint32(b []byte) uint32 {
	_ = "STUB: not implemented"
	// bounds check hint to compiler; see golang.org/issue/14808
	return 0
}

func (bigEndian) Uint64(b []byte) uint64 {
	_ = "STUB: not implemented"
	// bounds check hint to compiler; see golang.org/issue/14808
	return 0
}

// hostByteOrder returns littleEndian on little-endian machines and
// bigEndian on big-endian machines.
func hostByteOrder() byteOrder { _ = "STUB: not implemented"; return *new(byteOrder) }
