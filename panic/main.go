package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// t1()
	// t3(3)
	t4()
}

func t1() {
	defer fmt.Println("hello")
	panic("call panic") // panic之前defer的函数会被执行
}

// 过度防御性编程
// func t2(x *Buffer)  {
// 	if x==nil{
// 		panic("x is nil") // 这是没必要的，因为紧接着下面就对x的成员访问，会自动触发对空指针解引用,然后运行时panic。[除非这个报错提供了更多有价值的信息或者在访问x的成员之前做了很多不可逆的重要修改，把数据搞脏了]。总而言之，需要规范的使用[Panic]
// 	}
// 	x.element=nil
// }

// 对于大部分可以预料的错误应该使用go提供的错误机制，一个健壮的程序不要老是崩溃，也就是别滥用panic

func t3(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x) // 执行到x为0的时候，前一行panic了，不是说[panic之前会把defer语句执行]吗，发现没有打印`defer 0`, 这是因为根本没有执行到defer语句就panic了，根本没把这条defer语句记录下来，如果把defer语句与前一行交换位置就可以了
	t3(x - 1)

	/*
		f(3){
			f(2){
				f(1){
					f(0)
				}defer 1 执行
			}defer 2 执行
		}defer 3 执行
	*/
}

// / 对于panic机制来说，延迟函数（defer 后面的函数）的调用发生在释放堆栈信息之前
func printStack() {
	buf := make([]byte, 2096)
	n := runtime.Stack(buf, false)
	os.Stdout.Write(buf[:n])
} // 这里只是打印，打印不是释放，释放是指的是从系统删除了,找不到了

func myfunc1() {
	fmt.Println("hello")
}

func myfunc2() {
	panic("call panic")
}

func t4() {
	defer printStack()
	myfunc1() // 正常退出，立即释放堆栈信息，所以打印堆栈看不到这个函数
	myfunc2() // 发生panic后,异常退出，堆栈信息没有立即释放，所以能打印看到
}
