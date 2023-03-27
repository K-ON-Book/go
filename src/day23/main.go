package main

import "fmt"

func test1() {
	var arr1 [5]int
	for i := 0; i <= len(arr1)-1; i++ {
		arr1[i] = i
	}
	for iterm, value := range arr1 {
		fmt.Println("Array at index", iterm, "is", value)
		fmt.Println("test1")
	}
}
func test2() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
}
func test3() {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}
func test4_1(a [3]int) {
	fmt.Println(a)
}
func test4_2(a *[3]int) {
	a[0] = 1
	fmt.Println(*a)
}
func test5() {
	var arr1 [3]int
	arr2 := arr1
	fmt.Println(arr2)
	arr2[0] = 1
	fmt.Println(arr1)
}
func test6() {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrLazy = []int{5, 6, 7, 8, 22}	//注：初始化得到的实际上是切片slice
	var arrKeyValue = [5]string{1: "aaa", 3: "Chris", 4: "Ron"}
	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}	//注：初始化得到的实际上是切片slice
	for i := 0; i < len(arrAge); i++ {
		fmt.Printf("Person at %d is %v\n", i, arrAge[i])
	}

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
}
func test7() {
	type Vec [3]string
	var arr = Vec{"a", "b", "c"}
	fmt.Println(arr)
}
func test8() {
	var arr = [3][5]int{{}, {}, {}}
	fmt.Println(arr)
}
func test9() {
	const (
		WIDTH  = 1920
		HEIGHT = 1080
	)
	type pixel int
	var screen [WIDTH][HEIGHT]pixel
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
}
func test10(f func(a *[3]float64) (sum float64)) {
	arr := [3]float64{7.0, 8.5, 9.1}
	x := f(&arr)
	fmt.Printf("The sum of the array is: %f", x)
}
func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}
func main() {
	// test1()
	// test3()
	// var ar [3]int
	// test4_1(ar)
	// test4_2(&ar)
	// test5()
	// test6()
	// test7()
	// test8()
	// test9()
	test10(Sum)
}
