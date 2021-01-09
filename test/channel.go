package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var a string
var c = make(chan int, 1)

var l sync.Mutex

func Producer(begin, end int, queue chan<- int,wg *sync.WaitGroup) {
	for i:= begin ; i < end ; i++ {
		fmt.Println("produce:", i)
		queue <- i
		atomic.AddInt64(&count,1)
	}
	defer func() {
		wg.Done()
	}()
}
//消费者
func Consumer(queue <-chan int) {
	for val := range queue  { //当前的消费者循环消费
		fmt.Println("consume:", val)
	}

}
var count int64
func main() {
	queue := make(chan int)
	defer close(queue)
	w := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		w.Add(1)
		go Producer(i * 5, (i+1) * 5, queue,&w) //多个生产者
	}
	go Consumer(queue) //单个消费者
    w.Wait()
	fmt.Println("共生产：",count)

}
