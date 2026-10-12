package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"

	"struct/pkgs"
)

type Employee struct {
	ID        int
	Name      string
	DoB       time.Time
	Address   string
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee
	dilbert.Salary -= 5000
	position := &dilbert.Position // 这是对成员取地址，也就是&(dilbert.Position)，要和(&dilbert).Position 区分开, 指针.成员这里得到的值类型是string，而不是*string, 之前错误的认识："指针访问成员，那么得到的成员也是指针类型的"。go有个自动解引用的机制（“解引用”就是从“指针地址”到达“真实变量”的那一瞬间的动作），(&dilbert).Position 等价于 (*(&dilbert)).Position; 什么情况下会自动解引用？TODO:
	// var jone *Employee
	// fmt.Println(jone.Position) // 相当于(*jone).Position, 先对jone解引用，可是目前只是声明了jone没有分配内存，所以无法对nil进行解引用，运行会报错. 所以解决办法就是对jone初始化，写法：var jone &Employee{} 或者 var jone = new(Employee);new和make区别？见底部解释

	*position = "Senior " + *position
	fmt.Println(*position)

	// employeeByID1(1).Position="" // 是不可以的，因为返回值是一个值类型，不能放在等号左边;等号左边的需要是一个【可寻址】的变量
	/* 这样写可以
	e := employeeByID1(1)
	e.Position = ""
	*/
	employeeByID2(1).Salary = 500 // 正确的

	t1()
	t2()
	t3()
	t4()
	t5()
	t7()
}

func employeeByID1(id int) Employee {
	return Employee{ID: id}
}

func employeeByID2(id int) *Employee {
	return &Employee{ID: id}
}

type People struct {
	idcard int
	Name   string // 首字母大写表示这个成员是可导出的，外部可以访问（前提People这个类型必须是可导出的）
}

func t1() {
	type S1 struct {
		elem1 int
		elem2 int
	}
	type S2 struct {
		elem2 int
		elem1 int
	}
	var s1 S1
	var s2 S2
	// s1和s2 是不同的类型，字段顺序也是类型的要素之一
	fmt.Println(reflect.TypeOf(s1), reflect.TypeOf(s2), reflect.TypeOf(s1) == reflect.TypeOf(s2))
}

func t2() {
	/* 自包含，如果不是指针类型是不可以的，假设可以，那么占用内存计算：8 + (8 + (8 + (8 + …))) = ∞ 都撑爆了
	type node struct {
		val int64
		next node
	}
	*/

	type node struct {
		val  int64
		node *node // 指针只占用8字节
	} // 这样算下来一个节点其实只是占用16字节

	// 内存对齐: 是cpu硬件决定的，结构体总大小必须向最大成员对齐,编译器会实施对齐
	type Bad struct {
		A int8  // 1, 补7字节
		B int64 // 8
		C int8  // 1， 补7字节
	} // 占用24字节
	type Good1 struct {
		A int8  // 1
		C int8  // 1, 补6字节
		B int64 // 8
	} // 占用16字节
	type Good2 struct {
		B int64 // 8
		A int8  // 1
		C int8  // 1, 补6字节
	} // 占用16字节
	fmt.Println(unsafe.Sizeof(Bad{}), unsafe.Sizeof(Good1{}), unsafe.Sizeof(Good2{}))
}

func t3() {
	// 模拟set数据类型
	// go不带set数据类型，大道至简，可以用map实现
	// 什么是set：
	type nothing struct{}
	sets := make(map[string]struct{})
	// var sets1 map[string]struct{}
	// sets2 := map[string]struct{}{}
	// sets3 := map[string]nothing{}
	sets["apple"] = struct{}{}
	fmt.Println(unsafe.Sizeof(sets["apple"]), sets["apple"])

	if _, ok := sets["apple"]; ok {
		fmt.Println("Found it!")
	}
}

func t4() {
	type T struct{ a, b int }
	t1 := T{
		a: 1, b: 2,
	}
	t2 := T{
		1, 2,
	}
	/* 两种初始化方式，不能混合
	t3 := T{
		a: 1, 2,
	}
	*/

	modifySNoPtr := func(t T) {
		t.a = 99
	}
	modifySPtr := func(t *T) {
		t.a = 99
	}
	modifySNoPtr(t1)
	fmt.Println(t1) // {1,2} 没修改。因为go的函数都是值传递的，目前这个修改的副本离开函数就没了
	modifySPtr(&t2)
	fmt.Println(t2) // {99,2} 修改了。值传递，只不过是传入的地址，地址指向的值始终都是那一个地方，修改那个地方的值，就算函数退出了，依然有效
}

func t5() {
	// 如果结构体的成员是可比较的，那么两个结构体就可以比较。所以前提是两个实例必须是同一个结构体类型
	type point struct {
		x, y int
	}
	type pointAnother struct {
		x, y int
	}
	p1 := point{1, 2}
	p2 := point{2, 1}
	fmt.Println(p1 == p2) // false

	non := struct{ x, y int }{
		1, 2,
	}
	fmt.Println(p1 == non) // 也可以与匿名结构体比较，因为成员字段和顺序都一样
	/* p3 := pointAnother{1, 2}
	fmt.Println(p1 == p3) // 类型不匹配无法进行比较，对于具名的结构体类型，即使成员和字段顺序一样 */
}

// 结构体嵌入和匿名
func t6() {
	type circle struct {
		x, y, radius int
	}
	type wheel struct {
		x, y, radius, spokes int
	}
	// 发现上面两个结构体有重合支出，那么就可以适当拆开进行嵌入组合
	type point struct {
		x, y int
	}
	type newCircle struct {
		point
		radius int
	}
	type newWheel struct {
		newCircle newCircle // 发现成员名字和类型是可以写成一样的
		spokes    int
	}
	// 这样拆开看起来就和舒服了，也方便以后扩展其它图形
	// 但是有个问题就是去修改成员变量会有长链条调用不太美观
	w := newWheel{}
	w.newCircle.point.x = 1

	// 引出 匿名成员：没有名字只有类型的成员
	type 含有匿名成员的Wheel struct {
		newCircle
		int
	}
	w1 := 含有匿名成员的Wheel{}
	w1.x = 1
	w1.int = 3 // 这个类型能自动充当成员名;匿名成员也有一个隐式的名字，因此不能同时包含两个类型相同的匿名成员
	// 明显让调用链变短了，但是一定要注意二义性，如果一个结构体有两个匿名成员，他们分别的成员有一样的成员名，就会有二义性
	/* type A struct{
		Name string
	}
	type B struct{
		Name string
	}
	type C struct{
		A
		B
	}
	c:=C{}
	c.Name="ccc" // compiler: ambiguous selector c.Name 。 到底修改的哪个结构体的Name呢？？？ */

	// w2 := 含有匿名成员的Wheel{x: 1}//报错 字面值初始化就没有简短表示法了
	_ = 含有匿名成员的Wheel{
		newCircle{
			point{
				x: 1,
				y: 2,
			},
			3,
		},
		4,
	}

	p := pkgs.Outer{}
	// p.inner // 无法访问
	p.X = 1 // 可以访问, 虽然inner没有导出，依然可以通过Outer访问其内的X
	// p.y // 无法访问
}

type (
	nomember struct{}
	st1      struct {
		nomember
	}
)

func (nomember) say() {
	fmt.Println("nomember-func call")
}

func (nomember) spell() {
	fmt.Println("spell-func call")
}

func (st1) say() {
	fmt.Println("st1-func call")
}

func t7() {
	st := st1{}
	st.say()   // 就近原则，先调用类型st1的say方法
	st.spell() // st1类型可以调用成员nomember的方法
}

// TODO: new 和 make 区别
