package util

import "fmt"

const IsDebug bool = true

func Debug(s string) {
	if IsDebug {
		fmt.Println(s)
	}
}
