package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("go1/main/chapter_12/output.dat", "go1/main/main.go")
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
