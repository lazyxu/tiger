package semant

import (
	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/symbol"
	"github.com/MeteorKL/tiger/table"
	"github.com/MeteorKL/tiger/translate"
	"github.com/MeteorKL/tiger/types"
)

func transExp(level translate.Level, breakk translate.Exp, venv table.Table, tenv table.Table, e absyn.Exp) expTy {
	if e == nil {
		println("NoExp")
		return expTy{translate.NoExp(), &types.Tyvoid}
	}
	// var Var translate.Exp
	// var ExpTy expTy

	switch e.(type) {
	case *absyn.VarExp:
		println("VarExp")
		e := e.(*absyn.VarExp)
		return transVar(level, breakk, venv, tenv, e.Var)
	case *absyn.NilExp:
		println("NilExp")
		// e := e.(*absyn.NilExp)
	case *absyn.IntExp:
		println("IntExp")
		// e := e.(*absyn.IntExp)
	case *absyn.StringExp:
		println("StringExp")
		// e := e.(*absyn.StringExp)
	case *absyn.CallExp:
		println("CallExp")
		// e := e.(*absyn.CallExp)
	case *absyn.OpExp:
		println("OpExp")
		// e := e.(*absyn.OpExp)
	case *absyn.RecordExp:
		println("RecordExp")
		// e := e.(*absyn.RecordExp)
	case *absyn.SeqExp:
		println("SeqExp")
		// e := e.(*absyn.SeqExp)
	case *absyn.AssignExp:
		println("AssignExp")
		// e := e.(*absyn.AssignExp)
	case *absyn.IfExp:
		println("IfExp")
		// e := e.(*absyn.IfExp)
	case *absyn.WhileExp:
		println("WhileExp")
		// e := e.(*absyn.WhileExp)
	case *absyn.ForExp:
		println("ForExp")
		// e := e.(*absyn.ForExp)
	case *absyn.BreakExp:
		println("BreakExp")
		// e := e.(*absyn.BreakExp)
	case *absyn.LetExp:
		println("LetExp")
		e := e.(*absyn.LetExp)
		decList := e.Decs
		var expList translate.ExpList = nil

		symbol.BeginScope(venv)
		symbol.BeginScope(tenv)
		for decList != nil {
			translate.ExpList_prepend(transDec(level, breakk, venv, tenv, decList.Head), &expList)
			decList = decList.Tail
		}
		letBodyExpTy := transExp(level, breakk, venv, tenv, e.Body)
		translate.ExpList_prepend(letBodyExpTy.exp, &expList)

		symbol.EndScope(venv)
		symbol.EndScope(tenv)
		if level.Parent == nil {
			translate.ProcEntryExit(level, translate.SeqExp(expList), level.Formals)
		}
		return expTy{translate.SeqExp(expList), letBodyExpTy.ty}
	case *absyn.ArrayExp:
		println("ArrayExp")
		// e := e.(*absyn.ArrayExp)
	}
	// util.Assert(!true)
	return expTy{translate.NoExp(), &types.Tyvoid}
}
