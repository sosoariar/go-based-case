package main

import (
	"fmt"
)

func main() {
	case01()
}

func case01() {

	// 标准声明方式
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

}
