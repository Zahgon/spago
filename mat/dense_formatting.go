// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"fmt"

	"github.com/nlpodyssey/spago/mat/float"
)

// Format implements custom formatting for representing a Dense matrix.
// Thanks to this method, a Dense matrix satisfies the fmt.Formatter interface.
func (d *Dense[_]) Format(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

// TODO: print the shape

// TODO: print the shape

// %F (alias for %f) does not work with strconv.AppendFloat

// format formats a non-empty Dense matrix
func (d *Dense[_]) format(f fmt.State, c rune, precision int) { _ = "STUB: not implemented"; return }

func writeFormattedValue(f fmt.State, buf, spaceBuf []byte, maxW lrWidth) {
	_ = "STUB: not implemented"
	return
}

func makeSpaceBuffer(length int) []byte { _ = "STUB: not implemented"; return nil }

func (d *Dense[_]) formattingRowPrefixAndSuffix(rowIndex int) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

type lrWidth struct{ left, right int }

func (d *Dense[_]) formattingMaxColumnsWidth(
	f fmt.State,
	c rune,
	precision int,
) ([]lrWidth, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func formatValue[T float.DType](buf []byte, val T, c rune, precision int) []byte {
	_ = "STUB: not implemented"
	return nil
}

func maxInt(a, b int) int { _ = "STUB: not implemented"; return 0 }
