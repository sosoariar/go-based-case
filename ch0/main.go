package main

import "fmt"

func main() {
	printHellWorld()
}

func printHellWorld() {
	fmt.Println("Hello world")
	fmt.Println(sum(2, 1))
	printEscapeCharacter()
	variable()
}

func sum(num1 int, num2 int) (int, int) {

	sum := num1 + num2
	sub := num1 - num2

	return sum, sub
}

func printEscapeCharacter() {
	fmt.Println("escape\tCharacter")
	fmt.Println("escape\nCharacter")
	fmt.Println("escape\\Character")
	fmt.Println("escape\"Character")
	fmt.Println("escape\rCharacter")
}

func variable() {
	var i int
	i = 100
	fmt.Println("variables: %d", i)
}
