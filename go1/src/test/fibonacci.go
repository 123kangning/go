package main

import (
	"fmt"
	"time"
)

func main() {
	f := fibonacci()
	start := time.Now()
	for i := 0; i < 40; i++ {
		fmt.Println(f())
	}
	end := time.Now()
	fmt.Printf("speed time = %v", end.Sub(start))
}
func fibonacci() func() int {
	a := 0
	b := 1
	fmt.Println(0)
	return func() int {
		b = a + b
		a = b - a
		return b
	}
}
