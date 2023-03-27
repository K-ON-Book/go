package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	day := t.Day()
	month := t.Month()
	year := t.Year()
	fmt.Println(day, month, year)
	fmt.Println(t.Format("02 Jan 2006 15:04"))
}
