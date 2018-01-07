package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func executableDir() (string, error) {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}

	return filepath.Dir(pathAbs), nil
}

func GetProjectRoot() string {
	binDir, err := executableDir()
	if err != nil {
		return ""
	}

	return path.Dir(binDir)
}

func Welcome() {
	fmt.Println("***********************************")
	fmt.Println("*******欢迎来到Go语言中文网*******")
	fmt.Println("***********************************")
}

func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}

	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		if totalSize == asciiPos {
			return pos
		}
	}

	return pos
}