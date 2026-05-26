// Copyright 2022 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"sync"

	"github.com/nlpodyssey/spago/mat/float"
)

// A Dense matrix implementation.
type Dense[T float.DType] struct {
	gradMu       sync.RWMutex
	data         []T
	grad         *Dense[T]
	shape        []int
	requiresGrad bool // default: false
}

// makeDense returns a Dense matrix.
func makeDense[T float.DType](array []T, shape ...int) *Dense[T] {
	_ = "STUB: not implemented"
	return nil
}

func malloc[T float.DType](size int) []T { _ = "STUB: not implemented"; return nil }

// Shape returns the size in each dimension.
func (d *Dense[_]) Shape() []int {
	_ = "STUB: not implemented"

	// Dims returns the number of dimensions.
	return nil
}

func (d *Dense[_]) Dims() int {
	_ = "STUB: not implemented"
	// rows and columns
	return 0
}

// The Size of the matrix (rows*columns).
func (d *Dense[_]) Size() int {
	_ = "STUB: not implemented"

	// Data returns the underlying data of the matrix, as a raw one-dimensional
	// slice of values in row-major order.
	return 0
}

func (d *Dense[T]) Data() float.Slice { _ = "STUB: not implemented"; return *new(float.Slice) }

// SetData sets the content of the matrix, copying the given raw
// data representation as one-dimensional slice.
func (d *Dense[T]) SetData(data float.Slice) { _ = "STUB: not implemented"; return }

// ZerosLike returns a new matrix with the same dimensions of the
// receiver, initialized with zeroes.
func (d *Dense[T]) ZerosLike() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// OnesLike returns a new matrix with the same dimensions of the
// receiver, initialized with ones.
func (d *Dense[T]) OnesLike() Matrix {
	_ = "STUB: not implemented"
	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// avoid bounds check in loop

// Scalar returns the scalar value.
// It panics if the matrix does not contain exactly one element.
func (d *Dense[T]) Item() float.Float { _ = "STUB: not implemented"; return *new(float.Float) }

// Zeros sets all the values of the matrix to zero.
func (d *Dense[T]) Zeros() {
	_ = "STUB: not implemented"
	// avoid bounds check in loop
	return
}

// SetAt sets the value m at the given indices.
// It panics if the given indices are out of range.
func (d *Dense[T]) SetAt(m Tensor, indices ...int) { _ = "STUB: not implemented"; return }

// At returns the value at the given indices.
// It panics if the given indices are out of range.
func (d *Dense[T]) At(i ...int) Tensor { _ = "STUB: not implemented"; return *new(Tensor) }

// SetScalar sets the value v at the given indices.
// It panics if the given indices are out of range.
func (d *Dense[T]) SetScalar(v float.Float, indices ...int) { _ = "STUB: not implemented"; return }

// ScalarAt returns the value at the given indices.
// It panics if the given indices are out of range.
func (d *Dense[T]) ScalarAt(indices ...int) float.Float {
	_ = "STUB: not implemented"
	return *new(float.Float)
}

func (d *Dense[T]) set(v T, i ...int) { _ = "STUB: not implemented"; return }

func (d *Dense[T]) at(i ...int) T { _ = "STUB: not implemented"; return *new(T) }

// ExtractRow returns a copy of the i-th row of the matrix,
// as a row vector (1×cols).
func (d *Dense[T]) ExtractRow(i int) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.

// ExtractColumn returns a copy of the i-th column of the matrix,
// as a column vector (rows×1).
func (d *Dense[T]) ExtractColumn(i int) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.

// Slice returns a new matrix obtained by slicing the receiver across the
// given positions. The parameters "fromRow" and "fromCol" are inclusive,
// while "toRow" and "toCol" are exclusive.
func (d *Dense[T]) Slice(fromRow, fromCol, toRow, toCol int) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.

// exploiting append in loop

// Reshape returns a copy of the matrix.
// It panics if the dimensions are incompatible.
func (d *Dense[T]) Reshape(shape ...int) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func copySlice[T float.DType](src []T) []T { _ = "STUB: not implemented"; return nil }

// ReshapeInPlace changes the dimensions of the matrix in place and returns the
// matrix itself.
// It panics if the dimensions are incompatible.
func (d *Dense[T]) ReshapeInPlace(shape ...int) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

// Flatten creates a new row vector (1×size) corresponding to the
// "flattened" row-major ordered representation of the initial matrix.
func (d *Dense[T]) Flatten() Matrix {
	_ = "STUB: not implemented"
	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// FlattenInPlace transforms the matrix in place, changing its dimensions,
// obtaining a row vector (1×size) containing the "flattened" row-major
// ordered representation of the initial value.
// It returns the matrix itself.
func (d *Dense[T]) FlattenInPlace() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// ResizeVector returns a resized copy of the vector.
//
// If the new size is smaller than the input vector, the remaining tail
// elements are removed. If it's bigger, the additional tail elements
// are set to zero.
func (d *Dense[T]) ResizeVector(newSize int) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// T returns the transpose of the matrix.
func (d *Dense[T]) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.

// TransposeInPlace transposes the matrix in place, and returns the
// matrix itself.
func (d *Dense[T]) TransposeInPlace() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// Vector, scalar, or empty data

// Square matrix

// Rectangular matrix

// Add returns the addition between the receiver and another matrix.
func (d *Dense[T]) Add(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// AddInPlace performs the in-place addition with the other matrix.
func (d *Dense[T]) AddInPlace(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// AddScalar performs the addition between the matrix and the given value.
func (d *Dense[T]) AddScalar(n float64) Matrix {
	_ = "STUB: not implemented"
	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// AddScalarInPlace adds the scalar to all values of the matrix.
func (d *Dense[T]) AddScalarInPlace(n float64) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

// Sub returns the subtraction of the other matrix from the receiver.
func (d *Dense[T]) Sub(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// SubInPlace performs the in-place subtraction with the other matrix.
func (d *Dense[T]) SubInPlace(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// SubScalar performs a subtraction between the matrix and the given value.
func (d *Dense[T]) SubScalar(n float64) Matrix {
	_ = "STUB: not implemented"
	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// SubScalarInPlace subtracts the scalar from the receiver's values.
func (d *Dense[T]) SubScalarInPlace(n float64) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

// Prod performs the element-wise product between the receiver and the other matrix.
func (d *Dense[T]) Prod(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.

// Avoid bounds checks in loop

// ProdInPlace performs the in-place element-wise product with the other matrix.
func (d *Dense[T]) ProdInPlace(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// ProdScalar returns the multiplication between the matrix and the given value.
func (d *Dense[T]) ProdScalar(n float64) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// ProdScalarInPlace performs the in-place multiplication between the
// matrix and the given value.
func (d *Dense[T]) ProdScalarInPlace(n float64) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

// ProdMatrixScalarInPlace multiplies the given matrix with the value,
// storing the result in the receiver.
func (d *Dense[T]) ProdMatrixScalarInPlace(m Matrix, n float64) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

// Div returns the result of the element-wise division of the receiver by the other matrix.
func (d *Dense[T]) Div(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// DivInPlace performs the in-place element-wise division of the receiver by the other matrix.
func (d *Dense[T]) DivInPlace(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// Mul performs the multiplication row by column.
// If A is an i×j Matrix, and B is j×k, then the resulting Matrix
// C = AB will be i×k.
func (d *Dense[T]) Mul(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// aRows
// aCols
// bCols
// a
// b
// c

// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.

// aRows
// aCols
// bCols
// a
// b
// c

// m
// n
// alpha
// a
// lda
// x
// incX
// beta
// y
// incY

// MulT performs the matrix multiplication row by column.
// ATB = C, where AT is the transpose of A
// if A is an r x c Matrix, and B is j x k, r = j the resulting
// Matrix C will be c x k.
func (d *Dense[T]) MulT(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// DotUnitary returns the dot product of two vectors as a scalar Matrix.
func (d *Dense[T]) DotUnitary(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// ClipInPlace clips in place each value of the matrix.
func (d *Dense[T]) ClipInPlace(min, max float64) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

// Maximum returns a new matrix containing the element-wise maxima.
func (d *Dense[T]) Maximum(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.

// Minimum returns a new matrix containing the element-wise minima.
func (d *Dense[T]) Minimum(other Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.

// Abs returns a new matrix applying the absolute value function to all elements.
func (d *Dense[T]) Abs() Matrix {
	_ = "STUB: not implemented"
	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// Pow returns a new matrix, applying the power function with given exponent
// to all elements of the matrix.
func (d *Dense[T]) Pow(power float64) Matrix {
	_ = "STUB: not implemented"
	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// Sqrt returns a new matrix applying the square root function to all elements.
func (d *Dense[T]) Sqrt() Matrix {
	_ = "STUB: not implemented"
	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// Log returns a new matrix applying the natural logarithm function to each element.
func (d *Dense[T]) Log() Matrix {
	_ = "STUB: not implemented"
	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// Exp returns a new matrix applying the base-e exponential function to each element.
func (d *Dense[T]) Exp() Matrix {
	_ = "STUB: not implemented"
	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// Sigmoid returns a new matrix applying the sigmoid function to each element.
func (d *Dense[T]) Sigmoid() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// Sum returns the sum of all values of the matrix as a scalar Matrix.
func (d *Dense[T]) Sum() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (d *Dense[T]) sum() T { _ = "STUB: not implemented"; return *new(T) }

// Max returns the maximum value of the matrix as a scalar Matrix.
func (d *Dense[T]) Max() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (d *Dense[T]) max() T { _ = "STUB: not implemented"; return *new(T) }

// Min returns the minimum value of the matrix as a scalar Matrix.
func (d *Dense[T]) Min() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// ArgMax returns the index of the vector's element with the maximum value.
func (d *Dense[T]) ArgMax() int { _ = "STUB: not implemented"; return 0 }

// Softmax applies the softmax function to the vector, returning the
// result as a new column vector.
func (d *Dense[T]) Softmax() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// CumSum computes the cumulative sum of the vector's elements, returning
// the result as a new column vector.
func (d *Dense[T]) CumSum() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.

// Range creates a new vector initialized with data extracted from the
// matrix raw data, from start (inclusive) to end (exclusive).
func (d *Dense[T]) Range(start, end int) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// SplitV splits the vector in N chunks of given sizes,
// so that N[i] has size sizes[i].
func (d *Dense[T]) SplitV(sizes ...int) []Matrix { _ = "STUB: not implemented"; return nil }

// Augment places the identity matrix at the end of the original matrix.
func (d *Dense[T]) Augment() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// TODO: rewrite for better performance

// SwapInPlace swaps two rows of the matrix in place.
func (d *Dense[T]) SwapInPlace(r1, r2 int) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// TODO: rewrite for better performance

// PadRows returns a copy of the matrix with n additional tail rows.
// The additional elements are set to zero.
func (d *Dense[T]) PadRows(n int) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// PadColumns returns a copy of the matrix with n additional tail columns.
// The additional elements are set to zero.
func (d *Dense[T]) PadColumns(n int) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// AppendRows returns a copy of the matrix with len(vs) additional tail rows,
// being each new row filled with the values of each given vector.
//
// It accepts row or column vectors indifferently, virtually treating all of
// them as row vectors.
func (d *Dense[T]) AppendRows(vs ...Matrix) Matrix {
	_ = "STUB: not implemented"

	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// Norm returns the vector's norm. Use pow = 2.0 to compute the Euclidean norm.
// The result is a scalar Matrix.
func (d *Dense[T]) Norm(pow float64) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (d *Dense[T]) norm(pow float64) float64 { _ = "STUB: not implemented"; return 0 }

// Normalize2 normalizes an array with the Euclidean norm.
func (d *Dense[T]) Normalize2() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

// Apply creates a new matrix executing the unary function fn.
func (d *Dense[T]) Apply(fn func(r, c int, v float64) float64) Matrix {
	_ = "STUB: not implemented"
	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// ApplyInPlace executes the unary function fn over the matrix `a`,
// and stores the result in the receiver, returning the receiver itself.
func (d *Dense[T]) ApplyInPlace(fn func(r, c int, v float64) float64, a Matrix) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

// ApplyWithAlpha creates a new matrix executing the unary function fn,
// taking additional alpha.
func (d *Dense[T]) ApplyWithAlpha(fn func(r, c int, v float64, alpha ...float64) float64, alpha ...float64) Matrix {
	_ = "STUB: not implemented"
	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// ApplyWithAlphaInPlace executes the unary function fn over the matrix `a`,
// taking additional parameters alpha, and stores the result in the
// receiver, returning the receiver itself.
func (d *Dense[T]) ApplyWithAlphaInPlace(fn func(r, c int, v float64, alpha ...float64) float64, a Matrix, alpha ...float64) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

// TODO: rewrite for better performance

// DoNonZero calls a function for each non-zero element of the matrix.
// The parameters of the function are the element's indices and value.
func (d *Dense[T]) DoNonZero(fn func(r, c int, v float64)) { _ = "STUB: not implemented"; return }

// DoVecNonZero calls a function for each non-zero element of the vector.
// The parameters of the function are the element's index and value.
func (d *Dense[T]) DoVecNonZero(fn func(i int, v float64)) { _ = "STUB: not implemented"; return }

// Clone returns a new matrix, copying all its values from the receiver.
func (d *Dense[T]) Clone() Matrix {
	_ = "STUB: not implemented"
	// Note: Consider that for performance optimization, it's not necessary to initialize the underlying slice to zero.
	return *new(Matrix)
}

// Copy copies the data from the other matrix to the receiver.
// It panics if the matrices have different dimensions.
func (d *Dense[T]) Copy(other Matrix) { _ = "STUB: not implemented"; return }

// String returns a string representation of the matrix.
func (d *Dense[T]) String() string { _ = "STUB: not implemented"; return "" }

// NewMatrix creates a new matrix, of the same type of the receiver, of
// size rows×cols, initialized with a copy of raw data.
//
// Rows and columns MUST not be negative, and the length of data MUST be
// equal to rows*cols, otherwise the method panics.
func (d *Dense[T]) NewMatrix(opts ...OptionsFunc) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

func (d *Dense[T]) NewScalar(v float64, opts ...OptionsFunc) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

// NewConcatV creates a new column vector, of the same type of the receiver,
// concatenating two or more vectors "vertically"
// It accepts row or column vectors indifferently, virtually
// treating all of them as column vectors.
func (d *Dense[T]) NewConcatV(vs ...Matrix) Matrix {
	_ = "STUB: not implemented"
	return *

	// NewStack creates a new matrix, of the same type of the receiver, stacking
	// two or more vectors of the same size on top of each other; the result is
	// a new matrix where each row contains the data of each input vector.
	// It accepts row or column vectors indifferently, virtually treating all of
	// them as row vectors.
	new(Matrix)
}

func (d *Dense[T]) NewStack(vs ...Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }
