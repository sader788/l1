package main

import (
	"fmt"
	"sync"
)

var nums = []int{2, 4, 6, 8, 10}

func calcSquare(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		fmt.Print(num*num, " ")
	}
}

// С использованием канала
func sol1(gorutineNum int) {
	in := make(chan int, 2)
	wg := &sync.WaitGroup{}

	//запускаем gorutineNum гоурутин
	for i := 0; i < gorutineNum; i++ {
		wg.Add(1)
		go calcSquare(in, wg)
	}
	//записываем все элементы в канал по очереди
	for _, el := range nums {
		in <- el
	}
	// после того как записали все элементы массива в канал
	// закрываем канал, чтобы горутины закончили свою работу
	close(in)
	wg.Wait()
}

//Передача элемента массива как параметр функции
func sol2() {

	wg := &sync.WaitGroup{}
	//Передаем каждый элемент массива в анонимную функцию напрямую
	for _, el := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			print(num*num, " ")
		}(el)
	}

	wg.Wait()
}

func main() {
	sol1(2)
	fmt.Println()
	sol2()
}
