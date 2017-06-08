package absyn

import "github.com/MeteorKL/tiger/symbol"

type Oper int

const (
	PlusOp Oper = iota
	MinusOp
	TimesOp
	DivideOp
	EqOp
	NeqOp
	LtOp
	LeOp
	GtOp
	GeOp
)

var opnames = [...]string{
	"+",
	"-",
	"*",
	"/",
	"=",
	"<>",
	"<",
	"<=",
	">",
	">=",
}

type Node interface{}

/* Linked lists and nodes of lists */

type FieldList *FieldList_
type FieldList_ struct {
	Head Field
	Tail FieldList
}
type Field *Field_
type Field_ struct {
	Pos       Pos
	Name, Typ symbol.Symbol
	Escape    bool
}
type EfieldList *EfieldList_
type EfieldList_ struct {
	Head Efield
	Tail EfieldList
}
type Efield *Efield_
type Efield_ struct {
	Name symbol.Symbol
	Exp  Exp
}

func EfieldListInsert(head Efield, tail EfieldList) EfieldList {
	return &EfieldList_{
		Head: head,
		Tail: tail,
	}
}
func FieldListInsert(head Field, tail FieldList) FieldList {
	return &FieldList_{
		Head: head,
		Tail: tail,
	}
}
