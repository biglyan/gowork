package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e8)
	tick2 := time.Tick(2e8)
	boom := time.After(5e8)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-tick2:
			fmt.Println("tick2.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Print(".")
			time.Sleep(5e7)
		}
	}
}
