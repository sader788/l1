package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}

	i = "dsadsa"

	t := reflect.TypeOf(i)

	fmt.Println(t.Name())
}
