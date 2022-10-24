package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Reader(data <-chan int) {
	for d := range data {
		fmt.Println(d)
	}
}

func main() {
	//создаем контекст, который завершается по истичению заданного времени
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	c := make(chan int)

	before := time.Now()

	go Reader(c)

	defer cancel()
L:
	for {
		select {
		case <-ctx.Done():
			//по истичению времени закрываем канал для записи данных
			close(c)
			break L
		default:
			//пишмм в данные в канал
			c <- rand.Int()
		}

	}

	fmt.Println(time.Now().Sub(before))
}
