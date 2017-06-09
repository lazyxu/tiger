package util

func PanicErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Assert(b bool) {
	if !b {
		panic("assert")
	}
}
