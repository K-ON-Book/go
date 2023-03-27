package main

import (
	"fmt"
	"io"
	"log"
)

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}

func test1() {
	i := 0
	defer fmt.Println(i)
	i++
}
func test2() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	} // output: 4 3 2 1 0
}

/*
	func trace(s string) {
		fmt.Println("entering", s)
	}

	func untrace(s string) {
		fmt.Println("leavingf", s)
	}

	func a() {
		trace("a")
		defer untrace("a")
		fmt.Println("in a")
	}

	func b() {
		trace("b")
		defer untrace("b")
		fmt.Println("in b")
		a()
	}
*/
func trace(s string) string {
	fmt.Println("entering", s)
	return s
}
func un(s string) {
	fmt.Println("leaving", s)
}
func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}
func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}
func test4(s string) (n int, err error) {
	defer func() {
		log.Printf("func(%q) = %d,%v", s, n, err)
	}()
	return 7, io.EOF
}
func main() {
	// function1()
	test1()
	// test2()
	// b()
	// test4("Go")
}
