package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}
	s := arr[1:3:5]

	arr = append(arr, 888)
	arr = append(arr, 888)
	arr = append(arr, 888)
	fmt.Println(arr)
	fmt.Println("s = ", s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
