package main

import (
	"fmt"
	"reflect"
)

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
