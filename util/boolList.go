package util

type BoolList *BoolList_
type BoolList_ struct {
	Head bool
	Tail BoolList
}
