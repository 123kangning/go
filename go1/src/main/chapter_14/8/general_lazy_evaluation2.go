package main

//	func main() {
//		even := Fib()
//		for i := 0; i < 100; i++ {
//			fmt.Println(even())
//		}
//	}
func Fib() func() uint64 {
	var ret1, ret2 uint64
	ret1, ret2 = 1, 1
	retChannel := make(chan uint64)
	loop := func() {
		retChannel <- 1
		for {
			ret1, ret2 = ret2, ret1+ret2
			retChannel <- ret1
		}
	}
	go loop()
	return func() uint64 {
		return <-retChannel
	}
}
