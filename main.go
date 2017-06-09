package main

import (
	"os"

	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/semant"
	"github.com/MeteorKL/tiger/util"
	"github.com/MeteorKL/tiger/yacc"
)

//godoc -analysis=type -http=:6060
//go:generate go tool yacc -o yacc/y.tab.go yacc/tiger.y
//go:generate go build -o tiger main.go

func main() {
	// var key interface{}
	// println(&key)
	// // s := strings.NewReader("ABCEFG")
	// // br := bufio.NewReader(s)
	// b := bytes.NewBuffer(make([]byte, 0))
	// fmt.Fprintf(b, "%d", &key)
	// i := b.Bytes()
	// fmt.Printf("%d\n", i)
	// fmt.Printf("%s\n", b)
	args := os.Args
	if args == nil || len(args) != 2 {
		println("usage: tiger file.tig")
		return
	}
	absyn_root := yacc.YYParse(args[1])
	absyn.PrintExp(args[1]+".ast", absyn_root)
	util.Visualization(args[1] + ".ast")

	semant.SEM_transProg(absyn_root)
}
