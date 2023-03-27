package main

import (
	"fmt"
)

func test1() {
	fv := func() {
		fmt.Println("hello world")
	}
	fv()
	fmt.Printf("%T", fv)
}

func test2() (res int) {
	defer func() {
		res++
	}()
	return 1
}

func test3(i int) (ret int) {
	defer fmt.Println(i)
	return
}
func test() int { //这里返回值没有命名
	var i int
	defer func() {
		i++
		fmt.Println("defer1", i) //作为闭包引用的话，则会在defer函数执行时根据整个上下文确定当前的值。i=2
	}()
	defer func() {
		i++
		fmt.Println("defer2", i) //作为闭包引用的话，则会在defer函数执行时根据整个上下文确定当前的值。i=1
	}()
	return i
}
func test_() (i int) { //返回值命名i
	defer func() {
		i++
		fmt.Println("defer1", i)
	}()
	defer func() {
		i++
		fmt.Println("defer2", i)
	}()
	return
}

// 匿名函数作为函数返回值
func test4_1() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func test4_2(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
func test5() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
func main() {
	// func() {
	// 	fmt.Println("test")
	// }()
	// func(x, y int) int { return x + y }(3, 4)
	// test1()
	// fmt.Println(test2()) //1
	// fmt.Println(test3(3))
	// fmt.Println("return:", test())
	// fmt.Println("return:", test_())
	// p1 := test4_1()
	// fmt.Println(p1(2)) //4
	// p2 := test4_2(2)
	// fmt.Println(p2(3)) //5
	// f := test5()
	// v_1 := f(1)
	// fmt.Println(v_1) //1
	// v_2 := f(10)
	// fmt.Println(v_2) //11
	var g int
	func(i int) {
		s := 0
		for j := 0; j < i; j++ {
			s += j
		}
		g = s
		fmt.Println(g)
	}(1000)
}
