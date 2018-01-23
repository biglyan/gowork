package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func sum(id int) {
	var x int64
	for i := int64(0); i < math.MaxUint32; i++ {
		x += i
	}
	println(id, x)
}

func main() {
	t1 := time.Now()
	wg := new(sync.WaitGroup)
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}
	wg.Wait()

	t2 := time.Now()
	fmt.Println("time elapsed:", t2.Sub(t1), "s")
}
