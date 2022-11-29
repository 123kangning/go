package main

import (
	"fmt"
	"time"
)

func main() {
	prime := sieve()
	for {
		fmt.Println(<-prime)
	}
}

func sieve() chan int {
	ch := generate()
	out := make(chan int)
	go func() {
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
			time.Sleep(1e8)
		}
	}()
	return out
}
func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-in
			if i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}
