package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func HashFileMd5(filepath string) (string, error) {
	var res string
	file, err := os.Open(filepath)
	if err != nil {
		return res, nil
	}

	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return res, err
	}

	hashInBytes := hash.Sum(nil)[:16]
	res = hex.EncodeToString(hashInBytes)
	return res, nil
}

func main() {
	hash, err := HashFileMd5(os.Args[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)
}
