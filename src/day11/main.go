package main

import (
	"fmt"
	"math"
	"os"
	_ "os"
	"strconv"
)

func main() {
	var orig string = "ABC"
	// var an int
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Println(orig, "不能被转换为整数")
		// return
		// os.Exit(1)
	}
	fmt.Println(an + 2) //an == 0
	// fmt.Println(test())
	// test1()
	test2()
	var arr = [10]int{1, 2, 3, 0, 1, 4, 5, 7, 0, 0}
	fmt.Println((arr))
}

func test() error {
	var orig string = "ABC"
	// var an int
	_, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Println(orig, "不能被转换为整数")
		return err
	}
	return err
}

func test1() {
	var orig string = "567"
	// var an int
	var newS string
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// anInt, err = strconv.Atoi(origStr)
	if an, err := strconv.Atoi(orig); err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	} else {
		fmt.Printf("The integer is %d\n", an)
		an = an + 5
		newS = strconv.Itoa(an)
		fmt.Printf("The new string is: %s\n", newS)
	}
}

func test2() {
	t, ok := test2_1(25.0)
	if ok {
		fmt.Println(t)
	} else {
		os.Exit(1)
	}
}
func test2_1(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	}
	return math.Sqrt(f), true
}

func test3() {
	type binOP func(int) int
}

func Pop(st []int) int {
	v := 0
	for ix := len(st) - 1; ix >= 0; ix-- {
		if v = st[ix]; v != 0 {
			st[ix] = 0
			return v
		}
	}
	return 1
}
