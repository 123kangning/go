package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1e9)
	timer := time.NewTimer(7e9)
	after := time.After(1e10)
	for {
		select {
		case <-ticker.C:
			fmt.Println("ping pong")
		case <-timer.C:
			fmt.Println("快结束了")
		case <-after:
			fmt.Println("结束啦")
			return
		}
	}
}
