package main

import (
	"os"

	"github.com/MeteorKL/tiger_compiler/absyn"
	"github.com/MeteorKL/tiger_compiler/util"
	"github.com/MeteorKL/tiger_compiler/yacc"
)

//go:generate go tool yacc -o yacc/y.tab.go yacc/tiger.y
//go:generate go build -o tiger main.go

func main() {
	args := os.Args
	if args == nil || len(args) != 2 {
		println("usage: tiger file.tig")
		return
	}

	absyn.PrintExp(args[1]+".ast", yacc.YYParse(args[1]))
	util.Visualization(args[1] + ".ast")
}
