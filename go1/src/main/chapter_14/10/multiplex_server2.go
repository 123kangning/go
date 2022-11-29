package main

//c/s,使用quit管道结束server协程
import "fmt"

type Request struct {
	a, b   int
	replyc chan int //reply channel inside Request
}
type binOp func(a, b int) int

func (req *Request) String() string {
	return fmt.Sprintf("{a:%d;b:%d;replyc:%v}", req.a, req.b, req.replyc)
}
func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit:
			return
		}
	}
}

func startServer(op binOp) (service chan *Request, quit chan bool) {
	service = make(chan *Request)
	quit = make(chan bool)
	go server(op, service, quit)
	return service, quit
}

func main() {
	service, quit := startServer(func(a, b int) int { return a + b })
	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		service <- req
	}
	quit <- true
	for i := N - 1; i >= 0; i-- { // doesn't matter what order
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Printf("%s\n", &reqs[i])
		}
	}

	fmt.Println("done")

}
