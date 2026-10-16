package main

import (
	"fmt"
	"os"
)

// 可变参数：表示接收任意数量的该类型参数
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func errorf(linenum int, format string, args ...interface{}) { // 类型 interface{} 等价于 any, 可以通过 `go doc -u builtin.any` 查询到, 推荐使用any
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func main() {
	fmt.Println(sum(1, 2, 5))
	s := []int{1, 2, 5}
	fmt.Println(sum(s...))

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)
}
