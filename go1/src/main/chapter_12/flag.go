package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	add := flag.Bool("n", false, "add-add-add")
	flag.PrintDefaults()
	flag.Parse()

	s := ""
	for _, arg := range flag.Args() {
		s += arg + " "
		fmt.Println(*add)
		if *add {
			s += "add\n"
		}
	}
	//s += "\b"
	os.Stdout.WriteString(s)
}
