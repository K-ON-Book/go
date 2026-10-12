package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 <= 2)
	fmt.Println(5%-3, 5%3, -5%-3, -5%3) // 对于go，取模结果的符号只和被取模数的符号相同
	var b1 uint8 = 0b00001011
	var b2 uint8 = 0b0000_1_110
	// &^ 位清除：右边1是位置会让左边同样位置的bit位变成0
	// ^ 一元运算符是取反，作为二元运算符是异或,异或就是两个bit位一样，值就是0，否则1
	fmt.Printf("%08b %08b %08b\n", ^b1, b1^b2, b1&^b2)
	var b3 uint8 = 0b00000101
	// << [n] 左移：数字整体左移n位
	// >> [n] 右移：数字整体右移n位
	fmt.Printf("%08b %08b\n", b3<<8, b3>>2) // 00000000 00000001

	o := 8
	fmt.Printf("%08b %#[1]x %#8[1]b\n", o)

	// | 按位或，比如 1011 | 0100 => 1111
	// & 按位与，比如 1011 & 0100 => 0000
	fmt.Printf("%#08b %#08b\n", 0b1011|0b0100, 0b1011&0b0100)

	// || 逻辑或，&& 逻辑与
	flag1 := 1 != 2 && 5 > 6 // false
	flag2 := 1 != 2 || 5 > 6 // true
	fmt.Println(flag1, flag2)
}
