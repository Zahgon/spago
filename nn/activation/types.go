// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package activation

import (
	"github.com/nlpodyssey/spago/ag"
	"github.com/nlpodyssey/spago/mat"
)

// Activation is the enumeration-like type used for the set of built-in activations.
type Activation int

const (
	Identity Activation = iota
	Tan
	Tanh
	Sigmoid
	HardSigmoid
	HardTanh
	Softsign
	ReLU
	CELU
	GELU
	ELU
	PositiveELU
	SwishB
	Swish
	SiLU
	Mish
	LeakyReLU
	SELU
	SoftPlus
	SoftShrink
	Threshold
	Softmax
	LogSoftmax
	SparseMax
)

var activationsMap = map[Activation]string{
	Identity:    "Identity",
	Tan:         "Tan",
	Tanh:        "Tanh",
	Sigmoid:     "Sigmoid",
	HardSigmoid: "HardSigmoid",
	HardTanh:    "HardTanh",
	Softsign:    "Softsign",
	ReLU:        "ReLU",
	CELU:        "CELU",
	GELU:        "GELU",
	ELU:         "ELU",
	PositiveELU: "PositiveELU",
	SwishB:      "SwishB",
	Swish:       "Swish",
	SiLU:        "SiLU",
	Mish:        "Mish",
	LeakyReLU:   "LeakyReLU",
	SELU:        "SELU",
	SoftPlus:    "SoftPlus",
	SoftShrink:  "SoftShrink",
	Threshold:   "Threshold",
	Softmax:     "Softmax",
	LogSoftmax:  "LogSoftmax",
	SparseMax:   "SparseMax",
}

var activationFunctions = map[Activation]func(x mat.Tensor) mat.Tensor{
	Identity:    func(x mat.Tensor) mat.Tensor { return x },
	Tan:         ag.Tan,
	Tanh:        ag.Tanh,
	Sigmoid:     ag.Sigmoid,
	HardSigmoid: ag.HardSigmoid,
	HardTanh:    ag.HardTanh,
	Softsign:    ag.Softsign,
	ReLU:        ag.ReLU,
	GELU:        ag.GELU,
	PositiveELU: ag.PositiveELU,
	Swish:       ag.Swish,
	SiLU:        ag.SiLU,
	Mish:        ag.Mish,
	Softmax:     ag.Softmax,
	LogSoftmax:  ag.LogSoftmax,
	SparseMax:   ag.SparseMax,
}

var strActivationMap = strToActivationMap()

// strToName maps a string to a Activation.
func strToActivationMap() map[string]Activation { _ = "STUB: not implemented"; return nil }

// ParseActivation maps a string to an activation function.
// It returns an error if the string does not match any built-in activation (not even using lowercase).
func ParseActivation(str string) (Activation, error) {
	_ = "STUB: not implemented"
	return *new(Activation), nil
}

// MustParseActivation maps a string to an activation function.
// It panics if the string does not match any built-in activation (not even using lowercase).
func MustParseActivation(str string) Activation { _ = "STUB: not implemented"; return *new(Activation) }
