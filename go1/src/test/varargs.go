package main

import "fmt"

func main() {
	MyPrintln2(1, 2, true, false, "aaa", "sss", 2.9)
	//MyPrintln1(1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1)
}
func MyPrintln2(arg ...interface{}) {
	for _, value := range arg {
		switch v := value.(type) {
		case int:
			fmt.Printf("%T --> %v\n", v, value)
		case bool:
			fmt.Printf("%T --> %v\n", v, value)
		case string:
			fmt.Printf("%T --> %v\n", v, value)
		default:
			fmt.Printf("other type\n ")
		}
	}
}
func MyPrintln1(arg ...int) {
	if len(arg) == 0 {
		fmt.Println("---------------------------------")
		return
	}
	fmt.Println(arg[0])
	MyPrintln1(arg[1:]...)
}
