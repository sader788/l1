package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ex1(ctx context.Context, wg *sync.WaitGroup, in <-chan int) {
	defer wg.Done()
	for {
		select {
		case d := <-in:
			fmt.Println("ex1: ", d)
		case <-ctx.Done():
			fmt.Println("ex1: stop")
			return
		}
	}
}

var signal chan struct{}

func ex2(wg *sync.WaitGroup, in <-chan int) {
	defer wg.Done()
	for {
		select {
		case d := <-in:
			fmt.Println("ex2: ", d)
		case <-signal:
			fmt.Println("ex2: stop")
			return
		}
	}
}

func ex3(wg *sync.WaitGroup, in <-chan int) {
	defer wg.Done()
	for i := range in {
		fmt.Println("ex3: ", i)
	}
	fmt.Println("ex3: stop")
}

func main() {
	////////////////////////////example1//////////////////////////////////
	//Пример завершения работы горутины с помощью контекста
	c := make(chan int)
	//Создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go ex1(ctx, wg, c)
	//Иммитация работы горутины
	for i := 0; i < 4; i++ {
		c <- rand.Int()
		time.Sleep(time.Second)
	}
	//По завршению работы с горутиной вызываем cancel функцию
	//Которая закрывает канла Done данного контекста
	cancel()
	wg.Wait()

	////////////////////////////example2//////////////////////////////////
	//Пример с использованием сигнального канала
	signal = make(chan struct{})

	wg.Add(1)
	go ex2(wg, c)

	for i := 0; i < 4; i++ {
		c <- rand.Int()
		time.Sleep(time.Second)
	}
	//Отправляем структуру в канал, когда желаем завершить горутину
	signal <- struct{}{}
	wg.Wait()
	////////////////////////////example3//////////////////////////////////
	//Пример с закрытием канала
	wg.Add(1)
	go ex3(wg, c)

	for i := 0; i < 4; i++ {
		c <- rand.Int()
		time.Sleep(time.Second)
	}
	//Закрываем канал, тк в функции ex3 range цикл
	//Горутина выйдет из цикла и закончит работу
	close(c)
	wg.Wait()
}
