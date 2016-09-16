# List
[![Build Status](https://travis-ci.org/lucasvmiguel/list.svg?branch=master)](https://travis-ci.org/lucasvmiguel/list)
[![GoDoc](https://godoc.org/github.com/lucasvmiguel/list?status.svg)](https://godoc.org/github.com/lucasvmiguel/list)

## Description 

List is a library to handle with slice in golang easily, it's inspired by the Class List .Net.                            
Because sometimes you do not want to lose minutes to win nanoseconds ;)

https://msdn.microsoft.com/pt-br/library/6sh2ey19(v=vs.110).aspx

## Installation

```bash
go get github.com/lucasvmiguel/list
```

## How to use

``` go
import "github.com/lucasvmiguel/list"

func main() {
  l := list.Generic{}
  l.Add("1")
  l.Add("2")
  l.Add("3")

  l.Values() //[]interface{"1", "2", "3"}
  l.ToSliceString() //[]string{"1", "2", "3"}
}
```

## Benchmark

As you can see using the native code (benchmark with pure title) without library is faster than this library, but this library focus on usability.

```bash
BenchmarkAddPure-4   	       100	  15910984 ns/op
BenchmarkAdd-4       	         2	 667843794 ns/op
BenchmarkFilterPure-4	       200	  10268539 ns/op
BenchmarkFilter-4    	        30	  37735676 ns/op
BenchmarkFindPure-4  	    100000	     15686 ns/op
BenchmarkFind-4      	      2000	    525775 ns/op
BenchmarkGetPure-4   	2000000000	      0.77 ns/op
BenchmarkGet-4       	  30000000	      34.0 ns/op
```

## Methods

AddRange(vals []interface{})

Add(val interface{})

InsertAt(i int32, val interface{})

InsertRangeAt(i int32, vals []interface{})

InsertLast(val interface{})

GetAt(i int32) interface{}

GetRange(begin int32, end int32) []interface{}

GetFirst() interface{}

GetLast() interface{}

RemoveAt(i int32)

RemoveFirst()

RemoveLast()

Count() int

Clear()

Values() []interface{}

Filter(cond func(interface{}) bool) []interface{}

Exists(cond func(interface{}) bool) bool

Find(cond func(interface{}) bool) interface{}

FindAll(cond func(interface{}) bool) []interface{}

FindLast(cond func(interface{}) bool) interface{}

FindIndex(cond func(interface{}) bool) int

FindLastIndex(cond func(interface{}) bool) int

TrueForAll(cond func(interface{}) bool) bool

Reverse()

ToSliceString() []string

ToSliceInt() []int

ToSliceFloat64() []float64

## RoadMap

* Others type list(string, int...)
* Method DeleteRange

## License

  [MIT](LICENSE)
