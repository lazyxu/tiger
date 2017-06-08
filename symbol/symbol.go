package symbol

type Symbol *Symbol_
type Symbol_ struct {
	Name string
	Next Symbol
}

const SIZE = 109 /* should be prime */
var hashtable [SIZE]Symbol

func hash(bytes []byte) uint {
	var h uint = 0
	for _, b := range bytes {
		h = h*65599 + uint(b)
	}
	return h
}

func SymbolInsert(name string) Symbol {
	index := hash([]byte(name)) % SIZE
	syms := hashtable[index]
	var sym Symbol
	for sym = syms; sym != nil; sym = sym.Next {
		if sym.Name == name {
			return sym
		}
	}
	sym = &Symbol_{name, syms}
	hashtable[index] = sym
	return sym
}
