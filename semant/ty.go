package semant

import "github.com/MeteorKL/tiger/types"

func actual_ty(ty types.Ty) types.Ty {
	for {
		if ty == nil {
			break
		}
		Break := false
		switch ty.(type) {
		case *types.Name_:
			t := ty.(*types.Name_)
			ty = t.Ty
		default:
			Break = true
		}
		if Break {
			break
		}
	}
	return ty
}
