package basic

import "fmt"

func AddOne(num int) int {
	return num + 1
}

func AddOne2(num int) int {
	if 1 == 2 {
		fmt.Println("AddOne2 called with 2")
	}
	return num + 1
}
