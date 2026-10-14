package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
错误：是预期的一个值
异常：未被预料到的错误

假如当某个错误被当作异常处理后，这个错误会将堆栈跟踪信息返回给终端用户，这些信息复杂无用（对用户），无法帮助定位错误。所以正确的异常和错误处理很重要
*/

func main() {
	err := errors.New("这是一个错误")
	_ = fmt.Errorf("Test %s", err)  // 警告。开头首字母别大写
	_ = fmt.Errorf("test %s.", err) // 警告。末尾别加句号
	// 以上警告可以看官方文档：https://go.dev/wiki/CodeReviewComments#error-strings 在传播错误的过程中可能会有很长的错误链，就像“genesis: crashed: no parachute: G-switch failed: bad relay orientation” 为了清晰最好保持良好的规范

	if err := waitForServer("go.dev", time.Minute); err != nil {
		// fmt.Fprintf(os.Stderr, "site is down: %s", err)
		log.Fatalf("Site is down: %v", err) // 发现居然会自动换行,log包中的所有函数会为没有换行符的字符串增加换行符，
	} else {
		fmt.Println("success")
	}
}

func waitForServer(url string, timeout time.Duration) error {
	const maxSleepTime = 10 * time.Second
	ddl := time.Now().Add(timeout)
	for retry := 0; time.Now().Before(ddl); retry++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s);retrying...", url)
		sleepTime := time.Second << uint(retry)
		sleepTime = min(sleepTime, maxSleepTime)
		time.Sleep(sleepTime)
	}
	return fmt.Errorf("server %s failed to response after %s", url, timeout)
}
