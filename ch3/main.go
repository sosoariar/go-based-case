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
	case01()
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
	fmt.Println(var10)

	// 短声明方式
	var11 := "go Programer"
	fmt.Println(var11)

	// 常量
	fmt.Println(pi, e, n1, n4, n5)
}
