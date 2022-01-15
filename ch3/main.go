package main

import (
	"fmt"
)

// 常量定义方式
const (
	n1 = iota
	n2
	n3
	n4 = iota
)
const pi = 3.14159
const e = 2.7
const n5 = iota

func main() {
	case07()
}

func case01() {

	// 标准声明方式
	// 声明的变量都有默认初始值
	var var1 string
	var var2 int = 100 // 声明的同时指定初始值
	var var3 bool
	var var4 float64 = 100.0
	var var5 byte
	fmt.Println(var1, var2, var3, var4, var5)

	// 批量声明方式
	var (
		var6 string
		var7 int
		var8 bool
		var9 byte
	)
	fmt.Println(var6, var7, var8, var9)

	// 类型推断
	var var10 = "gohoer"
	var var11 = 10
	fmt.Printf("type of var10: %T \n", var10)
	fmt.Printf("type of var11: %T \n", var11)

	// 短声明方式
	var12 := "go Programer"
	fmt.Println(var12)

	// 常量
	fmt.Println(pi, e, n1, n4, n5)
}

func case02() {
	var var01 = 25
	fmt.Printf("二进制(binary)表示方式: %b \n", var01)
	fmt.Printf("八进制(octal)表示方式:  %o \n", var01)
	fmt.Printf("十进制(decimal)表示方式: %d \n", var01)
	fmt.Printf("十六进制(hex)表示方式: %x \n", var01)
}

// 转义字符
func case03() {
	var01 := "Hello World"
	fmt.Println(var01)
	var02 := "D:\\WorkSpace\\Go\\go-based-case"
	fmt.Println(var02)
}

func case04() {

	var var01 byte = 'c'
	var var02 rune = 'c'

	fmt.Printf("type of var01 : %T ,value of var01 %d \n", var01, var01)
	fmt.Printf("type of var02 : %T ,value of var01 %d \n", var02, var02)

	var var03 = "this a string"
	for i := 0; i < len(var03); i++ {
		fmt.Printf("%c ", var03[i])
	}
}

func case05() {
	if score := 99; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func case06() {

	/*	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	/*	var var01 = 0
		for ; var01 < 10; var01++ {
			fmt.Println(var01)
		}*/

	// break 跳出for循环
	/*	for var02 := 0; var02<5; var02++ {
		if var02 == 3{
			break;
		}
		fmt.Println(var02)
	}*/

}

func case07() {
	var01 := 8
	switch var01 {
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	case var01 >= 4:
		fmt.Println("D")
	default:
		fmt.Println("无效输入")
	}
}
