package symbol

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/MeteorKL/tiger/table"
	"github.com/MeteorKL/tiger/util"
)

type Symbol *Symbol_
type Symbol_ struct {
	Name string
	Next Symbol
}

const SIZE = 109 /* should be prime */
var hashtable [SIZE]Symbol

func Insert(name string) Symbol {
	index := util.Hash([]byte(name)) % SIZE
	syms := hashtable[index]
	var sym Symbol
	for sym = syms; sym != nil; sym = sym.Next {
		if sym.Name == name {
			return sym
		}
	}
	sym = &Symbol_{name, syms}
	hashtable[index] = sym
	return sym
}

func Empty() table.Table {
	return table.Empty()
}

func Enter(t table.Table, sym Symbol, value interface{}) {
	table.Enter(t, sym, value)
}

func Look(t table.Table, sym Symbol) interface{} {
	return table.Look(t, sym)
}

var marksym Symbol_ = Symbol_{"<mark>", nil}

func BeginScope(t table.Table) {
	Enter(t, &marksym, nil)
}

func EndScope(t table.Table) {
	s := table.Pop(t)
	for s != nil {
		v := reflect.ValueOf(s).Elem().UnsafeAddr()
		buf1 := bytes.NewBuffer(make([]byte, 0))
		fmt.Fprintf(buf1, "%d", v)
		buf2 := bytes.NewBuffer(make([]byte, 0))
		fmt.Fprintf(buf2, "%d", &marksym)
		println("&marksym: ", &marksym)
		println("v: ", v)
		println("buf1: ", string(buf1.Bytes()))
		println("buf2: ", string(buf2.Bytes()))
		if string(buf1.Bytes()) == string(buf2.Bytes()) {
			break
		}
		s = table.Pop(t)
	}
}
