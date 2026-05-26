// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exponential

// Exponential defines an exponential decay depending on the time step:
//
//	lr = exp((times - t) * log(lr) + log(final))
type Exponential struct {
	init  float64
	final float64
	times int
}

// New returns a new Exponential decay optimizer.
func New(init, final float64, iter int) *Exponential { _ = "STUB: not implemented"; return nil }

// Decay calculates the decay of the learning rate lr at time t.
func (d *Exponential) Decay(lr float64, t int) float64 { _ = "STUB: not implemented"; return 0 }
