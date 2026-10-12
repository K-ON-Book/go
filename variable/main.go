package main

import (
	"fmt"
)

type s struct {
	a int
	b string
}

func main() {
	// test1()
	// test2()
	v := 1
	incr(&v)
	fmt.Println(f() == f())
	test3()
	fmt.Println(*global)

	for _, v := range "aaa" {
		fmt.Println(v)
	}
}

func test1() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}

func test2() {
	p := new(int)
	q := new(int)
	fmt.Println(p, "\t", q) // 地址是不一样的
}

func incr(p *int) int {
	*p++
	return *p
}

func f() *int {
	v := 1
	return &v
}

var global *int

func test3() {
	i := 2
	global = &i
}

func delta(old, new int) int {
	// a := new(a) // new是预定义函数，已被用为变量名，所以无法使用new(T); https://go.dev/ref/spec#Predeclared_identifiers
	return new - old
}
