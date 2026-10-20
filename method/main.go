package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Distance 计算两个坐标之间的距离
// 包级别的函数
func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

// Distance 计算两个坐标之间的距离
// 参数 p 做方法的接收器（receiver）
// 类型 Point 的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := Point{1, 1}
	q := Point{4, 5}

	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // 虽然Point也是同名方法，但是因为前面的receiver不同所以不会命名冲突
}
