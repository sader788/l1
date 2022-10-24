package main

import (
	"fmt"
	"math"
)

func binarySearch(arr []int, n int, T int) int {
	L := 0
	R := n - 1

	for L <= R {
		m := int(math.Floor(float64((R + L)) / 2.0))
		if arr[m] < T {
			L = m + 1
		} else if arr[m] > T {
			R = m - 1
		} else {
			return m
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 6, 8, 12, 50, 100, 200, 205, 333, 544}

	fmt.Println(binarySearch(arr, 12, 3))

}
