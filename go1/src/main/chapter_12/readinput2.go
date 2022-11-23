package main

import "fmt"

func main() {
	name, age := "", 0
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("name =%s , age = %d", name, age)
}
