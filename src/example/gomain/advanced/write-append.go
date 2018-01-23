package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetCurPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	splitString := strings.Split(path, "\\")
	size := len(splitString)
	splitString = strings.Split(path, splitString[size-1])
	return strings.Replace(splitString[0], "\\", "/", size-1)
}

func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

func main() {
	fmt.Println("GetCurPath:", GetCurPath())
	fmt.Println("GetCurrPath:", GetCurrPath())
	os.Exit(1)
	file, err := os.OpenFile("input.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(" This is the text appended"); err != nil {
		panic(err)
	}

	fmt.Println("append success")
}
