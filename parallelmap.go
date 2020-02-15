package arrayutil

import (
	"reflect"
	"sync"
)

// ParallelMap an array of something into another thing using go routine
// Example:
//  Map([]int{1,2,3}, func(num int) int { return num+1 })
//	Results: []int{2,3,4}
func ParallelMap(arr interface{}, transform interface{}) (interface{}, error) {
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

	wg := &sync.WaitGroup{}
	wg.Add(arrV.Len())

	entryT := reflect.TypeOf(arr).Elem()
	result := reflect.MakeSlice(reflect.SliceOf(entryT), arrV.Len(), arrV.Cap())
	for i := 0; i < arrV.Len(); i++ {
		go func(idx int, entry reflect.Value) {
			//Call the transformation and store the result value
			tfResults := tv.Call([]reflect.Value{entry})

			//Store the transformation result into array of result
			resultEntry := result.Index(idx)
			if len(tfResults) > 0 {
				resultEntry.Set(tfResults[0])
			} else {
				resultEntry.Set(reflect.Zero(entryT))
			}

			//this go routine is done
			wg.Done()
		}(i, arrV.Index(i))
	}

	wg.Wait()
	return result.Interface(), nil
}
