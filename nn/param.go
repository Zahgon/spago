// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nn

import "github.com/nlpodyssey/spago/mat"

type Param struct {
	mat.Matrix
	State interface{} // support structure for the optimization algorithm
}

// NewParam returns a new param.
func NewParam(value mat.Matrix) *Param { _ = "STUB: not implemented"; return nil }

// WithGrad sets whether the param requires gradients (default true)
func (p *Param) WithGrad(value bool) *Param { _ = "STUB: not implemented"; return nil }

func (p *Param) ReplaceValue(value mat.Matrix) { _ = "STUB: not implemented"; return }
