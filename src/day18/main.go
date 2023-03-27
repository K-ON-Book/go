package main

import (
	"fmt"
	"strings"
)

func main() {
	// callback(1, add)
	// str := "this is an apple!"
	// iterm := strings.Index(str, "is")

	// fmt.Println(iterm)
	ans := strings.IndexFunc("Abc", checkRune)
	fmt.Println(ans) //
	asciiOnly := func(c rune) rune {
		if c > 127 {
			return '?'
		}
		return c
	}
	fmt.Println(strings.Map(asciiOnly, "Jérôme Österreic"))
}
func add(a, b int) {
	sum := a + b
	fmt.Println(a, "+", b, "=", sum)
}
func callback(y int, f func(int, int)) {
	f(y, 2)
}

func checkRune(r rune) bool {
	if r == 'c' {
		return r == 'c'
	}
	return false
}
