package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	stack, top := make([]int, 10), 0
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf(" err == %v", err)
		}
		input = strings.Replace(input, "\n", "", 1)
		//fmt.Printf("|%s|", input)
		if strings.Compare("q", input) == 0 {
			break
		}
		switch input[0] {
		case '+':
			stack[top-2] = stack[top-2] + stack[top-1]
			top--
		case '-':
			stack[top-2] = stack[top-2] - stack[top-1]
			top--
		case '*':
			stack[top-2] = stack[top-2] * stack[top-1]
			top--
		case '/':
			stack[top-2] = stack[top-2] / stack[top-1]
			top--
		default:
			stack[top], _ = strconv.Atoi(input)
			top++
		}
		//fmt.Println()
		/*for _, num := range stack {
			fmt.Printf("%d ", num)
		}*/
	}
	fmt.Println(stack[0])
}
