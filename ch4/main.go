package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	// method1
	p2 := Person{}
	p2.Name = "Go语言"
	p2.Age = 15
	fmt.Println(p2)

	p3 := Person{"Java语言", 25}
	fmt.Println(p3)

	var p4 *Person = new(Person)
	fmt.Println(p4)
}
