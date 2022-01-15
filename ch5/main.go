package main

import "fmt"

func main() {
	fmt.Println(addSum(1, 2, 3, 4, 5))
}

func addSum(argArr ...int) int {
	ret := 0
	for _, arg := range argArr {
		ret += arg
	}
	return ret
}
