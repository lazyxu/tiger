package tree

import (
	"github.com/MeteorKL/tiger/temp"
)

type Exp interface {
	T_exp()
}

func (*BINOP_) T_exp() {}
func (*MEM_) T_exp()   {}
func (*TEMP_) T_exp()  {}
func (*ESEQ_) T_exp()  {}
func (*NAME_) T_exp()  {}
func (*CONST_) T_exp() {}
func (*CALL_) T_exp()  {}

type BINOP *BINOP_
type BINOP_ struct {
	Op          BinOp
	Left, Right Exp
}
type MEM *MEM_
type MEM_ struct {
	Mem Exp
}
type TEMP *TEMP_
type TEMP_ struct {
	TEMP temp.Temp
}
type ESEQ *ESEQ_
type ESEQ_ struct {
	Stm Stm
	Exp Exp
}
type NAME *NAME_
type NAME_ struct {
	NAME temp.Label
}
type CONST *CONST_
type CONST_ struct {
	CONST int
}
type CALL *CALL_
type CALL_ struct {
	Fun  Exp
	Args ExpList
}

type ExpList *ExpList_
type ExpList_ struct {
	Head Exp
	Tail ExpList
}
