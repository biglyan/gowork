package main

import (
	"fmt"
	"mymath"
)

func main() {
	fmt.Printf("Hello, world. Sqrt(2) = %v\n", mymath.Sqrt(2))

	xs := []float64{1, 2, 3, 4}
	avg := mymath.Average(xs)
	fmt.Println(avg)
	fmt.Println(mymath.Min(xs))
	fmt.Println(mymath.Max(xs))
}
