package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	test1()
	// test2()
}

func test1() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func test2() {
	// 创建互斥锁
	mu := sync.Mutex{}
	var count int

	handler := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock() // 写需要加锁，如果同时请求/，拿到一样count初始值修改后并不准确，如果都是100变成101，这就不对了
		count++
		fmt.Fprintf(w, "URL.Path= %q\n", r.URL)
		mu.Unlock()
	}
	counter := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock() // 读操作也是需要加锁的，因为hander的修改count不是原子操作，如果同时请求/ 和/count，可能会读取到写的不完整的脏数据
		fmt.Println("count: ", count)
		mu.Unlock()
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.ListenAndServe("localhost:8000", nil)
}
