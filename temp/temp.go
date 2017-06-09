package temp

import (
	"strconv"

	"github.com/MeteorKL/tiger/table"
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
	temps++
	r := strconv.Itoa(p.Num)
	Enter(Name(), p, r)
	return p
}

func Look(m Map, t Temp) interface{} {
	util.Assert(m != nil && m.Tab != nil)
	s := table.Look(m.Tab, t)
	if s != nil {
		return s
	}
	if m.Under != nil {
		return Look(m.Under, t)
	}
	return nil
}
