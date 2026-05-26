// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"github.com/nlpodyssey/spago/mat/float"
)

// ShuffleInPlace pseudo-randomizes the order of elements, modifying the
// given slice in-place.
func ShuffleInPlace(xs []int, generator *LockedRand) []int { _ = "STUB: not implemented"; return nil }

// Warning: use global rand

// WeightedChoice performs a random generation of the indices based of the probability distribution itself.
// Please note that it uses the global random.
func WeightedChoice[T float.DType](dist []T) int { _ = "STUB: not implemented"; return 0 }

// Warning: use global rand

// Warning: use global rand

// GetUniqueRandomInt generates n mutually exclusive integers up to max, using the default random source.
// The callback checks whether a generated number can be accepted, or not.
func GetUniqueRandomInt(n, max int, valid func(r int) bool) []int {
	_ = "STUB: not implemented"
	return nil
}

// Warning: use global rand

// Warning: use global rand

// GetUniqueRandomIndices select n mutually exclusive indices, using the global random.
// The callback checks whether an extracted index can be accepted, or not.
func GetUniqueRandomIndices(n int, indices []int, valid func(r int) bool) []int {
	_ = "STUB: not implemented"
	return nil
}

// The generic type is irrelevant, since the given generator is nil.
// TODO: ugly API of ShuffleInPlace to be refactored
// Warning: use global rand

// Warning: use global rand

// contains returns whether the list contains the x-element, or not.
func contains(lst []int, x int) bool { _ = "STUB: not implemented"; return false }
