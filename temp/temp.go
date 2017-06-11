package temp

import (
	"strconv"
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
	println("new temp")
	p := new(Temp_)
	p.Num = temps
	temps++
	r := strconv.Itoa(p.Num)
	Enter(GetTempMap(), p, r)
	return p
}
