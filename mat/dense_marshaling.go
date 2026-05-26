// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"encoding/gob"
)

func init() {
	gob.Register(&Dense[float32]{})
	gob.Register(&Dense[float64]{})
}

// MarshalBinary marshals a Dense matrix into binary form.
func (d *Dense[T]) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary unmarshals a binary representation of a Dense matrix.
func (d *Dense[T]) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (d *Dense[T]) marshalBinaryFloat32() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Dense[T]) unmarshalBinaryFloat32(data []byte) error { _ = "STUB: not implemented"; return nil }

func (d *Dense[T]) marshalBinaryFloat64() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Dense[T]) unmarshalBinaryFloat64(data []byte) error { _ = "STUB: not implemented"; return nil }

func bytesToSlice[T any](b []byte, length int) []T { _ = "STUB: not implemented"; return nil }
