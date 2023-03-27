package main

import (
	"fmt"
	"unicode/utf8"
)

func test1() {
	str := "abcdefghijk"
	fmt.Printf("%c", str[0])
	fmt.Printf("%c", []byte(str))
}
func test2() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c\n", i, c)
	}
	fmt.Println(s)
}
func test3() {
	s1 := "にほんご你好abc"
	fmt.Println(len(s1), utf8.RuneCountInString(s1))
	fmt.Println(len([]byte(s1)), len([]int32(s1)))
	var s2 = "def"
	var b1 = []byte{'a'}
	b1 = append(b1, s2...)
	fmt.Printf("%c", b1)
	var b2 []byte
	b2 = append(b2, s2...)
	fmt.Printf("%c", b2)
}
func test4() {
	s := "aaaa"
	// s[i] = 'b' // 字符串是不可以修改的
	fmt.Println(s)
}
func test5() {
	s := "aaaa"
	c := []byte(s) // 需要将字符串转换为字节数组，字符串本质上来说是字节数组
	c[0] = 'b'
	fmt.Println(string(c)) // baaa
}
func test6() {
	c := []byte{'c', 1, 1, 1}
	fmt.Println(string(c))
	/* slice := []int{1, 1, 1, 1} // 改为byte即可
	fmt.Println(string(slice)) */
}
func test7() {

}
func main() {

	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()

}
