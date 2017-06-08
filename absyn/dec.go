package absyn

import "github.com/MeteorKL/tiger_compiler/symbol"

type Dec interface {
	Node
	decNode()
}

func (*FunctionDec) decNode() {}
func (*VarDec) decNode()      {}
func (*TypeDec) decNode()     {}

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
