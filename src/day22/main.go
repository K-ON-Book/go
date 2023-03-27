package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var n int
	fmt.Scanln(&n)
	res := longCalculation(n)
	fmt.Println("res=", res)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
func longCalculation(a int) (res int) {
	res = 0
	for i := 1; i <= a; i++ {
		res += i
	}
	return res
}
