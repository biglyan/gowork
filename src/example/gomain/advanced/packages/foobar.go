package foobar

func privateFunc() int {
	return 23
}

func PublicFunc(value int) int {
	return value + privateFunc()
}

func PlusOne(value int) int {
	return value + 1
}
