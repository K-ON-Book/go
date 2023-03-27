package main

import "fmt"

func test1() {
	var arr = [3][5]int{}
	for row := range arr {
		for col := range arr[row] {
			fmt.Print(arr[row][col])
		}
		fmt.Println()
	}
}

func test2() {
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
		fmt.Println(item)
	}
}
func test3() {
	var slice = make([]int, 3)
	var slice1 = new([3]int)[:]
	var slice2 = new([3]int)
	var slice3 = []int{1, 1, 1, 1}[:]
	// var arr = [3]int{1,1,1}[:]
	var arr = []int{1, 1, 1}
	slice4 := arr[:]
	fmt.Printf("%T\t%T\t%T\t%T\t%T", slice, slice1, slice2, slice3, slice4)
}

func sum(arr []float32) (sum float32) {
	for _, v := range arr {
		sum += v
	}
	return
}
func SumAndAverage(slice []int) (int, float32) {
	sum := 0
	average := 0
	for _, v := range slice {
		sum += v
	}
	average = sum / len(slice)
	return sum, float32(average)
}
func MinAndMax(slice []int) (min, max int) {
	for _, v := range slice {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return min, max
}
func test4() {
	slice1 := make([]int, 0, 10)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}
func test5() {
	var slice = []int{1, 1, 1, 1}
	fmt.Println(slice[:4], ",", slice[4:])
}
func test6() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	// slTo := []int{4, 5, 6, 7}
	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n)
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
	var arr = []int{4, 5, 6}
	sl3 = append(sl3, arr...)
	// fmt.Printf("%T\n", sl3)
	fmt.Println(sl3)
}
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
func test7(s []int, factor int) (res []int) {
	var newSlice = make([]int, len(s)*factor)
	copy(newSlice, s)
	res = newSlice
	return
}
func isDouble(i int) bool {
	if i%2 == 0 {
		return true
	} else {
		return false
	}
}
func Filter(slice []int, fn func(int) bool) (newSlice []int) {
	// index := -1
	for _, v := range slice {
		if fn(v) {
			// index++
			// newSlice[index] = v
			// copy(slice,slice)
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
func InsertStringSlice(s1, s2 []int, index int) (res []int) {
	s1 = append(s1, s2[index:]...)
	s2 = append(s2[:index], s1...)
	res = s2
	return
}

func RemoveStringSlice(slice []string, start, end int) []string {
	result := make([]string, len(slice)-(end-start+1))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end+1:])
	return result
}

//	func test8() {
//		var slice1 = []int{1, 1, 2, 2}
//		var slice2 = []int{2, 2}
//		res := copy(slice1, slice2)
//		fmt.Println(res, slice1)
//	}
func main() {
	// test1()
	// test2()
	// test3()
	/* arr := [4]float32{1, 1, 1, 1}
	fmt.Printf("%T\n", arr)
	arr1 := arr[:]
	fmt.Printf("%T\n", arr1)
	fmt.Println(sum(arr[:])) */
	/* slice := make([]int, 5)
	for i := 0; i < len(slice); i++ {
		slice[i] = i + 1
	}
	sum, average := SumAndAverage(slice)
	fmt.Println(sum, average)
	min, max := MinAndMax(slice)
	fmt.Println("min:", min, "max:", max) */
	// test4()
	// test6()
	/* var slice1 = []byte{1, 2, 3}
	var slice2 = []byte{4, 5, 6}
	res := AppendByte(slice1, slice2...)
	fmt.Println(res) */
	/* slice := new([10]int)[:]
	fmt.Printf("%v\t%v", test7(slice, 2), len(test7(slice, 2))) */
	// var s = []int{1, 2, 4, 7, 9}
	// fmt.Println(Filter(s, isDouble))
	// var s1 = []int{1, 3, 5, 7, 9}
	// var s2 = []int{2, 4, 6, 8}
	// fmt.Println(InsertStringSlice(s1, s2, 2))
	s := []string{"M", "N", "O", "P", "Q", "R"}
	fmt.Println(RemoveStringSlice(s, 1, 2))

	// test8()
}
