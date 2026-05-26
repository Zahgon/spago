// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"github.com/nlpodyssey/spago/mat/float"
)

type Options struct {
	RequiresGrad bool // default: false
	Shape        []int
	Slice        float.Slice
}

type OptionsFunc func(opt *Options)

func WithGrad(value bool) OptionsFunc { _ = "STUB: not implemented"; return *new(OptionsFunc) }

func WithShape(shape ...int) OptionsFunc { _ = "STUB: not implemented"; return *new(OptionsFunc) }

func WithBacking[T float.DType](data []T) OptionsFunc {
	_ = "STUB: not implemented"
	return *new(OptionsFunc)
}

func NewDense[T float.DType](opts ...OptionsFunc) *Dense[T] { _ = "STUB: not implemented"; return nil }

// TODO: do not panic

func Scalar[T float.DType](value T, opts ...OptionsFunc) *Dense[T] {
	_ = "STUB: not implemented"
	return nil
}

// TODO: do not panic

func newDense[T float.DType](opts ...OptionsFunc) (*Dense[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newDenseFromSlice[T float.DType](args *Options) (*Dense[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newDenseFromShape[T float.DType](args *Options) (*Dense[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func calculateSize(shape []int) int { _ = "STUB: not implemented"; return 0 }

func checkShape(shape ...int) error { _ = "STUB: not implemented"; return nil }

func adjustShape(shape ...int) []int { _ = "STUB: not implemented"; return nil }

func newScalar[T float.DType](value T, opts ...OptionsFunc) (*Dense[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateInitializedSlice[T float.DType](size int, v T) []T {
	_ = "STUB: not implemented"
	return nil
}

func CreateOneHotVector[T float.DType](size int, index int) []T {
	_ = "STUB: not implemented"
	return nil
}

func CreateIdentityMatrix[T float.DType](size int) []T { _ = "STUB: not implemented"; return nil }

func InitializeMatrix[T float.DType](rows, cols int, fn func(r, c int) T) []T {
	_ = "STUB: not implemented"
	return nil
}
