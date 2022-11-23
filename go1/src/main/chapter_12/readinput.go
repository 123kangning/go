package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString(' ')
	if err == nil {
		fmt.Printf("input = %s , err = %v", input, err)
	} else {
		fmt.Println(err)
	}
}
