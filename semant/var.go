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
			return expTy{Var, actual_ty(Env.Ty)}
		}
		Var = translate.NoExp()
		yacc.EM_error(v.Pos, "Undefined variable: "+v.Simple.Name)
		break
	case *absyn.FieldVar:
		println("FieldVar")
	case *absyn.SubscriptVar:
		println("SubscriptVar")
	}
	util.Assert(!true)
	return expTy{translate.NoExp(), &types.Tyvoid}
}
