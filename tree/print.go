package tree

import (
	"bufio"
	"os"
	"strconv"

	"github.com/MeteorKL/tiger/symbol"
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/util"
)

var bin_oper = [...]string{
	"PLUS", "MINUS", "TIMES", "DIVIDE",
	"AND", "OR", "LSHIFT", "RSHIFT", "ARSHIFT", "XOR",
}

var rel_oper = [...]string{
	"EQ", "NE", "LT", "GT", "LE", "GE", "ULT", "ULE", "UGT", "UGE",
}

var Node_count int

func printStm(w *bufio.Writer, stm Stm) {
	id := Node_count
	Node_count++
	switch stm.(type) {
	case *SEQ_:
		stm := stm.(*SEQ_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"SEQ|<1>left|<2>right\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(Node_count) + "\n")
		printStm(w, stm.Left)
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(Node_count) + "\n")
		printStm(w, stm.Right)
		break
	case *LABEL_:
		stm := stm.(*LABEL_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"LABEL\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + " -> " + strconv.Itoa(Node_count) + "\n")
		w.WriteString("\t" + strconv.Itoa(Node_count) + "[label=\"" + stm.Label.Name + "\"]\n")
		Node_count++
		break
	case *JUMP_:
		stm := stm.(*JUMP_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"JUMP|<1>exp\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(Node_count) + "\n")
		printExp(w, stm.Exp)
		break
	case *CJUMP_:
		stm := stm.(*CJUMP_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"CJUMP|<1>op|<2>left|<3>right|<4>true|<5>false\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(Node_count) + "\n")
		w.WriteString("\t" + strconv.Itoa(Node_count) + "[label=\"" + rel_oper[stm.Op] + "\"]\n")
		Node_count++
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(Node_count) + "\n")
		printExp(w, stm.Left)
		w.WriteString("\t" + strconv.Itoa(id) + ":3 -> " + strconv.Itoa(Node_count) + "\n")
		printExp(w, stm.Right)
		w.WriteString("\t" + strconv.Itoa(id) + ":4 -> " + strconv.Itoa(Node_count) + "\n")
		w.WriteString("\t" + strconv.Itoa(Node_count) + "[label=\"" + stm.True.Name + "\"]\n")
		Node_count++
		w.WriteString("\t" + strconv.Itoa(id) + ":5 -> " + strconv.Itoa(Node_count) + "\n")
		w.WriteString("\t" + strconv.Itoa(Node_count) + "[label=\"" + stm.False.Name + "\"]\n")
		Node_count++
		break
	case *MOVE_:
		stm := stm.(*MOVE_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"MOVE|<1>dst|<2>src\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(Node_count) + "\n")
		printExp(w, stm.Dst)
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(Node_count) + "\n")
		printExp(w, stm.Src)
		break
	case *EXP_:
		stm := stm.(*EXP_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"EXP|<1>exp\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(Node_count) + "\n")
		printExp(w, stm.Exp)
		break
	}
}
func printExpList(w *bufio.Writer, args ExpList) {
	id := Node_count
	Node_count++
	w.WriteString("\t" + strconv.Itoa(id) + "[label=\"expList|<1>head|<2>tail\"]\n")
	w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(Node_count) + "\n")
	printExp(w, args.Head)
	if args.Tail != nil {
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(Node_count) + "\n")
		printExpList(w, args.Tail)
	}
}
func printExp(w *bufio.Writer, exp Exp) {
	id := Node_count
	Node_count++
	switch exp.(type) {
	case *BINOP_:
		exp := exp.(*BINOP_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"BINOP|<1>op|<2>left|<3>right\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(Node_count) + "\n")
		w.WriteString("\t" + strconv.Itoa(Node_count) + "[label=\"" + bin_oper[exp.Op] + "\"]\n")
		Node_count++
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(Node_count) + "\n")
		printExp(w, exp.Left)
		w.WriteString("\t" + strconv.Itoa(id) + ":3 -> " + strconv.Itoa(Node_count) + "\n")
		printExp(w, exp.Right)
		break
	case *MEM_:
		exp := exp.(*MEM_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"MEM\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + " -> " + strconv.Itoa(Node_count) + "\n")
		printExp(w, exp.Mem)
		break
	case *TEMP_:
		exp := exp.(*TEMP_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"TEMP\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + " -> " + strconv.Itoa(Node_count) + "\n")
		w.WriteString("\t" + strconv.Itoa(Node_count) + "[label=\"" + temp.Look(temp.GetTempMap(), exp.TEMP) + "\"]\n")
		Node_count++
		break
	case *ESEQ_:
		exp := exp.(*ESEQ_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"ESEQ|<1>stm|<2>exp\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(Node_count) + "\n")
		printStm(w, exp.Stm)
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(Node_count) + "\n")
		printExp(w, exp.Exp)
		break
	case *NAME_:
		exp := exp.(*NAME_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"NAME\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + " -> " + strconv.Itoa(Node_count) + "\n")
		w.WriteString("\t" + strconv.Itoa(Node_count) + "[label=\"" + symbol.Symbol(exp.NAME).Name + "\"]\n")
		Node_count++
		break
	case *CONST_:
		exp := exp.(*CONST_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"CONST\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + " -> " + strconv.Itoa(Node_count) + "\n")
		w.WriteString("\t" + strconv.Itoa(Node_count) + "[label=\"" + strconv.Itoa(exp.CONST) + "\"]\n")
		Node_count++
		break
	case *CALL_:
		exp := exp.(*CALL_)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"CALL|<1>fun|<2>args\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(Node_count) + "\n")
		printExp(w, exp.Fun)
		if exp.Args != nil {
			w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(Node_count) + "\n")
			printExpList(w, exp.Args)
		}
		break
	} /* end of switch */
}

func PrintExp(filepath string, exp Exp) {
	var f *os.File
	var err error

	f, err = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	util.PanicErr(err)
	w := bufio.NewWriter(f)

	_, err = w.WriteString("digraph G{\n\tnode [shape = record,height=.1];\n")
	util.PanicErr(err)
	Node_count = 1
	printExp(w, exp)
	_, err = w.WriteString("}\n")
	util.PanicErr(err)

	w.Flush()
	f.Close()
}

func PrintStm(filepath string, stm Stm) {
	if util.CheckFileExist(filepath) {
		os.Remove(filepath)
	}
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	util.PanicErr(err)
	w := bufio.NewWriter(f)

	_, err = w.WriteString("digraph G{\n\tnode [shape = record,height=.1];\n")
	util.PanicErr(err)
	Node_count = 1
	printStm(w, stm)
	_, err = w.WriteString("}\n")
	util.PanicErr(err)

	w.Flush()
	f.Close()
}

func PPrintStm(w *bufio.Writer, stm Stm) {
	printStm(w, stm)
}

func PrintStmList(filepath string, s StmList) {
	if util.CheckFileExist(filepath) {
		os.Remove(filepath)
	}
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	util.PanicErr(err)
	w := bufio.NewWriter(f)

	_, err = w.WriteString("digraph G{\n\tnode [shape = record,height=.1];\n")
	util.PanicErr(err)
	Node_count = 1
	for ; s != nil; s = s.Tail {
		printStm(w, s.Head)
	}
	_, err = w.WriteString("}\n")
	util.PanicErr(err)

	w.Flush()
	f.Close()
}

func PPrintStmList(w *bufio.Writer, s StmList) {
	if s == nil {
		return
	}
	for ; s != nil; s = s.Tail {
		printStm(w, s.Head)
	}
}
