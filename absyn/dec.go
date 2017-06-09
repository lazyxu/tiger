package absyn

import "github.com/MeteorKL/tiger/symbol"

type Dec interface {
	Node
	A_dec()
}

func (*FunctionDec) A_dec() {}
func (*VarDec) A_dec()      {}
func (*TypeDec) A_dec()     {}

type FunctionDec struct {
	Pos      Pos
	Function FundecList
}
type VarDec struct {
	Pos    Pos
	Var    symbol.Symbol
	Typ    symbol.Symbol
	Init   Exp
	Escape bool
}
type TypeDec struct {
	Pos  Pos
	Type NametyList
}

/* Linked lists and nodes of lists */

type FundecList *FundecList_
type FundecList_ struct {
	Head Fundec
	Tail FundecList
}
type Fundec *Fundec_
type Fundec_ struct {
	Pos    Pos
	Name   symbol.Symbol
	Params FieldList
	Result symbol.Symbol
	Body   Exp
}

type DecList *DecList_
type DecList_ struct {
	Head Dec
	Tail DecList
}

func DecListInsert(head Dec, tail DecList) DecList {
	return &DecList_{
		Head: head,
		Tail: tail,
	}
}

func FundecListInsert(head Fundec, tail FundecList) FundecList {
	return &FundecList_{
		Head: head,
		Tail: tail,
	}
}
