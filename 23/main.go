package main

import (
	"fmt"
	"log"
)

func remove(slice []int, s int) ([]int, error) {
	if s < 0 || s >= len(slice) {
		return nil, fmt.Errorf("invalid index")
	}
	return append(slice[:s], slice[s+1:]...), nil
}

func main() {
	var slice = []int{5, 2, 3, 1, 6, 7}

	for _, n := range slice {
		fmt.Print(n, " ")
	}

	slice, err := remove(slice, 2)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println()
	for _, n := range slice {
		fmt.Print(n, " ")
	}
}
