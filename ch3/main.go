package main

import "fmt"

func main() {
	//fun01()
	fun02()
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
