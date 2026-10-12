package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }

	// s, sep := "", ""
	// for _, v := range os.Args[1:] {
	// 	s += sep + v
	// 	sep = " "
	// }
	// fmt.Println(s)

	// fmt.Println(strings.Join(os.Args[1:], " "))

	// fmt.Println(os.Args[1:])

	// funtest1()
	// funtest2()
	// funtest3()
	a, b := lengthOfLongestSubstring("1R1T7")
	fmt.Printf("%v,%c", a, b)
}

func lengthOfLongestSubstring(s string) (int, []rune) {
	block := []rune{}
	var maxLen int

	mmap := make(map[rune]int)

	for k, v := range s {
		if _, ok := mmap[v]; ok {
			block = []rune{}
			mmap = map[rune]int{}
			s = s[k:]
			continue
		}
		mmap[v] = k
		block = append(block, v)
		maxLen = max(len(block), maxLen)
	}
	return maxLen, block
}

func funtest1() {
	fmt.Println(os.Args[0])
}

func funtest2() {
	for i, arg := range os.Args {
		fmt.Println(i, " ", arg)
	}
}

func funtest3() {
	// 构造大量args
	args := []string{}
	for i := 1; ; i++ {
		args = append(args, strconv.Itoa(i))
		if i == 1000 {
			break
		}
	}

	start := time.Now()
	strings.Join(args, " ")
	end := time.Now()
	sub1 := end.Sub(start)
	fmt.Println("strings.Join take time： ", sub1)

	s, sep := "", ""
	start1 := time.Now()
	for _, v := range args[1:] {
		s += sep + v
		sep = ""
	}
	end1 := time.Now()
	sub2 := end1.Sub(start1)
	fmt.Println("+= take time: ", sub2)
}
