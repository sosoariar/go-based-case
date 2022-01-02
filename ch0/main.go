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

	fmt.Println("-----------------------------------------------")
	// 根据赋值自动判断变量类型
	var var2 = 10.0
	var var3 = "golang"
	fmt.Println(reflect.TypeOf(var2))
	fmt.Println(reflect.TypeOf(var3))

	fmt.Println("-----------------------------------------------")
	// 省略 var := 左侧的变量不应该是已经声明过的,否则会导致编译错误
	var4 := 10.0
	var5 := 10
	var6 := "golang"
	fmt.Println(reflect.TypeOf(var4))
	fmt.Println(reflect.TypeOf(var5))
	fmt.Println(reflect.TypeOf(var6))

	// 一次性声明多个变量
	fmt.Println("-----------------------------------------------")
	var var7, var8 int
	fmt.Println(reflect.TypeOf(var7))
	fmt.Println(reflect.TypeOf(var8))

	var var9, var10 = 10, "golang"
	fmt.Println(reflect.TypeOf(var9))
	fmt.Println(reflect.TypeOf(var10))

	fmt.Println("-----------------------------------------------")
	var (
		var11 int
		var12 string
	)
	fmt.Println(reflect.TypeOf(var11))
	fmt.Println(reflect.TypeOf(var12))
}
