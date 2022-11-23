package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile("go1/main/chapter_12/output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputwriter := bufio.NewWriter(outputFile)
	outputString := "hello world\n"
	for i := 0; i < 10; i++ {
		outputwriter.WriteString(outputString)
	}
	outputwriter.Flush()
	os.Stderr.WriteString("hello, world\n")
}
