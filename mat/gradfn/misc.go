// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gradfn

import (
	"github.com/nlpodyssey/spago/mat"
)

// Tan is an operator to perform element-wise tangent.
type Tan[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewTan returns a new UnaryElementwise tangent function.
func NewTan[O mat.Tensor](x O) *Tan[O] { _ = "STUB: not implemented"; return nil }

// Tanh is an operator to perform element-wise hyperbolic tangent.
type Tanh[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewTanh returns a new UnaryElementwise hyperbolic tangent function.
func NewTanh[O mat.Tensor](x O) *Tanh[O] { _ = "STUB: not implemented"; return nil }

// HardSigmoid is an operator to perform element-wise hard sigmoid.
type HardSigmoid[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewHardSigmoid returns a new UnaryElementwise hard sigmoid function.
func NewHardSigmoid[O mat.Tensor](x O) *HardSigmoid[O] { _ = "STUB: not implemented"; return nil }

// HardTanh is an operator to perform element-wise hard hyperbolic tangent.
type HardTanh[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewHardTanh returns a new UnaryElementwise hard hyperbolic tangent function.
func NewHardTanh[O mat.Tensor](x O) *HardTanh[O] { _ = "STUB: not implemented"; return nil }

// ReLU is an operator to perform element-wise Rectified Linear Unit (ReLU)
type ReLU[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewReLU returns a new UnaryElementwise Rectified Linear Unit (ReLU) function.
func NewReLU[O mat.Tensor](x O) *ReLU[O] { _ = "STUB: not implemented"; return nil }

// Softsign is an operator to perform element-wise softsign.
type Softsign[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewSoftsign returns a new UnaryElementwise softsign function.
func NewSoftsign[O mat.Tensor](x O) *Softsign[O] { _ = "STUB: not implemented"; return nil }

// Cos is an operator to perform element-wise cos.
type Cos[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewCos returns a new UnaryElementwise cos function.
func NewCos[O mat.Tensor](x O) *Cos[O] { _ = "STUB: not implemented"; return nil }

// Sin is an operator to perform element-wise sin.
type Sin[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewSin returns a new UnaryElementwise sine function.
func NewSin[O mat.Tensor](x O) *Sin[O] { _ = "STUB: not implemented"; return nil }

// Neg is an operator to perform element-wise f(x) = -x
type Neg[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewNeg returns a new UnaryElementwise f(x) = -x function.
func NewNeg[O mat.Tensor](x O) *Neg[O] { _ = "STUB: not implemented"; return nil }

// Reciprocal is an operator to perform element-wise reciprocal.
type Reciprocal[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewReciprocal returns a new UnaryElementwise reciprocal function.
func NewReciprocal[O mat.Tensor](x O) *Reciprocal[O] { _ = "STUB: not implemented"; return nil }

// Abs is an operator to perform element-wise absolute value function.
type Abs[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewAbs returns a new UnaryElementwise absolute value function.
func NewAbs[O mat.Tensor](x O) *Abs[O] { _ = "STUB: not implemented"; return nil }

// Mish is an operator to perform element-wise mish.
type Mish[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewMish returns a new UnaryElementwise Mish function.
//
// Mish is a self-regularized non-monotonic activation function which can be
// mathematically defined as f(x) = x * tanh(softplus(x)).
//
// Reference: "Mish: A Self Regularized Non-Monotonic Neural Activation Function"
// by Diganta Misra, 2019 (https://arxiv.org/pdf/1908.08681.pdf)
func NewMish[O mat.Tensor](x O) *Mish[O] { _ = "STUB: not implemented"; return nil }

// GELU is an operator to perform element-wise GELU.
type GELU[O mat.Tensor] struct {
	*UnaryElementwise[O]
}

// NewGELU returns a new UnaryElementwise Gaussian Error Linear Unit (GELU) function.
func NewGELU[O mat.Tensor](x O) *GELU[O] { _ = "STUB: not implemented"; return nil }

// NewSiLU (Sigmoid Linear Unit) returns a new function of the form f(x) = x * sigmoid(x).
// The function in an alias of NewSwish.
func NewSiLU[O mat.Tensor](x O) *Swish[O] { _ = "STUB: not implemented"; return nil }

func absDeriv(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

// undefined

func tan(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func tanDeriv(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func tanh(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func tanhDeriv(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func hardSigmoid(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func hardSigmoidDeriv(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func hardTanh(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func hardTanhDeriv(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func relu(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func reluDeriv(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func softsign(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func softsignDeriv(i, j int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func celu(_, _ int, v float64, alpha ...float64) float64 { _ = "STUB: not implemented"; return 0 }

func celuDeriv(_, _ int, v float64, alpha ...float64) float64 { _ = "STUB: not implemented"; return 0 }

func elu(_, _ int, v float64, alpha ...float64) float64 { _ = "STUB: not implemented"; return 0 }

func eluDeriv(_, _ int, v float64, alpha ...float64) float64 { _ = "STUB: not implemented"; return 0 }

func leakyReLU(_, _ int, v float64, alpha ...float64) float64 { _ = "STUB: not implemented"; return 0 }

// slope * v

func leakyReLUDeriv(_, _ int, v float64, alpha ...float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

// slope

// alpha[0] is the alpha
// alpha[1] is the scale
func selu(_, _ int, v float64, alpha ...float64) float64 { _ = "STUB: not implemented"; return 0 }

// alpha[0] is the alpha
// alpha[1] is the scale
func seluDeriv(_, _ int, v float64, alpha ...float64) float64 { _ = "STUB: not implemented"; return 0 }

func softPlus(_, _ int, v float64, alpha ...float64) float64 { _ = "STUB: not implemented"; return 0 }

func softPlusDeriv(_, _ int, v float64, alpha ...float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func softShrink(_, _ int, v float64, alpha ...float64) float64 { _ = "STUB: not implemented"; return 0 }

func softShrinkDeriv(_, _ int, v float64, alpha ...float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func threshold(_, _ int, v float64, alpha ...float64) float64 { _ = "STUB: not implemented"; return 0 }

func thresholdDeriv(_, _ int, v float64, alpha ...float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func swishB(_, _ int, v float64, beta ...float64) float64 { _ = "STUB: not implemented"; return 0 }

func swishBDeriv(_, _ int, v float64, beta ...float64) float64 { _ = "STUB: not implemented"; return 0 }

func swishBBetaDeriv(v, beta float64) float64 { _ = "STUB: not implemented"; return 0 }

// Reference: "Mish: A Self Regularized Non-Monotonic Neural Activation Function" by Diganta Misra, 2019.
// (https://arxiv.org/pdf/1908.08681.pdf)
func mish(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func mishDeriv(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func gelu(_, _ int, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func geluDeriv(_, _ int, x float64) float64 { _ = "STUB: not implemented"; return 0 }
