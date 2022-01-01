package main

import "fmt"

func main() {

	var map01 map[int]string

	if map01 == nil {
		fmt.Println("map01 is nil")
	}

	fmt.Println(len(map01))

	map02 := map[int]string{}

	if map02 == nil {
		fmt.Println("map02 is nil")
	}

	fmt.Println(len(map02))
}
