package temp

import (
	"strconv"

	"github.com/MeteorKL/tiger/symbol"
	"github.com/MeteorKL/tiger/util"
)

type Label symbol.Symbol

type LabelList *LabelList_
type LabelList_ struct {
	Head Label
	Tail LabelList
}

var labels int = 0

func Newlabel() Label {
	buf := "L" + strconv.Itoa(labels)
	util.Debug("Newlabel: L" + buf)
	labels++
	return Label(symbol.New(buf))
}

func StrToLabel(s string) Label {
	util.Debug("StrToLabel: " + s)
	return Label(symbol.New(s))
}
