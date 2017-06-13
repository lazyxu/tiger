package frame

import (
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/tree"
	"github.com/MeteorKL/tiger/util"
)

type Frame *Frame_
type Frame_ struct {
	LocalCount int
	Formals    AccessList
	Name       temp.Label
}

func NewFrame(name temp.Label, formals util.BoolList) Frame {
	frame := new(Frame_)
	frame.Name = name
	println("new frame: ", name.Name)
	frame.Formals = makeFormalAccessList(formals)
	frame.LocalCount = 0
	return frame
}

func AllocLocal(f Frame, escape bool) Access {
	new_count := f.LocalCount + 1
	f.LocalCount = new_count
	if escape {
		// one extra space for return
		return &FrameAccess{WORD_SIZE * (-(1 + f.LocalCount))}
	}
	return &RegAccess{temp.Newtemp()}
}

func Exp(acc Access, framePtr tree.Exp) tree.Exp {
	switch acc.(type) {
	case *FrameAccess:
		acc := acc.(*FrameAccess)
		return &tree.MEM_{&tree.BINOP_{tree.Plus, framePtr, &tree.CONST_{acc.Offset}}}
	case *RegAccess:
		acc := acc.(*RegAccess)
		return &tree.TEMP_{acc.Reg}
	}
	return nil
}

func ExternalCall(str string, args tree.ExpList) tree.Exp {
	fun_name := &tree.NAME_{temp.Namedlabel(str)}
	return &tree.CALL_{fun_name, args}
}

func callersaves() temp.TempList {
	/* assist-function of calldefs() */

	ebx := temp.Newtemp()
	ecx := temp.Newtemp()
	edx := temp.Newtemp()
	edi := temp.Newtemp()
	esi := temp.Newtemp()
	temp.Enter(TempMap(), ebx, "ebx")
	temp.Enter(TempMap(), ecx, "ecx")
	temp.Enter(TempMap(), edx, "edx")
	temp.Enter(TempMap(), edi, "edi")
	temp.Enter(TempMap(), esi, "esi")
	return &temp.TempList_{RV(), &temp.TempList_{ebx, &temp.TempList_{ecx, &temp.TempList_{edx, &temp.TempList_{edi, &temp.TempList_{esi, nil}}}}}}
}

var protected_regs temp.TempList = nil

func Calldefs() temp.TempList {
	/* some registers that may raise side-effect (caller procted, return-val-reg, return-addr-reg) */
	if protected_regs == nil {
		protected_regs = callersaves()
	}
	return protected_regs
}
