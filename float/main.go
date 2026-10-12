package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%g\n", 12.34000)
	fmt.Printf("%g\n", 0.000000123)
	fmt.Printf("%g\n", 123456789000000000000.0)

	var z float64
	// +Inf 表示正无穷，-Inf 表示负无穷，NaN表示不是数字
	fmt.Println(z, -z, -1/z, 1/z, z/z) // 0,-0,-Inf,+Inf,NaN

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}
