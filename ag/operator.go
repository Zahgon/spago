// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ag

import (
	"runtime"
	"sync"

	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
)

var (
	// forceSyncExecution, when set to true, forces operators to run synchronously, overriding any "async" flag in the Run() function.
	// This can be particularly useful for debugging.
	forceSyncExecution = false
)

// SetForceSyncExecution enables or disables the forcing of synchronous execution for all operators.
// When enabled, the operators will run synchronously, regardless of the "async" flag in the Run() function.
// This setting can be particularly useful for debugging.
func SetForceSyncExecution(enable bool) { _ = "STUB: not implemented"; return }

// backwardState is an enumeration type associated to an Operator, to keep
// track of its visited status among different backpropagation phases.
type backwardState = uint32

const (
	// idle reports that gradient propagation is not pending for an
	// operator node.
	//
	// It's the default zero-value state of an operator, and it's also the
	// final value set from the backward step once gradients have been
	// propagated.
	//
	// As soon as a backward operation is performed, the status will change to
	// pending.
	idle backwardState = iota
	// pending is set on an operator node from the preparatory phase
	// of the backward step.
	// It reports that the node has been marked as a candidate for gradients
	// propagation and the number of pendingGrads has been computed.
	//
	// The next logical state is ongoing.
	pending
	// ongoing is set on an operator node from the core phase of the
	// backward step. It reports that the node has been visited once for
	// performing its Operator.backward method.
	//
	// This status remains set until the gradients of all dependents have been
	// resolved, and the node's own gradients have been propagated too.
	// After that, the status is set back to idle.
	ongoing
)

// AutoGradFunction represents a function with automatic differentiation features.
// It's used to define a new operator.
type AutoGradFunction interface {
	// Forward computes the output of the function.
	Forward() (mat.Tensor, error)
	// Backward computes the backward pass given the gradient of the output.
	Backward(gy mat.Tensor) error
	// Operands returns the list of operands.
	Operands() []mat.Tensor
}

// forwardGuard is a buffered channel that acts as a semaphore to limit the concurrency
// of async forward operations in the Run function. Its buffer size determines the maximum
// number of forward operations that can run concurrently. Acquiring and releasing slots
// in the semaphore ensures that the concurrency level stays within the desired limit.
var forwardGuard chan struct{}

// Using runtime.NumCPU() * 2 is a common heuristic for setting the number of concurrent goroutines or the concurrency level in a Go program.
var concurrencyLimit = runtime.NumCPU() * 2

func init() {
	forwardGuard = make(chan struct{}, concurrencyLimit)
}

// Operator is a type of node.
// It's used to represent a function with automatic differentiation features.
type Operator struct {
	// value stores the results of a forward evaluation, as mat.Matrix.
	// It's set by executeForward() goroutine.
	// Use the Value() method to get the actual value.
	// It also contains the accumulated gradients. Use the Grad() method to get them.
	value mat.Tensor
	// onceOperands is used to initialize the operands only once.
	onceOperands sync.Once
	// AutoGradFunction's operands are memoized here after the first request.
	operands []mat.Tensor
	// backwardPass is the backward function to be executed.
	fn AutoGradFunction
	// broadcast is the channel used to broadcast the result of the forward pass.
	broadcast chan struct{}
	// broadcastGrad is the channel used to broadcast the result of the backward pass.
	// It is initialized only when the backward pass is performed.
	broadcastGrad chan struct{}
	// pendingGrads is the number of pending gradients to be accumulated. (default: 0)
	pendingGrads int64
	// onceRequiresGrad is used to initialize the requiresGrad only once.
	onceRequiresGrad sync.Once
	// requiresGrad is a flag that indicates whether the operator requires gradients.
	// Use the RequiresGrad() method to get the actual value.
	requiresGrad bool
	// backwardState is the state of the backward pass.
	backwardState backwardState
}

// NewOperator creates a new operator with the given AutoGradFunction.
// Note that the operator's Value() can only be accessed after calling the Run() function.
func NewOperator(f AutoGradFunction) *Operator { _ = "STUB: not implemented"; return nil }

// SetAt sets the value at the given indices.
// It panics if the given indices are out of range.
func (o *Operator) SetAt(m mat.Tensor, indices ...int) { _ = "STUB: not implemented"; return }

// At returns the value at the given indices.
// It panics if the given indices are out of range.
func (o *Operator) At(indices ...int) mat.Tensor {
	_ = "STUB: not implemented"
	return *new(mat.Tensor)
}

// Run starts the execution of the operator, performing the forward pass.
// If the optional async argument is set to true, the forward pass will be executed in a separate goroutine.
// The function returns a pointer to the Operator, allowing for method chaining.
func (o *Operator) Run(async ...bool) *Operator { _ = "STUB: not implemented"; return nil }

//lint:ignore S1019 explicitly set the buffer size to 0 as the channel is used as a signal

// forward executes the forward function and inform all goroutines that have been waiting for the result.
func (o *Operator) executeForward() { _ = "STUB: not implemented"; return }

// TODO: handle error

// if nil, it means that the operator is not async
// inform all goroutines that have been waiting for the result

// Value returns the result of the function.
func (o *Operator) Value() mat.Tensor {
	_ = "STUB: not implemented"
	return *
	// if nil, it means that the operator is not async
	new(mat.Tensor)
}

// wait for the forward goroutine to finish

func (o *Operator) Item() float.Float {
	_ = "STUB: not implemented"
	return *

	// Grad returns the gradients accumulated during the backward pass.
	new(float.Float)
}

func (o *Operator) Grad() mat.Tensor { _ = "STUB: not implemented"; return *new(mat.Tensor) }

// wait for the backward goroutine to finish

// HasGrad returns true if there are accumulated gradients.
func (o *Operator) HasGrad() bool { _ = "STUB: not implemented"; return false }

// safety wait for the backward goroutine to finish

// RequiresGrad returns true if the node requires gradients.
func (o *Operator) RequiresGrad() bool { _ = "STUB: not implemented"; return false }

// memoize the result

// Operands returns the operands of the operator.
func (o *Operator) Operands() []mat.Tensor { _ = "STUB: not implemented"; return nil }

// memoize the result

// ZeroGrad clears the gradients.
func (o *Operator) ZeroGrad() { _ = "STUB: not implemented"; return }

// AccGrad accumulates the gradients to the node itself.
func (o *Operator) AccGrad(grad mat.Tensor) { _ = "STUB: not implemented"; return }

// Don't decrement the counter if the backward pass is not running.

// notify all goroutines that have been waiting for the gradients

func (o *Operator) assignOutputGradient() error { _ = "STUB: not implemented"; return nil }

func (o *Operator) prepareBackwardPass() { _ = "STUB: not implemented"; return }

//lint:ignore S1019 explicitly set the buffer size to 0 as the channel is used as a signal

func (o *Operator) processBackwardPass(wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }

// decrement when the backward pass is done

func (o *Operator) executeBackward(wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }

// wait until the accumulated gradients are ready

// no gradients to propagate

// TODO: handle error

func (o *Operator) isBackwardIdle() bool { _ = "STUB: not implemented"; return false }

func (o *Operator) setBackwardIdle() { _ = "STUB: not implemented"; return }

func (o *Operator) trySetBackwardPending() bool { _ = "STUB: not implemented"; return false }

func (o *Operator) trySetBackwardOngoing() bool { _ = "STUB: not implemented"; return false }

// isNil returns true if the gradients are nil.
func isNil(grad any) bool { _ = "STUB: not implemented"; return false }
