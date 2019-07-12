package arrayutil

import "reflect"

//Intersect Check whether an array intersect with another array
func Intersect(arr1 interface{}, arr2 interface{}) bool {
	arrV := reflect.ValueOf(arr2)
	if arrV.Kind() != reflect.Slice {
		return false
	}

	//loop through arr2 and check whether arr1 contains entry of arr2
	for i := 0; i < arrV.Len(); i++ {
		entry := arrV.Index(i).Interface()
		if Contains(arr1, entry) {
			return true
		}
	}

	return false
}
