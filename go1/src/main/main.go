package main

import (
	"fmt"
	"os"
)

func main() {
	var user = os.Getenv("USER")
	fmt.Println(user)
}
