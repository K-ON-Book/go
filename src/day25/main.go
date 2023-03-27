package main

import (
	// "bytes"
	"bytes"
	"fmt"
)

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
func callback(f func([]int) int) {
	var arr = [5]int{0, 1, 2, 3, 4}
	res := f(arr[:])
	fmt.Println(res)
}
func test() {
	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v", arr1)
	fmt.Printf("%v", arr2)
	fmt.Printf("%v", arr3)
}

func test1() {
	var s1 = []int{1, 1, 1, 1, 1} //相关数组[5]{1,1,1,1,1}
	fmt.Println(s1)
	var s2 []int = []int{1, 2, 3} //等效于上
	fmt.Println(s2)
	//相关数组还没有定义时，我们可以使用 make() 函数来创建一个切片
	var slice1 = make([]int, 5)
	fmt.Printf("len:%d cap:%d\n", len(slice1), cap(slice1))
	var slice2 []int = make([]int, 5) //等效于上
	fmt.Printf("len:%d cap:%d\n", len(slice2), cap(slice2))
	var slice3 = make([]int, 2, 5)
	fmt.Printf("len:%d cap:%d\n", len(slice3), cap(slice3))
}
func test2() {
	var slice1 = make([]int, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}
	for k, v := range slice1 {
		fmt.Printf("Slice at %d is %d\n", k, v)
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}

func test3() {
	var slice1 = make([]int, 10, 50)
	slice2 := new([50]int)[0:10]
	fmt.Printf("%v\t%d\t%d\n", &slice1, len(slice1), cap(slice1))
	fmt.Printf("%v\t%d\t%d\n", &slice2, len(slice2), cap(slice2))
	slice3 := new([]int)
	fmt.Println(&slice3)
}
func test4() {
	var p1 *[]int = new([]int)
	// p := new([]int)
	fmt.Println(p1)
	p2 := make([]int, 0)
	fmt.Println(p2)
}

// 切片、数组还需要继续学习
func main() {
	// callback(sum)
	// test1()
	// test2()
	test3()
	// test4()
}
