package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type HandleFunc func(int)

func Handler(id int, handleFunc HandleFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	handleFunc(id)
}

func HasGoSched(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, i)
		if i == 3 || i == 8 {
			runtime.Gosched()
		}
	}
}

func SayHello(id int) {
	fmt.Println(id, "Hello World!")
}

func OtherProc(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, i)
		if i == 5 {
			runtime.Gosched()
		}
	}
}

func main() {
	t1 := time.Now()
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go Handler(0, HasGoSched, wg)
	go Handler(1, SayHello, wg)
	go Handler(2, OtherProc, wg)

	wg.Wait()

	t2 := time.Now()
	fmt.Println("elapsed time:", t2.Sub(t1), "s")
}
