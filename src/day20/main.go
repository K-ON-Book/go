package main

import "fmt"

func test1() {
	for i := 0; i < 4; i++ {
		g := func(i int) {
			fmt.Printf("%d", i)
		}
		g(i)
		fmt.Printf("- g is of type %T and has value %v\n", g, g)
	}
}
func test2() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func fib() func() int {
	a, b := 1, 1
	var tmp int
	return func() int {
		/*
		   a
		   b
		*/
		// a, b = b, a+b
		tmp = a
		a = b
		b = tmp + b
		return b
	}
}

//工厂函数

func main() {
	// test1()
	// var f = test2()
	// fmt.Print(f(1), "-")
	// fmt.Print(f(10), "-") // 11
	// fmt.Print(f(20), "-") // 31
	f := fib()
	println(f()) // tmp=1 a=1 b=1; tmp+b
	println(f()) // tmp=1 a=2  b=2;tmp+b
	println(f()) // tmp=2 a=3 b=3;tmp+b
}
