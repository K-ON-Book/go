package main

import (
	"fmt"
	"os"
)

func test1() {
LABLE1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABLE1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
func test2() {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		os.Exit(1)
	}
	goto HERE
}
func test3() {
	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")
}
func main() {

	// test1()
	// test2()
	test3()
}
