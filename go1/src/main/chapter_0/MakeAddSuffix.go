package main

import "fmt"

func main() {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	fmt.Println(addBmp("111"))
	fmt.Println(addJpeg("222"))
}
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		return name + suffix
	}
}
