package main

import (
	"fmt"
	"slices"
	"strings"
)

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

func main() {
	t1()
	t2()
	t3()

	t4()
}

func t1() {
	f := square // 等号右边是值，然后这个值是一个函数，所以叫函数值
	fmt.Println(f(3))
	fmt.Printf("%T\n", f)

	// g := negative
	// fmt.Println(f == g) //  compiler: invalid operation: f == g (func can only be compared to nil) [UndefinedOp]. 函数值是[不可以比较的],但是可以和nil比较
}

func t2() {
	var f func(x, y int) int
	fmt.Println(f == nil) // true, func value 的零值是nil
	f = product
	fmt.Println(f == nil) // false
	fmt.Println(f(3, 4))

	// mmap := map[func(int) int]string{
	// 	square: "square",
	// } //  compiler: invalid map key type func(int) int [IncomparableMapKey]. map的key必须是可比较的
}

func t3() {
	add1 := func(r rune) rune {
		return r + 1
	}
	fmt.Println(strings.Map(add1, "A")) // B
}

func t4() {
	f := func(name string) string {
		return name + "goo"
	}
	fmt.Println(expand("hello foo, nice to meet foo!", f))
}

func expand(s string, f func(target string) string) string {
	target := "foo"
	replacement := f(target)

	replace := func(e string) string {
		start := 0
		l := len(e)
		b := []byte(s)
		for start+l <= len(b) {
			if b[start] == e[0] {
				s1 := b[start : start+l]
				if string(s1) == e {
					b = slices.Replace(b, start, start+l, []byte(replacement)...)
					start = start + len(replacement)
					continue
				}
			}
			start++
		}
		return string(b)
	}

	return replace(target)
}
