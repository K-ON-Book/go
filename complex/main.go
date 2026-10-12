package main

import "fmt"

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(2, 3)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}
