package main

import (
	"day3/pa"
	"fmt"
)

func main() {

	var m, a, b int64
	fmt.Scanln(&m)
	a, b = test(m)
	fmt.Println(a, b)
	pa.TestA()
	pa.Testa_1()
	fmt.Println(pa.P)
}
func test(a int64) (int64, int64) {
	var m = a - 1
	var n = a + 1
	return m, n
}
