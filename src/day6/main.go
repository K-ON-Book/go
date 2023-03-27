package main

import (
	"fmt"
	"os"
	"runtime"
)

var a int
var b float32

// c :=5 //不能这样
func test(b *int) {
	print("================\n")
	print(&b, "\n")
	// b_1 := 3
	// b = &b_1
	*b = 3
	print(*b, "\n")
}
func test1(test2 int) {
	print("=========")
	print(test2)
}
func main() {
	// var test2 int = 2
	print(&a, "\n")
	var a = 0
	print(&a, "\n")
	// a:=1 // 已经声明过了a变量,报错
	a, e := 1, 2 //多变量的赋值至少有一个是新变量
	print(&a, "\t", a, "\n")
	print(e, "\n")
	print(&b, "\n")
	var b = 6
	test(&a)
	print("a=", a, "\n")
	// _ = test(&a)
	// test1(test2)
	print(&b, "\n")
	{
		var b = 50
		fmt.Print(b, "\n")
	}
	print(b, "\n") // 6
	var d *int
	d = &a
	fmt.Println(a, b) //1 6
	c := 5
	fmt.Println(c)
	var goos = runtime.GOOS
	fmt.Println("The operating system is: %s", goos)
	path := os.Getenv("PATH")
	fmt.Println("Path is %s", path)
	fmt.Println(*d)
	fmt.Print("\U00002002\U0000F7AA")
}
