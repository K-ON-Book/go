package main

import "math"

type myPoint struct {
	x, y float64
}

func (m *myPoint) ScaleBy(factor float64) {
	m.x *= factor
	m.y *= factor
}

func (m *myPoint) Distance(dot myPoint) float64 {
	return math.Hypot(m.x-dot.x, m.y-dot.y)
}

// 约定如果一个receiver的类型是指针类型那么即使其它方法不需要指针receiver也要统一使用指针receiver

type mySlice []int

// receiver值是nil也能调用方法，但是最好要对值是nil的情况处理
func (m *mySlice) sum() (sum int) {
	if m == nil {
		return 0
	}
	for _, v := range *m {
		sum += v
	}
	return
}
