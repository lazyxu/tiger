//line yacc/tiger.y:2
package yacc

import __yyfmt__ "fmt"

//line yacc/tiger.y:2
import (
	"github.com/MeteorKL/tiger_compiler/absyn"
	"github.com/MeteorKL/tiger_compiler/symbol"
)

const (
	EOF          = 0
	Debug        = 0
	ErrorVerbose = true
)

var Absyn_root absyn.Exp

//line yacc/tiger.y:15
type yySymType struct {
	yys        int
	ty         absyn.Ty
	namety     absyn.Namety
	sym        symbol.Symbol
	Var        absyn.Var
	exp        absyn.Exp
	expList    absyn.ExpList
	dec        absyn.Dec
	decList    absyn.DecList
	field      absyn.Field
	fieldList  absyn.FieldList
	fundec     absyn.Fundec
	fundecList absyn.FundecList
	nametyList absyn.NametyList
	efield     absyn.Efield
	efieldList absyn.EfieldList

	ival int
	sval string
}

const INT = 57346
const ID = 57347
const STRING = 57348
const COMMA = 57349
const COLON = 57350
const SEMICOLON = 57351
const LPAREN = 57352
const RPAREN = 57353
const LBRACK = 57354
const RBRACK = 57355
const LBRACE = 57356
const RBRACE = 57357
const DOT = 57358
const PLUS = 57359
const MINUS = 57360
const TIMES = 57361
const DIVIDE = 57362
const EQ = 57363
const NEQ = 57364
const LT = 57365
const LE = 57366
const GT = 57367
const GE = 57368
const AND = 57369
const OR = 57370
const ASSIGN = 57371
const reserved_word_beg = 57372
const WHILE = 57373
const FOR = 57374
const TO = 57375
const BREAK = 57376
const LET = 57377
const IN = 57378
const END = 57379
const FUNCTION = 57380
const VAR = 57381
const TYPE = 57382
const ARRAY = 57383
const IF = 57384
const THEN = 57385
const ELSE = 57386
const DO = 57387
const OF = 57388
const NIL = 57389
const reserved_word_end = 57390
const ILLEGAL = 57391
const LOW = 57392
const UMINUS = 57393

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"INT",
	"ID",
	"STRING",
	"COMMA",
	"COLON",
	"SEMICOLON",
	"LPAREN",
	"RPAREN",
	"LBRACK",
	"RBRACK",
	"LBRACE",
	"RBRACE",
	"DOT",
	"PLUS",
	"MINUS",
	"TIMES",
	"DIVIDE",
	"EQ",
	"NEQ",
	"LT",
	"LE",
	"GT",
	"GE",
	"AND",
	"OR",
	"ASSIGN",
	"reserved_word_beg",
	"WHILE",
	"FOR",
	"TO",
	"BREAK",
	"LET",
	"IN",
	"END",
	"FUNCTION",
	"VAR",
	"TYPE",
	"ARRAY",
	"IF",
	"THEN",
	"ELSE",
	"DO",
	"OF",
	"NIL",
	"reserved_word_end",
	"ILLEGAL",
	"LOW",
	"UMINUS",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 55,
	21, 0,
	22, 0,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 14,
	-1, 56,
	21, 0,
	22, 0,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 15,
	-1, 57,
	21, 0,
	22, 0,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 16,
	-1, 58,
	21, 0,
	22, 0,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 17,
	-1, 63,
	21, 0,
	22, 0,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 22,
	-1, 64,
	21, 0,
	22, 0,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 23,
}

const yyNprod = 72
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 337

var yyAct = [...]int{

	38, 3, 115, 114, 70, 66, 37, 47, 45, 124,
	103, 43, 51, 39, 52, 107, 40, 41, 7, 80,
	53, 54, 55, 56, 57, 58, 59, 60, 61, 62,
	63, 64, 67, 68, 122, 72, 42, 74, 25, 26,
	27, 28, 29, 30, 21, 23, 22, 24, 20, 19,
	140, 99, 71, 79, 73, 14, 81, 82, 84, 36,
	52, 49, 51, 35, 112, 100, 129, 91, 83, 134,
	85, 86, 98, 27, 28, 131, 34, 90, 94, 95,
	96, 125, 133, 93, 4, 14, 5, 97, 87, 67,
	9, 113, 104, 75, 102, 101, 127, 126, 12, 108,
	25, 26, 27, 28, 118, 14, 120, 121, 1, 69,
	50, 16, 17, 116, 13, 18, 123, 46, 109, 111,
	117, 44, 15, 130, 128, 65, 10, 6, 11, 135,
	137, 117, 2, 31, 138, 32, 8, 33, 71, 48,
	110, 141, 0, 132, 0, 117, 136, 0, 0, 0,
	0, 0, 0, 139, 25, 26, 27, 28, 29, 30,
	21, 23, 22, 24, 20, 19, 25, 26, 27, 28,
	29, 30, 21, 23, 22, 24, 20, 19, 0, 0,
	0, 0, 78, 25, 26, 27, 28, 29, 30, 21,
	23, 22, 24, 105, 25, 26, 27, 28, 29, 30,
	21, 23, 22, 24, 20, 19, 25, 26, 27, 28,
	29, 30, 21, 23, 22, 24, 20, 19, 119, 0,
	77, 0, 106, 0, 0, 0, 0, 0, 25, 26,
	27, 28, 29, 30, 21, 23, 22, 24, 20, 19,
	92, 0, 0, 0, 25, 26, 27, 28, 29, 30,
	21, 23, 22, 24, 20, 19, 89, 0, 0, 0,
	25, 26, 27, 28, 29, 30, 21, 23, 22, 24,
	20, 19, 88, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 25, 26, 27, 28, 29, 30, 21, 23,
	22, 24, 20, 19, 76, 0, 0, 0, 0, 0,
	0, 0, 25, 26, 27, 28, 29, 30, 21, 23,
	22, 24, 20, 19, 25, 26, 27, 28, 29, 30,
	21, 23, 22, 24, 20, 19, 25, 26, 27, 28,
	29, 30, 21, 23, 22, 24, 20,
}
var yyPact = [...]int{

	80, -1000, -1000, 297, -1000, -1000, -1000, 123, 47, 80,
	-1000, -1000, 80, -1000, -1000, 80, 80, 100, 22, 80,
	80, 80, 80, 80, 80, 80, 80, 80, 80, 80,
	80, 80, 80, 100, 80, 100, 80, 82, 285, -1000,
	177, 137, 24, -17, 22, -1000, -1000, -1000, -28, 100,
	-24, 100, 100, 309, 166, 83, 83, 83, 83, 54,
	54, -1000, -1000, 83, 83, 77, -1000, 265, 243, 62,
	-1000, 46, 297, -1000, 227, -1000, 80, 80, 80, 80,
	80, -1000, -1000, 43, -1000, 44, 85, -1000, 80, -36,
	-1000, 80, -1000, -1000, 149, 297, 189, -22, 80, 100,
	50, 100, -1000, 80, 211, 80, 80, -1000, 297, 5,
	-1000, -1000, 100, -37, 70, -1000, 90, 88, 297, 100,
	297, 21, 80, 60, 100, 61, 100, 100, -1000, 80,
	297, -1000, -1000, 80, 100, -1000, -1000, 297, 297, 29,
	80, 297,
}
var yyPgo = [...]int{

	0, 140, 139, 18, 136, 132, 0, 128, 126, 125,
	5, 6, 121, 117, 8, 7, 11, 113, 3, 2,
	110, 109, 4, 108,
}
var yyR1 = [...]int{

	0, 23, 5, 5, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 21,
	21, 22, 22, 7, 9, 9, 10, 10, 16, 16,
	12, 12, 12, 14, 14, 2, 1, 1, 1, 18,
	18, 19, 19, 17, 13, 13, 15, 15, 20, 20,
	3, 4, 4, 4, 4, 11, 11, 11, 8, 8,
	8, 8,
}
var yyR2 = [...]int{

	0, 1, 0, 1, 1, 1, 1, 4, 1, 3,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 6, 4, 3, 2, 1, 0,
	1, 3, 5, 5, 0, 1, 1, 3, 0, 2,
	1, 1, 1, 1, 2, 4, 1, 3, 3, 0,
	1, 1, 3, 3, 4, 6, 1, 2, 7, 9,
	1, 1, 3, 4, 4, 0, 1, 3, 6, 4,
	4, 8,
}
var yyChk = [...]int{

	-1000, -23, -5, -6, 4, 6, 47, -3, -4, 10,
	-8, -7, 18, 34, 5, 42, 31, 32, 35, 28,
	27, 23, 25, 24, 26, 17, 18, 19, 20, 21,
	22, 10, 12, 14, 29, 16, 12, -11, -6, -6,
	-6, -6, -3, -16, -12, -14, -13, -15, -2, 39,
	-20, 40, 38, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -9, -10, -6, -6, -21,
	-22, -3, -6, -3, -6, 11, 9, 43, 45, 29,
	36, -16, -14, -3, -15, -3, -3, 11, 7, 13,
	15, 21, 13, -11, -6, -6, -6, -11, 29, 8,
	21, 10, -10, 46, -6, 44, 33, 37, -6, -3,
	-1, -3, 14, 41, -18, -19, -17, -3, -6, 7,
	-6, -6, 29, -18, 46, 11, 7, 8, -22, 45,
	-6, 15, -3, 21, 8, -19, -3, -6, -6, -3,
	21, -6,
}
var yyDef = [...]int{

	2, -2, 1, 3, 4, 5, 6, 61, 8, 65,
	10, 11, 0, 28, 60, 0, 0, 0, 38, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 34, 0, 29, 0, 0, 0, 0, 66, 27,
	0, 0, 0, 0, 38, 40, 41, 42, 43, 0,
	56, 0, 0, 12, 13, -2, -2, -2, -2, 18,
	19, 20, 21, -2, -2, 0, 35, 36, 0, 0,
	30, 0, 26, 62, 0, 9, 65, 0, 0, 0,
	65, 39, 44, 0, 57, 0, 0, 7, 0, 63,
	25, 0, 64, 67, 69, 70, 0, 0, 0, 0,
	0, 49, 37, 0, 31, 0, 0, 33, 54, 0,
	45, 46, 49, 0, 0, 50, 51, 0, 24, 0,
	68, 0, 0, 0, 0, 0, 0, 0, 32, 0,
	55, 47, 48, 0, 0, 52, 53, 71, 58, 0,
	0, 59,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:74
		{
			Absyn_root = yyDollar[1].exp
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc/tiger.y:77
		{
			yyVAL.exp = nil
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:78
		{
			yyVAL.exp = yyDollar[1].exp
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:81
		{
			yyVAL.exp = &absyn.IntExp{EM_tokPos, yyDollar[1].ival}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:82
		{
			yyVAL.exp = &absyn.StringExp{EM_tokPos, yyDollar[1].sval}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:83
		{
			yyVAL.exp = &absyn.NilExp{EM_tokPos}
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc/tiger.y:84
		{
			yyVAL.exp = &absyn.CallExp{EM_tokPos, yyDollar[1].sym, yyDollar[3].expList}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:85
		{
			yyVAL.exp = &absyn.VarExp{EM_tokPos, yyDollar[1].Var}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:86
		{
			yyVAL.exp = &absyn.SeqExp{EM_tokPos, yyDollar[2].expList}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:87
		{
			yyVAL.exp = yyDollar[1].exp
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:88
		{
			yyVAL.exp = yyDollar[1].exp
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:89
		{
			yyVAL.exp = &absyn.IfExp{EM_tokPos, yyDollar[1].exp, &absyn.IntExp{EM_tokPos, 1}, yyDollar[3].exp}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:90
		{
			yyVAL.exp = &absyn.IfExp{EM_tokPos, yyDollar[1].exp, yyDollar[3].exp, &absyn.IntExp{EM_tokPos, 0}}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:91
		{
			yyVAL.exp = &absyn.OpExp{EM_tokPos, absyn.LtOp, yyDollar[1].exp, yyDollar[3].exp}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:92
		{
			yyVAL.exp = &absyn.OpExp{EM_tokPos, absyn.GtOp, yyDollar[1].exp, yyDollar[3].exp}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:93
		{
			yyVAL.exp = &absyn.OpExp{EM_tokPos, absyn.LeOp, yyDollar[1].exp, yyDollar[3].exp}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:94
		{
			yyVAL.exp = &absyn.OpExp{EM_tokPos, absyn.GeOp, yyDollar[1].exp, yyDollar[3].exp}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:95
		{
			yyVAL.exp = &absyn.OpExp{EM_tokPos, absyn.PlusOp, yyDollar[1].exp, yyDollar[3].exp}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:96
		{
			yyVAL.exp = &absyn.OpExp{EM_tokPos, absyn.MinusOp, yyDollar[1].exp, yyDollar[3].exp}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:97
		{
			yyVAL.exp = &absyn.OpExp{EM_tokPos, absyn.TimesOp, yyDollar[1].exp, yyDollar[3].exp}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:98
		{
			yyVAL.exp = &absyn.OpExp{EM_tokPos, absyn.DivideOp, yyDollar[1].exp, yyDollar[3].exp}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:99
		{
			yyVAL.exp = &absyn.OpExp{EM_tokPos, absyn.EqOp, yyDollar[1].exp, yyDollar[3].exp}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:100
		{
			yyVAL.exp = &absyn.OpExp{EM_tokPos, absyn.NeqOp, yyDollar[1].exp, yyDollar[3].exp}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line yacc/tiger.y:101
		{
			yyVAL.exp = &absyn.ArrayExp{EM_tokPos, yyDollar[1].sym, yyDollar[3].exp, yyDollar[6].exp}
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc/tiger.y:102
		{
			yyVAL.exp = &absyn.RecordExp{EM_tokPos, yyDollar[1].sym, yyDollar[3].efieldList}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:103
		{
			yyVAL.exp = &absyn.AssignExp{EM_tokPos, yyDollar[1].Var, yyDollar[3].exp}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yacc/tiger.y:104
		{
			yyVAL.exp = &absyn.OpExp{EM_tokPos, absyn.MinusOp, &absyn.IntExp{EM_tokPos, 0}, yyDollar[2].exp}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:105
		{
			yyVAL.exp = &absyn.BreakExp{EM_tokPos}
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc/tiger.y:108
		{
			yyVAL.efieldList = nil
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:109
		{
			yyVAL.efieldList = yyDollar[1].efieldList
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:112
		{
			yyVAL.efieldList = absyn.EfieldListInsert(&absyn.Efield_{yyDollar[1].sym, yyDollar[3].exp}, nil)
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yacc/tiger.y:113
		{
			yyVAL.efieldList = absyn.EfieldListInsert(&absyn.Efield_{yyDollar[1].sym, yyDollar[3].exp}, yyDollar[5].efieldList)
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yacc/tiger.y:116
		{
			yyVAL.exp = &absyn.LetExp{EM_tokPos, yyDollar[2].decList, &absyn.SeqExp{EM_tokPos, yyDollar[4].expList}}
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc/tiger.y:119
		{
			yyVAL.expList = nil
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:120
		{
			yyVAL.expList = yyDollar[1].expList
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:123
		{
			yyVAL.expList = absyn.ExpListInsert(yyDollar[1].exp, nil)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:124
		{
			yyVAL.expList = absyn.ExpListInsert(yyDollar[1].exp, yyDollar[3].expList)
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc/tiger.y:128
		{
			yyVAL.decList = nil
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yacc/tiger.y:129
		{
			yyVAL.decList = absyn.DecListInsert(yyDollar[1].dec, yyDollar[2].decList)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:135
		{
			yyVAL.dec = yyDollar[1].dec
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:136
		{
			yyVAL.dec = yyDollar[1].dec
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:137
		{
			yyVAL.dec = yyDollar[1].dec
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:140
		{
			yyVAL.dec = &absyn.TypeDec{EM_tokPos, absyn.NametyListInsert(yyDollar[1].namety, nil)}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yacc/tiger.y:141
		{
			yyVAL.dec = &absyn.TypeDec{EM_tokPos, absyn.NametyListInsert(yyDollar[1].namety, yyDollar[2].dec.(*absyn.TypeDec).Type)}
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc/tiger.y:145
		{
			yyVAL.namety = &absyn.Namety_{yyDollar[2].sym, yyDollar[4].ty}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:151
		{
			yyVAL.ty = &absyn.NameTy{EM_tokPos, yyDollar[1].sym}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:152
		{
			yyVAL.ty = &absyn.RecordTy{EM_tokPos, yyDollar[2].fieldList}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:153
		{
			yyVAL.ty = &absyn.ArrayTy{EM_tokPos, yyDollar[3].sym}
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc/tiger.y:158
		{
			yyVAL.fieldList = nil
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:159
		{
			yyVAL.fieldList = yyDollar[1].fieldList
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:162
		{
			yyVAL.fieldList = absyn.FieldListInsert(yyDollar[1].field, nil)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:163
		{
			yyVAL.fieldList = absyn.FieldListInsert(yyDollar[1].field, yyDollar[3].fieldList)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:166
		{
			yyVAL.field = &absyn.Field_{EM_tokPos, yyDollar[1].sym, yyDollar[3].sym, true}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc/tiger.y:171
		{
			yyVAL.dec = &absyn.VarDec{EM_tokPos, yyDollar[2].sym, nil, yyDollar[4].exp, true}
		}
	case 55:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line yacc/tiger.y:172
		{
			yyVAL.dec = &absyn.VarDec{EM_tokPos, yyDollar[2].sym, yyDollar[4].sym, yyDollar[6].exp, true}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:175
		{
			yyVAL.dec = &absyn.FunctionDec{EM_tokPos, absyn.FundecListInsert(yyDollar[1].fundec, nil)}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yacc/tiger.y:176
		{
			yyVAL.dec = &absyn.FunctionDec{EM_tokPos, absyn.FundecListInsert(yyDollar[1].fundec, yyDollar[2].dec.(*absyn.FunctionDec).Function)}
		}
	case 58:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line yacc/tiger.y:181
		{
			yyVAL.fundec = &absyn.Fundec_{EM_tokPos, yyDollar[2].sym, yyDollar[4].fieldList, nil, yyDollar[7].exp}
		}
	case 59:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line yacc/tiger.y:182
		{
			yyVAL.fundec = &absyn.Fundec_{EM_tokPos, yyDollar[2].sym, yyDollar[4].fieldList, yyDollar[7].sym, yyDollar[9].exp}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:185
		{
			yyVAL.sym = symbol.SymbolInsert(yyDollar[1].sval)
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:191
		{
			yyVAL.Var = &absyn.SimpleVar{EM_tokPos, yyDollar[1].sym}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:192
		{
			yyVAL.Var = &absyn.FieldVar{EM_tokPos, yyDollar[1].Var, yyDollar[3].sym}
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc/tiger.y:193
		{
			yyVAL.Var = &absyn.SubscriptVar{EM_tokPos, &absyn.SimpleVar{EM_tokPos, yyDollar[1].sym}, yyDollar[3].exp}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc/tiger.y:194
		{
			yyVAL.Var = &absyn.SubscriptVar{EM_tokPos, yyDollar[1].Var, yyDollar[3].exp}
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc/tiger.y:197
		{
			yyVAL.expList = nil
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc/tiger.y:198
		{
			yyVAL.expList = absyn.ExpListInsert(yyDollar[1].exp, nil)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc/tiger.y:199
		{
			yyVAL.expList = absyn.ExpListInsert(yyDollar[1].exp, yyDollar[3].expList)
		}
	case 68:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line yacc/tiger.y:202
		{
			yyVAL.exp = &absyn.IfExp{EM_tokPos, yyDollar[2].exp, yyDollar[4].exp, yyDollar[6].exp}
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc/tiger.y:203
		{
			yyVAL.exp = &absyn.IfExp{EM_tokPos, yyDollar[2].exp, yyDollar[4].exp, nil}
		}
	case 70:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc/tiger.y:204
		{
			yyVAL.exp = &absyn.WhileExp{EM_tokPos, yyDollar[2].exp, yyDollar[4].exp}
		}
	case 71:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line yacc/tiger.y:205
		{
			yyVAL.exp = &absyn.ForExp{EM_tokPos, yyDollar[2].sym, yyDollar[4].exp, yyDollar[6].exp, yyDollar[8].exp, true}
		}
	}
	goto yystack /* stack new state and value */
}
