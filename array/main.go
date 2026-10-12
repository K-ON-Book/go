package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	var r [2]int = [2]int{1, 2}
	fmt.Println(r[1])

	q := [...]int{1, 2, 3}
	fmt.Println(q)

	m := [...]int{99: 1}
	_ = m

	c1 := sha256.Sum256([]byte{'x'})
	c2 := sha256.Sum256([]byte{'X'})
	fmt.Printf("%x,%x,%t,%T\n", c1, c2, c1 == c2, c1)

	printSha()
}

func zero(ptr *[32]int) {
	for i := range ptr {
		(*ptr)[i] = 0
	}
}

func printSha() {
	mymap := map[string]bool{
		"sha256": true,
		"sha384": true,
		"sha512": true,
	}
	str := flag.String("str", "", "")
	sha := flag.String("sha", "sha256", "")
	flag.Parse()
	if _, ok := mymap[*sha]; ok {
		switch *sha {
		case "sha256":
			fmt.Printf("%x\n", sha256.Sum256([]byte(*str)))
		case "sha512":
			fmt.Printf("%x\n", sha512.Sum512([]byte(*str)))

		case "sha384":
			fmt.Printf("%x\n", sha512.Sum384([]byte(*str)))

		}
	} else {
		fmt.Printf("%s", "不支持的加密算法")
	}
}
