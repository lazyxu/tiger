package util

func PanicErr(e error) {
	if e != nil {
		panic(e)
	}
}
