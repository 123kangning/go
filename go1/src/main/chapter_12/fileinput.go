package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	/*	pwd, err := os.Getwd()
		if err == nil {
			fmt.Printf("user.dir = %s", pwd)
		}*/
	inputFile, inputError := os.Open("go1/main/main.go")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return
	}
	defer inputFile.Close()
	read2(inputFile)

}
func read1(inputFile *os.File) {
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readError := inputReader.ReadString('\n')
		if readError == io.EOF {
			return
		}
		fmt.Println(strings.Replace(inputString, "\n", "", -1))
	}
}
func read2(inputFile *os.File) {
	buf, err := ioutil.ReadFile(inputFile.Name())
	if err == nil {
		for _, byte := range buf {
			fmt.Printf("%c", byte)
		}
	}
}
func read3(inputFile *os.File) []byte {
	inputReader := bufio.NewReader(inputFile)
	buf := make([]byte, 1000)
	_, err := inputReader.Read(buf)
	if err == nil {
		return buf
	}
	return buf
}
