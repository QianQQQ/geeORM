package main

import (
	"fmt"
	"reflect"
)

func boost(f interface{}, in []reflect.Value) {
	fmt.Println("haha")
	if reflect.TypeOf(f).Kind() != reflect.Func {
		return
	}
	reflect.ValueOf(f).Call(in)
}

func main() {
	a := func() {
		fmt.Println("a")
	}
	boost(a, nil)
}
