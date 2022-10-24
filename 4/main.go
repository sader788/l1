package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

func Worker(ctx context.Context, wg *sync.WaitGroup, c <-chan int, num int) {
	defer wg.Done()

	for {
		select {
		case d := <-c:
			fmt.Printf("Worker %d: %d\n", num, d)
		case <-ctx.Done():
			fmt.Printf("Worker %d stop\n", num)
			return
		}
	}
}

func main() {
	c := make(chan int)

	workers := 5

	wg := sync.WaitGroup{}
	//создаем контекст, который завершает свою работу при ctrl+c
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go Worker(ctx, &wg, c, i+1)
	}

L:
	for {
		select {
		case <-ctx.Done():
			//использовано сочетние ctrl+c
			break L
		default:
			//пишем данные в канал
			d := rand.Int()
			c <- d
		}
	}
	wg.Wait()
}
