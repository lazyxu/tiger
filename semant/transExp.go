package semant

import (
	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/env"
	"github.com/MeteorKL/tiger/symbol"
	"github.com/MeteorKL/tiger/table"
	"github.com/MeteorKL/tiger/translate"
	"github.com/MeteorKL/tiger/types"
	"github.com/MeteorKL/tiger/util"
	"github.com/MeteorKL/tiger/yacc"
)

func transExp(level translate.Level, breakk translate.Exp, venv table.Table, tenv table.Table, e absyn.Exp) expTy {
	if e == nil {
		util.Debug("NoExp")
		return expTy{translate.NoExp(), &types.Tyvoid}
	}
	// var Var translate.Exp
	// var ExpTy expTy

	switch e.(type) {
	case *absyn.VarExp:
		util.Debug("VarExp")
		e := e.(*absyn.VarExp)
		return transVar(level, breakk, venv, tenv, e.Var)
	case *absyn.NilExp:
		util.Debug("NilExp")
		return expTy{translate.NilExp(), &types.Tynil}
	case *absyn.IntExp:
		util.Debug("IntExp")
		e := e.(*absyn.IntExp)
		return expTy{translate.IntExp(e.Int), &types.Tyint}
	case *absyn.StringExp:
		util.Debug("StringExp")
		e := e.(*absyn.StringExp)
		return expTy{translate.StringExp(e.String), &types.Tystring}
	case *absyn.CallExp:
		util.Debug("CallExp")
		e := e.(*absyn.CallExp)
		FunCall := symbol.Look(venv, e.Func)
		if FunCall, ok := FunCall.(*env.FunEntry); !ok {
			yacc.EM_error(e.Pos, "Function '"+e.Func.Name+"' is undefined!\n")
		} else {
			var ArgumentList translate.ExpList = nil
			formalList := FunCall.Formals
			Arguments := e.Args
			for Arguments != nil {
				translate.ExpList_prepend(transExp(level, breakk, venv, tenv, Arguments.Head).exp, &ArgumentList)
				Arguments = Arguments.Tail
			}
			Exp := translate.CallExp(FunCall.Label, FunCall.Level, level, &ArgumentList)
			Arguments = e.Args

			for Arguments != nil && formalList != nil {
				tempArg := transExp(level, breakk, venv, tenv, Arguments.Head)
				if !equalTy(tempArg.ty, formalList.Head, e.Pos) {
					yacc.EM_error(e.Pos, "Function parameters type failed to match in '"+e.Func.Name+"'\n")
					return expTy{Exp, &types.Tyvoid}
				}
				Arguments = Arguments.Tail
				formalList = formalList.Tail
			}
			if Arguments != nil && formalList == nil {
				yacc.EM_error(e.Pos, "Function '"+e.Func.Name+"' parameter redundant!\n")
			} else if Arguments == nil && formalList != nil {
				yacc.EM_error(e.Pos, "Function '"+e.Func.Name+"' parameter insufficient!\n")
			} else if FunCall.Result == nil {
				yacc.EM_error(e.Pos, "Function '"+e.Func.Name+"' return type undefined.\n")
			} else {
				return expTy{Exp, getActualTy(FunCall.Result)}
			}
		}
		return expTy{translate.NoExp(), &types.Tyvoid}
	case *absyn.OpExp:
		util.Debug("OpExp")
		e := e.(*absyn.OpExp)
		leftExp := transExp(level, breakk, venv, tenv, e.Left)
		rightExp := transExp(level, breakk, venv, tenv, e.Right)
		OpExp := translate.NoExp()
		switch e.Oper {
		case absyn.PlusOp:
			fallthrough
		case absyn.MinusOp:
			fallthrough
		case absyn.TimesOp:
			fallthrough
		case absyn.DivideOp:
			if _, ok := leftExp.ty.(*types.Int_); !ok {
				yacc.EM_error(e.Pos, "Left hand side is not int type!")
			} else if _, ok := rightExp.ty.(*types.Int_); !ok {
				yacc.EM_error(e.Pos, "Right hand side is not int type!")
			} else {
				OpExp = translate.ArithExp(e.Oper, leftExp.exp, rightExp.exp)
			}
		case absyn.EqOp:
			fallthrough
		case absyn.NeqOp:
			equalTy(leftExp.ty, rightExp.ty, e.Pos)
		case absyn.LtOp:
			fallthrough
		case absyn.LeOp:
			fallthrough
		case absyn.GtOp:
			fallthrough
		case absyn.GeOp:
			switch leftExp.ty.(type) {
			case *types.Int_:
				if _, ok := rightExp.ty.(*types.Int_); !ok {
					yacc.EM_error(e.Pos, "Right hand side type is expected to be int.")
				} else {
					OpExp = translate.RelExp(e.Oper, leftExp.exp, rightExp.exp)
				}
			case *types.String_:
				if _, ok := rightExp.ty.(*types.String_); !ok {
					yacc.EM_error(e.Pos, "Right hand side type is expected to be string.")
				} else {
					OpExp = translate.EqStringExp(e.Oper, leftExp.exp, rightExp.exp)
				}
			default:
				yacc.EM_error(e.Pos, "Unexpected type of left hand side.")
			}
		default:
			util.Assert(!true)
		}
		return expTy{OpExp, &types.Tyint}
	case *absyn.RecordExp:
		util.Debug("RecordExp")
		e := e.(*absyn.RecordExp)
		recordTy := getActualTy(symbol.Look(tenv, e.Typ).(types.Ty))
		if recordTy == nil {
			yacc.EM_error(e.Pos, "Array type '"+e.Typ.Name+"' is undefined!")
		} else if _, ok := recordTy.(*types.Record_); !ok {
			yacc.EM_error(e.Pos, "'"+e.Typ.Name+".' is not a record type!")
		} else {
			recordTy := recordTy.(*types.Record_)
			var RecordExp expTy
			tyList := recordTy.Record
			fieldList := e.Efields
			for tyList != nil && fieldList != nil {
				RecordExp = transExp(level, breakk, venv, tenv, fieldList.Head.Exp)
				if !equalTy(tyList.Head.Ty, RecordExp.ty, e.Pos) || tyList.Head.Name != fieldList.Head.Name {
					yacc.EM_error(e.Pos, "Type of '"+e.Typ.Name+"' in record does not match definition.")
					return expTy{translate.NoExp(), &types.Record_{nil}}
				}
				tyList = tyList.Tail
				fieldList = fieldList.Tail
			}
			if tyList != nil && fieldList == nil {
				yacc.EM_error(e.Pos, "Fields insufficient in '"+e.Typ.Name+"'.")
			} else if tyList == nil && fieldList != nil {
				yacc.EM_error(e.Pos, "Fields redundant in '"+e.Typ.Name+"'.")
			} else {
				var RecordExpList translate.ExpList = nil
				Offset := 0
				fieldList = e.Efields
				fieldList = e.Efields
				for fieldList != nil {
					translate.ExpList_prepend(transExp(level, breakk, venv, tenv, fieldList.Head.Exp).exp,
						&RecordExpList)
					fieldList = fieldList.Tail
					Offset++
				}
				return expTy{translate.RecordExp(Offset, RecordExpList), recordTy}
			}
		}
		/*return of error*/
		return expTy{translate.NoExp(), &types.Record_{nil}}
	case *absyn.SeqExp:
		util.Debug("SeqExp")
		e := e.(*absyn.SeqExp)
		var SeqExpTy expTy
		var Seq_TyList translate.ExpList = nil
		SeqExpList := e.Seq
		/*Void sequence*/
		if SeqExpList == nil {
			return expTy{translate.NoExp(), &types.Tyvoid}
		}
		for SeqExpList != nil {
			SeqExpTy = transExp(level, breakk, venv, tenv, SeqExpList.Head)
			translate.ExpList_prepend(SeqExpTy.exp, &Seq_TyList)
			SeqExpList = SeqExpList.Tail
		}
		return expTy{translate.SeqExp(Seq_TyList), SeqExpTy.ty}
	case *absyn.AssignExp:
		util.Debug("AssignExp")
		e := e.(*absyn.AssignExp)
		assignVar := transVar(level, breakk, venv, tenv, e.Var)
		AssignExp := transExp(level, breakk, venv, tenv, e.Exp)
		if !equalTy(assignVar.ty, AssignExp.ty, e.Pos) {
			yacc.EM_error(e.Pos, "Types of left and right side of assignment do not match!")
		}
		return expTy{translate.AssignExp(assignVar.exp, AssignExp.exp), &types.Tyvoid}
	case *absyn.IfExp:
		util.Debug("IfExp")
		e := e.(*absyn.IfExp)
		var testExpTy, thenExpTy expTy
		elseExpTy := expTy{nil, nil}
		test_exp := e.Test
		then_exp := e.Then
		else_exp := e.Else
		// access := translate.AllocLocal(level, false)
		// tmp := translate.SimpleVar(access, level)
		// assing_true := translate.AssignExp(tmp, translate.IntExp(1))
		// assing_false := translate.AssignExp(tmp, translate.IntExp(0))
		testExpTy = transExp(level, breakk, venv, tenv, test_exp)
		thenExpTy = transExp(level, breakk, venv, tenv, then_exp)
		if _, ok := testExpTy.ty.(*types.Int_); !ok {
			yacc.EM_error(e.Pos, "Test is required to be int type.")
			return expTy{translate.NoExp(), &types.Tyvoid}
		}
		if else_exp == nil {
			if _, ok := thenExpTy.ty.(*types.Void_); !ok {
				yacc.EM_error(e.Pos, "IF-ELSE: else exp should be void.")
				return expTy{translate.NoExp(), &types.Tyvoid}
			}
			return expTy{translate.IfExp(testExpTy.exp, thenExpTy.exp, nil), &types.Tyvoid}
		}
		elseExpTy = transExp(level, breakk, venv, tenv, else_exp)
		if !equalTy(thenExpTy.ty, elseExpTy.ty, e.Pos) {
			yacc.EM_error(e.Pos, "Then exp type and else exp type does not match.")
			return expTy{translate.NoExp(), &types.Tyvoid}
		}
		return expTy{translate.IfExp(testExpTy.exp, thenExpTy.exp, elseExpTy.exp), thenExpTy.ty}
	case *absyn.WhileExp:
		util.Debug("WhileExp")
		e := e.(*absyn.WhileExp)
		whileExpTy := transExp(level, breakk, venv, tenv, e.Test)
		if _, ok := whileExpTy.ty.(*types.Int_); ok {
			yacc.EM_error(e.Pos, "While test should be int type.")
		}
		return expTy{translate.WhileExp(
			whileExpTy.exp,
			transExp(level, translate.DoneExp(), venv, tenv, e.Body).exp,
			translate.DoneExp(),
		), &types.Tyvoid}
	case *absyn.ForExp:
		util.Debug("ForExp")
		e := e.(*absyn.ForExp)
		/*Transform for into while*/
		i_Dec := &absyn.VarDec{e.Pos, e.Var, nil, e.Lo, true}
		limit_Dec := &absyn.VarDec{e.Pos, symbol.New("limit"), nil, e.Hi, true}
		test_Dec := &absyn.VarDec{e.Pos, symbol.New("test"), nil, &absyn.IntExp{e.Pos, 1}, true}
		for_Dec := &absyn.DecList_{i_Dec, &absyn.DecList_{limit_Dec, &absyn.DecList_{test_Dec, nil}}}

		i_Exp := &absyn.VarExp{e.Pos, &absyn.SimpleVar{e.Pos, e.Var}}
		then_Exp := &absyn.AssignExp{e.Pos, &absyn.SimpleVar{e.Pos, e.Var},
			&absyn.OpExp{e.Pos, absyn.PlusOp, i_Exp, &absyn.IntExp{e.Pos, 1}}}
		limit_Exp := &absyn.VarExp{e.Pos, &absyn.SimpleVar{e.Pos, symbol.New("limit")}}
		else_Exp := &absyn.AssignExp{e.Pos, &absyn.SimpleVar{e.Pos, symbol.New("test")}, &absyn.IntExp{e.Pos, 0}}
		limit_test_Exp := &absyn.IfExp{e.Pos, &absyn.OpExp{e.Pos, absyn.LtOp, i_Exp, limit_Exp}, then_Exp, else_Exp}
		forSeq_Exp := &absyn.SeqExp{e.Pos, &absyn.ExpList_{e.Body, &absyn.ExpList_{limit_test_Exp, nil}}}
		test_Exp := &absyn.VarExp{e.Pos, &absyn.SimpleVar{e.Pos, symbol.New("test")}}
		while_Exp := &absyn.WhileExp{e.Pos, test_Exp, forSeq_Exp}
		if_Exp := &absyn.IfExp{e.Pos, &absyn.OpExp{e.Pos, absyn.LeOp, e.Lo, e.Hi}, while_Exp, nil}
		ifSeq_Exp := &absyn.SeqExp{e.Pos, &absyn.ExpList_{if_Exp, nil}}
		let_Exp := &absyn.LetExp{e.Pos, for_Dec, ifSeq_Exp}

		return transExp(level, breakk, venv, tenv, let_Exp)
	case *absyn.BreakExp:
		util.Debug("BreakExp")
		// e := e.(*absyn.BreakExp)
		if breakk == nil {
			return expTy{translate.NoExp(), &types.Tyvoid}
		}
		return expTy{translate.BreakExp(breakk), &types.Tyvoid}
	case *absyn.LetExp:
		util.Debug("LetExp")
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
		util.Debug("ArrayExp")
		e := e.(*absyn.ArrayExp)
		arrayTy := getActualTy(symbol.Look(tenv, e.Typ).(types.Ty))
		if arrayTy == nil {
			yacc.EM_error(e.Pos, "Array type '"+e.Typ.Name+"' is undefined!")
		} else if arrayTy, ok := arrayTy.(*types.Array_); !ok {
			yacc.EM_error(e.Pos, "'"+e.Typ.Name+"' is not a array type!")
		} else {
			sizeExpTy := transExp(level, breakk, venv, tenv, e.Size)
			initExpTy := transExp(level, breakk, venv, tenv, e.Init)
			if _, ok := sizeExpTy.ty.(*types.Int_); !ok {
				yacc.EM_error(e.Pos, "Array size '"+e.Typ.Name+"' should be int type!")
			} else if !equalTy(initExpTy.ty, arrayTy.Array, e.Pos) {
				yacc.EM_error(e.Pos, "Array type '"+e.Typ.Name+"' doesn't match!")
			} else {
				return expTy{translate.ArrayExp(sizeExpTy.exp, initExpTy.exp), arrayTy}
			}
		}
		/*return of error*/
		return expTy{translate.NoExp(), &types.Tyint}
	}
	// util.Assert(!true)
	return expTy{translate.NoExp(), &types.Tyvoid}
}
