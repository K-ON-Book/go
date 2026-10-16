package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// extractTitle 递归遍历 HTML 树，找到 <title> 文本并返回
// TODO: 后续理解算法
func extractTitle(n *html.Node) string {
	// 找到 <title> 标签，且它肚子里有子内容
	if n.Type == html.ElementNode && n.Data == "title" && n.
		FirstChild != nil {
		return n.FirstChild.Data // 核心：取出子节点（文本节点）里的文字！
	}

	// 深度优先遍历所有子节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if title := extractTitle(c); title != "" {
			return title
		}
	}
	return ""
}

func getTitle() (string, error) {
	const site = "https://go.dev"
	resp, err := http.Get(site)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if ct := resp.Header.Get("Content-Type"); !strings.HasPrefix(ct, "text/html") {
		return "", fmt.Errorf("content type %s not supported", ct) // 假如这里return了，那么就会在return之前先执行resp.Body.Close()
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}
	return extractTitle(doc), nil
}

func t1() {
	// 就像入栈出栈一样，执行defer的语句，最后执行（FILO）
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func t2(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}() // 在return之前执行，发现result是2x，也就是说先renun后面的x+x更新result值后，才执行的defer后面的函数
	return x + x // 等价于下面的写法，这样就理解上一句注释了
	/* result=x+x
	return */
}

func t3() {
	filenames := []string{"/path/to/A", "/path/to/B"}
	for _, filename := range filenames {
		f, _ := os.Open(filename)
		defer f.Close()
		b := make([]byte, 10)
		f.Read(b)
		fmt.Println(string(b)) // defer 语句是打印语句执行后接着执行吗？不是！记住defer是认函数的结束，也就是说 在do other things后面执行的。如果掉进这个误区，当filenames很多的时候就会导致打开的文件过多，系统报错“too many open files”, 所以正确的做法是将操作文件的流程封装成一个函数，循环调用这个函数就行
	}
	// do other things
}

func t3Fix() {
	filenames := []string{"/path/to/A", "/path/to/B"}
	rf := func(filename string) {
		f, _ := os.Open(filename)
		defer f.Close()
		b := make([]byte, 10)
		f.Read(b)
		fmt.Println(string(b))
	}
	for _, filename := range filenames {
		rf(filename)
	}
}

func main() {
	res, _ := getTitle()
	fmt.Println(res)

	t1()
	t2(3)
}
