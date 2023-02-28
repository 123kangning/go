package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := new([]int)
	*s = make([]int, 2, 5)
	//s[1] = 3
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}
func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	ans := make([]int, 1001)
	for _, v := range items1 {
		ans[v[0]] += v[1]
	}
	for _, v := range items2 {
		ans[v[0]] += v[1]
	}
	res := make([][]int, 0, 1)
	for i, v := range ans {
		res = append(res, []int{i, v})
	}
	return res
}
