package table

import (
	"bytes"
	"fmt"

	"github.com/MeteorKL/tiger/util"
)

const TABSIZE uint = 127

type binder *binder_
type binder_ struct {
	key     interface{}
	value   interface{}
	next    binder
	prevtop interface{}
}
type Table *Table_
type Table_ struct {
	Table [TABSIZE]binder
	Top   interface{}
}

func Empty() Table {
	t := new(Table_)
	t.Top = nil
	var i uint
	for i = 0; i < TABSIZE; i++ {
		t.Table[i] = nil
	}
	return t
}

func Enter(t Table, key interface{}, value interface{}) {
	util.Assert(t != nil && key != nil)
	buf := bytes.NewBuffer(make([]byte, 0))
	fmt.Fprintf(buf, "%d", key)
	// println("Enter: ", string(buf.Bytes()))
	index := util.Hash(buf.Bytes()) % TABSIZE
	t.Table[index] = &binder_{key, value, t.Table[index], t.Top}
	t.Top = key
}

func Look(t Table, key interface{}) interface{} {
	util.Assert(t != nil && key != nil)
	buf := bytes.NewBuffer(make([]byte, 0))
	fmt.Fprintf(buf, "%d", key)
	index := util.Hash(buf.Bytes()) % TABSIZE
	for b := t.Table[index]; b != nil; b = b.next {
		if b.key == key {
			return b.value
		}
	}
	return nil
}

func Pop(t Table) interface{} {
	util.Assert(t != nil)
	key := t.Top
	util.Assert(key != nil)
	buf := bytes.NewBuffer(make([]byte, 0))
	fmt.Fprintf(buf, "%d", key)
	// println("Pop: ", string(buf.Bytes()))
	index := util.Hash(buf.Bytes()) % TABSIZE
	b := t.Table[index]
	util.Assert(b != nil)
	t.Table[index] = b.next
	t.Top = b.prevtop
	return b.key
}
