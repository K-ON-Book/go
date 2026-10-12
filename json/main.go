package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	t1()
	t2()
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"` // omitempty 表示当结构体成员为空或者零值(bool类型零值是false)时不生成该json对象
	Actors []string
}

func t1() {
	movies := []Movie{
		{
			Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
		},
		{
			Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"},
		},
		{
			Title: "Bullitt", Year: 1968,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
		},
	}
	// data, err := json.Marshal(movies) // 打印出来可读性不好
	data, err := json.MarshalIndent(movies, "", "  ") // 只有导出的成员才会被编码,把某个成员改成小写发现打印出来就没有了(试过将Movie改成movie没对打印结果造成影响)，对于未打tag的成员名字会作为json key的名字
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	os.WriteFile("test.json", data, 0o600) // 0600 & umask 才是最终的文件权限; 0o(ctal) 八进制，0b(inary) 二进制，0x(hexadecimal) 十六进制
	//////////////////
	f1, err := os.Create("test1.json")
	if err != nil {
		log.Fatal(err)
	}

	f1.Write(data)
	f1.Close()
	/////////////
	f2, _ := os.OpenFile("test2.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o666)
	f2.Write(data)
	f1.Close()

	fmt.Printf("%s\n", data)

	// 编码的操作是解码，这里就是反序列化
	binddata := []struct{ Title string }{} // 反序列化的成员也必须是导出的才行, 改成小写解出来是空值[{} {} {}]
	type td struct {
		Title string
	}
	var binddata1 []td
	if err := json.Unmarshal(data, &binddata); err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(data, &binddata1) // 可以解码出来
	fmt.Println(binddata, binddata1)
}

func t2() {
	const issueURL = "https://api.github.com/search/issues"
	type user struct {
		Login   string
		HTMLURL string `json:"html_url"`
	}
	type issue struct {
		Number    int
		HTMLURL   string `json:"html_url"`
		Title     string
		State     string
		User      *user
		CreatedAt time.Time `json:"created_at"`
		Body      string    // in Markdown format
	}
	type issuesSearchResult struct {
		TotalCount int `json:"total_count"`
		Items      []*issue
	}

	searchIssues := func(terms []string) (*issuesSearchResult, error) {
		q := url.QueryEscape(strings.Join(terms, " "))
		resp, err := http.Get(issueURL + "?q=" + q)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("search query failed: %s", resp.Status)
		}
		var result issuesSearchResult
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, err
		}
		return &result, nil
	}

	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title) // -5 表示左对齐，最小宽度是5，比如"123"会表示成"123  "; 9.9 小数点前9表示最下宽度是9，小数点后9表示如果超过9个字符后面内容会被截断；.55表示超过55个字符后面内容会被截断
	}
}
