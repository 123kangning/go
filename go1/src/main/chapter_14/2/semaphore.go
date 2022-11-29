package main

type Empty interface{}
type semaphore chan Empty

func main() {
	N, n := 10, 10
	sem := make(semaphore, N)
	go func(s semaphore) {
		e := new(Empty)
		for i := 0; i < n; i++ {
			s <- e
		}
	}(sem)

	// release n resources
	go func(s semaphore) {
		for i := 0; i < n; i++ {
			<-s
		}
	}(sem)
}
