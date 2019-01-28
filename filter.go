package arrayutil

import (
	"fmt"
	"reflect"
)

//FilterFunc representation
type FilterFunc func(interface{}) bool

//Filter an array
func Filter(arr interface{}, filterf FilterFunc) ([]interface{}, error) {
	arrV := reflect.ValueOf(arr)
	if arrV.Kind() != reflect.Slice {
		return nil, fmt.Errorf("Input value is not an array")
	}

	result := make([]interface{}, 0, 0)
	for i := 0; i < arrV.Len(); i++ {
		entry := arrV.Index(i).Interface()
		exists := filterf(entry)
		if exists {
			result = append(result, entry)
		}
	}
	return result, nil
}
