package main

import (
	"fmt"
	"strconv"
	// "math/rand"
	"strings"
	"unicode/utf8"
)

func main() {
	// var a = 10 * (rand.Float32())
	// fmt.Println(a)
	// type_alias()
	// string_learn()
	// count_number_of_characters()
	// prefuffix()
	// contains_learn()
	// string_index()
	// string_repalce()
	// string_rpeat()
	// string_to_upper_lower()
	// string_trim()
	// string_splitjoin()
	strconv_learn()
}

func type_alias() {
	type Rope string
	var a Rope = "test tset"
	fmt.Printf("%s\n", a)
}

func string_learn() {
	a := `This is a raw string \n`
	fmt.Println(a, "\t", a[1], "\t", a[3])
}

func count_number_of_characters() {
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	fmt.Printf("The number of characters in string str2 is %d", utf8.RuneCountInString(str2))
}

func prefuffix() {
	var str = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("T/F? Does the string \"%s\" have suffix %s?", str, "gn")
	fmt.Printf("%t\n", strings.HasSuffix(str, "gn"))
}

func contains_learn() {
	var a string = "this is an apple!"
	fmt.Println(strings.Contains(a, "s a"))
}

func string_index() {
	var a string = "this is an apple!"
	fmt.Println(strings.Index(a, "s"))
	fmt.Println(strings.Index(a, "a"))     //8
	fmt.Println(strings.Index(a, "q"))     //-1
	fmt.Println(strings.LastIndex(a, "a")) //11
	b := "asSASA ddd dsjkdsjsこん dk"
	fmt.Println(strings.IndexRune(b, 'こ'))
	fmt.Println(strings.Index(b, "ds")) //11
}

func string_repalce() {
	var a string = "this is an apple!"
	fmt.Println(strings.Replace(a, "this is", "that", 1))
}

func string_rpeat() {
	var origs string = "Hi there! "
	fmt.Println(strings.Repeat(origs, 5))
}
func string_to_upper_lower() {
	var orig string = "Hey, how are you George?"
	var lower string
	var upper string
	fmt.Printf("The original string is: %s\n", orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n", upper)
}

func string_trim() {
	var a string = " this is an apple! "
	fmt.Println(a)
	fmt.Println(strings.TrimSpace(a))
	var b string = "this is an apple! "
	fmt.Println(strings.Trim(b, "th"))
	fmt.Println(strings.TrimRight(b, "! "))
}

func string_splitjoin() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %s", sl)
	for key, val := range sl {
		fmt.Printf("%d_%s - ", key, val)
	}
	fmt.Print("\n==================\n")
	str2 := "GO1*The ABC of Go*25"
	sl2 := strings.Split(str2, "*")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for key, val := range sl2 {
		fmt.Printf("%d_%s - ", key, val)
	}
	fmt.Print("\n==================\n")
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s", str3)
}

func strconv_learn() {
	const (
		m = iota
		n
	)
	var a int = 50
	var b = 10 << n                                // 20
	fmt.Println(strconv.Itoa(a) + strconv.Itoa(b)) //5020
	var c float64 = 5.024284
	str := strconv.FormatFloat(c, 'g', 3, 64)
	fmt.Println(str)
	var str1 string = "10"
	var str2 string = "10"
	d1, _ := strconv.Atoi(str1)
	d2, _ := strconv.Atoi(str2)
	fmt.Println(d1 + d2)
	var str3 string = "10.254"
	var str4 string = "10.254"
	d3, _ := strconv.ParseFloat(str3, 64)
	d4, _ := strconv.ParseFloat(str4, 64)
	fmt.Println(d3 + d4)
}
