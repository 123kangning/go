package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("三秒没有输入将结束程序")
	timeout := time.NewTimer(3e9)
	ch := make(chan int)
	/*go func() {
		for {
			var num int
			fmt.Scanf("%d\n", &num)
			ch <- num
		}
	}()*/
	go func() {
		var num int
		fmt.Scanf("%d\n", &num)
		ch <- num
	}()
	select {
	case v := <-ch:
		fmt.Println("\n======== ", v, " =======")
	case <-timeout.C:
		fmt.Println("\nfinish")
		break
	}
}
