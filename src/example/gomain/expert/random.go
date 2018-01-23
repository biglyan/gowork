package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		fmt.Println(random(1000, 9999))
	}
}
