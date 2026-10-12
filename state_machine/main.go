package main

import "fmt"

type State int

const (
	Locked State = iota
	Unlocked
)

func (s State) String() string {
	if s == Locked {
		return "【已锁定 🔴】"
	}
	return "【已解锁 🟢】"
}

type Turnstile struct {
	state      State
	passCount  int // 累计通过多少人
	totalCoins int // 累计投币/刷卡次数
}

func NewTurnstile() *Turnstile {
	return &Turnstile{
		state: Locked,
	}
}

func (t *Turnstile) Card() {
	fmt.Printf("\n--> 用户执行了 [刷卡]操作，当前状态：%s\n", t.state)
	switch t.state {
	case Locked:
		t.totalCoins++
		t.state = Unlocked
		fmt.Println("    [响应] 嘀！刷卡成功，开锁放行！")
	case Unlocked:
		// 规则：已经开门了还刷卡 -> 提醒用户
		fmt.Println("    [响应] 门已经是开着的，请直接通过，不重复扣费。")

	}

	fmt.Printf("    当前最新状态变为：%s\n", t.state)
}

func (t *Turnstile) Push() {
	fmt.Printf("\n--> 用户执行了 [推门]操作，当前状态：%s\n", t.state)
	switch t.state {
	case Locked:
		fmt.Println("请刷卡")
	case Unlocked:
		t.passCount++
		t.state = Locked
	}
}

func main() {
	t1 := NewTurnstile()
	t1.Push()

	t1.Card()
	t1.Push()
	t1.Push()
	fmt.Println("通过人数：", t1.passCount, "刷币次数：", t1.totalCoins)

	t2 := NewTurnstile()
	t2.Card()
	t2.Card()
	t2.Push()
	fmt.Println("通过人数：", t2.passCount, "刷币次数：", t2.totalCoins)
}
