package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ch := futureFib(80)
	ch1 := futureFib(80)
	ch2 := futureFib(80)
	fmt.Println(<-ch, <-ch1, <-ch2)
	end := time.Now()
	fmt.Println("speed time is ", end.Sub(start))
}
func futureFib(n int) chan int {
	resv := Fib()
	retChannel := make(chan int, 1)
	//go func() {
	for i := 0; i < n-1; i++ {
		resv()
	}
	retChannel <- int(resv())
	//}()
	return retChannel
}
