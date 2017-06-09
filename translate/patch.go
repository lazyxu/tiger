package translate

import "github.com/MeteorKL/tiger/temp"

type PatchList *PatchList_
type PatchList_ struct {
	Head *temp.Label
	Tail PatchList
}

func doPatch(pList PatchList, label temp.Label) {
	for ; pList != nil; pList = pList.Tail {
		*(pList.Head) = label
	}
}

func joinPatch(fList PatchList, sList PatchList) PatchList {
	if fList == nil {
		return sList
	}
	for fList.Tail != nil {
		fList = fList.Tail
	}
	fList.Tail = sList
	return fList
}
