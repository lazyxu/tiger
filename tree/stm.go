package tree

import (
	"github.com/MeteorKL/tiger/temp"
)

type Stm interface {
	T_stm()
}

func (*SEQ_) T_stm()   {}
func (*LABEL_) T_stm() {}
func (*JUMP_) T_stm()  {}
func (*CJUMP_) T_stm() {}
func (*MOVE_) T_stm()  {}
func (*EXP_) T_stm()   {}

type SEQ *SEQ_
type SEQ_ struct {
	Left, Right Stm
}
type LABEL *LABEL_
type LABEL_ struct {
	Label temp.Label
}
type JUMP *JUMP_
type JUMP_ struct {
	Exp   Exp
	Jumps temp.LabelList
}
type CJUMP *CJUMP_
type CJUMP_ struct {
	Op          RelOp
	Left, Right Exp
	True, False temp.Label
}
type MOVE *MOVE_
type MOVE_ struct {
	Des, Src Exp
}
type EXP *EXP_
type EXP_ struct {
	Exp Exp
}

type StmList *StmList_
type StmList_ struct {
	Head Stm
	Tail StmList
}
