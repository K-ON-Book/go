package main

import (
	"fmt"
)

func test1() {
	switch 4 {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
		fallthrough
	case 3:
		fmt.Println(3)
	default:
		{
			fmt.Println(4)
			fmt.Println(4)
			fmt.Println(4)
		}
	}
}
func test2() {
	switch {
	case 1 > 2:
		fmt.Println("madness")
	}
}
func test3() {
	var a = 2
	var b = 3
	a, b = b, a
	fmt.Println(a, b)
	switch a := 2; {
	case a > 0:
		fmt.Println(a)
	}
}
func test4() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
func main() {
	// test1()
	// test2()
	// test3()
	test4()
}
