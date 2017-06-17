package util

import "fmt"

const IsDebug bool = false

func Debug(s string) {
	if IsDebug {
		fmt.Println(s)
	}
}
