package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(cwd)
}

func main() {
	fmt.Println("main" + cwd)
}
