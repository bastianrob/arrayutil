package arrayutil

import (
	"reflect"
	"sync"
)

// ParallelFilter an array using go routine
// This function will not guarantee order of results
func ParallelFilter(arr interface{}, filterf FilterFunc) (interface{}, error) {
	arrV := reflect.ValueOf(arr)
	kind := arrV.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, ErrMapSourceNotArray
	}

	wg := &sync.WaitGroup{}
	wg.Add(arrV.Len())
	queue := make(chan interface{}, 3)

	for i := 0; i < arrV.Len(); i++ {
		go func(idx int, entry interface{}) {
			exists := filterf(entry)
			if exists {
				queue <- entry
			} else {
				queue <- nil
			}
		}(i, arrV.Index(i).Interface())
	}

	//This whole debacle is just to convert []things into *[]things
	entryT := reflect.TypeOf(arr).Elem()
	sliceOfT := reflect.MakeSlice(reflect.SliceOf(entryT), 0, 0)
	ptrToSliceOfT := reflect.New(sliceOfT.Type())
	ptrToSliceOfT.Elem().Set(sliceOfT)

	slicePtr := reflect.ValueOf(ptrToSliceOfT.Interface())
	sliceValuePtr := slicePtr.Elem()

	go func() {
		for entry := range queue {
			if entry != nil {
				appendResult := reflect.Append(sliceValuePtr, reflect.ValueOf(entry))
				sliceValuePtr.Set(appendResult)
			}
			wg.Done()
		}
	}()

	wg.Wait()    //wait for all filter to be done, and results appended to sliceValuePtr
	close(queue) //close the queue channel so the goroutine can exit
	return sliceValuePtr.Interface(), nil
}
