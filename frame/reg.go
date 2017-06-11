package frame

import (
	"github.com/MeteorKL/tiger/temp"
)

var tempMap temp.Map = nil

func TempMap() temp.Map {
	if tempMap != nil {
		return tempMap
	}
	tempMap = temp.Empty()
	return tempMap
}

func initReg(reg *temp.Temp, name string) temp.Temp {
	if *reg != nil {
		return *reg
	}
	*reg = temp.Newtemp()
	temp.Enter(TempMap(), *reg, name)
	return *reg
}

var fp temp.Temp = nil

func FP() temp.Temp {
	return initReg(&fp, "ebp")
}

var sp temp.Temp = nil

func SP() temp.Temp {
	return initReg(&sp, "esp")
}

var rv temp.Temp = nil

func RV() temp.Temp {
	return initReg(&rv, "eax")
}

var ra temp.Temp = nil

func RA() temp.Temp {
	return initReg(&ra, "---")
}
