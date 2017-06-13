package main

import (
	"bufio"
	"os"

	"strconv"

	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/assem"
	"github.com/MeteorKL/tiger/canon"
	"github.com/MeteorKL/tiger/codegen"
	"github.com/MeteorKL/tiger/frame"
	"github.com/MeteorKL/tiger/semant"
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/tree"
	"github.com/MeteorKL/tiger/util"
	"github.com/MeteorKL/tiger/yacc"
)

/* clean:rm -rf testcases/*.png testcases/*.Linearize testcases/*.ir testcases/*.ast testcases/*.TraceSchedule testcases/*.s */

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

	f, err := os.OpenFile(args[1]+".s", os.O_WRONLY|os.O_CREATE, 0666)
	util.PanicErr(err)
	w := bufio.NewWriter(f)
	for i := 0; frags != nil; frags = frags.Tail {
		switch frags.Head.(type) {
		case *frame.ProcFrag_:
			println("*frame.ProcFrag_")
			tree.Node_count = 1
			procFrag := frags.Head.(*frame.ProcFrag_)
			tree.PrintStm(args[1]+strconv.Itoa(i)+".ir", procFrag.Body)
			util.Visualization(args[1] + strconv.Itoa(i) + ".ir")

			stmList := canon.Linearize(procFrag.Body)
			tree.PrintStmList(args[1]+strconv.Itoa(i)+".Linearize", stmList)
			util.Visualization(args[1] + strconv.Itoa(i) + ".Linearize")

			stmList = canon.TraceSchedule(canon.BasicBlocks(stmList))
			tree.PrintStmList(args[1]+strconv.Itoa(i)+".TraceSchedule", stmList)
			util.Visualization(args[1] + strconv.Itoa(i) + ".TraceSchedule")

			iList := codegen.Codegen(procFrag.Frame, stmList) /* 9 */

			w.WriteString("BEGIN " + procFrag.Frame.Name.Name + "\n")
			assem.PrintInstrList(w, iList, temp.LayerMap(frame.TempMap(), temp.GetTempMap()))
			w.WriteString("END " + procFrag.Frame.Name.Name + "\n\n")

		case *frame.StringFrag_:
			println("*frame.StringFrag_")
			stringFrag := frags.Head.(*frame.StringFrag_)
			w.WriteString(stringFrag.Label.Name + ": " + stringFrag.Str + "\n")
		}
	}
	w.Flush()
	f.Close()
}
