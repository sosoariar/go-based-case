package main

import (
	"fmt"
	"math"
	"time"
)

/*
	标准库中方法的常规使用方法
*/
func main() {
	//fun01()
	fun02()
}

func fun01() {
	const day = 24 * time.Hour
	fmt.Println(day.Seconds())
}

func fun02() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q)) // p.Distance 称为选择子(selector)
}

/*
	Go 语言中，接受者不使用特殊名(比如this 或者self),而是开发者自己选择接收的名字,就像其他的参数变量一样
	最常用的方法是取类型名称首字母
*/
type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// p为方法的接收者,
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

// Distance 方法返回路径的长度
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
