package main

import (
	"fmt"
	"reflect"
)

func swap(a *int, b *int) {
	a, b = b, a
}

func main() {
	arr := []int{1, 3, 5, 6, 7}

	fmt.Println(arr)

	//Первый способ
	arr[0], arr[4] = arr[4], arr[0]

	fmt.Println(arr)
	//Второй способ
	reflect.Swapper(arr)(0, 4)

	fmt.Println(arr)
}
