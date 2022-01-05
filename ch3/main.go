package main

import (
	"fmt"
	"math"
)

const e = 2.71828
const Avogadro = 6.02214129e23
const Planck = 6.62606957e23

func main() {
	//fun01()
	//fun02()
	fun03()
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
