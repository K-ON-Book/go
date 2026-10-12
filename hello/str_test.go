package main

import (
	"strconv"
	"strings"
	"testing"
)

func generateArgs(n int) []string {
	args := make([]string, n)
	for i := 0; i < n; i++ {
		args[i] = strconv.Itoa(i)
	}
	return args
}

func BenchmarkStringJoin(b *testing.B) {
	args := generateArgs(1000)

	for b.Loop() {
		strings.Join(args, " ")
	}
}

func BenchmarkStrinngPlus(b *testing.B) {
	args := generateArgs(1000)

	for b.Loop() {
		s, sep := "", ""
		for _, v := range args {
			s += v + sep
			sep = " "
		}
	}
}
