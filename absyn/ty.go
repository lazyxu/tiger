package absyn

import "github.com/MeteorKL/tiger_compiler/symbol"

type Ty interface {
	Node
	tyNode()
}

func (*NameTy) tyNode()   {}
func (*RecordTy) tyNode() {}
func (*ArrayTy) tyNode()  {}

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
