package main

import (
	"fmt"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"

	var s string
	fmt.Println(s)

	var i, j, k int
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	var b, f, l = true, 2.3, "four"
	fmt.Println(b)
	fmt.Println(f)
	fmt.Println(l)

	fmt.Println(&b)
	fmt.Println(&f)

	x := &f
	fmt.Println(*x)

	fmt.Println("==========================")
	var m, n int
	fmt.Println(&m == &n)
	fmt.Println(&m == nil)

	fmt.Println("==========================")
	v := 1
	incr(&v)
	fmt.Println(incr(&v))

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func incr(p *int) int {
	*p++
	return *p
}
