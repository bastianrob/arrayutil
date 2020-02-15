# Collection of Array utilities

## `Map` and `ParallelMap`

Just like javascript's map

* `Map` is done on sequentially
* `ParallelMap`
  * is done concurrently using go routine on each entry
  * will be faster on large amount of data with slow `map/transformation`

```go
func Map(arr interface{}, transform interface{}) (interface{}, error)
```

### Accepts

* `arr interface{}` have to be either array or slice of something
* `transform interface{}` have to be a `func` that accepts element type of `arr` and returns either a new type, or the same as element type of `arr`

### Returns

* `interface{}` is a slice of the return type of `transform`
* `error` if
  * `arr` is not array or slice
  * `transform` is nil or not a function

### Example

```go
source := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
addOne := func(num int) int {
    time.Sleep(50 * time.Millisecond) //simulate slow operation
    return num+1
}

start := time.Now()
resultMap, err := arrayutil.Map(source, addOne)
fmt.Println("SRC:", source)
fmt.Println("RES:", resultMap.([]int))
fmt.Println("ERR:", err)
fmt.Println("DUR:", float64(time.Since(start).Nanoseconds())/1000000., "ms")
//SRC: [1 2 3 4 5 6 7 8 9 10]
//RES: [2 3 4 5 6 7 8 9 10 11]
//ERR: <nil>
//DUR: 536.282464 ms

start = time.Now()
resultPMap, err := arrayutil.ParallelMap(source, addOne)
fmt.Println("SRC:", source)
fmt.Println("RES:", resultPMap.([]int))
fmt.Println("ERR:", err)
fmt.Println("DUR:", float64(time.Since(start).Nanoseconds())/1000000., "ms")
//SRC: [1 2 3 4 5 6 7 8 9 10]
//RES: [2 3 4 5 6 7 8 9 10 11]
//ERR: <nil>
//DUR: 54.597015 ms
```

---

## Reduce

Just like javascript's reduce

```go
func Reduce(arr interface{}, initialValue interface{}, transform interface{}) (interface{}, error)
```

### Accepts

* `arr interface{}` have to be either array or slice of something
* `initialValue interface{}` initial value of the reduce result
* `transform interface{}` reducer function
  * Must accepts 3 parameters
    * `accumulator` in which type must be the same of `initialValue` type
    * `entry` in which type must be the same of `arr` element type
    * `idx` integer index
  * And returns
    * `T` in which `T` is the type of `initialValue`

### Returns

* `interface{}` is the return type of `transform`
* `error` if
  * `arr` is not array or slice
  * `transform` is nil or not a function

### Example

```go
source := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
sum := func(accumulator, entry, _ int) int {
    return accumulator + entry
}

result, err := arrayutil.Reduce(source, 0, sum)
fmt.Println("SRC:", source)
fmt.Println("RES:", result.(int))
fmt.Println("ERR:", err)
//SRC: [1 2 3 4 5 6 7 8 9 10]
//RES: 55
//ERR: <nil>
```
