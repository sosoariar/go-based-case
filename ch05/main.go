package main

import "fmt"

func main() {
	//helloWorld := func(){
	//	fmt.Println("Hello World")
	//}
	//helloWorld()

	var01 := printHello()
	var01()
}

// 闭包,如果一个函数只会使用一次,但是这个函数需要使用环境变量
func printHello() func() {
	name := "Golang"
	return func() {
		fmt.Println("Hello ", name)
	}
}

func case01() {
	defer fmt.Println("function end")
	fmt.Println(addSum(1, 2, 3, 4, 5))
	fmt.Println(addSum(10, 20, 30, 40, 50))

	// 函数也是一种类型
	var01 := addSum(1, 2, 3)
	fmt.Println(var01)

	// 函数作为一种类型
	fmt.Println(calc(100, 200, add))
}

func addSum(argArr ...int) int {
	ret := 0
	for _, arg := range argArr {
		ret += arg
	}
	return ret
}

// 函数不能重载,且多个参数时,可变参数一定要放在最后
/*func addSum(arg01 int,argArr ... int) int{

}*/

// 函数作为参数
func add(x, y int) int {
	return x + y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
