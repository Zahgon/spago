// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nn

import (
	"reflect"
	"sync"
)

// ParamsTraverser allows you to define a custom procedure to traverse the parameters of a model.
// If a model implements this procedure, it will take precedence over the regular parameters visit.
type ParamsTraverser interface {
	// TraverseParams visit each Param.
	TraverseParams(callback func(param *Param))
}

// paramsTraversal allows the traversal of Model parameters.
// The given paramsFunc is invoked for each parameter of the Model.
// If exploreSubModels is true, every nested Model and its parameters are
// also visited.
type paramsTraversal struct {
	paramsFunc       func(param *Param)
	modelsFunc       func(model Model)
	exploreSubModels bool
}

// walk iterates through all the parameters of m.
func (pt paramsTraversal) walk(m any) { _ = "STUB: not implemented"; return }

func (pt paramsTraversal) walkStructOrPtr(item any, name string) bool {
	_ = "STUB: not implemented"
	return false
}

// skip

func (pt paramsTraversal) walkSyncMap(i *sync.Map, name string) { _ = "STUB: not implemented"; return }

// skip map if the key is not a string or an int

// skip

func (pt paramsTraversal) walkSlice(v reflect.Value, name string) {
	_ = "STUB: not implemented"
	return
}

// skip

func (pt paramsTraversal) walkMap(v reflect.Value, name string) { _ = "STUB: not implemented"; return }

// skip map if the key is not a string or an int

// skip

// forEachField calls the paramsFunc for each field of the struct i.
func forEachField(i any, callback func(field any, name string)) { _ = "STUB: not implemented"; return }
