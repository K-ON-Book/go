package main

import (
	"fmt"
)

func main() {
	b := false
	fmt.Printf("%v %v", b, btoi(b))
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
