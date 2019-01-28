package arrayutil

import (
	"fmt"
	"reflect"
)

//Map an array of something into another thing
//Example:
//	Map([]int{1,2,3}, func(num int) int { return num+1 })
//	Results: []int{2,3,4}
func Map(arr interface{}, transform interface{}) ([]interface{}, error) {
	arrV := reflect.ValueOf(arr)
	if arrV.Kind() != reflect.Slice {
		return nil, fmt.Errorf("Input value is not an array")
	}

	if transform == nil {
		return nil, fmt.Errorf("Transform function cannot be nil")
	}

	tv := reflect.ValueOf(transform)
	if tv.Kind() != reflect.Func {
		return nil, fmt.Errorf("Transform argument must be a function")
	}

	result := make([]interface{}, arrV.Len())
	for i := 0; i < arrV.Len(); i++ {
		entry := arrV.Index(i)
		tfResults := tv.Call([]reflect.Value{entry})
		if len(tfResults) <= 0 {
			result[i] = nil
		} else {
			result[i] = tfResults[0].Interface()
		}
	}
	return result, nil
}
