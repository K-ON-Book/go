package main

import (
	"day1/b"
	"fmt"
)

func main() {
	fmt.Println(p)
	// fmt.Println(a)
	b.B()
	var m, n int64
	fmt.Println("Please enter two integer numbers:")
	fmt.Scanln(&m, &n)
	m, n = reserve(m, n)
	fmt.Println("absolute value:")
	fmt.Println(m, n)
}

func test3() {
	fmt.Println("aa")
}
func reserve(a, b int64) (int64, int64) {
	var reserve_a, reserve_b int64
	if a > 0 || a == 0 {
		reserve_a = a
	} else {
		reserve_a = -a
	}
	if b > 0 || b == 0 {
		reserve_b = b
	} else {
		reserve_b = -b
	}
	return reserve_a, reserve_b
}

var p float64 = 6.3589
