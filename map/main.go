package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"os"
	"slices"
	"sort"
	"unicode"
)

func main() {
	// ages := make(map[string]int)
	ages := map[string]int{
		"alice":   19,
		"charlie": 34,
	}

	delete(ages, "alice") // 跳转后发现[delete]只是对map有用，注意到parameter的Type和Type1声明为了int，是个假的，其实可以接受任意类型的map；delete掉后map任然占用着着资源，后面增加元素，也是优先占用之前delete掉的位置，如果需要清空整个map可以使用[clear]函数，可以清理占用的资源

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	clear(ages)
	fmt.Println("after clear,len(ages): ", len(ages)) //0
	ages["carol"] = 21
	fmt.Println(ages) // 都clear了，按理来说不能增加元素carol的，会panic，难道clear不会让[ages]变成nil？不会变成nil，底层数据结构还是完整保留，只是清理完成后变成空的map; 如何通过key找到value的？nvim $(go env GOROOT)/src/internal/runtime/maps/runtime.go:runtime_mapaccess1 读懂这个函数的实现

	// mmap:=map[int]int{
	// 	1:2,
	// 	2:3,
	// }
	// i:=0
	// _=mmap[i++]  // i++是语句，key得是表达式

	// var mmap map[int]int// 不允许对nil map通过下标的方式进行新增元素
	mmap := map[int]int{} // 初始化后就能下标的方式存入元素
	mmap[1] = 1
	fmt.Println(mmap)

	// 顺带提一下slice. 对slice初始化后不能下标的方式去新增元素，因为下标只是去访问，对于空元素的切片是找不到的
	s := []int{}
	// s[0] = 1
	// fmt.Println(s) // 报错: index out of range [0] with length 0
	s = append(s, 1) // append 可以新增元素, [append]只是用于切片的的一个函数
	fmt.Println(s)

	// map是不能通过==比较的
	// 可以通过[maps.Equal] 比较两个map: key都能找得到不，找到了value是否相等
	m1 := map[int]int{1: 1}
	m2 := map[int]int{1: 0, 2: 1}
	// m1 == m2 // 非法操作
	fmt.Println(maps.Equal(m1, m2))
	// 如果value是比如切片类型这种不可比较类型的(Equal的实现是==比较value的)，则可以用maps.EqulFunc(s1,s1,slices.Equal)
	m3 := map[int][]int{1: {1, 1}}
	m4 := map[int][]int{0: {1, 1}}
	// fmt.Println(maps.Equal(m3, m4)) // 非法操作 []int not comparable
	fmt.Println(maps.EqualFunc(m3, m4, slices.Equal))
	// 什么叫做不可比较类型：不能通过==或!=比较运算符比较的类型

	fmt.Println(k([]string{"11", "22"})) // ["11" "22"]

	countRune()

	// 测试函数是否可以对外部map进行修改
	fmt.Println(m_)
	modifyM(m_)
	fmt.Println(m_)
	// 测试函数是否可以对外部slice进行修改
	fmt.Println(s_)
	modifyS(s_)
	fmt.Println(s_)
}

func equal(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, xv := range m1 {
		if yv, ok := m2[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}

func k(list []string) string { return fmt.Sprintf("%q", list) }

func countRune() {
	counts := make(map[rune]int)
	rd := bufio.NewReader(os.Stdin)
	for {
		r, n, e := rd.ReadRune()
		if e == io.EOF {
			break
		}
		if e != nil {
			fmt.Fprintf(os.Stdout, "%v", e)
		}
		if r == unicode.ReplacementChar && n == 1 {
			continue
		} // 这是担心非法输入然后go会替换成\uFFFD这个菱形问号，为什么还要n==1判断呢，因为如果输入的是合法的\uFFFD是占用3字节的
		counts[r]++
	}
	ru := '\uFFFD'
	byte := []byte(string(ru))
	fmt.Printf("%q,%#[1]b\n", ru)
	for _, v := range byte {
		fmt.Printf("%#b\n", v)
	}
	for k, v := range counts {
		fmt.Fprintf(os.Stdout, "%q %d\n", k, v)
	}
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

var (
	m_ = map[string]int{}
	s_ = []int{1, 2}
)

func modifyM(m map[string]int) {
	m["apple"] = 1
}

func modifyS(s []int) {
	s[0] = 99
	s = append(s, 100)
}
