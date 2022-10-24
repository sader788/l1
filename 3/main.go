package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// решение с использованием атомарных примитивов
func sol1() {
	var nums = []int32{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	var sum int32 = 0

	for _, num := range nums {
		wg.Add(1)
		go func(delta int32) {
			defer wg.Done()
			atomic.AddInt32(&sum, delta*delta)
		}(num)
	}

	wg.Wait()
	fmt.Println("sol 1: ", sum)
}

// Решение с Mutex
func sol2() {
	var nums = []int32{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	var sum int32 = 0

	for _, num := range nums {
		wg.Add(1)
		go func(delta int32) {
			defer wg.Done()
			//критическая секция
			mtx.Lock()
			sum += delta * delta
			mtx.Unlock()
		}(num)
	}

	wg.Wait()
	fmt.Println("sol 2: ", sum)
}

func main() {
	sol1()
	sol2()
}
