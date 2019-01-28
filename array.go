package arrayutil

import (
	"fmt"
	"reflect"
)

//Contains find whether array contains the search clause or not
func Contains(arr interface{}, clause interface{}) bool {
	arrV := reflect.ValueOf(arr)
	if arrV.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < arrV.Len(); i++ {
		entry := arrV.Index(i).Interface()
		equal := reflect.DeepEqual(entry, clause)
		if equal {
			return true
		}
	}
	return false
}

//Reduce an array of something into another thing
func Reduce(arr interface{}, initialValue interface{}, transform interface{}) (interface{}, error) {
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

	accV := reflect.ValueOf(initialValue)
	for i := 0; i < arrV.Len(); i++ {
		entry := arrV.Index(i)
		tfResults := tv.Call([]reflect.Value{accV, entry, reflect.ValueOf(i)})
		accV = tfResults[0]
	}

	return accV.Interface(), nil
}
