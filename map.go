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
)

// Map an array of something into another thing
// Example:
//  Map([]int{1,2,3}, func(num int) int { return num+1 })
//	Results: []int{2,3,4}
func Map(arr interface{}, transform interface{}) (interface{}, error) {
	arrV := reflect.ValueOf(arr)
	kind := arrV.Kind()
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

	entryT := reflect.TypeOf(arr).Elem()
	result := reflect.MakeSlice(reflect.SliceOf(entryT), arrV.Len(), arrV.Cap())
	for i := 0; i < arrV.Len(); i++ {
		//Call the transformation and store the result value
		tfResults := tv.Call([]reflect.Value{arrV.Index(i)})

		resultEntry := result.Index(i)
		if len(tfResults) > 0 {
			resultEntry.Set(tfResults[0])
		} else {
			resultEntry.Set(reflect.Zero(entryT))
		}
	}

	return result.Interface(), nil
}
