package main

import "fmt"

func quicksort(arr []int, low int, high int) {
	if low >= high || low < 0 {
		return
	}

	p := partion(arr, low, high)

	quicksort(arr, low, p-1)
	quicksort(arr, p+1, high)
}

func partion(arr []int, low int, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	i += 1
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func main() {
	arr := []int{8, 3, 5, 1}

	fmt.Println(arr)

	quicksort(arr, 0, 3)

	fmt.Println(arr)

}
