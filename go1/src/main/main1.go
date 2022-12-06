package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	i, len := 0, len(nums)
	for ; i <= 3 && i < len; i++ {
		//if i == 3 {
		//	break
		//}
	}
	fmt.Println(" i =", i)
}
