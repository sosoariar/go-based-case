package main

import "fmt"

func main() {
	var a int = 10
	var p *int = &a
	fmt.Println("a = ", *p)
	fmt.Println("address: ", p)

}
