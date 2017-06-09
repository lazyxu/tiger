package frame

import (
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/tree"
)

type Frag interface {
	F_frag()
}

func (*StringFrag_) F_frag() {}
func (*ProcFrag_) F_frag()   {}

type StringFrag *StringFrag_
type StringFrag_ struct {
	Label temp.Label
	Str   string
}
type ProcFrag *ProcFrag_
type ProcFrag_ struct {
	Body  tree.Stm
	Frame Frame
}

type FragList *FragList_
type FragList_ struct {
	Head Frag
	Tail FragList
}
