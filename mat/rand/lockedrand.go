// Copyright 2020 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"sync"

	"github.com/nlpodyssey/spago/mat/internal/rand"
)

// LockedRand is an implementation of rand.Rand that is concurrency-safe.
// It is just a wrap of the standard rand.Rand with its operations protected by a sync.Mutex.
type LockedRand struct {
	lk sync.Mutex
	r  *rand.Rand
}

// NewLockedRand creates a new LockedRand that implements all Rand functions that is safe
// for concurrent use.
func NewLockedRand(seed uint64) *LockedRand { _ = "STUB: not implemented"; return nil }

// Seed uses the provided seed value to initialize the generator to a deterministic state.
// Seed should not be called concurrently with any other Rand method.
func (lr *LockedRand) Seed(seed uint64) { _ = "STUB: not implemented"; return }

// TwoInt63 generates 2 random int64 without locking twice.
func (lr *LockedRand) TwoInt63() (n1, n2 int64) { _ = "STUB: not implemented"; return 0, 0 }

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (lr *LockedRand) Int63() (n int64) { _ = "STUB: not implemented"; return 0 }

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (lr *LockedRand) Uint32() (n uint32) { _ = "STUB: not implemented"; return 0 }

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (lr *LockedRand) Uint64() (n uint64) { _ = "STUB: not implemented"; return 0 }

// Uint64n returns, as a uint64, a pseudo-random number in [0,n).
// It is guaranteed more uniform than taking a Source value mod n
// for any n that is not a power of 2.
func (lr *LockedRand) Uint64n(n uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func (lr *LockedRand) Int31() (n int32) { _ = "STUB: not implemented"; return 0 }

// Int returns a non-negative pseudo-random int.
func (lr *LockedRand) Int() (n int) { _ = "STUB: not implemented"; return 0 }

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (lr *LockedRand) Int63n(n int64) (r int64) { _ = "STUB: not implemented"; return 0 }

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (lr *LockedRand) Int31n(n int32) (r int32) { _ = "STUB: not implemented"; return 0 }

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (lr *LockedRand) Intn(n int) (r int) { _ = "STUB: not implemented"; return 0 }

// NormFloat64 returns a normally distributed value in the range
// [-math.MaxFloat64, +math.MaxFloat64] with standard normal
// distribution (mean = 0, stddev = 1).
func (lr *LockedRand) NormFloat64() (n float64) { _ = "STUB: not implemented"; return 0 }

// Float32 returns a pseudo-random number in [0.0,1.0).
func (lr *LockedRand) Float32() (n float32) { _ = "STUB: not implemented"; return 0 }

// Float64 returns a pseudo-random number in [0.0,1.0).
func (lr *LockedRand) Float64() (n float64) { _ = "STUB: not implemented"; return 0 }

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func (lr *LockedRand) Perm(n int) (r []int) { _ = "STUB: not implemented"; return nil }

// Read generates len(p) random bytes and writes them into p. It
// always returns len(p) and a nil error.
// Read should not be called concurrently with any other Rand method.
func (lr *LockedRand) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Shuffle pseudo-randomizes the order of elements using the default Source.
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func (lr *LockedRand) Shuffle(i int, swap func(i int, j int)) { _ = "STUB: not implemented"; return }
