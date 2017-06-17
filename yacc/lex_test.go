package yacc

import (
	"flag"
	"strconv"
	"testing"

	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/util"

	"fmt"
)

func Test_Lex(t *testing.T) {
	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf(`
Usage:
	go test -v -run='Test_Lex' -args ../testcases/test4.tig
`)
		return
	}
	testfile := args[0]

	fmt.Printf("+--------------------------------------\n")
	fmt.Printf("| Test_Lex")
	text := util.ReadFile(testfile)
	l := NewLex(string(text))
	yylval := &yySymType{
		ival: 0,
		sval: "",
	}
	fmt.Printf("+--------------------------------------\n")
	fmt.Printf("| off\ttoken\t\tlit\n")
	fmt.Printf("+--------------------------------------\n")
	for {
		tok := l.Lex(yylval)
		if tok == EOF {
			break
		}
		position := absyn.PositionFor(EM_old_tokPos)
		if tok == EOF {
			fmt.Printf("| %d:%d\tEOF\n", position.Line, position.Column)
			fmt.Printf("+--------------------------------------\n")
		} else if tok == INT {
			fmt.Printf("| %d:%d\tINT(%d)\t\t%s\n", position.Line, position.Column, yylval.ival, Em_tokStr)
		} else if tok == STRING {
			fmt.Printf("| %d:%d\tSTRING(%s)\t\t%s\n", position.Line, position.Column, yylval.sval, yylval.sval)
		} else {
			fmt.Printf("| %d:%d\t%s\t\t%s\n", position.Line, position.Column, yyToknames[tok-INT+3], Em_tokStr)
		}
	}
	fmt.Printf("+--------------------------------------\n")
}

func Test_LexAll(t *testing.T) {
	for i := 1; i < 50; i++ {
		testfile := "../testcases/test" + strconv.Itoa(i) + ".tig"
		text := util.ReadFile(testfile)
		l := NewLex(string(text))
		yylval := &yySymType{
			ival: 0,
			sval: "",
		}
		fmt.Printf("lex " + testfile)
		for {
			tok := l.Lex(yylval)
			if tok == EOF {
				break
			}
			position := absyn.PositionFor(EM_old_tokPos)
			if tok == ILLEGAL {
				fmt.Printf("%s:%d:%d\tILLEGAL\n", position.Line, position.Column)
			}
		}
	}
}
