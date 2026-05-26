// Copyright 2021 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"github.com/nlpodyssey/spago/mat/float"
)

// Float returns, as a T, a pseudo-random number in [0.0,1.0)
// from the default Source.
func Float[T float.DType]() T { _ = "STUB: not implemented"; return *new(T) }
