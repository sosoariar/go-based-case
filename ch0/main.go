package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func main() {

	//printHellWorld()
	//fmt.Println(sumAndSub(2, 1))
	//printEscapeCharacter()
	//variable()
	//fmt.Printf("sum : %d", sum())
	//sizeOfint()
	//sizeOfFloat()
	//sizeOfByte()
	//caseOfBool()
	//caseOfString()
	//tranType()
	//pointCase()
	caseOfMethod()
}

func init() {
	fmt.Println("init")
}
func caseOfMethod() {
	var1 := getSum
	fmt.Println(var1)
	fmt.Printf("%T \n", var1)

	var2 := var1(10, 10)
	fmt.Println(var2)

	var3 := funCase(getSum, 10, 20)
	fmt.Println(var3)
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func funCase(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func pointCase() {
	var var1 int
	fmt.Println("var1 memory address: ", &var1)
	fmt.Println("var1 value: ", var1)

	// 指针变量存储内存地址
	fmt.Println("--------------------------------------------------------")
	var var2 *int
	var2 = &var1
	fmt.Println("var2 value: ", var2)
	fmt.Println("var2 memory address: ", &var2)

	// 通过地址获取变量后,修改值
	fmt.Println("--------------------------------------------------------")
	*var2 = 10
	fmt.Println("var1 value: ", var1)

}

func tranType() {
	var var1 int32 = 1
	var var2 float32 = float32(var1)
	var var3 int8 = int8(var1)
	var var4 int64 = int64(var1)

	fmt.Printf("var1= %T var2 = %T var3 = %T var4 = %T \n", var1, var2, var3, var4)

	// 转换溢出处理方式
	var var5 int64 = 999999
	var var6 int8 = int8(var5)
	fmt.Println(var6)

	// 基本数据类型转成 string
	var var7 int = 99
	var var8 float64 = 1.234
	var var9 bool = true
	//var var10 byte = 'a'
	var str string

	str = fmt.Sprintf("%d \n", var7)
	fmt.Printf("the type of \"str\" : %T , the value of str: %v \n", str, str)

	str = fmt.Sprintf("%f \n", var8)
	fmt.Printf("the type of \"str\" : %T , the value of str: %v \n", str, str)

	str = fmt.Sprintf("%t \n", var9)
	fmt.Printf("the type of \"str\" : %T , the value of str: %v \n", str, str)

	// 第二种类型转换
	var var10 int = 99
	str = strconv.FormatInt(int64(var10), 10)
	fmt.Printf("the type of \"str\" : %T , the value of str: %v \n", str, str)

	// 第三种类型转换
	var var11 string = "true"
	var var12 bool
	var12, _ = strconv.ParseBool(var11) // 只想获得 value bool , 不想获得 err 使用下划线忽略返回值
	fmt.Printf("the type of \"var12\" : %T , the value of var12: %v \n", var12, var12)

}

func caseOfString() {
	var var1 = `\n \t \o \\ \"`
	fmt.Println(var1)

	var var2 = "Hello"
	var2 += "World"
	fmt.Println(var2)
}

func caseOfBool() {

	var var1 bool
	// 默认值
	fmt.Println(var1)

	var var2 = true
	fmt.Println(var2)
	fmt.Printf("%T size of int :  %d  \n", var2, unsafe.Sizeof(var2))

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

func sizeOfByte() {

	var var1 byte = 'a'
	var var2 byte = '0'
	var var3 int = '陈'
	var var4 = '陈'

	fmt.Println("var1 = : ", var1)
	fmt.Println("var2 = : ", var2)
	fmt.Printf("%T size of int :  %c  \n", var3, var3)
	fmt.Printf("%T size of int :  %c  \n", var4, var4)

	// 输出对应的unicode 对应的字符
	var var5 int = 22269
	fmt.Printf("var5 : %c \n", var5)

	var var6 = 10 + 'a'
	fmt.Printf("var6 : %c \n", var6)

	var var7 = 10 + 'a'
	fmt.Println("var7 : ", var7)
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
