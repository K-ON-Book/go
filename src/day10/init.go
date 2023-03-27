package main

import (
	"fmt"
	"runtime"
)

var prompt string = "Enter %s to quit!"

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z,Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}
