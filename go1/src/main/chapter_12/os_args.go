package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice "
	fmt.Println("len = ", len(os.Args))
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
