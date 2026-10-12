package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	// 0 1 2 3 s[0:]
	// 1 2 3 s[1:]
	//

	s1 := []int{1, 2, 3, 4}
	fmt.Println(remove(s1, 1))
	reserve(s1)
	fmt.Println(s1)

	a1 := [3]int{1, 2, 3}
	reserveA(&a1)
	fmt.Println(a1)
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func reserveA(arr *[3]int) {
	left, right := 0, len(arr)-1
	for left <= right {
		fmt.Println(arr[left])
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func reserve(slice []int) {
	left, right := 0, len(slice)-1
	for left <= right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
}
