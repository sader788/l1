package main

import (
	"context"
	"fmt"
	"time"
)

func Sleep1(duration time.Duration) {
	<-time.After(duration)
}

func Sleep2(duration time.Duration) {
	ctx := context.Background()
	timeout, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	<-timeout.Done()
}

func main() {
	seconds := 2

	before := time.Now()
	fmt.Println("Waiting for ", seconds, " seconds (Sleep1)")
	Sleep1(time.Second * time.Duration(seconds))
	fmt.Println("success ", time.Now().Sub(before))

	before = time.Now()
	fmt.Println("Waiting for ", seconds, " seconds (Sleep2)")
	Sleep2(time.Second * time.Duration(seconds))
	fmt.Println("success ", time.Now().Sub(before))
}
