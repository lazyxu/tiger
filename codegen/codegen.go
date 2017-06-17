package codegen

import (
	"bytes"
	"fmt"

	"github.com/MeteorKL/tiger/assem"
	"github.com/MeteorKL/tiger/frame"
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/tree"
	"github.com/MeteorKL/tiger/util"
)

var Frame frame.Frame /* current function frame */

func matchOp(I tree.BinOp, Op *string) {
	switch I {
	case tree.Plus:
		*Op = "add"
	case tree.Minus:
		*Op = "sub"
	case tree.Mul:
		*Op = "mul"
	case tree.Div:
		*Op = "div"
	default:
		util.Assert(!true)
	}
}

func writeAsmStr(Str string, Args ...interface{}) string {
	buf := bytes.NewBuffer(make([]byte, 0))
	fmt.Fprintf(buf, Str, Args[0:]...)
	return string(buf.Bytes())
}

func munchExp(e tree.Exp) temp.Temp {
	// char assem_string[100];
	var asmStr string
	r := temp.Newtemp() /* return value */

	switch e.(type) {
	case *tree.BINOP_:
		e := e.(*tree.BINOP_)
		op := new(string)
		left := e.Left
		right := e.Right
		matchOp(e.Op, op)
		if l, ok := e.Left.(*tree.CONST_); ok {
			asmStr = writeAsmStr("%s $%x, `d0", *op, l.CONST)
			r = munchExp(right)
			emit(
				&assem.Oper{
					Assem: asmStr,
					Dst: &temp.TempList_{
						Head: r,
						Tail: nil},
					Src:   nil,
					Jumps: nil,
				},
			)
		} else if ri, ok := e.Right.(*tree.CONST_); ok { /* BINOP(op, e, CONST) */
			asmStr = writeAsmStr("%s $%x, `d0", *op, ri.CONST)
			r = munchExp(left)
			emit(&assem.Oper{asmStr, &temp.TempList_{r, nil}, nil, nil})
		} else { /* BINOP(op, e, e) */
			asmStr = writeAsmStr("%s `s0, `d0", *op)
			r1 := munchExp(right)
			r = munchExp(left)
			emit(&assem.Oper{asmStr, &temp.TempList_{r1, nil}, &temp.TempList_{r, nil}, nil})
		}
		return r
	case *tree.MEM_:
		e := e.(*tree.MEM_)
		mem := e.Mem
		if binOp, ok := mem.(*tree.BINOP_); ok && binOp.Op == tree.Plus {
			left := binOp.Left
			right := binOp.Right
			if l, ok := left.(*tree.CONST_); ok { /* MEM(BINOP(+, CONST, e)) */
				asmStr = writeAsmStr("mov %d(`s0), `d0", l.CONST)
				emit(&assem.Move{asmStr, &temp.TempList_{r, nil}, &temp.TempList_{munchExp(right), nil}})
			} else if rig, ok := right.(*tree.CONST_); ok { /**/
				asmStr = writeAsmStr("mov %d(`s0), `d0", rig.CONST)
				emit(&assem.Move{asmStr, &temp.TempList_{r, nil}, &temp.TempList_{munchExp(left), nil}})
			} else {
				util.Assert(!true) /*??? this shouldnot occur */
			}
		} else if c, ok := mem.(*tree.CONST_); ok { /* MEM(CONST) */
			asmStr = writeAsmStr("mov ($0x%x), `d0", c.CONST)
			emit(&assem.Move{asmStr, &temp.TempList_{r, nil}, nil})
		} else { /* MEM(e) */
			emit(&assem.Move{"mov (`s0), `d0", &temp.TempList_{r, nil}, &temp.TempList_{munchExp(mem), nil}})
		}
		return r
	case *tree.TEMP_:
		e := e.(*tree.TEMP_)
		return e.TEMP
	case *tree.ESEQ_:
		e := e.(*tree.ESEQ_)
		munchStm(e.Stm)
		return munchExp(e.Exp)
	case *tree.NAME_:
		e := e.(*tree.NAME_)
		temp.Enter(frame.TempMap(), r, e.NAME.Name)
		return r
	case *tree.CONST_:
		e := e.(*tree.CONST_)
		asmStr = writeAsmStr("mov $0x%x, `d0", e.CONST)
		emit(&assem.Move{asmStr, &temp.TempList_{r, nil}, nil})
		return r
	case *tree.CALL_:
		e := e.(*tree.CALL_)
		r = munchExp(e.Fun)
		emit(&assem.Oper{"call `s0", frame.Calldefs(), &temp.TempList_{r, munchArgs(0, e.Args)}, nil})
		return r /* return value unsure */
	}
	util.Assert(!true)
	return nil
}

func ASSEM_MOVE_MEM_PLUS(asmStr *string, Dst tree.Exp, Src tree.Exp, Constt int) {
	buf := bytes.NewBuffer(make([]byte, 0))
	fmt.Fprintf(buf, "mov `s1, %d(`s0)", Constt)
	*asmStr = string(buf.Bytes())
	emit(&assem.Move{*asmStr, nil, &temp.TempList_{munchExp(Dst), &temp.TempList_{munchExp(Src), nil}}})
}

func MATCH_CMP(I tree.RelOp, Op *string) {
	switch I {
	case tree.Eq:
		*Op = "je"
	case tree.Ne:
		*Op = "jne"
	case tree.Lt:
		*Op = "jl"
	case tree.Gt:
		*Op = "jg"
	case tree.Le:
		*Op = "jle"
	case tree.Ge:
		*Op = "jge"
	default:
		util.Assert(!true)
	}
}

func munchStm(s tree.Stm) {
	// char assem_string[100];
	var asmStr string

	switch s.(type) {
	case *tree.MOVE_:
		s := s.(*tree.MOVE_)
		dst := s.Dst
		src := s.Src
		switch dst.(type) {
		case *tree.MEM_:
			dst := dst.(*tree.MEM_)
			if binOp, ok := dst.Mem.(*tree.BINOP_); ok && binOp.Op == tree.Plus {
				if right, ok := binOp.Right.(*tree.CONST_); ok { /* MOVE (MEM(BINOP(+, e, CONST)), e) */
					ASSEM_MOVE_MEM_PLUS(&asmStr, binOp.Left, src, right.CONST)
				}
				if left, ok := binOp.Left.(*tree.CONST_); ok { /* MOVE (MEM(BINOP(+, CONST, e)), e) */
					ASSEM_MOVE_MEM_PLUS(&asmStr, binOp.Right, src, left.CONST)
				}
			} else if c, ok := dst.Mem.(*tree.CONST_); ok { /* MOVE(MEM(CONST), e) */
				asmStr = writeAsmStr("mov `s0, (%d)", c.CONST)
				emit(&assem.Move{asmStr, nil, &temp.TempList_{munchExp(src), nil}})
			} else if src, ok := src.(*tree.MEM_); ok { /* MOVE(MEM(e), MEM(e)) */
				emit(&assem.Move{"mov `s1, (`s0)", nil, &temp.TempList_{munchExp(dst.Mem), &temp.TempList_{munchExp(src.Mem), nil}}})
			} else { /* MOVE(MEM(e), e) */
				emit(&assem.Move{"mov `s1, (`s0)", nil, &temp.TempList_{munchExp(dst.Mem), &temp.TempList_{munchExp(src), nil}}})
			}
		case *tree.TEMP_: /* MOVE(TEMP(e), e) */
			emit(&assem.Move{"mov `s0, `d0", &temp.TempList_{munchExp(dst), nil}, &temp.TempList_{munchExp(src), nil}})
		}
	case *tree.SEQ_:
		s := s.(*tree.SEQ_)
		munchStm(s.Left)
		munchStm(s.Right)
	case *tree.LABEL_:

		s := s.(*tree.LABEL_)
		asmStr = writeAsmStr("%s", s.Label.Name)
		emit(&assem.Label{asmStr, s.Label})
	case *tree.JUMP_:
		s := s.(*tree.JUMP_)
		r := munchExp(s.Exp)
		emit(&assem.Oper{"jmp `d0", &temp.TempList_{r, nil}, nil, &assem.Targets_{s.Jumps}})
	case *tree.CJUMP_:
		s := s.(*tree.CJUMP_)
		var cmp string
		left := munchExp(s.Left)
		right := munchExp(s.Right)
		emit(&assem.Oper{"cmp `s0, `s1", nil, &temp.TempList_{left, &temp.TempList_{right, nil}}, nil})
		MATCH_CMP(s.Op, &cmp)
		asmStr = writeAsmStr("%s `j0", cmp)
		emit(&assem.Oper{asmStr, nil, nil, &assem.Targets_{&temp.LabelList_{s.True, nil}}})
	case *tree.EXP_:
		s := s.(*tree.EXP_)
		munchExp(s.Exp)
	default:
		util.Assert(!true)
	}
}

var reg_names = [...]string{"eax", "ebx", "ecx", "edx", "edi", "esi"}
var reg_count int = 0

func munchArgs(i int, args tree.ExpList /*, F_accessList formals*/) temp.TempList {
	/* pass params to function
	 * actually use all push stack, no reg pass paras
	 */

	/* get args register-list */
	if args == nil {
		return nil
	}

	tlist := munchArgs(i+1, args.Tail)
	rarg := munchExp(args.Head)
	// char assem_string[100];
	// var asmStr string

	emit(&assem.Oper{"push `s0", nil, &temp.TempList_{rarg, nil}, nil})
	return &temp.TempList_{rarg, tlist}
}

var instrList assem.InstructionList = nil
var last assem.InstructionList = nil

func emit(instr assem.Instruction) {
	if instrList == nil {
		instrList = &assem.InstructionList_{instr, nil}
		last = instrList
	} else {
		last.Tail = &assem.InstructionList_{instr, nil}
		last = last.Tail
	}
}

func Codegen(f frame.Frame, s tree.StmList) assem.InstructionList {
	/* interface */
	var al assem.InstructionList = nil
	sl := s
	Frame = f
	for ; sl != nil; sl = sl.Tail {
		munchStm(sl.Head)
	}
	al = instrList
	last = nil
	instrList = last
	return al
}
