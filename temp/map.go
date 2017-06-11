package temp

import (
	"bytes"
	"encoding/binary"

	"github.com/MeteorKL/tiger/table"
	"github.com/MeteorKL/tiger/util"
)

type Map *Map_
type Map_ struct {
	Tab   table.Table
	Under Map
}

var m Map = nil

func Empty() Map {
	return &Map_{table.Empty(), nil}
}

func GetTempMap() Map {
	if m == nil {
		m = &Map_{table.Empty(), nil}
	}
	return m
}

func Enter(m Map, t Temp, s string) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, t)
	table.Enter(m.Tab, t, s)
}

func Look(m Map, t Temp) string {
	util.Assert(m != nil && m.Tab != nil)
	s := table.Look(m.Tab, t)
	if s != nil {
		switch s.(type) {
		case string:
			return s.(string)
		}
	}
	if m.Under != nil {
		return Look(m.Under, t)
	}
	return ""
}

func LayerMap(over Map, under Map) Map {
	if over == nil {
		return under
	}
	return &Map_{over.Tab, LayerMap(over.Under, under)}
}
