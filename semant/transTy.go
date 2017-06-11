package semant

import (
	"github.com/MeteorKL/tiger/absyn"
	"github.com/MeteorKL/tiger/symbol"
	"github.com/MeteorKL/tiger/table"
	"github.com/MeteorKL/tiger/types"
	"github.com/MeteorKL/tiger/util"
	"github.com/MeteorKL/tiger/yacc"
)

func getActualTy(ty types.Ty) types.Ty {
	for {
		if ty == nil {
			break
		}
		if _, ok := ty.(*types.Name_); !ok {
			break
		}
		ty = ty.(*types.Name_).Ty
	}
	return ty
}

func equalTy(leftTy types.Ty, rightTy types.Ty, p absyn.Pos) bool {
	leftTy = getActualTy(leftTy)
	rightTy = getActualTy(rightTy)
	if _, ok := leftTy.(*types.Void_); ok {
		println(leftTy.(*types.Void_))
		yacc.EM_error(p, "Left hand side type can not be void.")
		return false
	}
	if _, ok := rightTy.(*types.Void_); ok {
		yacc.EM_error(p, "Right hand side type can not be void.")
		return false
	}
	if _, ok := leftTy.(*types.Nil_); ok {
		switch rightTy.(type) {
		case *types.Nil_:
			yacc.EM_error(p, "Left and right hand side types can not be all nil.")
			return false
		case *types.Record_:
			return true
		default:
			return false
		}
	}
	if _, ok := leftTy.(*types.Record_); ok {
		switch rightTy.(type) {
		case *types.Nil_:
			return true
		case *types.Record_:
			return true
		default:
			yacc.EM_error(p, "Left and right hand side types do not match.")
			return false
		}
	}
	if _, ok := leftTy.(*types.Int_); ok {
		if _, ok := rightTy.(*types.Int_); ok {
			return true
		}
		yacc.EM_error(p, "Left and right hand side types do not match.")
		return false
	}
	if _, ok := leftTy.(*types.String_); ok {
		if _, ok := rightTy.(*types.String_); ok {
			return true
		}
		yacc.EM_error(p, "Left and right hand side types do not match.")
		return false
	}
	if _, ok := leftTy.(*types.Array_); ok {
		if _, ok := rightTy.(*types.Array_); ok {
			return true
		}
		yacc.EM_error(p, "Left and right hand side types do not match.")
		return false
	}
	return false
}

func transTy(tenv table.Table, t absyn.Ty) types.Ty {
	switch t.(type) {
	case *absyn.NameTy:
		t := t.(*absyn.NameTy)
		nameTy := symbol.Look(tenv, t.Name).(types.Ty)
		if nameTy == nil {
			yacc.EM_error(t.Pos, "Type '"+t.Name.Name+"' is undefined.")
			return &types.Tyint
		}
		return nameTy
	case *absyn.RecordTy:
		t := t.(*absyn.RecordTy)
		recordFieldList := t.Record
		var TyList types.FieldList = nil
		var tyHead types.FieldList = nil
		var tyRecordHead types.Ty
		var tempTy types.Field
		for recordFieldList != nil {
			tyRecordHead = symbol.Look(tenv, recordFieldList.Head.Typ).(types.Ty)
			if tyRecordHead == nil {
				yacc.EM_error(recordFieldList.Head.Pos, "Type '"+recordFieldList.Head.Typ.Name+"' is undefined.")
			} else {
				tempTy = &types.Field_{recordFieldList.Head.Name, tyRecordHead}
				if TyList == nil {
					TyList = &types.FieldList_{tempTy, nil}
					tyHead = TyList
				} else {
					TyList.Tail = &types.FieldList_{tempTy, nil}
					TyList = TyList.Tail
				}
			}
			recordFieldList = recordFieldList.Tail
		}
		return &types.Record_{tyHead}
	case *absyn.ArrayTy:
		t := t.(*absyn.ArrayTy)
		arrayTy := symbol.Look(tenv, t.Array).(types.Ty)
		if arrayTy == nil {
			yacc.EM_error(t.Pos, "Type '"+t.Array.Name+"' is undefined.")
		}
		return &types.Array_{arrayTy}

	}
	util.Assert(!true)
	return nil
}
