package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(prompt)
	var a int
	fmt.Scanln(&a)
	var str string = "我是我是黑而黑"
	fmt.Println(utf8.RuneCountInString(str))
	b := 5
	if b := 1; b > 0 {
		a = 6
		println(b)
	}
	fmt.Println(b)
	if_else()
}

func if_else() {
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Println("a,b的绝对值:")
	if a < 0 && a*b < 0 {
		reserve_a := -a
		reserve_b := b
		fmt.Println(reserve_a, reserve_b)
	} else if a > 0 && a*b < 0 {
		reserve_a := a
		reserve_b := -b
		fmt.Println(reserve_a, reserve_b)
	} else if a*b == 0 {
		if a == 0 {
			reserve_a := a
			reserve_b := -b
			fmt.Println(reserve_a, reserve_b)
		} else {
			reserve_a := -a
			reserve_b := b
			fmt.Println(reserve_a, reserve_b)
		}
	}
}

func test() {
	var first int = 10
	var cond int

	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {
		fmt.Printf("first is between 0 and 5\n")
	} else {
		fmt.Printf("first is 5 or greater\n")
	}
	if cond = 5; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	} else {
		fmt.Printf("cond is not greater than 10\n")
	}
}
