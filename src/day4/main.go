package main

import (
	"fmt"
)

func main() {
	var a int64 = 954
	var b float64 = 55.65446878
	var c float64 = 64.999999999999999999999999
	var d int64 = 6
	// c = float64(d)
	d = int64(c)
	// fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Println(int64(b), float64(a))

	type IZ int64
	var e IZ = 5
	f := int64(e)
	fmt.Println(f)
	// var 3_a = 1
	// fmt.Println(3_a)
	var _a = 1
	fmt.Println(_a)
}
