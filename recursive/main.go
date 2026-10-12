package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(contEven([]int{1, 2, 3, 4}))

	fmt.Println(myMax([]int{-1, 1, 2, 3, 4}))
}

func contEven(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	rest := contEven(arr[1:])
	if arr[0]%2 == 0 {
		return 1 + rest
	}
	return rest
}

func myMax(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("空数组，无最大值")
	}
	if len(arr) == 1 {
		return arr[0], nil
	}
	rest, err := myMax(arr[1:])
	if err != nil {
		return 0, err
	}
	return max(arr[0], rest), nil
}
