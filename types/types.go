package types

import "github.com/MeteorKL/tiger/symbol"

type Ty interface {
	Ty_ty()
}

func (*Record_) Ty_ty() {}
func (*Nil_) Ty_ty()    {}
func (*Int_) Ty_ty()    {}
func (*String_) Ty_ty() {}
func (*Name_) Ty_ty()   {}
func (*Array_) Ty_ty()  {}
func (*Void_) Ty_ty()   {}

var Tynil Nil_
var Tyint Int_
var Tystring String_
var Tyvoid Void_

type Record *Record_
type Record_ struct {
	Record FieldList
}
type Nil *Nil_
type Nil_ struct {
}
type Int *Int_
type Int_ struct {
}
type String *String_
type String_ struct {
}
type Array *Array_
type Array_ struct {
	Array Ty
}
type Name *Name_
type Name_ struct {
	Sym symbol.Symbol
	Ty  Ty
}
type Void *Void_
type Void_ struct {
}

type Field *Field_
type Field_ struct {
	Name symbol.Symbol
	Ty   Ty
}

type FieldList *FieldList_
type FieldList_ struct {
	Head Field
	Tail FieldList
}

type TyList *TyList_
type TyList_ struct {
	Head Ty
	Tail TyList
}
