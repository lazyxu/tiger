package absyn

import (
	"github.com/MeteorKL/tiger/symbol"
)

type Var interface {
	ASTNode
	A_var()
}

func (*SimpleVar) A_var()    {}
func (*FieldVar) A_var()     {}
func (*SubscriptVar) A_var() {}

type SimpleVar struct {
	Pos    Pos
	Simple symbol.Symbol
}
type FieldVar struct {
	Pos Pos
	Var Var
	Sym symbol.Symbol
}
type SubscriptVar struct {
	Pos Pos
	Var Var
	Exp Exp
}
