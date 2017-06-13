package tree

import "github.com/MeteorKL/tiger/util"

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

func NotRel(r RelOp) RelOp {
	switch r {
	case Eq:
		return Ne
	case Ne:
		return Eq
	case Lt:
		return Ge
	case Ge:
		return Lt
	case Gt:
		return Le
	case Le:
		return Gt
	case Ult:
		return Uge
	case Uge:
		return Ult
	case Ule:
		return Ugt
	case Ugt:
		return Ule
	}
	util.Assert(!true)
	return 0
}
