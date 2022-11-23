package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (error MyError) Error() string {
	return fmt.Sprintf("at %v , %s", error.When, error.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it don't work",
	}
}
func main() {
	if error := run(); error != nil {
		fmt.Println(error.Error())
	}
}
