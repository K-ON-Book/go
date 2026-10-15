package main

import (
	"fmt"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return // 等价于 return z
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}

func main() {
	// 四个函数的签名相同，类型相同，但是和map，slice一样是不可比较类型，但是可以和nil比较
	fmt.Printf("%T | %T | %T | %T\n", add, sub, first, zero)

	// lsp跳转到定义的时候发现库里面的有的函数只有函数声明，没有函数体,说明实现不是通过go实现的
	fmt.Println(len([]int{}))

	f2()
}

// 实参是通过值的方式传递，其实就是[浅拷贝]，传入给函数的是副本，但是如果实参是引用类型（指针 slice map function channel等）那么通过这个[引用所指向的外部底层数据]，可能会被间接修改

// 理解深浅拷贝
func shallowDeepCopy() {
	type house struct {
		city string
	}
	type person struct {
		age  int
		home *house
	}
	p1 := person{
		age:  10,
		home: &house{city: "beijin"},
	}
	p2 := p1
	p2.home.city = "shanghai" // 浅拷贝。p2 拿到的home是p1的home地址，修改city后，p1 的city 也会变更
	fmt.Println(p1.home.city)
	//================
	val := 10
	v1 := &val
	// 深拷贝：v2去申请新的内存，拷贝的是v1地址指向的值10
	v2 := new(int)
	*v2 = *v1 // 左值是[定位],右值是[取值]
	// go没有深拷贝语法
}

// 透传
func transparentSend() {
	A := func(x, y, z int) {
	}
	B := func(m, n, q int) {
		// ...
		A(m, n, q) // B函数为了能够调用A，不得不增加一个除A以外需要的形参q，B全程不动q，只是为了传递给A，即为[B透传了q给A]
		// ...
	}
	B(0, 0, 0)

	// 假如透传链条很长
	// 那么到时候要新增一个参数，需要到处找到后修改函数签名这很繁琐，所以会引入go的[Context]，后续只需要修改这一处ctx就行

	// 一股脑所有参数都放ctx是不是就更好？
	// 不要这么做，原因：
	// 1. 可读性极差：比如看到 func Pay(ctx) 完全猜不到需要传什么，必须看内部代码才知道需要读哪些key，可读性极差。
	// 所以要遵循一些规则：
	// 1. 缺了它函数就无法工作的核心业务数据，必须显式声明在形参里
	// 2. 整趟请求大家都在同一条船上、人人都要遵守的背景信息(比如Timeout)

	// 如果透传链很长，这个参数又是核心业务数据？
	// 那为了可维护，可以改用结构体打包这些业务参数
}

// 回调和闭包
// 回调：函数A作为函数B的参数，A就是回调函数
func f1() {
	callme := func() {
		fmt.Println("callme 被调用")
	}

	notification := func(after time.Duration, cb func()) {
		time.Sleep(after)
		cb()
	}

	notification(1, callme)
}

// 闭包(closure)：“拐走”外部变量的函数
func closure1(a int) func() int {
	count := 0
	return func() int {
		count++
		a++
		return count + a
	} // return 后面的这个函数就是[闭包],是特殊的[匿名函数], a count 就是外部变量
}

func f2() {
	a := 0
	fmt.Println("==========")
	fmt.Println(closure1(a)()) // 2
	fmt.Println(closure1(a)()) // 2
	fmt.Println(closure1(a)()) // 2
	fmt.Println("==========")
	cl := closure1(a)     // 当用变量存起来后，a count一直都是在闭包中, 所以下面每一次调用后打印的结果是变化的--[私有持久化]
	fmt.Println(cl())     // 2
	fmt.Println(cl())     // 4
	fmt.Println(cl())     // 6
	fmt.Println("a= ", a) // 0, 发现闭包没有修改到a,因为 closure1 拿到的是a的副本，记住函数传参都是值传递(浅拷贝)
}
