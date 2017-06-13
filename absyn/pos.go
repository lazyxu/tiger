package absyn

type Pos int

type Position struct {
	Line, Column int
}

var lineNum int

var linePosList LinePosList

type LinePosList *LinePosList_
type LinePosList_ struct {
	LinePos Pos
	Next    LinePosList
}

func EM_newline(pos Pos) {
	lineNum++
	linePosList = &LinePosList_{pos, linePosList}
}

func LinePosListInit() {
	lineNum = 1
	linePosList = &LinePosList_{-1, nil}
}

func PositionFor(p Pos) (pos Position) {
	lines := linePosList
	num := lineNum
	for lines != nil && lines.LinePos > p {
		lines = lines.Next
		num--
	}
	pos.Line = num
	if lines != nil {
		pos.Column = int(p - lines.LinePos)
	}
	// IfExpression()
	return
}
