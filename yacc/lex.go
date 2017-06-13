package yacc

import (
	"fmt"

	"github.com/MeteorKL/tiger/absyn"
)

var EM_old_tokPos absyn.Pos
var EM_tokPos absyn.Pos
var Em_ch byte
var Em_src string
var Em_tokStr string

const EOFCH byte = 0xff

type TigerLex struct {
	yyLexer
}

func (l *TigerLex) next() byte {
	EM_tokPos++
	if EM_tokPos < absyn.Pos(len(Em_src)) {
		Em_ch = Em_src[EM_tokPos]
		if Em_ch == '\n' {
			absyn.EM_newline(EM_tokPos)
		}
	} else {
		EM_tokPos = absyn.Pos(len(Em_src))
		Em_ch = EOFCH
	}
	return Em_ch
}

func (l *TigerLex) skipWhitespace() {
	for Em_ch == ' ' || Em_ch == '\t' || Em_ch == '\n' || Em_ch == '\r' {
		l.next()
	}
}

func (l *TigerLex) LookAhead(offset int) byte {
	if int(EM_tokPos)+offset >= len(Em_src) {
		return 0
	}
	return Em_src[int(EM_tokPos)+offset]
}

func IsLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func IsDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func (l *TigerLex) scanIdentifier() string {
	offs := EM_tokPos
	for IsLetter(Em_ch) || IsDigit(Em_ch) {
		l.next()
	}
	return string(Em_src[offs:EM_tokPos])
}

func (l *TigerLex) scanNumber(yylval *yySymType) string {
	offs := EM_tokPos
	(*yylval).ival = 0
	base := 10
	for num := Digit(Em_ch); num < base; num = Digit(Em_ch) {
		l.next()
		(*yylval).ival = (*yylval).ival*10 + num
	}
	return string(Em_src[offs:EM_tokPos])
}

func (l *TigerLex) scanMantissaT(base int, times int) (value int) {
	value = 0
	i := 0
	for i < times {
		t := Digit(Em_ch)
		if t >= base {
			break
		}
		l.next()
		value = value*base + t
		i++
	}
	return
}

func NewLex(src string) *TigerLex {
	l := new(TigerLex)
	EM_tokPos = -1
	Em_src = src
	l.next()
	return l
}

func (l *TigerLex) Lex(yylval *yySymType) int {
	var tok int
	var lit string
	l.skipWhitespace()
	EM_old_tokPos = EM_tokPos
	switch c := Em_ch; {
	case IsLetter(c):
		lit = l.scanIdentifier()
		if len(lit) > 1 {
			// token长度大于1才查找否则肯定是标识符
			tok = Lookup(lit)
		} else {
			tok = ID
		}
		yylval.sval = Em_src[EM_old_tokPos:EM_tokPos]
		Em_tokStr = string(Em_src[EM_old_tokPos:EM_tokPos])
	case IsDigit(c):
		lit = l.scanNumber(yylval)
		tok = INT
		Em_tokStr = string(Em_src[EM_old_tokPos:EM_tokPos])
	default:
		l.next()
		Em_tokStr = string(Em_src[EM_old_tokPos:EM_tokPos])
		switch c {
		case EOFCH:
			tok = EOF
		case '"':
			tok = STRING
			str := []byte{}
			start := int(EM_tokPos)
			var end int
			for {
				c := Em_ch
				if c == '\n' || c == EOFCH {
					l.Error("string literal not terminated")
					end = int(EM_tokPos)
					break
				}
				if c == '"' {
					l.next()
					end = int(EM_tokPos) - 1
					break
				}
				str = append(str, Em_ch)
				l.next()
			}
			yylval.sval = Em_src[start:end]
		case ',':
			tok = COMMA
		case ':':
			if Em_ch == '=' {
				l.next()
				Em_tokStr = string(Em_src[EM_old_tokPos:EM_tokPos])
				tok = ASSIGN
			} else {
				tok = COLON
			}
		case ';':
			tok = SEMICOLON
		case '(':
			tok = LPAREN
		case ')':
			tok = RPAREN
		case '[':
			tok = LBRACK
		case ']':
			tok = RBRACK
		case '{':
			tok = LBRACE
		case '}':
			tok = RBRACE
		case '.':
			tok = DOT
		case '+':
			tok = PLUS
		case '-':
			tok = MINUS
		case '*':
			tok = TIMES
		case '/':
			if Em_ch == '*' {
				l.next()
				for Em_ch != EOFCH {
					ch := Em_ch
					l.next()
					if ch == '*' && Em_ch == '/' {
						l.next()
						return l.Lex(yylval)
					}
				}
				l.Error("comment not terminated")
			} else if Em_ch == '/' {
				l.next()
				for Em_ch != '\n' && Em_ch != EOFCH {
					l.next()
				}
				return l.Lex(yylval)
			} else {
				tok = DIVIDE
			}
		case '=':
			tok = EQ
		case '<':
			if Em_ch == '>' {
				l.next()
				Em_tokStr = string(Em_src[EM_old_tokPos:EM_tokPos])
				tok = NEQ
			} else if Em_ch == '=' {
				l.next()
				Em_tokStr = string(Em_src[EM_old_tokPos:EM_tokPos])
				tok = LE
			} else {
				tok = LT
			}
		case '>':
			if Em_ch == '=' {
				l.next()
				Em_tokStr = string(Em_src[EM_old_tokPos:EM_tokPos])
				tok = GE
			} else {
				tok = GT
			}
		case '&':
			tok = AND
		case '|':
			tok = OR
		default:
			tok = ILLEGAL
		}
	}
	return tok
}

func (l TigerLex) Error(s string) {
	EM_error(EM_tokPos, s)
}

func EM_error(p absyn.Pos, s string) {
	position := absyn.PositionFor(p)
	fmt.Printf("%d:%d: error: %s\n", position.Line, position.Column, s)

}
