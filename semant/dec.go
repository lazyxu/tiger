package semant

import (
	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/table"
	"github.com/MeteorKL/tiger/translate"
)

func transDec(level translate.Level, breakk translate.Exp, venv table.Table, tenv table.Table, d absyn.Dec) translate.Exp {
	switch d.(type) {
	case *absyn.FunctionDec:
		println("FunctionDec")
	case *absyn.VarDec:
		println("VarDec")
	case *absyn.TypeDec:
		println("TypeDec")
	}
	return translate.NoExp()
}
