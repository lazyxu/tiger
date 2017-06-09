package temp

import "github.com/MeteorKL/tiger/symbol"
import "strconv"

type Label symbol.Symbol

type LabelList *LabelList_
type LabelList_ struct {
	Head Label
	Tail LabelList
}

var labels int = 0

func Newlabel() Label {
	buf := "L" + strconv.Itoa(labels)
	labels++
	return Label(symbol.Insert(buf))
}
func Namedlabel(s string) Label {
	return Label(symbol.Insert(s))
}
