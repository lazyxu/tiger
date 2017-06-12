package absyn

import "github.com/MeteorKL/tiger/symbol"

// Oper define
type Oper int

// Op define
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
	"+", // fsd
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

// 一共有 Exp, Dec, Ty, Var 四种抽象语法树的结点
type ASTNode interface{}

/* Linked lists and nodes of lists */

type FieldList *FieldList_
type FieldList_ struct {
	Head Field
	Tail FieldList
}

// Field is "id COLON id"
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

// Efield id EQ exp
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
