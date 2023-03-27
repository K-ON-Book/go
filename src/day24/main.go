package main

import "fmt"

func test1() {
	arr := []byte{1, 2, 3}
	// var slice1 = arr[0:3] // 左闭右开
	fmt.Println(fmt.Println(arr))
}
func test2() {
	arr := [3]byte{}
	slice1 := &arr
	fmt.Println(&arr)
	fmt.Println(slice1)
}
func test3() {
	var slice1 []int
	arr1 := [3]int{1, 2, 3}
	slice1 = arr1[:]
	fmt.Printf("%v-%T\t%v%T", arr1, arr1, slice1, slice1)
}
func test4() {
	// slice1 := [3]int{1, 2, 3}[:]
	arr1 := [3]int{1, 2, 3}
	slice1 := arr1[:]
	fmt.Println(slice1)
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)
}
func test5() {
	s := []string{"a", "b", "c"}
	// fmt.Println(s[0:3])
	fmt.Println(s[3:])
	// s == s[:3]+s[3:]
	// fmt.Println(s)
}
func test6() {
	s1 := []int8{1, 2, 3}
	s2 := s1[:]
	fmt.Println(s2)
}
func test7() {
	s1 := []int8{1, 2, 3}
	fmt.Println("len", len(s1), "cap", cap(s1))
	x := []int{2, 3, 5, 7, 11}
	y := x[1:3]
	fmt.Println("len", len(x), "cap", cap(x))
	fmt.Println("len", len(y), "cap", cap(y))
}
func test8() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	for k := range slice1 {
		fmt.Println(k, "\t", slice1[k])
	}
	// fmt.Printf("The length of arr1 is %d\n", len(arr1))
	// fmt.Printf("The length of slice1 is %d\n", len(slice1))
	// fmt.Printf("The capacity of slice1 is %d\n", cap(slice1)) //4
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice1 at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))   //4
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1)) //4
}
func test9() {
	var x = []int{1, 1, 1, 1} //相关数组[4]int{1,1,1,1}
	fmt.Println(&x)
	for i := 0; i < len(x); i++ {
		fmt.Println(&x[i], "\t", x[i])
	}
	for k, v := range x {
		fmt.Println(&x[k], "\t", v)
	}
	var arr = [4]int{2, 2, 2, 2}
	var y = arr[:] //相关数组[4]int{1,1,1,1}
	for k, v := range arr {
		fmt.Println(&y[k], "\t", v)
	}
	for k, v := range y {
		fmt.Println(&y[k], "\t", v)
	}
}
func test10() {
	var s2 = []int{1, 1, 1, 1}
	s2 = s2[1:]
	// s2 = s2[-1:]
	fmt.Println(s2)
}
func test11() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Printf("%c\t%c\t%c\t%c", b[1:4], b[:2], b[2:], b[:])
	//out: [o,l,a] [g,o] [l.a.n.g] [g,o,l,a,n,g]
}
func main() {
	// test1()
	// test2()
	// test3()
	// test5()
	// test7()
	// test8()
	// test9()
	// test10()
	test11()
}
