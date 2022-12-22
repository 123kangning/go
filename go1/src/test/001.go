package main

//继承测试
import "fmt"

type book struct {
	english int
	math    int
	box     int
}
type store struct {
	box int
}

type shop struct {
	book
	pen int
	//box int
	store
}

func main() {
	s := shop{book{3, 4, 5}, 2, store{6}}
	fmt.Println(s.book.box)
}
