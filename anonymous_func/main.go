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
		os.Mkdir(d, 0o755)
		rmdir = append(rmdir, func() { os.RemoveAll(d) })
	}

	dirs := tempDir()
	for i := 0; i < len(dirs); i++ {
		os.Mkdir(dirs[i], 0o755)
		rmdir = append(rmdir, func() {
			os.RemoveAll(dirs[i])
		})
	}
	for _, clean := range rmdir {
		clean()
	}
}
