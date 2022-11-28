package main

import "fmt"

// integer producer:
func numGen(start, count int, out chan int) {
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
	}
	close(out) //关闭管道写端之后，从管道中读完之后继续读会非阻塞返回类型零值
}

// integer consumer:
func numEchoRange(in chan int, done chan bool) {
	for {
		num, ok := <-in
		if !ok {
			break
		}
		fmt.Printf("%d\n", num)
	}
	done <- true
	close(done)
}

func main() {
	numChan := make(chan int)
	done := make(chan bool)
	go numGen(0, 10, numChan)
	go numEchoRange(numChan, done)
	<-done
	<-done
}
