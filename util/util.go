package util

func checked_malloc() {

}

func Hash(bytes []byte) uint {
	var h uint = 0
	for _, b := range bytes {
		h = h*65599 + uint(b)
	}
	return h
}
