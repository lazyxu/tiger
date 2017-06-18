package main

import (
	"bufio"
	"fmt"
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

/* clean:rm -rf testcases/*.png testcases/*.dot testcases/*.s */

//go doc -u xxx 查看某个函数的头文件
//godoc -analysis=type -http=:6060
//go:generate go tool yacc -o yacc/y.tab.go yacc/tiger.y
//go:generate go build -o tiger main.go

func main() {
	args := os.Args
	if args == nil || len(args) != 2 {
		fmt.Printf("usage: tiger file.tig")
		return
	}
	absyn_root := yacc.YYParse(args[1])
	absyn.PrintExp(args[1]+".ast", absyn_root)
	util.Visualization(args[1] + ".ast")

	procFrag, stringFrag := semant.SEM_transProg(absyn_root)

	tree.Node_count = 1
	f, err := os.OpenFile(args[1]+".s", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	util.PanicErr(err)
	w := bufio.NewWriter(f)
	w.WriteString("SECTION .data\n")
	for ; stringFrag != nil; stringFrag = stringFrag.Tail {
		util.Debug("*frame.StringFrag_")
		w.WriteString("\t" + stringFrag.Head.Label.Name + ": db " + stringFrag.Head.Str + ", 0x0a\n")
	}
	w.WriteString("\n")
	w.WriteString("SECTION .text\n")
	w.WriteString("\tglobal start\n\n")
	w.WriteString("start:\n")
	for i := 0; procFrag != nil; procFrag = procFrag.Tail {
		util.Debug("*frame.ProcFrag_")
		tree.Node_count = 1

		out := args[1] + strconv.Itoa(i)
		tree.PrintStm(out+".ir.dot", procFrag.Head.Body)
		util.Visualization(out + ".ir.dot")

		stm := canon.DoStm(procFrag.Head.Body)
		tree.PrintStm(out+".DoStm.dot", stm)
		util.Visualization(out + ".DoStm.dot")
		stmList := canon.Linearize(stm)
		tree.PrintStmList(out+".Linearize.dot", stmList)
		util.Visualization(out + ".Linearize.dot")

		BasicBlocks := canon.BasicBlocks(stmList)
		tree.PrintStmList(out+"."+BasicBlocks.Label.Name, BasicBlocks.StmLists.Head)
		util.Visualization(out + "." + BasicBlocks.Label.Name)
		stmList = canon.TraceSchedule(BasicBlocks)
		tree.PrintStmList(out+".TraceSchedule.dot", stmList)
		util.Visualization(out + ".TraceSchedule.dot")

		iList := codegen.Codegen(procFrag.Head.Frame, stmList)

		assem.PrintInstrList(w, iList, temp.LayerMap(frame.TempMap(), temp.GetTempMap()))
		i++
	}
	w.WriteString("\texit\n")
	w.Flush()
	f.Close()
}
