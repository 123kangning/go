package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var newLine = flag.Bool("n", false, "")
var line = 1

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if *newLine {
			fmt.Fprintf(os.Stdout, "%d ", line)
			line++
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
		if err == io.EOF {
			break
		}
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}
