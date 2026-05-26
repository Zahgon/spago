// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bernulli

import (
	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/mat/rand"
)

// Distribution creates a new matrix initialized with Bernoulli distribution.
func Distribution[T float.DType](r, c int, prob T, generator *rand.LockedRand) mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}
