package main

import "fmt"

const Pi = -3.14
const hardEight = (1 << 100) >> 97
const (
	a, b, c, d, e, f = 1, 2, 3, 4, 5, 6
)
const (
	g, h = iota, iota
	i    = iota // 1
	j    = iota // 2
	k    = 5
	l    = iota // 4
)
const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)
const (
	Open = 1 << iota // 0001
	Close
	Pending
)

func main() {
	fmt.Print(Pi)
	fmt.Println(hardEight)
	fmt.Println(a, b, c, d)
	fmt.Println(g, h)
	fmt.Println(i, j)
	fmt.Println(k, l)
	fmt.Print(Cherimoya, Durian, Elderberry, Fig, "\n")
	fmt.Println(Open, Close, Pending)
}
