package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64 = 0
	count := 0

	for i := 0; i < 100; i++ {
		go func() {
			for {
				count++
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
	fmt.Println("count:", count)
}
