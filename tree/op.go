package tree

type BinOp int

const (
	Plus BinOp = iota
	Minus
	Mul
	Div
	And
	Or
	Lshift
	Rshift
	Arshift
	Xor
)

type RelOp int

const (
	Eq RelOp = iota
	Ne
	Lt
	Gt
	Le
	Ge
	Ult
	Ule
	Ugt
	Uge
)
