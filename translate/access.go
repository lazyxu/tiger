package translate

import "github.com/MeteorKL/tiger/frame"

type Access *Access_
type Access_ struct {
	Level  Level
	Access frame.Access
}

type AccessList *AccessList_
type AccessList_ struct {
	Head Access
	Tail AccessList
}

func makeFormalAccessList(level Level) AccessList {
	var accessHead AccessList = nil
	var accessList AccessList
	// the first item is discarded because of static links
	frameList := level.Frame.Formals.Tail
	for frameList != nil {
		access := &Access_{level, frameList.Head}
		if accessHead != nil {
			accessList.Tail = &AccessList_{access, nil}
			accessList = accessList.Tail
		} else {
			accessHead = &AccessList_{access, nil}
			accessList = accessHead
		}
		frameList = frameList.Tail
	}
	return accessHead
}
