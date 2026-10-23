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

type (
	myt1 *int
	myt2 interface{}
	myt3 any
)

// receiver 的底层类型不能是指针或者interface
// 宪法：https://go.dev/ref/spec#Method_declarations "Its type must be a defined type T....."
// 不能是指针：
// 不能是interface：interface是方法的集合，再给interface绑定方法就语义混乱
// /compiler: invalid receiver type myt1 (pointer or interface type) [InvalidRecv]
// func (m myt1) print() { fmt.Println() }
//
// func (m myt2) print() {}
//
// func (m myt3) print() {}

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

	var s1 mySlice = []int{1, 2, 3}
	var s2 *mySlice
	fmt.Println((&s1).sum()) // 功能上等价于写 s1.sum()，因为会自动加&,自动加的前提是s1是可寻址的
	// mySlice{1,2,3}).sum()// 报错，字面量是临时值，无法寻址
	(&mySlice{1, 2, 3}).sum()      // 修复后不报错， 主动加&，会拿到字面量的地址符合receiver类型，值得注意的是不加小括号的话，因为调用方法的优先级高于取地址，所以会报错
	fmt.Printf("%p\n", &[1]int{1}) // 可以打印出来这个临时地址,复合字面量可以取地址，但是比如&1就不行. Doc: https://go.dev/ref/spec#Address_operators "As an exception to the addressability requirement, x may also be a (possibly parenthesized) composite literal..."
	// 反思：为什么go不也自动对字面量加取地址符？TODO:
	fmt.Println(s2.sum()) // s2 零值是nil，nil也能调用方法，但是一定要注意方法最好要对receiver值是nil的情况处理
}
