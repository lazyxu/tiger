package absyn

import (
	"bufio"
	"os"

	"strconv"

	"github.com/MeteorKL/tiger/symbol"
	"github.com/MeteorKL/tiger/util"
)

var node_count int

func printPos(w *bufio.Writer, pos Pos) {
	id := node_count
	node_count++
	position := PositionFor(pos)
	w.WriteString("\t" + strconv.Itoa(id) + "[label=\"" + strconv.Itoa(position.Line) + ":" + strconv.Itoa(position.Column) + "\"]\n")
}

func printSymbol(w *bufio.Writer, sym symbol.Symbol) {
	id := node_count
	node_count++
	w.WriteString("\t" + strconv.Itoa(id) + "[label=\"Symbol|<1>Name|<2>Next\"]\n")
	w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
	w.WriteString("\t" + strconv.Itoa(node_count) + "[label=\"" + sym.Name + "\"]\n")
	node_count++
	if sym.Next != nil {
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printSymbol(w, sym.Next)
	}
}

func printOper(w *bufio.Writer, oper Oper) {
	id := node_count
	node_count++
	w.WriteString("\t" + strconv.Itoa(id) + "[label=\"" + opnames[oper] + "\"]\n")
}

func printField(w *bufio.Writer, field Field) {
	id := node_count
	node_count++
	w.WriteString("\t" + strconv.Itoa(id) + "[label=\"Field|<1>Pos|<2>Name|<3>Typ|<4>Escape\"]\n")
	w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
	printPos(w, field.Pos)
	w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
	printSymbol(w, field.Name)
	w.WriteString("\t" + strconv.Itoa(id) + ":3 -> " + strconv.Itoa(node_count) + "\n")
	printSymbol(w, field.Typ)
	w.WriteString("\t" + strconv.Itoa(id) + ":4 -> " + strconv.Itoa(node_count) + "\n")
	if field.Escape {
		w.WriteString("\t" + strconv.Itoa(node_count) + "[label=\"true\"]\n")
	} else {
		w.WriteString("\t" + strconv.Itoa(node_count) + "[label=\"false\"]\n")
	}
	node_count++
}

func printFieldList(w *bufio.Writer, fieldList FieldList) {
	id := node_count
	node_count++
	w.WriteString("\t" + strconv.Itoa(id) + "[label=\"FieldList|<1>Head|<2>Tail\"]\n")
	w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
	printField(w, fieldList.Head)
	if fieldList.Tail != nil {
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printFieldList(w, fieldList.Tail)
	}
}

func printFundec(w *bufio.Writer, fundec Fundec) {
	id := node_count
	node_count++
	w.WriteString("\t" + strconv.Itoa(id) + "[label=\"Fundec|<1>Pos|<2>Name|<3>Params|<4>Result|<5>Body\"]\n")
	w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
	printPos(w, fundec.Pos)
	w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
	printSymbol(w, fundec.Name)
	w.WriteString("\t" + strconv.Itoa(id) + ":3 -> " + strconv.Itoa(node_count) + "\n")
	printFieldList(w, fundec.Params)
	w.WriteString("\t" + strconv.Itoa(id) + ":4 -> " + strconv.Itoa(node_count) + "\n")
	printSymbol(w, fundec.Result)
	w.WriteString("\t" + strconv.Itoa(id) + ":5 -> " + strconv.Itoa(node_count) + "\n")
	printExp(w, fundec.Body)
}

func printFundecList(w *bufio.Writer, fundecList FundecList) {
	id := node_count
	node_count++
	w.WriteString("\t" + strconv.Itoa(id) + "[label=\"FundecList|<1>Head|<2>Tail\"]\n")
	w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
	printFundec(w, fundecList.Head)
	if fundecList.Tail != nil {
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printFundecList(w, fundecList.Tail)
	}
}

func printDec(w *bufio.Writer, dec Dec) {
	id := node_count
	node_count++
	switch dec.(type) {
	case *FunctionDec:
		dec := dec.(*FunctionDec)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"FunctionDec|<1>Pos|<2>Function\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, dec.Pos)
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printFundecList(w, dec.Function)
	case *VarDec:
		dec := dec.(*VarDec)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"VarDec|<1>Pos|<2>Var|<3>Typ|<4>Init|<5>Escape\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, dec.Pos)
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printSymbol(w, dec.Var)
		w.WriteString("\t" + strconv.Itoa(id) + ":3 -> " + strconv.Itoa(node_count) + "\n")
		printSymbol(w, dec.Typ)
		w.WriteString("\t" + strconv.Itoa(id) + ":4 -> " + strconv.Itoa(node_count) + "\n")
		printExp(w, dec.Init)
		w.WriteString("\t" + strconv.Itoa(id) + ":5 -> " + strconv.Itoa(node_count) + "\n")
		if dec.Escape {
			w.WriteString("\t" + strconv.Itoa(node_count) + "[label=\"true\"]\n")
		} else {
			w.WriteString("\t" + strconv.Itoa(node_count) + "[label=\"false\"]\n")
		}
		node_count++
	case *TypeDec:
		dec := dec.(*TypeDec)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"TypeDec|<1>Pos\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, dec.Pos)
		//
	}
}

func printDecList(w *bufio.Writer, decList DecList) {
	id := node_count
	node_count++
	w.WriteString("\t" + strconv.Itoa(id) + "[label=\"DecList|<1>Head|<2>Tail\"]\n")
	w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
	printDec(w, decList.Head)
	if decList.Tail != nil {
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printDecList(w, decList.Tail)
	}
}

func printVar(w *bufio.Writer, Var Var) {
	id := node_count
	node_count++
	switch Var.(type) {
	case *SimpleVar:
		Var := Var.(*SimpleVar)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"SimpleVar|<1>Pos|<2>Var\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, Var.Pos)
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printSymbol(w, Var.Simple)
	case *FieldVar:
		Var := Var.(*FieldVar)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"FieldVar|<1>Pos\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, Var.Pos)
		//
	case *SubscriptVar:
		Var := Var.(*SubscriptVar)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"SubscriptVar|<1>Pos\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, Var.Pos)
		//
	}
}

func printExp(w *bufio.Writer, exp Exp) {
	id := node_count
	node_count++
	switch exp.(type) {
	case *VarExp:
		exp := exp.(*VarExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"VarExp|<1>Pos|<2>Var\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, exp.Pos)
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printVar(w, exp.Var)
	case *NilExp:
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"NilExp\"]\n")
	case *IntExp:
		exp := exp.(*IntExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"IntExp|" + strconv.Itoa(exp.Int) + "\"]\n")
	case *StringExp:
		exp := exp.(*StringExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"StringExp|" + exp.String + "\"]\n")
	case *CallExp:
		exp := exp.(*CallExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"CallExp|<1>Pos|<2>Func|<3>Args\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, exp.Pos)
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printSymbol(w, exp.Func)
		w.WriteString("\t" + strconv.Itoa(id) + ":3 -> " + strconv.Itoa(node_count) + "\n")
		printExpList(w, exp.Args)
	case *OpExp:
		exp := exp.(*OpExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"OpExp|<1>Pos|<2>Oper|<3>Left|<4>Right\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, exp.Pos)
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printOper(w, exp.Oper)
		w.WriteString("\t" + strconv.Itoa(id) + ":3 -> " + strconv.Itoa(node_count) + "\n")
		printExp(w, exp.Left)
		w.WriteString("\t" + strconv.Itoa(id) + ":4 -> " + strconv.Itoa(node_count) + "\n")
		printExp(w, exp.Right)
	case *RecordExp:
		exp := exp.(*RecordExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"RecordExp|<1>Pos|<2>Typ|<3>Efields\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, exp.Pos)
		// w.WriteString("\t"+ strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) +"\n")
		// printOper(w, exp.Typ)
		// w.WriteString("\t"+ strconv.Itoa(id) + ":3 -> " + strconv.Itoa(node_count) +"\n")
		// printExp(w, exp.Efields)
		//
	case *SeqExp:
		exp := exp.(*SeqExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"SeqExp|<1>Pos|<2>Oper|<3>Left|<4>Right\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, exp.Pos)
		if exp.Seq != nil {
			w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
			printExpList(w, exp.Seq)
		}
	case *AssignExp:
		exp := exp.(*AssignExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"AssignExp|<1>Pos|<2>Typ|<3>Efields\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, exp.Pos)
		// w.WriteString("\t"+ strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) +"\n")
		// printOper(w, exp.Typ)
		// w.WriteString("\t"+ strconv.Itoa(id) + ":3 -> " + strconv.Itoa(node_count) +"\n")
		// printExp(w, exp.Efields)
		//
	case *IfExp:
		exp := exp.(*IfExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"IfExp|<1>Pos|<2>Test|<3>Then|<4>Else\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, exp.Pos)
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printExp(w, exp.Test)
		w.WriteString("\t" + strconv.Itoa(id) + ":3 -> " + strconv.Itoa(node_count) + "\n")
		printExp(w, exp.Then)
		if exp.Else != nil {
			w.WriteString("\t" + strconv.Itoa(id) + ":4 -> " + strconv.Itoa(node_count) + "\n")
			printExp(w, exp.Else)
		}
	case *WhileExp:
		exp := exp.(*WhileExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"WhileExp|<1>Pos|<2>Test|<3>Body\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, exp.Pos)
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printExp(w, exp.Body)
	case *ForExp:
		exp := exp.(*ForExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"ForExp|<1>Pos\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, exp.Pos)
		//
	case *BreakExp:
		exp := exp.(*BreakExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"BreakExp|<1>Pos\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, exp.Pos)
	case *LetExp:
		exp := exp.(*LetExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"LetExp|<1>Pos|<2>Decs|<3>Body\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, exp.Pos)
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printDecList(w, exp.Decs)
		w.WriteString("\t" + strconv.Itoa(id) + ":3 -> " + strconv.Itoa(node_count) + "\n")
		printExp(w, exp.Body)
	case *ArrayExp:
		exp := exp.(*ArrayExp)
		w.WriteString("\t" + strconv.Itoa(id) + "[label=\"ArrayExp|<1>Pos\"]\n")
		w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
		printPos(w, exp.Pos)
		//
	}
}

func printExpList(w *bufio.Writer, expList ExpList) {
	id := node_count
	node_count++
	w.WriteString("\t" + strconv.Itoa(id) + "[label=\"ExpList|<1>Head|<2>Tail\"]\n")
	w.WriteString("\t" + strconv.Itoa(id) + ":1 -> " + strconv.Itoa(node_count) + "\n")
	printExp(w, expList.Head)
	if expList.Tail != nil {
		w.WriteString("\t" + strconv.Itoa(id) + ":2 -> " + strconv.Itoa(node_count) + "\n")
		printExpList(w, expList.Tail)
	}
}

func PrintExp(filepath string, exp Exp) {
	if util.CheckFileExist(filepath) {
		os.Remove(filepath)
	}
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	util.PanicErr(err)
	w := bufio.NewWriter(f)

	_, err = w.WriteString("digraph G{\n\tnode [shape = record,height=.1];\n")
	util.PanicErr(err)
	node_count = 1
	printExp(w, exp)
	_, err = w.WriteString("}\n")
	util.PanicErr(err)

	w.Flush()
	f.Close()
}
