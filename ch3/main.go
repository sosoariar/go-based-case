package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

const e = 2.71828
const Avogadro = 6.02214129e23
const Planck = 6.62606957e23

func main() {
	//fun01()
	//fun02()
	//fun03()
	//fun04()
	fun05()
}

/*
	字符串是不可变的字节序列,包含任意数据
	文本字符串被解读成按 UTF-8 编码的 Unicode 序列
	字符串中 第i个字节不一定就是第 i 个字符,因为 UTF-8 是变长字符集, 英文占两个字节,中文占三个字节
	天然适合保存单个文字的数据类型是 int32, 为 Go 使用, rune 类型作为 int32 类型的别名
*/
func fun05() {
	var1 := "Hello World"

	fmt.Println(len(var1))
	fmt.Println(var1[0], var1[7])
	// : 两边的默认值
	fmt.Println(var1[:5])
	fmt.Println(var1[7:])
	fmt.Println(var1[:])

	var2 := "left foot"
	var3 := var2
	var2 += ", right foot"
	fmt.Println(var2)
	fmt.Println(var3)

	// 转义字符
	fmt.Println("\a")
	fmt.Println("\v")
	// 访问字符下标越界
	//fmt.Println(var1[len(var1)])

}

// Go具备两种大小的复数complex64和complex128，二者分别由float32和float64构成。
func fun04() {

	var var1 complex128 = complex(1, 2)
	var var2 complex128 = complex(3, 4)
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(real(var1 * var2))
	fmt.Println(imag(var1 * var2))

	// math/cmplx 包提供了复数运算所需的库函数
	var3 := 1 + 2i
	var4 := 3 + 4i
	fmt.Println(var3)
	fmt.Println(var4)
	fmt.Println(cmplx.Sqrt(-1))
	fmt.Println(!true)

}

func fun03() {

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	var var1 float32 = 16777216
	fmt.Println(var1 == var1+1)

	for x := 0; x < 8; x++ {
		fmt.Printf("x=%d e = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	// 判断其参数是否为非数值
	fmt.Println(math.IsNaN(z))

	// 用于返回非数值
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)

}

func fun02() {

	var1 := 011
	fmt.Printf("十进制 %d \n", var1)
	fmt.Printf("二进制 %b \n", var1)
	fmt.Printf("二进制 %b \n", var1)

	var2 := int64(0xdeadbeef)
	fmt.Printf("%d \n", var2)
}

func fun01() {

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

}
