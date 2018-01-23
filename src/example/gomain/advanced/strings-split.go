package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) //根据,进行拆分，返回的是[]string
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))

	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("  foo1;bar2,baz3...", f))

	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) //连接

	fmt.Println("ba" + strings.Repeat("na", 2)) //重复
}
