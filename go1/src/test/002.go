package main

//多态测试
import "fmt"

type A interface {
	say()
}
type B struct {
	word string
}
type C struct {
	word string
}

func main() {
	b, c := B{"I'm B"}, C{"I'm C"}
	var a A = c
	isB(a)
	Hello(b)
	Hello(c)
}
func (b B) say() {
	fmt.Println(b.word)
}
func (c C) say() {
	fmt.Println(c.word)
}
func Hello(a A) {
	a.say()
}
func isB(a A) {
	b, ok := a.(B)
	fmt.Println(" ok = ", ok)
	fmt.Println(b)
}
