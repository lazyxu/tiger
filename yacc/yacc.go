package yacc

import (
	"io/ioutil"
	"os"

	"github.com/MeteorKL/tiger/absyn"
)

func YYParse(filepath string) absyn.Exp {
	absyn.LinePosListInit()
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	src, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	f.Close()
	yyParse(NewLex(string(src)))
	return Absyn_root
}
