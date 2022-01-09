package main

import (
	"crypto/sha256"
	"fmt"
)

/*
符合数据类型是由基本数据类型以各种方式组合而构成
1. 数组
2. slice
3. map
4. 结构体
*/
func main() {
	//fun01()
	//fun02()
	fun03()
}

func fun03() {
	arr1 := []int{0, 1, 2, 3, 4, 5}
	reverse(arr1[:])
	fmt.Println(arr1)
}

func fun01() {

	var arr1 [3]int

	fmt.Println(arr1[0])
	fmt.Println(arr1[len(arr1)-1])

	var arr2 [3]int = [3]int{4, 5, 6}
	for i, v := range arr2 {
		fmt.Printf("index: %d , value: %d \n", i, v)
	}

	arr3 := [...]int{7, 8, 9, 10}
	fmt.Printf("%T \n", arr3)
	for i, v := range arr3 {
		fmt.Printf("index: %d , value: %d \n", i, v)
	}

}

func fun02() {

	var1 := sha256.Sum256([]byte("x"))
	var2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n%t\n%T\n", var1, var2, var1 == var2, var1)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
