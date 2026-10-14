package main

import (
	"log"
	"os"
)

func main() {
	f, _ := t1() // blank identifier "_" 可以忽略返回值
	log.Println(f)
	log.Println(t1())
}

func t1() (f *os.File, err error) { // *os.File 等价于 *(os.File)
	f, err = os.OpenFile("test.txt", os.O_RDONLY, 0o000)
	return // bare return. 等价于 return f,err
}
