package canon

import (
	"github.com/MeteorKL/tiger/symbol"
	"github.com/MeteorKL/tiger/table"
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/tree"
	"github.com/MeteorKL/tiger/util"
)

/*
 * canon.c - Functions to convert the IR trees into basic blocks and traces.
 * ESEQ & CALL & SEQ &CJUMP may raise trouble (because must compute by order)
 */

type ExpRefList *ExpRefList_
type ExpRefList_ struct {
	Head *tree.Exp
	Tail ExpRefList
}

/* is *tree.CONST */
func isNop(x tree.Stm) bool {
	switch x.(type) {
	case *tree.EXP_:
		x := x.(*tree.EXP_)
		switch x.Exp.(type) {
		case *tree.CONST_:
			return true
		}

	}
	return false
}

/* simple from seq return a stm and remove the unuse item */
func seq(x tree.Stm, y tree.Stm) tree.Stm {
	if isNop(x) {
		return y
	}
	if isNop(y) {
		return x
	}
	return &tree.SEQ_{x, y}
}

/* is alter-able */
// 判断一个 Stm 是否可与一个 Exp 交换
func commute(x tree.Stm, y tree.Exp) bool {

	if isNop(x) {
		return true
	}
	switch y.(type) {
	case *tree.NAME_:
		return true
	case *tree.CONST_:
		return true

	}
	return false
}

type StmExp struct {
	s tree.Stm
	e tree.Exp
}

/* return stm that in sub-exp */
func reorder(rlist ExpRefList) tree.Stm {
	if rlist == nil {
		return &tree.EXP_{&tree.CONST_{0}} /* nop */
	}
	if _, ok := (*rlist.Head).(*tree.CALL_); ok {
		t := temp.Newtemp()
		(*rlist.Head) = &tree.ESEQ_{&tree.MOVE_{&tree.TEMP_{t}, (*rlist.Head)}, &tree.TEMP_{t}}
		return reorder(rlist)
	}

	hd := do_exp((*rlist.Head))
	s := reorder(rlist.Tail)
	/* the real change part */
	if commute(s, hd.e) {
		(*rlist.Head) = hd.e
		return seq(hd.s, s)
	}
	t := temp.Newtemp()
	(*rlist.Head) = &tree.TEMP_{t}
	return seq(hd.s, seq(&tree.MOVE_{&tree.TEMP_{t}, hd.e}, s))
}

func get_Call_rlist(exp tree.Exp) ExpRefList {
	e := exp.(*tree.CALL_)
	args := e.Args
	rlist := &ExpRefList_{&e.Fun, nil}
	curr := rlist
	for ; args != nil; args = args.Tail {
		curr.Tail = &ExpRefList_{&args.Head, nil}
		curr = curr.Tail
	}
	return rlist
}

/* change exp-stm order stm-order */
func do_exp(exp tree.Exp) StmExp {

	switch exp.(type) {
	case *tree.BINOP_:
		exp := exp.(*tree.BINOP_)
		return StmExp{reorder(&ExpRefList_{&exp.Left, &ExpRefList_{&exp.Right, nil}}),
			exp}
	case *tree.MEM_:
		exp := exp.(*tree.MEM_)
		return StmExp{reorder(&ExpRefList_{&exp.Mem, nil}), exp}
	case *tree.ESEQ_:
		exp := exp.(*tree.ESEQ_)
		x := do_exp(exp.Exp)
		return StmExp{seq(DoStm(exp.Stm), x.s), x.e}
	case *tree.CALL_:
		exp := exp.(*tree.CALL_)
		return StmExp{reorder(get_Call_rlist(exp)), exp}
	default:
		return StmExp{reorder(nil), exp}
	}
}

/* change stm1-stm2 to stm2-stm1 */
func DoStm(stm tree.Stm) tree.Stm {

	switch stm.(type) {
	case *tree.SEQ_:
		stm := stm.(*tree.SEQ_)
		return seq(DoStm(stm.Left), DoStm(stm.Right))
	case *tree.JUMP_:
		stm := stm.(*tree.JUMP_)
		return seq(reorder(&ExpRefList_{&stm.Exp, nil}), stm)
	case *tree.CJUMP_:
		stm := stm.(*tree.CJUMP_)
		return seq(reorder(&ExpRefList_{&stm.Left, &ExpRefList_{&stm.Right, nil}}),
			stm)
	case *tree.MOVE_:
		stm := stm.(*tree.MOVE_)
		switch stm.Dst.(type) {
		case *tree.TEMP_:
			switch stm.Src.(type) {
			case *tree.CALL_:
				return seq(reorder(get_Call_rlist(stm.Src)), stm)
			}
			return seq(reorder(&ExpRefList_{&stm.Src, nil}), stm)
		case *tree.MEM_:
			dst := stm.Dst.(*tree.MEM_)
			return seq(reorder(&ExpRefList_{&dst.Mem, &ExpRefList_{&stm.Src, nil}}), stm)
		case *tree.ESEQ_:
			dst := stm.Dst.(*tree.ESEQ_)
			s := dst.Stm
			stm.Dst = dst.Exp
			return DoStm(&tree.SEQ_{s, stm})

		}
		util.Assert(!true)

	case *tree.EXP_:
		stm := stm.(*tree.EXP_)
		switch stm.Exp.(type) {
		case *tree.CALL_:
			return seq(reorder(get_Call_rlist(stm.Exp)), stm)
		}
		return seq(reorder(&ExpRefList_{&stm.Exp, nil}), stm)
	default:
		return stm
	}
	return stm
}

/* linear gets rid of the top-level SEQ's, producing a list */
func linear(stm tree.Stm, right tree.StmList) tree.StmList {
	switch stm.(type) {
	case *tree.SEQ_:
		stm := stm.(*tree.SEQ_)
		return linear(stm.Left, linear(stm.Right, right))
	}
	return &tree.StmList_{stm, right}
}

/* From an arbitrary Tree statement, produce a list of cleaned trees
   satisfying the following properties:
      1.  No SEQ's or ESEQ's
      2.  The parent of every CALL is an EXP(..) or a MOVE(TEMP t,..) */
func Linearize(stm tree.Stm) tree.StmList {
	return linear(stm, nil)
}

type StmListList *StmListList_
type Block struct {
	StmLists StmListList
	Label    temp.Label
}
type StmListList_ struct {
	Head tree.StmList
	tail StmListList
}

/* Go down a list looking for end of basic block */
func next(prevstms tree.StmList, stms tree.StmList, done temp.Label) StmListList {
	/* the end of the stmlist add the JUMP */
	if stms == nil {
		return next(prevstms,
			&tree.StmList_{&tree.JUMP_{&tree.NAME_{done}, &temp.LabelList_{done, nil}}, nil},
			done)
	}
	switch stms.Head.(type) {
	case *tree.JUMP_:
		prevstms.Tail = stms
		stmLists := mkBlocks(stms.Tail, done)
		stms.Tail = nil
		return stmLists
	case *tree.CJUMP_:
		prevstms.Tail = stms
		stmLists := mkBlocks(stms.Tail, done)
		stms.Tail = nil
		return stmLists
	case *tree.LABEL_:
		lab := stms.Head.(*tree.LABEL_).Label
		return next(prevstms,
			&tree.StmList_{&tree.JUMP_{&tree.NAME_{lab}, &temp.LabelList_{lab, nil}}, stms},
			done)

	}
	prevstms.Tail = stms
	return next(stms, stms.Tail, done)
}

/* Create the beginning of a basic block */
func mkBlocks(stms tree.StmList, done temp.Label) StmListList {
	if stms == nil {
		return nil
	}
	if _, ok := stms.Head.(*tree.LABEL_); !ok {
		return mkBlocks(&tree.StmList_{&tree.LABEL_{temp.Label(temp.Newlabel())}, stms}, done)
	}
	/* else there already is a label */
	return &StmListList_{stms, next(stms, stms.Tail, done)}
}

/* basicBlocks : Tree.stm list -> (Tree.stm list list * Tree.label)
	       From a list of cleaned trees, produce a list of
	       basic blocks satisfying the following properties:
	       1. and 2. as above;
	       3.  Every block begins with a LABEL;
           4.  A LABEL appears only at the beginning of a block;
           5.  Any JUMP or CJUMP is the last stm in a block;
           6.  Every block ends with a JUMP or CJUMP;
           Also produce the "label" to which control will be passed
           upon exit.
*/
func BasicBlocks(stmList tree.StmList) Block {
	var b Block
	b.Label = temp.Newlabel() /*done label*/
	b.StmLists = mkBlocks(stmList, b.Label)
	return b
}

var block_env table.Table
var global_block Block

func getLast(list tree.StmList) tree.StmList {
	last := list
	for last.Tail != nil && last.Tail.Tail != nil {
		last = last.Tail
	}
	return last
}

func trace(list tree.StmList) {
	last := getLast(list)
	lab := list.Head
	s := last.Tail.Head
	symbol.Enter(block_env, symbol.Symbol(lab.(*tree.LABEL_).Label), nil)
	switch s.(type) {
	case *tree.JUMP_:
		s := s.(*tree.JUMP_)
		target := symbol.Look(block_env, symbol.Symbol(s.Jumps.Head))
		if s.Jumps.Tail == nil && target != nil {
			target := target.(tree.StmList)
			last.Tail = target /* merge the 2 lists removing JUMP stm */
			trace(target)
		} else {
			last.Tail.Tail = getNext() /* merge and keep JUMP stm */
		}
	case *tree.CJUMP_:
		/* we want false label to follow CJUMP */
		s := s.(*tree.CJUMP_)
		True := symbol.Look(block_env, symbol.Symbol(s.True)).(tree.StmList)
		False := symbol.Look(block_env, symbol.Symbol(s.False)).(tree.StmList)
		if False != nil {
			last.Tail.Tail = False
			trace(False)
		} else if True != nil { /* convert so that existing label is a false label */
			last.Tail.Head = &tree.CJUMP_{tree.NotRel(s.Op), s.Left,
				s.Right, s.False,
				s.True}
			last.Tail.Tail = True
			trace(True)
		} else {
			False := temp.Newlabel()
			last.Tail.Head = &tree.CJUMP_{s.Op, s.Left,
				s.Right, s.True, False}
			last.Tail.Tail = &tree.StmList_{&tree.LABEL_{temp.Label(False)}, getNext()}
		}
	default:
		util.Assert(!true)
	}
}

/* get the next block from the list of stmLists, using only those that have
 * not been traced yet */
func getNext() tree.StmList {
	if global_block.StmLists == nil {
		return &tree.StmList_{&tree.LABEL_{temp.Label(global_block.Label)}, nil}
	} else {
		s := global_block.StmLists.Head
		if symbol.Look(block_env, symbol.Symbol(s.Head.(*tree.LABEL_).Label)) != nil { /* label exists in the table */
			trace(s)
			return s
		}
		global_block.StmLists = global_block.StmLists.tail
		return getNext()
	}
}

/* traceSchedule : Tree.stm list list * Tree.label -> Tree.stm list
            From a list of basic blocks satisfying properties 1-6,
            along with an "exit" label,
	        produce a list of stms such that:
	        1. and 2. as above;
            7. Every CJUMP(_,t,f) is immediately followed by LABEL f.
            The blocks are reordered to satisfy property 7; also
	        in this reordering as many JUMP(T.NAME(lab)) statements
            as possible are eliminated by falling through into T.LABEL(lab).
*/

func TraceSchedule(b Block) tree.StmList {
	var sList StmListList
	block_env = symbol.Empty()
	global_block = b
	for sList = global_block.StmLists; sList != nil; sList = sList.tail {
		symbol.Enter(block_env, symbol.Symbol(sList.Head.Head.(*tree.LABEL_).Label), sList.Head)
	}
	return getNext()
}
