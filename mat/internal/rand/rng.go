// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

// PCGSource is an implementation of a 64-bit permuted congruential
// generator as defined in
//
//	PCG: A Family of Simple Fast Space-Efficient Statistically Good
//	Algorithms for Random Number Generation
//	Melissa E. O’Neill, Harvey Mudd College
//	http://www.pcg-random.org/pdf/toms-oneill-pcg-family-v1.02.pdf
//
// The generator here is the congruential generator PCG XSL RR 128/64 (LCG)
// as found in the software available at http://www.pcg-random.org/.
// It has period 2^128 with 128 bits of state, producing 64-bit values.
// Is state is represented by two uint64 words.
type PCGSource struct {
	low  uint64
	high uint64
}

const (
	multiplier = 47026247687942121848144207491837523525
	mulHigh    = multiplier >> 64
	mulLow     = multiplier & maxUint64

	increment = 117397592171526113268558934119004209487
	incHigh   = increment >> 64
	incLow    = increment & maxUint64

	// TODO: Use these?
	//initializer = 245720598905631564143578724636268694099
	//initHigh    = initializer >> 64
	//initLow     = initializer & maxUint64
)

// Seed uses the provided seed value to initialize the generator to a deterministic state.
func (pcg *PCGSource) Seed(seed uint64) { _ = "STUB: not implemented"; return }

// TODO: What is right?

// Uint64 returns a pseudo-random 64-bit unsigned integer as a uint64.
func (pcg *PCGSource) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

// XOR high and low 64 bits together and rotate right by high 6 bits of state.

func (pcg *PCGSource) add() { _ = "STUB: not implemented"; return }

func (pcg *PCGSource) multiply() { _ = "STUB: not implemented"; return }

// MarshalBinary returns the binary representation of the current state of the generator.
func (pcg *PCGSource) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary sets the state of the generator to the state represented in data.
func (pcg *PCGSource) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }
