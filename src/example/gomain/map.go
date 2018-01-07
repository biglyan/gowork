package main

import (
	"fmt"
	"sort"
)

func modify(dict map[string]int) {
	dict["张三"] = 10
}

func main() {
	dict := map[string]int{"王五":60, "张三":43}
	var names []string
	for name, _ := range dict {
		names = append(names, name)
	}

	modify(dict)
	sort.Strings(names)
	for _, key := range names {
		fmt.Println(key, dict[key])
	}
}