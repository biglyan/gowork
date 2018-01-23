package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, world")
	loopworker()
}

func loopworker() {
	i := 0
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			i++
			doxx(i)
		}
	}
}

func doxx(i int) {
	// time.Sleep(3 * time.Second)
	fmt.Println("do", i, time.Now().Unix())
}
