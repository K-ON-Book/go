package main

import (
	"fmt"
	"math"
	"os"
)

func mult_returnval_1(a, b int) (Sum, Product, Difference int) {
	Sum = a + b
	Product = a * b
	Difference = a - b
	return Sum, Product, Difference
}
func mult_returnval_2(a, b int) (int, int, int) {
	Sum := a + b
	Product := a * b
	Difference := a - b
	return Sum, Product, Difference
}
func MySqrt_1(f float64) (v float64) {
	v = math.Sqrt(f)
	return
}
func MySqrt_2(f float64) float64 {
	v := math.Sqrt(f)
	return v
}
func minmax(a, b int) (min, max int) {
	if a > b {
		min = b
		max = a
	} else {
		min = a
		max = b
	}
	return
}

func ptr(ptra *int) int{
	*ptra = 10
  return *ptra
}

func task1_get_max(s ...int) int {
	if len(s) == 0 {
		os.Exit(1)
	}
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}
func varargs(s ...int) {
	for _, value := range s {
		fmt.Println(value)
	}
}
func test_c() {
	sourceArr := []int{0, 1, 2, 3}
	targetArr := []*int{}

	for _, num := range sourceArr {
		targetArr = append(targetArr, &num)
	}

	for _, num_p := range targetArr {
		fmt.Printf("num: %v\n", *num_p)
	}
}
func main() {
	// var m, n int
	// fmt.Scanln(&m, &n)
	// a1, b1, c1 := mult_returnval_1(m, n)
	// fmt.Println(a1, b1, c1)
	// a2, b2, c2 := mult_returnval_2(m, n)
	// fmt.Println(a2, b2, c2)
	// var f float64
	// fmt.Scanln(&f)
	// fmt.Println(MySqrt_1(f))
	// fmt.Println(MySqrt_2(f))
	// fmt.Println(minmax(m, n))
	// var a int = 5
	// ptr(&a)
	// fmt.Println(a) // 10
	// x := task1_get_max(1, 3, 2, 0)
	// fmt.Printf("The minimum is: %d\n", x)
	// slice := []int{7, 9, 3, 5, 1}
	// x := task1_get_max(slice...)
	// fmt.Println(x)
	slice := []int{7, 9, 3, 5, 1}
	varargs(slice...)
	// for _, v := range slice {
	// 	fmt.Println()
	// }
	// test_c()
}
