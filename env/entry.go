package env

import (
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/translate"
	"github.com/MeteorKL/tiger/types"
)

type Entry interface {
	E_enventry()
}

func (*VarEntry) E_enventry() {}
func (*FunEntry) E_enventry() {}

type VarEntry struct {
	Access translate.Access
	Ty     types.Ty
}

type FunEntry struct {
	Level   translate.Level
	Label   temp.Label
	Formals types.TyList
	Result  types.Ty
}
