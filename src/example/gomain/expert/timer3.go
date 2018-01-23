package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		timer := time.NewTimer(5 * time.Second) // --A
		<-timer.C
		time.Sleep(2 * time.Second) // -- B
		fmt.Println("now:", time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Time: A + B
