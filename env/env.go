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

	symbol.Enter(tenvBaseTable, symbol.Insert("int"), &types.Tyint)
	symbol.Enter(tenvBaseTable, symbol.Insert("string"), &types.Tystring)

	return tenvBaseTable
}

func Base_venv() table.Table {
	venvBaseTable := table.Empty()

	symbol.Enter(venvBaseTable, symbol.Insert("print"),
		&FunEntry{translate.Outermost(), temp.Newlabel(), &types.TyList_{&types.Tystring, nil}, &types.Tyvoid})

	symbol.Enter(venvBaseTable, symbol.Insert("flush"),
		&FunEntry{translate.Outermost(), temp.Newlabel(), nil, &types.Tyvoid})

	symbol.Enter(venvBaseTable, symbol.Insert("getchar"),
		&FunEntry{translate.Outermost(), temp.Newlabel(), nil, &types.Tystring})

	symbol.Enter(venvBaseTable, symbol.Insert("ord"),
		&FunEntry{translate.Outermost(), temp.Newlabel(), &types.TyList_{&types.Tystring, nil}, &types.Tyint})

	symbol.Enter(venvBaseTable, symbol.Insert("chr"),
		&FunEntry{translate.Outermost(), temp.Newlabel(), &types.TyList_{&types.Tyint, nil}, &types.Tystring})

	symbol.Enter(venvBaseTable, symbol.Insert("size"),
		&FunEntry{translate.Outermost(), temp.Newlabel(), &types.TyList_{&types.Tystring, nil}, &types.Tyint})

	symbol.Enter(venvBaseTable, symbol.Insert("substring"),
		&FunEntry{translate.Outermost(), temp.Newlabel(),
			&types.TyList_{&types.Tystring, &types.TyList_{&types.Tyint, &types.TyList_{&types.Tyint, nil}}}, &types.Tystring})

	symbol.Enter(venvBaseTable, symbol.Insert("concat"),
		&FunEntry{translate.Outermost(), temp.Newlabel(), &types.TyList_{&types.Tystring, &types.TyList_{&types.Tystring, nil}}, &types.Tystring})

	symbol.Enter(venvBaseTable, symbol.Insert("not"),
		&FunEntry{translate.Outermost(), temp.Newlabel(), &types.TyList_{&types.Tyint, nil}, &types.Tyint})

	symbol.Enter(venvBaseTable, symbol.Insert("exit"),
		&FunEntry{translate.Outermost(), temp.Newlabel(), &types.TyList_{&types.Tyint, nil}, &types.Tyvoid})

	return venvBaseTable
}
