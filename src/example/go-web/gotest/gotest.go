package gotest

import "errors"

func Add(x, y int) (z int) {
	z = x + y
	return
}

type ForTest struct {
	num int
}

func (this *ForTest) Loops() {
	for i := 0; i < 10000; i++ {
		this.num++
	}
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	return a / b, nil
}