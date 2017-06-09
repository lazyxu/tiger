package translate

import (
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/tree"
	"github.com/MeteorKL/tiger/util"
)

type Exp interface {
	Tr_exp()
}

func (*Ex_) Tr_exp() {}
func (*Nx_) Tr_exp() {}
func (*Cx_) Tr_exp() {}

type Ex *Ex_
type Ex_ struct {
	Ex tree.Exp
}
type Nx *Nx_
type Nx_ struct {
	Nx tree.Stm
}
type Cx *Cx_
type Cx_ struct {
	Trues  PatchList
	Falses PatchList
	Stm    tree.Stm
}

type ExpList *ExpList_
type ExpList_ struct {
	Head Exp
	Tail ExpList
}

//convert other kind to Ex
func unEx(e Exp) tree.Exp {
	switch e.(type) {
	case *Ex_:
		e := e.(*Ex_)
		return e.Ex
	case *Nx_:
		e := e.(*Nx_)
		return &tree.ESEQ_{e.Nx, &tree.CONST_{0}}
	case *Cx_:
		e := e.(*Cx_)
		r := temp.Newtemp()
		t := temp.Newlabel()
		f := temp.Newlabel()
		doPatch(e.Trues, t)
		doPatch(e.Falses, f)
		return &tree.ESEQ_{&tree.MOVE_{&tree.TEMP_{r}, &tree.CONST_{1}},
			&tree.ESEQ_{e.Stm,
				&tree.ESEQ_{&tree.LABEL_{f},
					&tree.ESEQ_{&tree.MOVE_{&tree.TEMP_{r}, &tree.CONST_{0}},
						&tree.ESEQ_{&tree.LABEL_{t}, &tree.TEMP_{r}}}}}}
	default:
		util.Assert(!true)
	}
	return nil
}

//convert other kind to Nx
func unNx(e Exp) tree.Stm {
	switch e.(type) {
	case *Ex_:
		e := e.(*Ex_)
		return &tree.EXP_{e.Ex}
	case *Nx_:
		e := e.(*Nx_)
		return e.Nx
	case *Cx_:
		e := e.(*Cx_)
		r := temp.Newtemp()
		t := temp.Newlabel()
		f := temp.Newlabel()
		doPatch(e.Trues, t)
		doPatch(e.Falses, f)
		return &tree.EXP_{&tree.ESEQ_{&tree.MOVE_{&tree.TEMP_{r}, &tree.CONST_{1}},
			&tree.ESEQ_{e.Stm,
				&tree.ESEQ_{&tree.LABEL_{f},
					&tree.ESEQ_{&tree.MOVE_{&tree.TEMP_{r}, &tree.CONST_{0}},
						&tree.ESEQ_{&tree.LABEL_{t}, &tree.TEMP_{r}}}}}}}

	default:
		util.Assert(!true)
	}
	return nil
}

//convert other kind to Cx
func unCx(e Exp) Cx {
	switch e.(type) {
	case *Ex_:
		e := e.(*Ex_)
		cx := new(Cx_)
		cx.Stm = &tree.CJUMP_{tree.Eq, e.Ex, &tree.CONST_{0}, nil, nil}
		stm := cx.Stm.(*tree.CJUMP_)
		cx.Trues = &PatchList_{&(stm.True), nil}
		cx.Falses = &PatchList_{&(stm.False), nil}
		return cx
	case *Nx_:
		util.Assert(!true) // no such case
	case *Cx_:
		e := e.(*Cx_)
		return e
	default:
		util.Assert(!true)
	}
	return nil
}
