package main

import (
	"fmt"
	"reflect"
)

func CompositeLiteral复合字面量() {
	s1 := make([]int, 3, 5)
	fmt.Printf("make初始化：%p\n", s1) // 没使用&表示切片底层数组的地址
	s1 = []int{11, 22, 33, 44}
	fmt.Printf("复合字面量赋值：%p\n", s1)

	fmt.Println("s1=", s1, "len(s1)=", len(s1), "cap(s1)=", cap(s1))

	mmap := make(map[string]int, 2)
	fmt.Printf("make初始化：%p\n", mmap)
	mmap = map[string]int{"1": 1}
	fmt.Printf("复合字面量赋值：%p\n", mmap)
	fmt.Println(mmap)

	// 类型+{} 这种复合字面量语法糖，会申请一块新的内存，比如上面的map和slice一开始都make完成了初始化，后面接着通过符合字面量赋值，lsp触发了"SA4006: this value of s1 is never used [default]"在初始化那一行的警告

	// 验证方法：打印切片和 Map所指向的【底层数据内存地址】（而非变量自身的&s1）。两次打印地址完全不同，证实复合字面量重新分配了底层堆内存并替换了内部指针。
}

func init() {
	CompositeLiteral复合字面量()
	// Offcial doc：https://go.dev/ref/spec#Slice_expressions :Full slice expressions
	// slice expression: s[low:high:max], max 不写默认等于cap(s)
	// len(s[low:hight])的值为 high-low
	// cap(s[low:high])的值 max-low

	s1 := []int{11, 22, 33, 44, 55}
	t1 := s1
	s2 := s1
	fmt.Println("s1=", s1, "len(s1)=", len(s1), "cap(s1)=", cap(s1))
	s1 = s1[1:3] // [22,33]
	fmt.Println("s1=", s1, "len(s1)=", len(s1), "cap(s1)=", cap(s1))
	s1 = s1[2:3] //[44] ??? 上面截断后只有两个元素，按理来说下标索引不可能有2，3才对，这是一个坑。第一次截断后cap为5-1=4，3<=4合法，是可以看到后面的所谓的“幽灵元素”
	fmt.Println("s1=", s1, "len(s1)=", len(s1), "cap(s1)=", cap(s1))

	fmt.Println(t1[1:3:4], cap(t1[1:3:4])) // [22,33] (4-1)
	t1 = t1[:3]
	fmt.Println(t1, cap(t1)) //[11 22 33] 5
	t1 = append(t1, 66)
	fmt.Println(t1, cap(t1), s2) // [11 22 33 66] 5 [11 22 33 66 55], 66不是在55后面, s2 66 后面的44 也不在了???
	/*
				底层数组一开始长这样
				底层数组:  [ 11 | 22 | 33 | 44 | 55 ]
				下标:        0    1    2    3    4
										 ------------
												t1[:3]

		append 这个被截断的t1, 发生的是先检查cap是5，len是3，那就还是在原始数组进行append新增 66 这个元素, 放在下标是3的位置，这导致原本是44的元素被覆盖了，所以就会导致s2 看到的结果
	*/

	mys := []int{1, 2, 3}
	mys1 := mys
	mys = mys[2:2] // 3(max)-2(low)=1(cap)
	mys = append(mys, 4)
	fmt.Println(mys, mys1) //[4] [1,2,4] 上面理解了，这个也是同理
	// 那么如何避免这种底层数组被修改的情况？通过max限制
	m := []int{1, 2, 3}
	m1 := m
	m = m[1:2:2]       // 2-1=1(cap)
	m = append(m, 4)   // 新申请内存
	fmt.Println(m, m1) // [2,4] [1,2,3]
}

func main() {
	s1 := []int{1, 2, 3}
	fmt.Printf("cap=%d, len=%d, sptr=%p, sptr_arr=%p\n", cap(s1), len(s1), &s1, s1)
	s1 = s1[1:]
	// 发现和 sptr_arr 相差了 8。
	// 这是因为切片未重新分配内存，依然复用原底层数组，但切片指针后移指向了原底层数组的第 2 个元素。
	// 内存是按字节寻址的（1 个地址代表 1 字节），64 位的 int 占用 8 字节（64/8=8），
	// 所以第 2 个元素的地址正好是第 1 个元素地址 + 8。
	fmt.Printf("%p\n", s1)
	s1 = append(s1, 4) // 扩容后cap变了，底层数组的地址也就变化了

	// 切片是不能通过==比较的
	// 如果可以比较，那么就能作为map的key，可是如果这个key被修改了，那么就找不到Value了
	// mmap := map[[]int]int{} // 这种写法也是语法报错的
	// 如果包含自身后，那么这个slice本身就会报错何谈比较
	s := []any{nil}
	// s[0] = s
	// fmt.Println(s) // 报错

	// 如果只是比较切片里面的元素可以用slices.Equal
	// fmt.Println(slices.Equal(s, s)) // 报错， 因为Equal实现也是每个元素进行==比较，== 只能对可比较类型进行比较，s的里面的元素是一个[]interface{}类型的
	fmt.Println(reflect.DeepEqual(s, s)) // true, 反射可以比较，原理：？？？需要看实现，和ai帮助下进行理解，现在写了也会忘记，而且也挺麻烦的样子说起来,后面学到反射再回顾
}
