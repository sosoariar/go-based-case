package main

import "fmt"

/*
符合数据类型是由基本数据类型以各种方式组合而构成
1. 数组
2. slice
3. map
4. 结构体
*/
func main() {
	fun01()
}

func fun01() {

	var arr1 [3]int

	fmt.Println(arr1[0])
	fmt.Println(arr1[len(arr1)-1])

	var arr2 [3]int = [3]int{4, 5, 6}
	for i, v := range arr2 {
		fmt.Printf("index: %d , value: %d \n", i, v)
	}
}
