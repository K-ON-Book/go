package main

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	const GoUsage = `Go is a tool for managing Go source code.

Usage:
    go command [arguments]
...`
	fmt.Println(GoUsage)

	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		ru, si := utf8.DecodeRuneInString(s[i:]) // ru 不是rune么，那么si不就是4字节？其实不是的，rune确实是4字节，但是si表示的是这个字符在utf-8规则编码后，在内存中占用的[实际字节数]
		fmt.Printf("%d\t%c\n", i, ru)
		i += si
	}
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	var r rune = '\uFFFD'
	fmt.Printf("%c\n", r)

	fmt.Println(intsToString([]int{1, 2, 3}))
	fmt.Println(comma("12345"))
	intToString()
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	fmt.Println(buf)
	return buf.String()
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	first := n % 3
	if first == 0 {
		first = 3
	}
	buf.WriteString(s[:first])
	for i := first; i < n; i = i + 3 {
		buf.WriteString(", ")
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

func intToString() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2))

	a, _ := strconv.Atoi("789")
	b, _ := strconv.ParseInt("-789", 10, 16)
	fmt.Println(a, b)
	// fmt.Scanln(&x)
	// fmt.Scan(&x)
	fmt.Scanf("x=%d", &x)
	fmt.Scanf("x=%d", &x)
	fmt.Println(x)
}

/*
utf-8的编码规则:
1byte 0xxxxxxx
2byte 110xxxxx 10xxxxxx
3byte 1110xxxx 10xxxxxx 10xxxxxx
4byte 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx

通过前面的第一字节就知道在哪截断是一个完整的字符了
*/

func init() {
	c := "世"
	fmt.Printf("%b, 总计byte: %d\n", []byte(c), len(c)) // [11100100 10111000 10010110], 总计byte: 3
}
