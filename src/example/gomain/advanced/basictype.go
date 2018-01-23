package main

import (
	"fmt"
	"reflect"
)

func main() {
	var b byte = 'D'
	fmt.Printf("output: %v\n", reflect.TypeOf(b))
	fmt.Printf("output: %v\n", reflect.TypeOf(b).Kind())
}
