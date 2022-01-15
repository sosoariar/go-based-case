package main

import "fmt"

func main() {
	fmt.Println(addSum(1, 2, 3, 4, 5))
	fmt.Println(addSum(10, 20, 30, 40, 50))
}

func addSum(argArr ...int) int {
	ret := 0
	for _, arg := range argArr {
		ret += arg
	}
	return ret
}
