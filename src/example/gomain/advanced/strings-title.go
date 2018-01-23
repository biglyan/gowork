package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var s string = "Welcome to The WORld of go!"
	fmt.Println("original:", s)

	fmt.Println("upper:", strings.ToUpper(s))
	fmt.Println("lower:", strings.ToLower(s))
	fmt.Println("title:", strings.ToTitle(s))

	var SC unicode.SpecialCase
	fmt.Println("upper special:", strings.ToUpperSpecial(SC, s))
	fmt.Println("lower special:", strings.ToLowerSpecial(SC, s))
	fmt.Println("title special:", strings.ToTitleSpecial(SC, s))
}
