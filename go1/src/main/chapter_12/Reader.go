package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("123456789")
	temp := make([]byte, 4)
	for {
		n, err := r.Read(temp)
		if n > 0 {
			fmt.Printf("size = %v , |%c|\n", n, temp[:n])
		}
		if err == io.EOF {
			break
		}
	}
}
