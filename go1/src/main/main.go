package main

import "fmt"

func main() {
	//myMap := map[string]int{
	//	"Java":   11,
	//	"Perl":   8,
	//	"Python": 13,
	//	"Shell":  23,
	//}
	//
	//println(len(myMap)) // 4
	//
	//delete(myMap, "Perl")
	//
	//println(len(myMap)) // 3
	fmt.Println(numDifferentIntegers("234iujhk1243kujhih1243iuuh2134ikujh23i41ujh123ik4ujhi23uj4hiu"))
}
func numDifferentIntegers(word string) int {
	nums := map[int]int{}
	len := len(word)
	for i := 0; i < len; i++ {
		if word[i] <= '9' && word[i] >= '0' {
			num := getNum(word, &i)
			//fmt.Println(num)
			nums[num] = 1
		}
	}
	return 10
}
func getNum(word string, start *int) int {
	i, len, ans := *start, len(word), 0
	for ; i < len && word[i] <= '9' && word[i] >= '0'; i++ {
		ans = 10*ans + int(word[i]-'0')
	}
	*start = i
	return ans
}
