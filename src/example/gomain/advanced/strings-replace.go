package main

import (
	"fmt"
	"strings"
)

func main() {
	value := "You cat is cute"
	fmt.Println(value)

	result := strings.Replace(value, "cat", "dog", -1)
	fmt.Println(result)

	value = "bird bird bird"
	result = strings.Replace(value, "bird", "fish", 1)
	fmt.Println(result)

	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}

	fmt.Println(strings.Map(rot13, "Twas brillig and the slithy gopher..."))
}
