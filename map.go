package arrayutil

import (
	"errors"
	"reflect"
)

//MapError
var (
	ErrMapSourceNotArray   = errors.New("Input value is not an array")
	ErrMapTransformNil     = errors.New("Transform function cannot be nil")
	ErrMapTransformNotFunc = errors.New("Transform argument must be a function")
	ErrMapResultTypeNil    = errors.New("Map result type cannot be nil")
)

// Map an array of something into another thing
// Example:
//  Map([]int{1,2,3}, func(num int) int { return num+1 })
//	Results: []int{2,3,4}
func Map(source interface{}, transform interface{}, T reflect.Type) (interface{}, error) {
	srcV := reflect.ValueOf(source)
	kind := srcV.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, ErrMapSourceNotArray
	}

	if transform == nil {
		return nil, ErrMapTransformNil
	}

	tv := reflect.ValueOf(transform)
	if tv.Kind() != reflect.Func {
		return nil, ErrMapTransformNotFunc
	}

	if T == nil {
		return nil, ErrMapResultTypeNil
	}

	// kinda equivalent to = make([]T, srcv.Len())
	result := reflect.MakeSlice(reflect.SliceOf(T), srcV.Len(), srcV.Cap())
	for i := 0; i < srcV.Len(); i++ {
		//Call the transformation and store the result value
		tfResults := tv.Call([]reflect.Value{srcV.Index(i)})

		resultEntry := result.Index(i)
		if len(tfResults) > 0 {
			resultEntry.Set(tfResults[0])
		} else {
			resultEntry.Set(reflect.Zero(T))
		}
	}

	return result.Interface(), nil
}
