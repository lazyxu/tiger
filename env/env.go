package env

import (
	"github.com/MeteorKL/tiger/symbol"
	"github.com/MeteorKL/tiger/table"
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/translate"
	"github.com/MeteorKL/tiger/types"
)

func Base_tenv() table.Table {
	tenvBaseTable := table.Empty()

	symbol.Enter(tenvBaseTable, symbol.New("int"), &types.Tyint)
	symbol.Enter(tenvBaseTable, symbol.New("string"), &types.Tystring)

	return tenvBaseTable
}

func Base_venv() table.Table {
	venvBaseTable := table.Empty()

	symbol.Enter(venvBaseTable, symbol.New("print"),
		&FunEntry{translate.Outermost(), temp.StrToLabel("print"), &types.TyList_{&types.Tystring, nil}, &types.Tyvoid})

	symbol.Enter(venvBaseTable, symbol.New("flush"),
		&FunEntry{translate.Outermost(), temp.StrToLabel("flush"), nil, &types.Tyvoid})

	symbol.Enter(venvBaseTable, symbol.New("getchar"),
		&FunEntry{translate.Outermost(), temp.StrToLabel("getchar"), nil, &types.Tystring})

	symbol.Enter(venvBaseTable, symbol.New("ord"),
		&FunEntry{translate.Outermost(), temp.StrToLabel("ord"), &types.TyList_{&types.Tystring, nil}, &types.Tyint})

	symbol.Enter(venvBaseTable, symbol.New("chr"),
		&FunEntry{translate.Outermost(), temp.StrToLabel("chr"), &types.TyList_{&types.Tyint, nil}, &types.Tystring})

	symbol.Enter(venvBaseTable, symbol.New("size"),
		&FunEntry{translate.Outermost(), temp.StrToLabel("size"), &types.TyList_{&types.Tystring, nil}, &types.Tyint})

	symbol.Enter(venvBaseTable, symbol.New("substring"),
		&FunEntry{translate.Outermost(), temp.StrToLabel("substring"),
			&types.TyList_{&types.Tystring, &types.TyList_{&types.Tyint, &types.TyList_{&types.Tyint, nil}}}, &types.Tystring})

	symbol.Enter(venvBaseTable, symbol.New("concat"),
		&FunEntry{translate.Outermost(), temp.StrToLabel("concat"), &types.TyList_{&types.Tystring, &types.TyList_{&types.Tystring, nil}}, &types.Tystring})

	symbol.Enter(venvBaseTable, symbol.New("not"),
		&FunEntry{translate.Outermost(), temp.StrToLabel("not"), &types.TyList_{&types.Tyint, nil}, &types.Tyint})

	symbol.Enter(venvBaseTable, symbol.New("exit"),
		&FunEntry{translate.Outermost(), temp.StrToLabel("exit"), &types.TyList_{&types.Tyint, nil}, &types.Tyvoid})

	return venvBaseTable
}
