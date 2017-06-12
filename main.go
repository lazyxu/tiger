package main

import (
	"bufio"
	"os"

	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/frame"
	"github.com/MeteorKL/tiger/semant"
	"github.com/MeteorKL/tiger/tree"
	"github.com/MeteorKL/tiger/util"
	"github.com/MeteorKL/tiger/yacc"
)

//go doc -u xxx 查看某个函数的头文件
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

	frags := semant.SEM_transProg(absyn_root)

	tree.Node_count = 1

	var f *os.File
	var err error
	filepath := args[1] + ".ir"
	f, err = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	util.PanicErr(err)
	w := bufio.NewWriter(f)

	_, err = w.WriteString("digraph G{\n\tnode [shape = record,height=.1];\n")
	util.PanicErr(err)
	for ; frags != nil; frags = frags.Tail {
		switch frags.Head.(type) {
		case *frame.ProcFrag_:
			f := frags.Head.(*frame.ProcFrag_)
			tree.PPrintStm(w, f.Body)
		}
	}

	_, err = w.WriteString("}\n")
	util.PanicErr(err)
	w.Flush()
	f.Close()

	util.Visualization(args[1] + ".ir")
}
