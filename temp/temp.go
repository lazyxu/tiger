package temp

import (
	"strconv"

	"github.com/MeteorKL/tiger/util"
)

type Temp *Temp_
type Temp_ struct {
	Num int
}

type TempList *TempList_
type TempList_ struct {
	Head Temp
	Tail TempList
}

var temps int = 0

func Newtemp() Temp {
	p := new(Temp_)
	p.Num = temps
	buf := strconv.Itoa(p.Num)
	util.Debug("Newtemp: r" + buf)
	Enter(GetTempMap(), p, "r"+buf)
	temps++
	return p
}
