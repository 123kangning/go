package main

import "fmt"

type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println(t.S)
	}
}
func (t *T) String() string {
	return fmt.Sprintf("hahaha---%q---\n", t.S)
}
func main() {
	var j I
	//i = &T{"hello"}

	if sv, ok := j.(I); ok {
		fmt.Printf("%s\n", sv)
	}

}
func describe(i I) {
	fmt.Printf("val = %v , type = %T\n", i, i)
	fmt.Printf("val = %#v , type = %T\n", i, i)
}
