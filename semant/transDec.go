package semant

import (
	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/env"
	"github.com/MeteorKL/tiger/symbol"
	"github.com/MeteorKL/tiger/table"
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/translate"
	"github.com/MeteorKL/tiger/types"
	"github.com/MeteorKL/tiger/util"
	"github.com/MeteorKL/tiger/yacc"
)

func transDec(level translate.Level, breakk translate.Exp, venv table.Table, tenv table.Table, d absyn.Dec) translate.Exp {
	switch d.(type) {
	case *absyn.FunctionDec:
		util.Debug("FunctionDec")
		d := d.(*absyn.FunctionDec)
		fundecList := d.Function
		var paramFieldList absyn.FieldList
		var formalTyLists types.TyList
		var resultTy types.Ty
		for fundecList != nil {
			/*--------------------------------------Examine return type------------------------------------------*/
			if fundecList.Head.Result == nil {
				resultTy = &types.Tyvoid
			} else {
				resultTy = symbol.Look(tenv, fundecList.Head.Result).(types.Ty)
				if resultTy == nil {
					yacc.EM_error(fundecList.Head.Pos, "Return type '"+fundecList.Head.Result.Name+"' is undefined.")
					resultTy = &types.Tyvoid
				}
			}
			/*----------------------------Push function decelaration into venv----------------------------------*/
			var paramTy types.Ty
			var formalTyList types.TyList = nil
			formalTyLists = formalTyList
			paramFieldList = fundecList.Head.Params
			for paramFieldList != nil {
				paramTy = symbol.Look(tenv, paramFieldList.Head.Typ).(types.Ty)
				if paramTy == nil {
					yacc.EM_error(paramFieldList.Head.Pos, "Parameter type '"+paramFieldList.Head.Typ.Name+"' is undefined!")
					paramTy = &types.Tyint
				} else if formalTyList == nil {
					formalTyList = &types.TyList_{paramTy, nil}
					formalTyLists = formalTyList
				} else {
					formalTyList.Tail = &types.TyList_{paramTy, nil}
					formalTyList = formalTyList.Tail
				}
				paramFieldList = paramFieldList.Tail
			}

			/*------------------------------------------Create new level---------------------------------------------*/
			label := temp.Newlabel()
			var boolList_head util.BoolList = nil
			var boolList_tail util.BoolList = nil
			paramFieldList = fundecList.Head.Params
			for paramFieldList != nil {
				if boolList_head != nil {
					boolList_tail.Tail = &util.BoolList_{true, nil}
					boolList_tail = boolList_tail.Tail
				} else {
					boolList_head = &util.BoolList_{true, nil}
					boolList_tail = boolList_head
				}
				paramFieldList = paramFieldList.Tail
			}
			newLevel := translate.NewLevel(level, label, boolList_head)
			symbol.Enter(venv, fundecList.Head.Name, &env.FunEntry{newLevel, label, formalTyLists, resultTy})
			fundecList = fundecList.Tail
		}
		/*--------------------------------------Function body part--------------------------------------------------*/
		var formalTy types.TyList
		fundecList = d.Function
		for fundecList != nil {
			funEntry := symbol.Look(venv, fundecList.Head.Name).(*env.FunEntry)
			/*-----------------------------------symbol.BeginScope---------------------------------------------*/
			symbol.BeginScope(venv)

			formalTyLists = funEntry.Formals
			formalTy = formalTyLists
			paramFieldList = fundecList.Head.Params
			accessList := funEntry.Level.Formals
			var bodyExpTy expTy
			for accessList != nil && paramFieldList != nil && formalTy != nil {
				symbol.Enter(venv, paramFieldList.Head.Name, &env.VarEntry{accessList.Head, formalTy.Head})
				accessList = accessList.Tail
				paramFieldList = paramFieldList.Tail
				formalTy = formalTy.Tail
			}
			// 这一行好像是多余的
			funName := symbol.Look(venv, fundecList.Head.Name).(*env.FunEntry)
			bodyExpTy = transExp(funEntry.Level, breakk, venv, tenv, fundecList.Head.Body)
			if !equalTy(funName.Result, bodyExpTy.ty, fundecList.Head.Pos) {
				yacc.EM_error(fundecList.Head.Pos, "Return type of function '"+fundecList.Head.Name.Name+"' is incorrect.")
			}
			translate.ProcEntryExit(funEntry.Level, bodyExpTy.exp, accessList)
			symbol.EndScope(venv)
			/*--------------------------------------symbol.EndScope------------------------------------------*/
			fundecList = fundecList.Tail
		}
		return translate.NoExp()
	case *absyn.VarDec:
		util.Debug("VarDec")
		d := d.(*absyn.VarDec)
		initExpTy := transExp(level, breakk, venv, tenv, d.Init)
		access := translate.AllocLocal(level, d.Escape)
		if d.Typ != nil { /*var with type decelaration, e.g.: var a:int := 0*/
			varTy := symbol.Look(tenv, d.Typ).(types.Ty)
			// util.Debug(getActualTy(varTy).(*types.Int_))
			// util.Debug(getActualTy(initExpTy.ty).(*types.Int_))
			if varTy == nil {
				yacc.EM_error(d.Pos, "Type '"+d.Typ.Name+"' is undefined.")
			} else if !equalTy(varTy, initExpTy.ty, d.Pos) {
				yacc.EM_error(d.Pos, "Type '"+d.Typ.Name+"' doesn't match definition.")
				symbol.Enter(venv, d.Var, &env.VarEntry{access, varTy})
			} else {
				symbol.Enter(venv, d.Var, &env.VarEntry{access, varTy})
			}
		} else { /*var with our type decelaration, e.g.: var a:= 0.*/
			switch initExpTy.ty.(type) {
			case *types.Void_:
				yacc.EM_error(d.Pos, "Initial value type is invalid for '"+d.Var.Name+"'")
				symbol.Enter(venv, d.Var, &env.VarEntry{access, &types.Tyint})
			case *types.Nil_:
				yacc.EM_error(d.Pos, "Initial value type is invalid for '"+d.Var.Name+"'")
				symbol.Enter(venv, d.Var, &env.VarEntry{access, &types.Tyint})
			default:
				symbol.Enter(venv, d.Var, &env.VarEntry{access, initExpTy.ty})
			}
		}
		return translate.AssignExp(translate.SimpleVar(access, level), initExpTy.exp)
	case *absyn.TypeDec:
		util.Debug("TypeDec")
		d := d.(*absyn.TypeDec)
		var nameTy types.Ty
		Cycle := true
		typeList := d.Type
		for typeList != nil {
			symbol.Enter(tenv, typeList.Head.Name, &types.Name_{typeList.Head.Name, nil})
			typeList = typeList.Tail
		}
		typeList = d.Type
		for typeList != nil {
			typeTy := transTy(tenv, typeList.Head.Ty)
			if Cycle {
				if _, ok := typeTy.(*types.Name_); !ok {
					Cycle = false
				}
			}
			nameTy = symbol.Look(tenv, typeList.Head.Name).(types.Ty)
			nameTy.(*types.Name_).Ty = typeTy
			typeList = typeList.Tail
		}
		if Cycle {
			yacc.EM_error(d.Pos, "Illegal Cycle.")
		}
	}
	return translate.NoExp()
}
