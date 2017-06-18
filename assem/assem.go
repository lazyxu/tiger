package assem

import (
	"bufio"
	"strconv"

	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/util"
)

type Targets *Targets_
type Targets_ struct {
	Labels temp.LabelList
}

type Instruction interface {
	insteructionNode()
}

func (*Oper) insteructionNode()  {}
func (*Label) insteructionNode() {}
func (*Move) insteructionNode()  {}

type Oper struct {
	Assem    string
	Dst, Src temp.TempList
	Jumps    Targets
}
type Label struct {
	Assem string
	Label temp.Label
}
type Move struct {
	Assem    string
	Dst, Src temp.TempList
}
type InstructionList *InstructionList_
type InstructionList_ struct {
	Head Instruction
	Tail InstructionList
}

/* put list b at the end of list a */
func Splice(a InstructionList, b InstructionList) InstructionList {
	var p InstructionList
	if a == nil {
		return b
	}
	for p = a; p.Tail != nil; p = p.Tail {

	}
	p.Tail = b
	return a
}

func nthTemp(list temp.TempList, i int) temp.Temp {
	util.Assert(list != nil)
	if i == 0 {
		return list.Head
	}
	return nthTemp(list.Tail, i-1)
}

func nthLabel(list temp.LabelList, i int) temp.Label {
	util.Assert(list != nil)
	if i == 0 {
		return list.Head
	}
	return nthLabel(list.Tail, i-1)
}

/* first param is string created by this function by reading 'assem' string
 * and replacing `d `s and `j stuff.
 * Last param is function to use to determine what to do with each temp.
 */
func format(
	assem string,
	dst temp.TempList,
	src temp.TempList,
	jumps Targets,
	m temp.Map) string {
	result := ""
	for index := 0; index < len(assem); index++ {
		if assem[index] == '`' {
			index++
			switch assem[index] {
			case 's': // src
				index++
				n, _ := strconv.Atoi(assem[index:])
				s := temp.Look(m, nthTemp(src, n))
				result = result + s
			case 'd': // dst
				index++
				n, _ := strconv.Atoi(assem[index:])
				s := temp.Look(m, nthTemp(dst, n))
				result = result + s
			case 'j': // jump
				index++
				util.Assert(jumps != nil)
				n, _ := strconv.Atoi(assem[index:])
				s := nthLabel(jumps.Labels, n).Name
				result = result + s
			case '`':
				result = result + "`"
			default:
				util.Assert(!true)
			}
		} else {
			result = result + string(assem[index])
		}
	}
	return result
}

func Print(w *bufio.Writer, i Instruction, m temp.Map) {
	switch i.(type) {
	case *Oper:
		i := i.(*Oper)
		s := format(i.Assem, i.Dst, i.Src, i.Jumps, m)
		w.WriteString("\t" + s + "\n")
		break
	case *Label:
		i := i.(*Label)
		s := format(i.Assem, nil, nil, nil, m)
		w.WriteString(s + ":\n")
		break
	case *Move:
		i := i.(*Move)
		s := format(i.Assem, i.Dst, i.Src, nil, m)
		w.WriteString("\t" + s + "\n")
		break
	}
}

/* c should be COL_color; temporarily it is not */
func PrintInstrList(w *bufio.Writer, iList InstructionList, m temp.Map) {
	for ; iList != nil; iList = iList.Tail {
		Print(w, iList.Head, m)
	}
}
