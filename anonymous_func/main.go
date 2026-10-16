package main

import (
	"fmt"
	"os"
)

func main() {
	i := 3
	increment := func() int {
		i++
		return i
	} // 等号右边的东西其实就是[匿名函数]，因为func后面没写名字，也叫[闭包]因为“拐走”了i，也叫[函数值],因为能赋值给等号左边的变量
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(i) // 6 。本质：函数值中记录的是i的内存地址, 所以每次修改值都是修改到i地址指向的具体值.涉及到[逃逸分析]，increment函数执行完i的值到哪里去了,当闭包修改i时会把i从栈移动到堆上分配内存,所以即使函数退出了，这个i的值还能找到

	t1()
}

// 陷阱：捕获迭代变量
func tempDir() []string {
	return []string{"A", "B", "C"}
}

func t1() {
	var rmdir []func()
	for _, d := range tempDir() {
		fmt.Println(&d) // go 1.21以前会是同一个地址，由于闭包记录的是循环变量的地址，d地址指向的值就会是最后一次循环的值，这会导致下面只能删除同一个目录. 1.21以后 for range，和普通的for后面的循环变量当真正需要内存地址的时候都会开辟新的内存，为了性能，当不明确需要这个地址的时候编译器会优化（https://github.com/golang/go/issues/60078 https://go.dev/blog/loopvar）。修改go.mod 的go版本可以验证
		os.Mkdir(d, 0o755)
		rmdir = append(rmdir, func() { os.RemoveAll(d) })
	}

	dirs := tempDir()
	// i := 0
	for i := 0; i < len(dirs); i++ { // 如果这个i声明放到for的外面，那么1.21以后也还是会是打印同一处地址
		fmt.Println(&i) // 和上面同理
		os.Mkdir(dirs[i], 0o755)
		rmdir = append(rmdir, func() {
			os.RemoveAll(dirs[i])
		})
	}
	for _, clean := range rmdir {
		clean()
	}
}
