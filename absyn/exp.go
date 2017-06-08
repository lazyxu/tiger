package absyn

import "github.com/MeteorKL/tiger_compiler/symbol"

type Exp interface {
	Node
	expNode()
}

func (*VarExp) expNode()    {}
func (*NilExp) expNode()    {}
func (*IntExp) expNode()    {}
func (*StringExp) expNode() {}
func (*CallExp) expNode()   {}
func (*OpExp) expNode()     {}
func (*RecordExp) expNode() {}
func (*SeqExp) expNode()    {}
func (*AssignExp) expNode() {}
func (*IfExp) expNode()     {}
func (*WhileExp) expNode()  {}
func (*ForExp) expNode()    {}
func (*BreakExp) expNode()  {}
func (*LetExp) expNode()    {}
func (*ArrayExp) expNode()  {}

type VarExp struct {
	Pos Pos
	Var Var
}
type NilExp struct {
	Pos Pos
}
type IntExp struct {
	Pos Pos
	Int int
}
type StringExp struct {
	Pos    Pos
	String string
}
type CallExp struct {
	Pos  Pos
	Func symbol.Symbol
	Args ExpList
}
type OpExp struct {
	Pos   Pos
	Oper  Oper
	Left  Exp
	Right Exp
}
type RecordExp struct {
	Pos     Pos
	Typ     symbol.Symbol
	Efields EfieldList
}
type SeqExp struct {
	Pos Pos
	Seq ExpList
}
type AssignExp struct {
	Pos Pos
	Var Var
	Exp Exp
}
type IfExp struct {
	Pos              Pos
	Test, Then, Else Exp
}
type WhileExp struct {
	Pos        Pos
	Test, Body Exp
}
type ForExp struct {
	Pos          Pos
	Var          symbol.Symbol
	Lo, Hi, Body Exp
	Escape       bool
}
type BreakExp struct {
	Pos Pos
}
type LetExp struct {
	Pos  Pos
	Decs DecList
	Body Exp
}
type ArrayExp struct {
	Pos        Pos
	Typ        symbol.Symbol
	Size, Init Exp
}

/* Linked lists and nodes of lists */

type ExpList *ExpList_
type ExpList_ struct {
	Head Exp
	Tail ExpList
}

func ExpListInsert(head Exp, tail ExpList) ExpList {
	return &ExpList_{
		Head: head,
		Tail: tail,
	}
}
