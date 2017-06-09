package temp

import (
	"bytes"
	"encoding/binary"

	"github.com/MeteorKL/tiger/table"
)

type Map *Map_
type Map_ struct {
	Tab   table.Table
	Under Map
}

func Empty() Map {
	return &Map_{table.Empty(), nil}
}

func Enter(m Map, t Temp, s string) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, t)
	table.Enter(m.Tab, t, s)
}

var m Map = nil

func Name() Map {
	if m == nil {
		m = &Map_{table.Empty(), nil}
	}
	return m
}

func LayerMap(over Map, under Map) Map {
	if over == nil {
		return under
	}
	return &Map_{over.Tab, LayerMap(over.Under, under)}
}
