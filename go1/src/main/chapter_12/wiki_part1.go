package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	testLoad()

}
func testSave() {
	page := new(Page)
	page.Title = "go1/main/chapter_12/output.dat"
	inputFile, inputError := os.Open("go1/main/main.go")
	if inputError != nil {
		return
	}
	inputReader := bufio.NewReader(inputFile)
	buf := make([]byte, 1000)
	n, _ := inputReader.Read(buf)
	defer inputFile.Close()
	buf = buf[:n]
	page.Body = buf
	ret := page.save()
	if ret >= 0 {
		fmt.Printf("write Body size is %d\n", ret)
	} else {
		fmt.Println("write false")
	}
}
func testLoad() {
	page := new(Page)
	page.Title = "go1/main/chapter_12/output.dat"
	page.load()
	println(string(page.Body))
}

// return write size when success,return -1 when error
func (page *Page) save() int {
	outputFile, outputError := os.OpenFile(page.Title, os.O_WRONLY|os.O_CREATE, 0x664)
	if outputError != nil {
		fmt.Printf("page.Title == %s is error ,err = %v", page.Title, outputError)
		return -1
	}
	defer outputFile.Close()
	outputWrite := bufio.NewWriter(outputFile)
	n, writeErr := outputWrite.Write(page.Body)
	outputWrite.Flush()
	if writeErr == nil {
		return n
	} else {
		return -1
	}
}
func (page *Page) load() int {
	inputFile, inputError := os.OpenFile(page.Title, os.O_RDONLY, 0x664)
	if inputError != nil {
		fmt.Println("err is ", inputError)
		return -1
	}
	buf := make([]byte, 1000)
	inputRead := bufio.NewReader(inputFile)
	sum := 0
	n, err := inputRead.Read(buf)
	sum += n
	page.Body = buf
	for ; n > 0 && err != io.EOF; n, err = inputRead.Read(buf) {
		page.Body = append(page.Body, buf...)
		sum += n
	}
	page.Body = page.Body[:sum]
	return sum
}
