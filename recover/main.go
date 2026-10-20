package main

import "fmt"

func main() {
	t1()
	fmt.Println("=========")
	t2()
	fmt.Println("=========")
	t3()
}

func t1() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic("panic call")
}

func t2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	defer fmt.Println(1)
	defer fmt.Println(2)
	panic("panic call")

	// recover在最后才会执行，发生panic了，居然也会打印“2” “1”，说明发生panic不影响前面“注册”进来的defer语句执行
}

func t3() {
	defer fmt.Println(1)
	panic("panic call") // 把panic提前，发现只是执行了前面“注册”进来的"defer xxx"。defer是当前t3函数结束之前执行，这里的panic就是t3执行结束的“信号”,所以打印1之后，后面的语句不会再执行
	defer fmt.Println(2)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
}
