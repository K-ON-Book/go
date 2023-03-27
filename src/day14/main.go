package main

import (
	"fmt"
)

func test1() {
	str1 := "Go is a beautiful language!"
	fmt.Println("The length of str is:", len(str1))
	for ix := 0; ix < len(str1); ix++ {
		fmt.Printf("%d,%c\n", ix, str1[ix])
	}
	str2 := "日本語"
	fmt.Println("The length of str is:", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Println(ix, "--", str2[ix])
	}
}
func test2_1() {
	var count = 0
	for a := 0; a < 15; a++ {
		count++
		if count == 15 {
			fmt.Println(a)
		}
	}
}
func test2_2() {
	i := 0
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	if i < 15 {
		goto START
	}
}
func test3_1() {
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
}
func test3_2() {
	str1 := ""
	for i := 1; i <= 25; i++ {
		str1 = str1 + "G"
		fmt.Printf("%v\n", str1)
	}
}
func test4() {
	for i := 0; i < 11; i++ {
		fmt.Printf("%v", ^i)
	}
}
func test5() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("%s-", "FizzBuzz")
		case i%3 == 0 && i%5 != 0:
			fmt.Printf("%s-", "Buzz")
		case i%3 != 0 && i%5 == 0:
			fmt.Printf("%s-", "Buzz")
		default:
			fmt.Printf("%d-", i)
		}
	}
}
func test6() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 20; j++ {
			fmt.Printf("%c", '*')
		}
		fmt.Println()
	}
}
func test7() {
	var i int = 5
	for i >= 0 {
		i--
		fmt.Printf("The variable i is now: %d\n", i)
	}
}
func test8() {
	str1 := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d", len(str1))
	for key, char := range str1 {
		fmt.Printf("Character on position %d is: %c \n", key, char)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for key, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, key)
	}
	fmt.Println()
	fmt.Println("index\tint(rune)\trune\tchar\tbytes")
	for index, rune := range str2 {
		fmt.Printf("%d\t%d\t\t%U\t'%c'\t%X\n", index, rune, rune, rune, []byte(string(rune)))
	}
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("%d\t%c\n ", ix, str2[ix])
	}
}
func test9() {
	// for i := 0; i < 3; {
	// 	fmt.Println("Value of i:", i)
	// }
	// s := ""
	// for s != "aaaaa" {
	// 	fmt.Println("Value of s:", s)
	// 	s = s + "a"
	// }
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
		/*
		   0 5 a
		   1 6 aa
		   2  7 aaa
		*/
	}

}
func test10() {
	// var i int = 5
	// for {
	// 	i = i - 1
	// 	if i < 0 {
	// 		break
	// 	}
	// }
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			print(j)
		}
		print("  ")
	}
}
func test11() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		print(i)
		print(" ")
	}
} // 0 1 2 3 4 6 7 8 9
func main() {
	// test1()
	// test2_1()
	// test3_1()
	// test3_2()
	// test4()
	// test5()
	// test6()
	// test7()
	// test8()
	// test9()
	// test10()
	test11()

}
