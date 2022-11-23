package main

import (
	"fmt"
	"reflect"
)

type proc struct {
	// 结构中只有被导出字段（首字母大写）才是可设置的
	A1 int
	A2 int
	A3 int
}

func main() {
	p := proc{1, 2, 3}
	fmt.Println(p)
	newp := reflect.ValueOf(p)
	num := newp.NumField()
	for i := 0; i < num; i++ {
		fmt.Printf("%d feld is %v\n", i, newp.Field(i))
	}
	change := reflect.ValueOf(&p)
	change = change.Elem()
	num1 := change.NumField()
	change.Field(0).SetInt(10)
	for i := 0; i < num1; i++ {
		fmt.Printf("%d feld is %v\n", i, change.Field(i))
	}
}
