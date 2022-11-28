package main

import (
	"fmt"
	"time"
)

func main() {
	stream := facPump() //从工厂方法中接收管道
	facSuck(stream)
	time.Sleep(1e9)
}

func facPump() chan int { //工厂方法返回管道
	ch := make(chan int)
	go func() {
		for i := 0; ; i += 2 {
			ch <- i
		}
	}()
	go func() {
		for i := 1; ; i += 2 {
			ch <- i
		}
	}()
	return ch
}

func facSuck(ch chan int) {
	func() {
		for i := 0; i < 100000000; i++ {
			<-ch
		}
	}()
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
}
