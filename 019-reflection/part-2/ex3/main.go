package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}
	x = &struct{ name string }{}
	t0 := reflect.TypeOf(x)
	v0 := reflect.ValueOf(x)
	fmt.Printf("x: type = %v, value = %v\n", t0, v0)

	x = new(string)

	t1 := reflect.TypeOf(x)
	v1 := reflect.ValueOf(x)
	fmt.Printf("x: type = %v, value = %v\n", t1, v1)

	// what kind or category is the type a member of?
	// Think of Kind as super-type
	fmt.Printf("t0: type = %v, kind = %v\n", t0, t0.Kind())
	fmt.Printf("t1: type = %v, kind = %v\n", t1, t1.Kind())

	// a0 := [5]int{}
	// a1 := [10]float{}
	// a0 and a1 are different types, but both are kind: array

}
