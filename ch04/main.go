package main

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"time"
)

/*
符合数据类型是由基本数据类型以各种方式组合而构成
1. 数组
2. slice
3. map
4. 结构体
*/

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	fun11()
}

func fun11() {

	var var01 map[string]int
	fmt.Println(var01 == nil)

	fmt.Println("--------- make 初始化方式 一 -----------")
	var01 = make(map[string]int, 10)
	fmt.Println(var01 == nil)

	var01["java01"] = 69
	var01["java02"] = 99
	fmt.Println(var01)

	fmt.Println("--------- make 初始化方式 二 -----------")
	var02 := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Println(var02)

	fmt.Println("--------- 判断某个键是否存在 -----------")
	value, ok := var02[3]
	if ok {
		fmt.Println("容器中存在:", value)
	} else {
		fmt.Println("不存在此元素")
	}

}

func fun09() {

	var var01 []int //此时并没有申请内存
	fmt.Println(var01)

	var01 = append(var01, 1)
	var01 = append(var01, 1)
	var01 = append(var01, 1)
	var01 = append(var01, 1)

	for i := 5; i < 10; i++ {
		var01 = append(var01, i)
	}

	var01 = append(var01, 2, 3, 4)
	fmt.Println(var01)

}

func fun10() {

	var01 := []int{1, 2, 3, 4, 5}
	var02 := make([]int, 5, 5)
	copy(var02, var01)
	fmt.Println(var01)
	fmt.Println(var02)

	var var03 = [...]int{8, 7, 6, 5, 4, 3, 2, 1, 0}
	sort.Ints(var03[:])
	fmt.Println(var03)

}

func fun08() {

	var var01 = []string{}
	var var02 = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("the type of []string{} %T,  \n", var01)
	fmt.Printf("the type of []string{} %T,  \n", var02)

	var03 := var02[1:4] //左闭右开的结果
	fmt.Println(var03)

	var03 = var02[1:]
	fmt.Println(var03)

	var03 = var02[:4]
	fmt.Println(var03)

	var03 = var02[0:len(var02)]
	fmt.Println(var03)

	fmt.Println("----------------- 通过 make 函数构造切片 -----------------")
	// 容量为10,实际存储为 5, 类型为 []int 的切片
	var04 := make([]int, 5, 10)
	fmt.Println(var04)
	fmt.Printf("%T \n", var04)

}

func fun01() {

	var arr1 [3]int
	var arr2 [3]int

	fmt.Println(arr1)
	fmt.Println(arr2)

	// 定义时使用初始值列表的方式初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(cityArray)

	var boolArray = [...]bool{true, false, true, false, true}
	fmt.Println(boolArray)

	var mapArray = [...]int{1: 11, 3: 13, 7: 17}
	fmt.Println(mapArray)

	fmt.Println("------------- 索引遍历方式 -------------")
	for i := 0; i < len(mapArray); i++ {
		fmt.Printf("%d ", mapArray[i])
	}

	for index, value := range mapArray {
		fmt.Printf("index: %d, value: %d \n", index, value)
	}

}

func fun02() {

	var1 := sha256.Sum256([]byte("x"))
	var2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n%t\n%T\n", var1, var2, var1 == var2, var1)

}

func fun07() {

	var dilbert Employee
	dilbert.ID = 1101
	dilbert.Name = "Go 语言程序设计"
	dilbert.Address = "北京清华大学"
	fmt.Println(dilbert)

	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Println(dilbert)

}

func fun06() {

	args1 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(args1)

	args2 := make(map[string]int)
	args2["alice"] = 31
	args2["charlie"] = 34
	fmt.Println(args2)

}

func fun05() {
	var runes []rune
	for _, r := range "Hello,World" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

func fun04() {
	months := [...]string{1: "jan", 2: "fer", 3: "march", 4: "april", 5: "may", 6: "june", 7: "july", 8: "august", 9: "september", 10: "october", 11: "november", 12: "december"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Printf("%T , cap %d \n", Q2, cap(Q2))
	fmt.Printf("%T , cap %d \n", summer, cap(summer))
	fmt.Println(summer)
}

func fun03() {
	arr1 := []int{0, 1, 2, 3, 4, 5}
	reverse(arr1[:])
	fmt.Println(arr1)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
