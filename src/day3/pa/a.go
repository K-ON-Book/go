package pa

import "day3/pa/pb"

var P int

func TestA() {
	// var P int = 1 // 该P变量不能被其他包使用，定义再函数外部就可以
	// print(P)
	print("package a")
	pb.TestB()
}
