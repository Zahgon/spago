// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ag

import (
	"github.com/nlpodyssey/spago/mat"
)

// Abs returns a new operator node as a result of the `Abs` function.
func Abs(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Add returns a new operator node as a result of the gradfn.Add function.
// As special case, the first node may be null.
// This help to keep the code as concise as possible e.g. during accumulation.
func Add(x1 mat.Tensor, x2 mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// return a copy of `x2` as is

// AddScalar returns a new operator node as a result of the gradfn.AddScalar function.
func AddScalar(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Affine returns a new operator node as a result of the gradfn.Affine function.
func Affine(b, w1, x1 mat.Tensor, wxPairs ...mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// AppendRows returns a new operator node as a result of the gradfn.AppendRows function.
func AppendRows(x mat.Tensor, vs ...mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// At returns a new operator node as a result of the gradfn.At function.
func At(x mat.Tensor, indices ...int) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// CELU returns a new operator node as a result of the gradfn.CELU function.
func CELU(x, alpha mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// ColView returns a new operator node as a result of the gradfn.ColView function.
func ColView(x mat.Tensor, column int) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// Concat returns a new operator node as a result of the gradfn.Concat function.
func Concat(xs ...mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Cos returns a new operator node as a result of the `Cos` function.
func Cos(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Div returns a new operator node as a result of the gradfn.Div function.
func Div(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// DivScalar returns a new operator node as a result of the gradfn.DivScalar function.
func DivScalar(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Dot returns a new operator node as a result of the gradfn.Dot function.
func Dot(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// DropoutFunc returns a function to create a Dropout operator working with the given dropout probability.
func DropoutFunc(p float64) func(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return nil }

// Dropout returns a new operator node as a result of the gradfn.Dropout function.
// If the dropout probability is zero, the operator will not be created,
// so the input itself is returned directly.
func Dropout(x mat.Tensor, p float64) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// ELU returns a new operator node as a result of the gradfn.ELU function.
func ELU(x, alpha mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Exp returns a new operator node as a result of the `Exp` function.
func Exp(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Flatten returns a new operator node as a result of the gradfn.Flatten function.
func Flatten(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// GELU returns a new operator node as a result of the gradfn.GELU function.
func GELU(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// HardSigmoid returns a new operator node as a result of the `HardSigmoid` function.
func HardSigmoid(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// HardTanh returns a new operator node as a result of the `HardTanh` function.
func HardTanh(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Copy returns a new operator node as a result of the gradfn.Copy function.
func Copy(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// LeakyReLU returns a new operator node as a result of the gradfn.LeakyReLU function.
func LeakyReLU(x, alpha mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Log returns a new operator node as a result of the `Log` function.
func Log(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Max returns a new operator node as a result of the gradfn.Max function.
func Max(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// MaxPooling returns a new operator node as a result of the gradfn.MaxPooling function.
func MaxPooling(x mat.Tensor, rows, columns int) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// Min returns a new operator node as a result of the gradfn.Min function.
func Min(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Mish returns a new operator node as a result of the `Mish` function.
func Mish(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Mul returns a new operator node as a result of the gradfn.Mul function.
func Mul(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

func MulT(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Neg returns a new operator node as a result of the `Neg` function.
func Neg(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Pow returns a new operator node as a result of the gradfn.Pow function.
func Pow(x mat.Tensor, power float64) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// Prod returns a new operator node as a result of the gradfn.Prod function.
func Prod(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// ProdScalar returns a new operator node as a result of the gradfn.ProdScalar function.
func ProdScalar(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Reciprocal returns a new operator node as a result of the `Reciprocal` function.
func Reciprocal(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// ReduceMax returns a new operator node as a result of the gradfn.ReduceMax function.
func ReduceMax(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// ReduceMean returns a new operator node as a result of the gradfn.ReduceMean function.
func ReduceMean(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// ReduceSum returns a new operator node as a result of the gradfn.ReduceSum function.
func ReduceSum(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// ReLU returns a new operator node as a result of the `ReLU` function.
func ReLU(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Reshape returns a new operator node as a result of the gradfn.Reshape function.
func Reshape(x mat.Tensor, rows, columns int) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// ReverseSub returns a new operator node as a result of the fn.ReverseSub function.
func ReverseSub(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// ReverseSubOne returns a new operator node as a result of applying reverse subtraction with 1.0 to the input using the fn.ReverseSub function.
func ReverseSubOne(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// RotateR performs the right circular shift.
// `i` is the number of places by which the elements are shifted.
func RotateR(x mat.Tensor, i int) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// RowView returns a new operator node as a result of the gradfn.RowView function.
func RowView(x mat.Tensor, row int) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// ScalarMax returns a new operator node as a result of the gradfn.ScalarMax function.
func ScalarMax(xs []mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// SELU returns a new operator node as a result of the gradfn.SELU function.
func SELU(x, alpha mat.Tensor, scale mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// Sigmoid returns a new operator node as a result of the `Sigmoid` function.
func Sigmoid(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// SiLU returns a new operator node as a result of the fn.SiLU function.
func SiLU(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Sin returns a new operator node as a result of the `Sin` function.
func Sin(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Slice returns a new operator node as a result of the gradfn.Slice function.
func Slice(x mat.Tensor, fromRow, fromCol, toRow, toCol int) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// Softmax returns a new operator node as a result of the gradfn.Softmax function.
func Softmax(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// SoftPlus returns a new operator node as a result of the gradfn.SoftPlus function.
func SoftPlus(x, beta, threshold mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// SoftShrink returns a new operator node as a result of the gradfn.SoftShrink function.
func SoftShrink(x, lambda mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// Softsign returns a new operator node as a result of the `SoftSign` function.
func Softsign(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// SparseMax returns a new operator node as a result of the gradfn.SparseMax function.
func SparseMax(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// SparseMaxLoss returns a new operator node as a result of the gradfn.SparseMaxLoss function.
func SparseMaxLoss(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Sqrt returns a new operator node as a result of the `Sqrt` function.
func Sqrt(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Square returns a new operator node as a result of the gradfn.Prod(x, x) function.
func Square(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Stack returns a new operator node as a result of the gradfn.Stack function.
func Stack(xs ...mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Sub returns a new operator node as a result of the gradfn.Sub function.
func Sub(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// SubScalar returns a new operator node as a result of the gradfn.SubScalar function.
func SubScalar(x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Swish returns a new operator node as a result of the gradfn.Swish function.
func Swish(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// SwishB returns a new operator node as a result of the gradfn.SwishB function.
func SwishB(x, beta mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// T returns a new operator node as a result of the fn.T function.
func T(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Tan returns a new operator node as a result of the `Tan` function.
func Tan(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Tanh returns a new operator node as a result of the `Tanh` function.
func Tanh(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Threshold returns a new operator node as a result of the gradfn.Threshold function.
func Threshold(x, threshold, k mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// Map returns a transformed version of xs with all its components modified according to the mapping function.
// It is useful for applying an operator to a sequence of nodes. Keep in mind that using this function has an overhead
// because of the callback, however insignificant compared to mathematical computations.
func Map(mapping func(mat.Tensor) mat.Tensor, xs []mat.Tensor) []mat.Tensor {
	_ = "STUB: not implemented"
	return nil
}

// Map2 takes two arguments and applies a mapping function (that must take two arguments) to the items from the two node-slices in parallel.
// It panics if one slice is shorter than the other.
func Map2(mapping func(a mat.Tensor, b mat.Tensor) mat.Tensor, xs1 []mat.Tensor, xs2 []mat.Tensor) []mat.Tensor {
	_ = "STUB: not implemented"
	return nil
}

// Pad down/up samples the input to the given size.
func Pad(xs []mat.Tensor, seqLen int, padding func(i int) mat.Tensor) []mat.Tensor {
	_ = "STUB: not implemented"
	return nil
}

// SeparateMatrix returns a matrix of Node(s) represented as a slice of slice containing the elements extracted from the input.
// The dimensions of the resulting matrix are the same of the input.
func SeparateMatrix(x mat.Tensor) [][]mat.Tensor { _ = "STUB: not implemented"; return nil }

// SeparateVec returns a slice of Node(s) containing the elements extracted from the input.
// The size of the vector equals the number of input elements.
// You can think of this method as the inverse of the ag.Concat operator.
func SeparateVec(x mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

// SplitVec splits the x Node into multiple chunks.
func SplitVec(x mat.Tensor, chunks int) []mat.Tensor { _ = "STUB: not implemented"; return nil }

// Sum returns the value that describes the sum of the sample.
// It panics if the input is empty.
func Sum(xs ...mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Mean returns the value that describes the average of the sample.
func Mean(xs []mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Maximum returns the value that describes the maximum of the sample.
func Maximum(xs []mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// Minimum returns the value that describes the minimum of the sample.
func Minimum(xs []mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// BiLinear performs a bilinear transformation of the type (x_1 W x_2)
func BiLinear(w, x1, x2 mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// BiAffine performs a biaffine transformation.
func BiAffine(w, u, v, b, x1, x2 mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// PositiveELU returns a new operator node as a result of ELU(x) + 1.
func PositiveELU(x mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// LogSoftmax returns a new operator node as a result of Log(Softmax(x)).
func LogSoftmax(x mat.Tensor) mat.Tensor {
	_ = "STUB: not implemented"
	return *

	// LogSumExp "trick" computes the log of the sum of exponentials of input elements.
	// When the input is one, this must be a vector. Alternatively, the calculation
	// is conducted on a list of scalars.
	new(mat.Tensor)
}

func LogSumExp(xs ...mat.Tensor) mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// RowViews calls RowView for each row of x, returning a new slice
// of row-view Nodes.
func RowViews(x mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }

// ColViews calls ColView for each column of x, returning a new slice
// of column-view Nodes.
func ColViews(x mat.Tensor) []mat.Tensor { _ = "STUB: not implemented"; return nil }
