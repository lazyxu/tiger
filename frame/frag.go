package frame

import (
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/tree"
)

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

type StringFragList *StringFragList_
type StringFragList_ struct {
	Head StringFrag
	Tail StringFragList
}

type ProcFragList *ProcFragList_
type ProcFragList_ struct {
	Head ProcFrag
	Tail ProcFragList
}
