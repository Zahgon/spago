// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ag

import (
	"github.com/nlpodyssey/spago/mat/rand"
)

var globalGenerator = rand.NewLockedRand(12345)

// Seed sets the seed for generating random numbers to the current time (converted to uint64).
func Seed() *rand.LockedRand { _ = "STUB: not implemented"; return nil }

// ManualSeed sets the seed for generating random numbers.
func ManualSeed(seed uint64) *rand.LockedRand { _ = "STUB: not implemented"; return nil }

// Rand returns the global random number generator.
func Rand() *rand.LockedRand { _ = "STUB: not implemented"; return nil }
