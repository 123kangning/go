package main

import "fmt"

type book struct {
	english int
	math    int
}

type shop struct {
	pen int
	box int
	book
}

func main() {
	s := shop{1, 2, book{3, 4}}
	fmt.Println(s.math)
}
