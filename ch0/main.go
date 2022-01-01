package main

import "fmt"

func main() {

	var map01 map[int]string

	if map01 == nil {
		fmt.Println("map01 is nil")
	}

	fmt.Println(len(map01))
	fmt.Println("------------------------end------------------------")

	map02 := map[int]string{}

	if map02 == nil {
		fmt.Println("map02 is nil")
	}

	fmt.Println(len(map02))
	fmt.Println("------------------------end------------------------")

	map03 := make(map[int]string)
	if map03 == nil {
		fmt.Println("map03 is nil")
	}
	fmt.Println(len(map03))
	fmt.Println("------------------------end------------------------")

	map04 := make(map[int]string, 10)
	if map04 == nil {
		fmt.Println("map04 is nil")
	}
	fmt.Println(len(map04))
	fmt.Println("------------------------end------------------------")

	map05 := make(map[int]string, 1)
	if map05 == nil {
		fmt.Println("map05 is nil")
	}

	map05[1] = "value01"
	map05[2] = "value02"
	map05[3] = "valye03"
	fmt.Println("map05:", map05)
}
