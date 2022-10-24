package main

import (
	"fmt"
	"sync"
)

// Возьмем некоторые сравниваемые типы для ключа
type ComparissionTypes interface {
	int | int64 | int32 | int8 | int16 |
		float64 | float32 | string |
		complex64 | complex128 | bool
}

// Для доступа к критической секции будем использовать RWmutex
type ConcurrentMap[keyT ComparissionTypes, valT any] struct {
	mtx     sync.RWMutex
	concMap map[keyT]valT
}

func NewConcMap[keyT ComparissionTypes, valT any]() *ConcurrentMap[keyT, valT] {
	return &ConcurrentMap[keyT, valT]{concMap: make(map[keyT]valT)}
}

// Для создания критческой секции используем методы мьютекса Lock и Unlock
func (cm *ConcurrentMap[keyT, valT]) Set(key keyT, val valT) {
	cm.mtx.Lock()
	cm.concMap[key] = val
	cm.mtx.Unlock()
}

func (cm *ConcurrentMap[keyT, valT]) Get(key keyT) (valT, bool) {
	cm.mtx.RLock()
	v, ok := cm.concMap[key]
	cm.mtx.RLock()
	return v, ok
}

func main() {

	cm := NewConcMap[int, int]()

	goroutines := 10

	wg := &sync.WaitGroup{}

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(key int, val int, cm *ConcurrentMap[int, int]) {
			cm.Set(key, val)
			wg.Done()
		}(i, i*i, cm)
	}

	wg.Wait()
	fmt.Println(cm)
}
