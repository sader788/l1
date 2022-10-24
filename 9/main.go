package main

import (
	"fmt"
	"sync"
)

var arr = []int{2, 4, 6, 8, 10}

func Reader(wg *sync.WaitGroup, c chan<- int) {
	defer wg.Done()

	for n := range arr {
		c <- n
	}

	close(c)
}

func Conv(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * n
	}
	close(out)
}

func Squares(wg *sync.WaitGroup, c <-chan int) {
	defer wg.Done()

	for n := range c {
		fmt.Println(n)
	}
}

func main() {
	//создаем два канала, один для горутины, которая читает из массива,
	//второй горутины возводящей в квадрат
	cReader := make(chan int)
	cSquares := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	//Читает из массива
	go Reader(wg, cReader)
	//Получает данные из Reader и передает в Squares
	go Conv(cReader, cSquares)
	//Возводит в квадрат
	go Squares(wg, cSquares)

	wg.Wait()
}
