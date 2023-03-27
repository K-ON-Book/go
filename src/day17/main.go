package main

import "fmt"

func fibonacci(n int) (iterm, res int) {
	if n <= 2 {
		res = 1
	} else {
		_, v1 := fibonacci(n - 1)
		_, v2 := fibonacci(n - 2)
		res = v1 + v2
	}
	iterm = n
	return iterm, res
}
func testfib() {

	for i := 1; i <= 10; i++ {
		iterm, res := fibonacci(i)
		fmt.Println(iterm, res)
	}
}
func RevSign(nr int) int {
	if nr <= 0 {
		return -nr
	}
	return nr
}
func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}
func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}
func printrec_1(i int) {
	if i > 10 {
		return
	}
	printrec_1(i + 1)
	fmt.Printf("%d ", i)
}
func printrec_2(i int) {
	if i < 1 {
		return
	}
	printrec_2(i - 1)
	fmt.Println(i)
}
func test3(n int) (out int) {
	if n < 1 {
		return 1
	}
	return n * test3(n-1)
}
func test4(n int) int {
	if n <= 2 {
		return 1
	}
	return test4(n-1) + test4(n-2)
}
func main() {
	// testfib()
	// fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	// fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// // 17 is odd: is true
	// fmt.Printf("%d is odd: is %t\n", 18, odd(18))
	// out := test1(11)
	// fmt.Println(out)
	// printrec_2(10)
	// test2(1)
	// fmt.Println(test3(12))
	fmt.Println(test4(5))
}
