package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Счетчик с критической секцией, создаваемой с помощью мьютексов
type CounterMutex struct {
	mtx sync.Mutex
	c   int
}

func (counter *CounterMutex) Inc() {
	counter.mtx.Lock()
	counter.c++
	counter.mtx.Unlock()
}

func (counter *CounterMutex) getValue() int {
	return counter.c
}

func NewCounterMtx() *CounterMutex {
	return &CounterMutex{}
}

// Счетчик с использованием атомарных операций
type CounterAtom struct {
	c int32
}

func (counter *CounterAtom) Inc() {
	atomic.AddInt32(&counter.c, 1)
}

func (counter *CounterAtom) getValue() int32 {
	return counter.c
}

func NewCounterAtom() *CounterAtom {
	return &CounterAtom{}
}

// функция для горутин
func Incrementer(wg *sync.WaitGroup, cntMtx *CounterMutex, cntAtom *CounterAtom) {
	cntMtx.Inc()
	cntAtom.Inc()
	wg.Done()
}

func main() {
	counterMtx := NewCounterMtx()
	counterAtom := NewCounterAtom()
	n := 55555
	wg := &sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go Incrementer(wg, counterMtx, counterAtom)
	}

	wg.Wait()
	fmt.Println("Mutex counter: ", counterMtx.getValue())
	fmt.Println("Atomic counter: ", counterAtom.getValue())

}
