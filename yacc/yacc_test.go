package yacc

import (
	"flag"
	"fmt"
	"testing"

	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/util"
)

func Test_Yacc(t *testing.T) {
	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf(`
Usage:
	go test -v -run='Test_Yacc' -args ../testcases/test4.tig
`)
		return
	}
	testfile := args[0]

	absyn.PrintExp(testfile+".ast", YYParse(testfile))
	util.Visualization(testfile + ".ast")
}
