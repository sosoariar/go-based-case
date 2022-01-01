package main

import "fmt"

func main() {

	var map01 map[int]string

	if map01 == nil {
		fmt.Println("map is nil")
	}

	fmt.Println(len(map01))
}
