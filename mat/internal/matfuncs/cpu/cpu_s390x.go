// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

const cacheLineSize = 256

func initOptions() { _ = "STUB: not implemented"; return }

// bitIsSet reports whether the bit at index is set. The bit index
// is in big endian order, so bit index 0 is the leftmost bit.
func bitIsSet(bits []uint64, index uint) bool { _ = "STUB: not implemented"; return false }

// facility is a bit index for the named facility.
type facility uint8

const (
	// mandatory facilities
	zarch  facility = 1  // z architecture mode is active
	stflef facility = 7  // store-facility-list-extended
	ldisp  facility = 18 // long-displacement
	eimm   facility = 21 // extended-immediate

	// miscellaneous facilities
	dfp    facility = 42 // decimal-floating-point
	etf3eh facility = 30 // extended-translation 3 enhancement

	// cryptography facilities
	msa  facility = 17  // message-security-assist
	msa3 facility = 76  // message-security-assist extension 3
	msa4 facility = 77  // message-security-assist extension 4
	msa5 facility = 57  // message-security-assist extension 5
	msa8 facility = 146 // message-security-assist extension 8
	msa9 facility = 155 // message-security-assist extension 9

	// vector facilities
	vx   facility = 129 // vector facility
	vxe  facility = 135 // vector-enhancements 1
	vxe2 facility = 148 // vector-enhancements 2
)

// facilityList contains the result of an STFLE call.
// Bits are numbered in big endian order so the
// leftmost bit (the MSB) is at index 0.
type facilityList struct {
	bits [4]uint64
}

// Has reports whether the given facilities are present.
func (s *facilityList) Has(fs ...facility) bool { _ = "STUB: not implemented"; return false }

// function is the code for the named cryptographic function.
type function uint8

const (
	// KM{,A,C,CTR} function codes
	aes128 function = 18 // AES-128
	aes192 function = 19 // AES-192
	aes256 function = 20 // AES-256

	// K{I,L}MD function codes
	sha1     function = 1  // SHA-1
	sha256   function = 2  // SHA-256
	sha512   function = 3  // SHA-512
	sha3_224 function = 32 // SHA3-224
	sha3_256 function = 33 // SHA3-256
	sha3_384 function = 34 // SHA3-384
	sha3_512 function = 35 // SHA3-512
	shake128 function = 36 // SHAKE-128
	shake256 function = 37 // SHAKE-256

	// KLMD function codes
	ghash function = 65 // GHASH
)

// queryResult contains the result of a Query function
// call. Bits are numbered in big endian order so the
// leftmost bit (the MSB) is at index 0.
type queryResult struct {
	bits [2]uint64
}

// Has reports whether the given functions are present.
func (q *queryResult) Has(fns ...function) bool { _ = "STUB: not implemented"; return false }

func doinit() {
	_ = "STUB: not implemented"

	// We need implementations of stfle, km and so on
	// to detect cryptographic features.
	return
}

// optional cryptographic functions

// cipher message

// compute message digest
// intermediate (no padding)
// last (padding)

// KLMD-GHASH does not exist
