package translate

import (
	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/frame"
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/tree"
	"github.com/MeteorKL/tiger/util"
)

// level is the accessed level, access is the position of the variable
//
func SimpleVar(access Access, level Level) Exp {
	var fp tree.Exp
	fp = &tree.TEMP_{frame.FP()}
	for l := level; l != access.Level.Parent; l = l.Parent {
		static_link := level.Frame.Formals.Head
		fp = frame.Exp(static_link, fp)
	}
	return &Ex_{frame.Exp(access.Access, fp)}
}

func FieldVar(base Exp, offset int) Exp {
	addr := &tree.BINOP_{tree.Plus, unEx(base), &tree.CONST_{offset * frame.WORD_SIZE}}
	value := &tree.MEM_{addr}
	return &Ex_{value}
}

func SubscriptVar(base Exp, index Exp) Exp {
	addr := &tree.BINOP_{tree.Plus, unEx(base), &tree.BINOP_{tree.Mul, unEx(index), &tree.CONST_{frame.WORD_SIZE}}}
	newaddr := &tree.BINOP_{tree.Plus, addr, &tree.CONST_{0}}
	value := &tree.MEM_{newaddr}
	return &Ex_{value}
}

func ArithExp(op absyn.Oper, left_exp Exp, right_exp Exp) Exp {
	var t_binop tree.BinOp
	switch op {
	case absyn.PlusOp:
		t_binop = tree.Plus
		break
	case absyn.MinusOp:
		t_binop = tree.Minus
		break
	case absyn.TimesOp:
		t_binop = tree.Mul
		break
	case absyn.DivideOp:
		t_binop = tree.Div
		break
	default:
		break
	}
	return &Ex_{&tree.BINOP_{t_binop, unEx(left_exp), unEx(right_exp)}}
}

//default signed int
func RelExp(op absyn.Oper, left_exp Exp, right_exp Exp) Exp {
	var t_relop tree.RelOp
	switch op {
	case absyn.EqOp:
		t_relop = tree.Eq
		break
	case absyn.NeqOp:
		t_relop = tree.Ne
		break
	case absyn.LtOp:
		t_relop = tree.Lt
		break
	case absyn.LeOp:
		t_relop = tree.Le
		break
	case absyn.GtOp:
		t_relop = tree.Gt
		break
	case absyn.GeOp:
		t_relop = tree.Ge
		break
	default:
		util.Assert(!true) /* should never happen*/
	}
	stm := &tree.CJUMP_{t_relop, unEx(left_exp), unEx(right_exp), nil, nil}

	trues := &PatchList_{&stm.True, nil}
	falses := &PatchList_{&stm.False, nil}
	return &Cx_{trues, falses, stm}
}

func IntExp(i int) Exp {
	int_exp := &tree.CONST_{i}
	return &Ex_{int_exp}
}

func NoExp() Exp {
	no_exp := &tree.CONST_{0}
	return &Ex_{no_exp}
}

var nil_temp temp.Temp = nil

func NilExp() Exp {
	if nil_temp == nil {
		nil_temp = temp.Newtemp()
		t_malloc := &tree.MOVE_{&tree.TEMP_{nil_temp}, frame.ExternalCall("malloc", &tree.ExpList_{&tree.CONST_{0}, nil})}
		return &Ex_{&tree.ESEQ_{t_malloc, &tree.TEMP_{nil_temp}}}
	}
	return &Ex_{&tree.TEMP_{nil_temp}}
}

func ExpList_prepend(head Exp, tr_explist *ExpList) {
	tr_list := &ExpList_{head, nil}
	tr_list.Tail = *tr_explist
	*tr_explist = tr_list
}

//function call
func CallExp(fun_name temp.Label, fun_def Level, fun_call Level, arg_list *ExpList) Exp {
	var fp tree.Exp
	fp = &tree.TEMP_{frame.FP()}
	for fun_call != nil && fun_call != fun_def.Parent {
		f_acc := fun_call.Frame.Formals.Head
		fp = frame.Exp(f_acc, fp)
		fun_call = fun_call.Parent
	}
	addr := &Ex_{fp}
	ExpList_prepend(addr, arg_list)

	var args tree.ExpList = nil
	var list tree.ExpList = nil
	for *arg_list != nil {
		node := unEx((*arg_list).Head)
		if args != nil {
			list.Tail = &tree.ExpList_{node, nil}
			list = list.Tail
		} else {
			args = &tree.ExpList_{node, nil}
			list = args
		}
		(*arg_list) = (*arg_list).Tail
	}

	return &Ex_{&tree.CALL_{&tree.NAME_{fun_name}, args}}
}

func RecordExp(size int, initList ExpList) Exp {
	r := temp.Newtemp()
	args := &tree.ExpList_{&tree.CONST_{size * frame.WORD_SIZE}, nil}
	t_malloc := &tree.MOVE_{&tree.TEMP_{r}, frame.ExternalCall("malloc", args)}
	var t_seq tree.Stm
	t_seq = &tree.MOVE_{&tree.MEM_{&tree.BINOP_{tree.Plus, &tree.TEMP_{r}, &tree.CONST_{(size - 1) * frame.WORD_SIZE}}}, unEx(initList.Head)}
	var t_mv tree.Stm
	initList = initList.Tail
	for i := size - 1; initList != nil; i-- {
		t_mv = &tree.MOVE_{&tree.MEM_{&tree.BINOP_{tree.Plus, &tree.TEMP_{r}, &tree.CONST_{(i - 1) * frame.WORD_SIZE}}}, unEx(initList.Head)}
		t_seq = &tree.SEQ_{t_mv, t_seq}
		initList = initList.Tail
	}
	return &Ex_{&tree.ESEQ_{&tree.SEQ_{t_malloc, t_seq}, &tree.TEMP_{r}}}
}

func ArrayExp(size Exp, init Exp) Exp {
	args := &tree.ExpList_{unEx(size), &tree.ExpList_{unEx(init), nil}}
	func_call := frame.ExternalCall("initArray", args)
	return &Ex_{func_call}
}

func IfExp(e1 Exp, e2 Exp, e3 Exp) Exp {
	t := temp.Newlabel()
	f := temp.Newlabel()
	join := temp.Newlabel()
	c1 := unCx(e1)
	doPatch(c1.Trues, t)
	doPatch(c1.Falses, f)

	if e3 != nil {
		var e2_stm, e3_stm tree.Stm
		join_stm := &tree.JUMP_{&tree.NAME_{join}, &temp.LabelList_{join, nil}}
		switch e2.(type) {
		case *Ex_:
			e2 := e2.(*Ex_)
			e2_stm = &tree.EXP_{e2.Ex}
			break
		case *Nx_:
			e2 := e2.(*Nx_)
			e2_stm = e2.Nx
			break
		case *Cx_:
			e2 := e2.(*Cx_)
			e2_stm = e2.Stm
			break
		}
		switch e3.(type) {
		case *Ex_:
			e3 := e3.(*Ex_)
			e3_stm = &tree.EXP_{e3.Ex}
			break
		case *Nx_:
			e3 := e3.(*Nx_)
			e3_stm = e3.Nx
			break
		case *Cx_:
			e3 := e3.(*Cx_)
			e3_stm = e3.Stm
			break
		}
		return &Nx_{&tree.SEQ_{c1.Stm, &tree.SEQ_{&tree.LABEL_{t}, &tree.SEQ_{e2_stm, &tree.SEQ_{join_stm, &tree.SEQ_{&tree.LABEL_{f}, &tree.SEQ_{e3_stm,
			&tree.SEQ_{join_stm,
				&tree.LABEL_{join}}}}}}}}}
	} else {
		switch e2.(type) {
		case *Ex_:
			{
				e2 := e2.(*Ex_)
				return &Nx_{&tree.SEQ_{c1.Stm, &tree.SEQ_{&tree.LABEL_{t}, &tree.SEQ_{&tree.EXP_{unEx(e2)}, &tree.LABEL_{f}}}}}
				break
			}
		case *Nx_:
			{
				e2 := e2.(*Nx_)
				return &Nx_{&tree.SEQ_{c1.Stm, &tree.SEQ_{&tree.LABEL_{t}, &tree.SEQ_{e2.Nx, &tree.LABEL_{f}}}}}
				break
			}
		case *Cx_:
			{
				e2 := e2.(*Cx_)
				return &Nx_{&tree.SEQ_{c1.Stm, &tree.SEQ_{&tree.LABEL_{t}, &tree.SEQ_{e2.Stm, &tree.LABEL_{f}}}}}
				break
			}

		}

	}
	return nil
}

func DoneExp() Exp {
	done_label := &tree.NAME_{temp.Newlabel()}
	return &Ex_{done_label}
}

func WhileExp(test Exp, body Exp, done Exp) Exp {
	test_label := temp.Newlabel()
	body_label := temp.Newlabel()
	return &Ex_{&tree.ESEQ_{&tree.JUMP_{&tree.NAME_{test_label}, &temp.LabelList_{test_label, nil}},
		&tree.ESEQ_{&tree.LABEL_{body_label},
			&tree.ESEQ_{unNx(body),
				&tree.ESEQ_{&tree.LABEL_{test_label},
					&tree.ESEQ_{&tree.CJUMP_{tree.Eq, unEx(test), &tree.CONST_{0}, unEx(done).(*tree.NAME_).NAME,
						body_label},
						&tree.ESEQ_{&tree.LABEL_{unEx(done).(*tree.NAME_).NAME}, &tree.CONST_{0}}}}}}}}
}

func EseqExp(e1 Exp, e2 Exp) Exp {
	return &Ex_{&tree.ESEQ_{unNx(e1), unEx(e2)}}
}

func AssignExp(left_value Exp, value Exp) Exp {
	assign_exp := &tree.MOVE_{unEx(left_value), unEx(value)}
	return &Nx_{assign_exp}
}

func BreakExp(done_label Exp) Exp {
	jump_stm := &tree.JUMP_{&tree.NAME_{unEx(done_label).(*tree.NAME_).NAME}, &temp.LabelList_{unEx(done_label).(*tree.NAME_).NAME, nil}}
	return &Nx_{jump_stm}
}

func SeqExp(tr_explist ExpList) Exp {
	seq := unEx(tr_explist.Head)
	tr_explist = tr_explist.Tail
	for tr_explist != nil {
		seq = &tree.ESEQ_{&tree.EXP_{unEx(tr_explist.Head)}, seq}
		tr_explist = tr_explist.Tail
	}
	return &Ex_{seq}
}

func EqStringExp(op absyn.Oper, left_exp Exp, right_exp Exp) Exp {
	arg_list := &tree.ExpList_{unEx(left_exp), &tree.ExpList_{unEx(right_exp), nil}}
	res := frame.ExternalCall("stringEqual", arg_list)
	switch op {
	case absyn.EqOp:
		return &Ex_{res}
	case absyn.NeqOp:
		switch res.(type) {
		case *tree.CONST_:
			res := res.(*tree.CONST_)
			if res.CONST == 1 {
				return &Ex_{&tree.CONST_{0}}
			}
		}
		return &Ex_{&tree.CONST_{1}}
	default:
		return nil
	}
}

func EqExp(op absyn.Oper, left Exp, right Exp) Exp {
	var opp tree.RelOp
	if op == absyn.EqOp {
		opp = tree.Eq
	} else {
		opp = tree.Ne
	}
	cond := &tree.CJUMP_{opp, unEx(left), unEx(right), nil, nil}
	trues := &PatchList_{&cond.True, nil}
	falses := &PatchList_{&cond.False, nil}
	return &Cx_{trues, falses, cond}
}

func EqRef(op absyn.Oper, left Exp, right Exp) Exp {
	var opp tree.RelOp
	if op == absyn.EqOp {
		opp = tree.Eq
	} else {
		opp = tree.Ne
	}
	cond := &tree.CJUMP_{opp, unEx(left), unEx(right), nil, nil}
	trues := &PatchList_{&cond.True, nil}
	falses := &PatchList_{&cond.False, nil}
	return &Cx_{trues, falses, cond}
}

var fragList frame.FragList = nil //proc list
func ProcEntryExit(level Level, body Exp, formals AccessList) {
	proc_frag := &frame.ProcFrag_{unNx(body), level.Frame}
	fragList = &frame.FragList_{proc_frag, fragList}
}

var stringFragList frame.FragList = nil //string list
func StringExp(s string) Exp {
	label := temp.Newlabel()
	string_frag := &frame.StringFrag_{label, s}
	stringFragList = &frame.FragList_{string_frag, stringFragList}
	return &Ex_{&tree.NAME_{label}}
}

func GetResult() frame.FragList {
	var frag_tail frame.FragList
	frag := stringFragList
	for frag != nil {
		frag_tail = frag
		frag = frag.Tail
	}
	if frag_tail != nil {
		frag_tail.Tail = fragList
	}
	if stringFragList != nil {
		return stringFragList
	}
	return fragList
}
