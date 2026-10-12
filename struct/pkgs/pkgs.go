package pkgs

type inner struct { // 小写：未导出
	X int // 大写：导出
	y int // 小写：未导出
}

type Outer struct { // 大写：公开导出的结构体
	inner //  匿名嵌入了未导出的 inner
}
