// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nn

import (
	"io"
)

// Dump saves a serialized object to a stream. This function uses Gob utility for serialization.
// Models, matrices, and all kinds of Gob serializable objects can be saved using this function.
func Dump(obj any, w io.Writer) error { _ = "STUB: not implemented"; return nil }

// DumpToFile saves a serialized object to a file.
// See Dump for further details.
func DumpToFile[T any](obj T, filename string) error { _ = "STUB: not implemented"; return nil }

// Load uses Gob to deserialize objects to memory.
func Load[T any](r io.Reader) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

// LoadFromFile uses Gob to deserialize objects files to memory.
// See Load for further details.
func LoadFromFile[T any](filename string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
