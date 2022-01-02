package main

import (
	"fmt"
	"reflect"
)

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
	// 指定变量类型,如果没有赋值,使用各种类型的默认值
	var var1 int
	var1 = 100
	fmt.Println("variables: %d", var1)

	// 根据赋值自动判断变量类型
	var var2 = 10.0
	var var3 = "golang"
	fmt.Println(reflect.TypeOf(var2))
	fmt.Println(reflect.TypeOf(var3))
}
