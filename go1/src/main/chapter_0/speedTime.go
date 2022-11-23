package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		time.Sleep(1)
	}
	end := time.Now()
	fmt.Printf("speedTime = %v", end.Sub(start))
}
