package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)

func main() {
	flag.Parse() // 必须在使用前调用，用于更新flag的默认值
	fmt.Print(strings.Join(flag.Args(), *s))
	if !*n {
		fmt.Println()
	}
}
