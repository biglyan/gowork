package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.Split("127.0.0.1:5432", ":")
	ip, port := s[0], s[1]
	fmt.Println(ip, port)
}
