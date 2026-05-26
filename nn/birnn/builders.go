// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package birnn

import (
	"github.com/nlpodyssey/spago/mat/float"
)

// NewBiLSTM returns a new Bidirectional LSTM Model.
func NewBiLSTM[T float.DType](input, hidden int, merge MergeType) *Model {
	_ = "STUB: not implemented"
	return nil
}

// NewBiGRU returns a new Bidirectional GRU Model.
func NewBiGRU[T float.DType](input, hidden int, merge MergeType) *Model {
	_ = "STUB: not implemented"
	return nil
}

// NewBiBiLSTM returns a new Bidirectional BiLSTM Model.
func NewBiBiLSTM[T float.DType](input, hidden int, merge MergeType) []*Model {
	_ = "STUB: not implemented"
	return nil
}
