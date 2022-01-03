package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//printHellWorld()
	//fmt.Println(sumAndSub(2, 1))
	//printEscapeCharacter()
	//variable()
	//fmt.Printf("sum : %d", sum())
	//sizeOfint()
	sizeOfFloat()
}

// 变量占用空间大小
func sizeOfint() {

	var var1 int
	var var2 int8
	var var3 int16
	var var4 int32
	var var5 int64
	var var6 = 1
	var var7 = 128
	var var8 = 65537

	fmt.Printf("%T size of int :  %d  \n", var1, unsafe.Sizeof(var1))
	fmt.Printf("%T size of int :  %d  \n", var2, unsafe.Sizeof(var2))
	fmt.Printf("%T size of int :  %d  \n", var3, unsafe.Sizeof(var3))
	fmt.Printf("%T size of int :  %d  \n", var4, unsafe.Sizeof(var4))
	fmt.Printf("%T size of int :  %d  \n", var5, unsafe.Sizeof(var5))
	fmt.Printf("%T size of int :  %d  \n", var6, unsafe.Sizeof(var6))
	fmt.Printf("%T size of int :  %d  \n", var7, unsafe.Sizeof(var7))
	fmt.Printf("%T size of int :  %d  \n", var8, unsafe.Sizeof(var8))

}

func sizeOfFloat() {
	var var1 float32
	fmt.Printf("%T size of int :  %d  \n", var1, unsafe.Sizeof(var1))

	// 精度丢失
	var var2 float32 = -123.0000001
	var var3 float64 = -123.0000001
	fmt.Println(var2)
	fmt.Println(var3)

	var4 := 1.23
	var5 := .123
	fmt.Println(var4)
	fmt.Println(var5)

	var6 := 5.1234e2
	var7 := 5.1234e2
	var8 := 5.1234e-2
	fmt.Println("var: ", var6)
	fmt.Println("var: ", var7)
	fmt.Println("var: ", var8)
}

func sum() int {

	var num1 = 1
	var num2 = 1

	var str1 = "Hello"
	var str2 = "World"
	fmt.Println(str1 + " " + str2)

	return num1 + num2
}

func printHellWorld() {
	fmt.Println("Hello world")
}

func sumAndSub(num1 int, num2 int) (int, int) {

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
