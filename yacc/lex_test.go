package yacc

import (
	"flag"
	"testing"

	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/util"

	"fmt"
)

func Test_Lex(t *testing.T) {
	args := flag.Args()
	if len(args) < 1 {
		println(`
Usage:
	go test -v -run='Test_Lex' -args ../testcases/test4.tig
`)
		return
	}
	testfile := args[0]

	fmt.Printf("+--------------------------------------\n")
	println("| Test_Lex")
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
