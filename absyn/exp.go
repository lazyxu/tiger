package absyn

import "github.com/MeteorKL/tiger/symbol"

type Exp interface {
	ASTNode
	A_exp()
}

func (*VarExp) A_exp()    {}
func (*NilExp) A_exp()    {}
func (*IntExp) A_exp()    {}
func (*StringExp) A_exp() {}
func (*CallExp) A_exp()   {}
func (*OpExp) A_exp()     {}
func (*RecordExp) A_exp() {}
func (*SeqExp) A_exp()    {}
func (*AssignExp) A_exp() {}
func (*IfExp) A_exp()     {}
func (*WhileExp) A_exp()  {}
func (*ForExp) A_exp()    {}
func (*BreakExp) A_exp()  {}
func (*LetExp) A_exp()    {}
func (*ArrayExp) A_exp()  {}

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

// IfExp f
type IfExp struct {
	Pos              Pos
	Test, Then, Else Exp
}

// IfExpression : There are two forms of variable declarations in Tiger.
// __if exp1 then exp2 else exp3__
// exp1 is typed as an integer, exp2 and exp3 must have the same type which will be the type of the entire structure. The resulting type cannot be that of nil.
// __if exp1 then exp2__
// exp1 is typed as an integer, and exp2 must have no value. The whole expression has no value either.
func IfExpression(e Exp) {
}

func (*IfExp) Positon() int {
	return 1
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
