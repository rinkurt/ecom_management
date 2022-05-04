package main

import (
	"fmt"
	"reflect"
)

type Foo[T any] struct {
	val T
}

func main() {
	a := Foo[int]{}
	b := Foo[string]{}
	
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

}
