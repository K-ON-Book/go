package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	test1()
	where()
	test2()
	where()
	test3()

}
func test1() {
	fmt.Println("test1")
}
func test2() {
	fmt.Println("test2")
}
func test3() {
	fmt.Println("test3")
}
