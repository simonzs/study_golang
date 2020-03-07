package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	var n int
	var ok bool
	for {
		if n, ok = <-c; !ok {
			break
		}
		// time.Sleep(3 * time.Microsecond)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int { // 往chan发数据
	// func createWorker(id int) <-chan int { // 从chan收数据
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	// var c chan int // c == nil
	var channels [10]chan<- int
	// c := make(chan int)

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)

	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	// chanDemo()
	// bufferedChannel()
	channelClose()
}
