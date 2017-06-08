package util

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	src, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	f.Close()
	return src
}

func CheckFileExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
