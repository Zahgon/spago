// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hyperbolic

// Hyperbolic defines an hyperbolic decay depending on the time step
//
//	lr = lr / (1 + rate*t).
type Hyperbolic struct {
	init  float64
	final float64
	rate  float64
}

// New returns a new Hyperbolic decay optimizer.
func New(init, final, rate float64) *Hyperbolic { _ = "STUB: not implemented"; return nil }

// Decay calculates the decay of the learning rate lr at time t.
func (d *Hyperbolic) Decay(lr float64, t int) float64 { _ = "STUB: not implemented"; return 0 }
