package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type Point struct {
	X, Y float64
}

type ColorPoint struct {
	Point // 这叫匿名字段(成员)，不是匿名结构体，匿名结构体就是没有"type xxx struct" 声明的
	Color color.RGBA
}

type ColorPointMemberIsPoint struct {
	*Point
	Color color.RGBA
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

// Q: 假如把 * 去掉，外部的 Point 会被修改成功吗？
// A: 不会！
// 1. 不带 * (值接收器)：发生结构体值拷贝，方法内修改的仅仅是栈上的副本，原件纹丝不动。
// 2. 带 * (指针接收器)：虽然指针本身也被复制了一份，但两个指针指向同一个内存地址，因此能原地修改成功
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	var cp ColorPoint
	cp.X = 1
	cp.Point.Y = 2
	fmt.Printf("%+v\n", cp)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColorPoint{Point{1, 2}, red}
	q := ColorPoint{Point{3, 6}, blue}
	p.ScaleBy(2) // 能调用内嵌的结构体的方法,因为内嵌字段会指导编译器生成额外的包装方法，如下：
	/*
	   func (p *ColoredPoint) ScaleBy(factor float64) {
	       p.Point.ScaleBy(factor)
	   }
	*/
	q.ScaleBy(2) // 等价于(&q).ScaleBy(2)
	fmt.Printf("%+v %+v\n", p, q)

	// p.Distance(q) // 报错，发现q参数不能传入，因为类型不匹配，这里没用发生上述争对receiver类似的对参数的处理，所以需要显示指定传入正确的成员
	fmt.Println(p.Distance(q.Point))

	m := ColorPointMemberIsPoint{
		&Point{1, 2},
		red,
	}
	n := ColorPointMemberIsPoint{
		&Point{3, 5},
		red,
	}
	n.Point = m.Point
	m.Point.X = 10
	fmt.Printf("%+v %+v\n", n.Point, m.Point)
	n.Point.Y = 50
	fmt.Printf("%+v %+v\n", n.Point, m.Point)
}

// 匿名结构体不能定义方法,不过可以通过下面这种内嵌(匿名)成员的方法进行方法的调用
// func (s struct{}) anythingMethod() {
// }

func t1() {
	anonymous := struct {
		sync.Mutex
		mapping map[string]string
	}{
		mapping: make(map[string]string),
	}
	anonymous.Lock() // 匿名结构体拥有匿名成员的方法
}
