package semant

import (
	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/env"
	"github.com/MeteorKL/tiger/frame"
	"github.com/MeteorKL/tiger/translate"
	"github.com/MeteorKL/tiger/types"
)

type expTy struct {
	exp translate.Exp
	ty  types.Ty
}

func SEM_transProg(exp absyn.Exp) frame.FragList {
	tenv := env.Base_tenv()
	venv := env.Base_venv()
	transExp(translate.Outermost(), nil, venv, tenv, exp)

	return translate.GetResult()
}
