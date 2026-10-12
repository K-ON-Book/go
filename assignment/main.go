package main

import "fmt"

func main() {
	v := 1
	v++ // 不是表达式，是语句
	v--
	for i := range 9 {
		fmt.Println(fib(i + 1))
	}

	myMap := map[string]int{
		"1": 1,
		"3": 3,
	}
	var x any = 1
	c, ok := x.(string)
	fmt.Printf("%q %v\n", c, ok)
	v, ok = myMap["2"]
	fmt.Println(v, "\t", ok)
}

func fib(n int) int {
	if n < 3 {
		return n - 1
	}
	return fib(n-2) + fib(n-1)
}
