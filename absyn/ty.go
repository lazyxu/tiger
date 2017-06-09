package absyn

import "github.com/MeteorKL/tiger/symbol"

type Ty interface {
	Node
	A_ty()
}

func (*NameTy) A_ty()   {}
func (*RecordTy) A_ty() {}
func (*ArrayTy) A_ty()  {}

type NameTy struct {
	Pos  Pos
	Name symbol.Symbol
}
type RecordTy struct {
	Pos    Pos
	Record FieldList
}
type ArrayTy struct {
	Pos   Pos
	Array symbol.Symbol
}

/* Linked lists and nodes of lists */
type NametyList *NametyList_
type NametyList_ struct {
	Head Namety
	Tail NametyList
}

type Namety *Namety_
type Namety_ struct {
	Name symbol.Symbol
	Ty   Ty
}

func NametyListInsert(head Namety, tail NametyList) NametyList {
	return &NametyList_{
		Head: head,
		Tail: tail,
	}
}
