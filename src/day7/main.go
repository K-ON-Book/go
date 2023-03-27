package main

import (
	"fmt"
	"math"
)

func main() {
	var b = 10
	var c = 20
	var a *int = &b
	a = &c
	fmt.Println(*a)
	var d float32 = 2.254
	fmt.Printf("%f\t", d)
	fmt.Printf("%g\t", d)
	fmt.Printf("%e\t", d)
	fmt.Println(math.MaxInt8)
	var f float64
	fmt.Scanln(&f)
	m := IntFromFloat64(f)
	fmt.Println(m)
	fmt.Print("=======================\n")
	complex_learn()
	bit_operation_learn()
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}

func complex_learn() {
	var c1 complex64 = 5 + 10i
	fmt.Printf("The value is: %v\n", c1)
	c := complex(5, 10)
	fmt.Println(real(c), imag(c))
}

func bit_operation_learn() {
	var x uint8 = 22           // 10110
	var y uint8 = 4            // 00100
	fmt.Printf("%08b\n", x&^y) // 10010
	var test1 uint8 = 10
	fmt.Println(^test1) //245 不懂！！！
	var test2 int8 = 10
	fmt.Println(^test2) //-11 不懂！！！
	fmt.Println(1 >> 2) //0.25???
	type BitFlag int
	const (
		Active  BitFlag = 1 << iota
		Send            // 1<<1
		Receive         //1<<2
	)
	flag := Active | Send
	fmt.Println(flag)
}
