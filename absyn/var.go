package absyn

import (
	"github.com/MeteorKL/tiger/symbol"
)

type Var interface {
	Node
	varNode()
}

func (*SimpleVar) varNode()    {}
func (*FieldVar) varNode()     {}
func (*SubscriptVar) varNode() {}

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
