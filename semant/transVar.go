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

func transVar(level translate.Level, breakk translate.Exp, venv table.Table, tenv table.Table, v absyn.Var) expTy {
	if v == nil {
		println("NoExp")
		return expTy{translate.NoExp(), &types.Tyvoid}
	}
	var Var translate.Exp
	// var ExpTy expTy

	switch v.(type) {
	case *absyn.SimpleVar:
		println("SimpleVar")
		v := v.(*absyn.SimpleVar)
		Env := symbol.Look(venv, v.Simple)
		switch Env.(type) {
		case *env.VarEntry:
			Env := Env.(*env.VarEntry)
			Var = translate.SimpleVar(Env.Access, level)
			return expTy{Var, getActualTy(Env.Ty)}
		}
		Var = translate.NoExp()
		yacc.EM_error(v.Pos, "Undefined variable: "+v.Simple.Name)
		break
	case *absyn.FieldVar:
		println("FieldVar")
		v := v.(*absyn.FieldVar)
		ExpTy := transVar(level, breakk, venv, tenv, v.Var)
		Var = translate.NoExp()
		if ty, ok := ExpTy.ty.(*types.Record_); !ok {
			yacc.EM_error(v.Pos, "'"+v.Sym.Name+"' is not a record type.")
			break
		} else {
			Offset := 0
			fieldList := ty.Record
			for fieldList != nil {
				if fieldList.Head.Name == v.Sym {
					Var = translate.FieldVar(ExpTy.exp, Offset)
					return expTy{Var, getActualTy(fieldList.Head.Ty)}
				}
				fieldList = fieldList.Tail
				Offset++
			}
			yacc.EM_error(v.Pos, "Can't find field '"+v.Sym.Name+"' in record type")
			break
		}
	case *absyn.SubscriptVar:
		println("SubscriptVar")
		v := v.(*absyn.SubscriptVar)
		ExpTy := transVar(level, breakk, venv, tenv, v.Var)
		Var := translate.NoExp()
		if ty, ok := ExpTy.ty.(*types.Array_); !ok {
			yacc.EM_error(v.Pos, "Not a array type")
			break
		} else {
			ExpTy_Subscript := transExp(level, breakk, venv, tenv, v.Exp)
			if _, ok := ExpTy.ty.(*types.Int_); !ok {
				yacc.EM_error(v.Pos, "Subscript is not int type!")
				break
			} else {
				Var = translate.SubscriptVar(ExpTy.exp, ExpTy_Subscript.exp)
				return expTy{Var, getActualTy(ty.Array)}
			}
		}
	}
	util.Assert(!true)
	return expTy{translate.NoExp(), &types.Tyvoid}
}
