package main

import (
	"fmt"
	"net"
)

func main() {
	host, port, err := net.SplitHostPort("127.0.0.1:5432")
	fmt.Println(host, port, err)
}
