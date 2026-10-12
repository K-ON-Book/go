package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"unicode/utf8"
)

func main() {
	// dup1()
	// dup2()
	// dup3()
	// testDecodeRuneInString()
}

// 标准输入读取重复的行（第一种方法）
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		// if n > 1 && line != "\n" { //不能过滤回车存入的换行符，因为bufio.Scanner默认按行读取，但是会把换行符去掉，
		// 就会导致每次回车会拿到""，所以真正需要过滤的是""
		if counts[line] > 1 && line != "" {
			fmt.Printf("%d\t%s\n", counts[line], line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "读取输入时发生错误:", err)
		os.Exit(1)
	}
}

// 从文件读取重复的行
func dup2() {
	files := os.Args[1:]
	counts := make(map[string]int)
	countlines := func(f *os.File, counts map[string]int) {
		input := bufio.NewScanner(f)
		for input.Scan() {
			counts[input.Text()]++
		}
		if err := input.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "input err: %v", err)
		}
	}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		}
		countlines(f, counts)
		f.Close()

	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup3() {
	counts := make(map[string]int)
	file := os.Args[1]
	data, _ := os.ReadFile(file)
	content := string(data)

	sep := ""
	switch runtime.GOOS {
	case "linux":
		sep = "\n"
	case "windows":
		sep = "\r\n"
	case "darwin":
		sep = "\n"
	}
	for v := range strings.SplitSeq(content, sep) {
		counts[v]++
	}
	fmt.Println(runtime.GOOS)
	for k, v := range counts {
		if v > 1 {
			fmt.Printf("%s %d\n", k, v)
		}
	}
}

func testDecodeRuneInString() {
	s := "A你"
	res, size := utf8.DecodeRuneInString(s) // 只看第一个字符
	fmt.Println(res, "\t", size)
}
