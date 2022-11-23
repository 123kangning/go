package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5}
	fmt.Println(num)
	/*removeSlice(&num, 2, 3)*/
	/*insertSlice(&num, []int{9, 9, 9}, 2)*/
	/*	filterSlice(&num,
		func(a int) bool {
			return a%2 == 0
		})*/
	/*magnifySlice(&num, 3)*/
	fmt.Println(num)
}
func removeSlice(s1 *[]int, start int, end int) {
	s := *s1
	len := len(s)
	if start < 0 {
		panic("start < 0")
	}
	if end > len {
		end = len
	}
	last := s[end:len]
	s = s[:start]
	s = append(s, last...)
	*s1 = s
}
func insertSlice(s1 *[]int, sub []int, index int) {
	s := *s1
	len := len(s)
	if index > len {
		index = len
	}
	if index < 0 {
		panic("index < 0")
	}
	last := make([]int, len-index)
	copy(last, s[index:])
	s = s[:index]
	s = append(s, sub...)
	s = append(s, last...)
	*s1 = s
}
func filterSlice(s1 *[]int, fn func(int) bool) {
	s := *s1
	len := len(s)
	for i := 0; i < len; i++ {
		if fn(s[i]) {
			s = append(s[:i], s[i+1:]...)
			i--
			len--
		}
	}
	*s1 = s
}
func magnifySlice(s1 *[]int, factor int) {
	s := *s1
	len := len(s)
	len *= factor - 1
	for i := 0; i < len; i++ {
		s = append(s, 0)
	}
	*s1 = s
}
