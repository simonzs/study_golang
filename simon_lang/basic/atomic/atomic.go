package main

import (
	"fmt"
	"sync"
	"time"
)

// atomicInt 原子化的Int
type atomicInt struct {
	value int
	lock  sync.Mutex
}

// increment ...
func (a *atomicInt) increment() {

	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(1 * time.Millisecond)
	fmt.Println(a.get())
}
