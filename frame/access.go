package frame

import (
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/util"
)

type Access interface {
	F_access()
}

func (*FrameAccess) F_access() {}
func (*RegAccess) F_access()   {}

type FrameAccess struct {
	Offset int
}
type RegAccess struct {
	Reg temp.Temp
}

type AccessList *AccessList_
type AccessList_ struct {
	Head Access
	Tail AccessList
}

const WORD_SIZE int = 4

// formals
// false: RegAccess
// true:  FrameAccess
func makeFormalAccessList(formals util.BoolList) AccessList {
	formal := formals
	accessList := new(AccessList_)
	var accessListHead AccessList = nil
	var access Access = nil
	InFrame_cnt := 0
	for formal != nil {
		if !formal.Head {
			access = &RegAccess{temp.Newtemp()}
		} else {
			//extra one space for control link
			access = &FrameAccess{(1 + InFrame_cnt) * WORD_SIZE}
			InFrame_cnt++
		}

		if accessListHead != nil {
			accessList.Tail = &AccessList_{access, nil}
			accessList = accessList.Tail
		} else {
			accessListHead = &AccessList_{access, nil}
			accessList = accessListHead
		}
		formal = formal.Tail
	}
	return accessListHead
}
