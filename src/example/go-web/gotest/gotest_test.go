package gotest

import "testing"

type AddArray struct {
	result int
	add1 int
	add2 int
}

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除法函数测试没通过")
	} else {
		t.Log("第一个测试通过了")
	}
}

func Test_Division_2(t *testing.T) {
	if _, e := Division(6, 0); e == nil {
		t.Error("Division did not work as expected.")
	} else {
		t.Log("one test passed.", e)
	}
}

func TestAdd(t *testing.T) {
	var testData = [3]AddArray {
		{2, 1, 1},
		{5, 2, 3},
		{4, 2, 2},
	}

	for _, v := range testData {
		if v.result != Add(v.add1, v.add2) {
			t.Errorf("Add(%d, %d) != %d \n", v.add1, v.add2, v.result)
		}
	}
}