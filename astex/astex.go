package astex

import (
	"go/ast"
	"iter"
)

func ExprGetFullTypeName(fieldType ast.Expr) (string, bool) {
	switch ftype := fieldType.(type) {
	case *ast.Ident:
		return ftype.Name, true
	case *ast.StarExpr:
		return ExprGetFullTypeName(ftype.X)
	case *ast.ArrayType:
		return ExprGetFullTypeName(ftype.Elt)
	case *ast.SelectorExpr:
		if x, ok := ExprGetFullTypeName(ftype.X); ok {
			if sel, ok := ExprGetFullTypeName(ftype.Sel); ok {
				if len(x) != 0 {
					return x + "." + sel, true
				}
				return sel, true
			}
		}
	case *ast.IndexExpr:
		if x, ok := ExprGetFullTypeName(ftype.X); ok {
			if index, ok := ExprGetFullTypeName(ftype.Index); ok {
				return x + "[" + index + "]", true
			}
		}
	}

	return "", true
}

func ExprGetCallTypeName(fieldType ast.Expr) (string, bool) {
	switch ftype := fieldType.(type) {
	case *ast.Ident:
		return ftype.Name, true
	case *ast.StarExpr:
		return ExprGetFullTypeName(ftype.X)
	case *ast.ArrayType:
		return ExprGetFullTypeName(ftype.Elt)
	case *ast.SelectorExpr:
		return ExprGetCallTypeName(ftype.Sel)
	case *ast.IndexExpr:
		return ExprGetCallTypeName(ftype.X)
	}

	return "", false
}

func GetFieldDeclTypeName(fieldType ast.Expr) (string, bool) {
	switch ftype := fieldType.(type) {
	case *ast.Ident:
		return ftype.Name, true
	case *ast.StarExpr:
		if x, ok := GetFieldDeclTypeName(ftype.X); ok {
			return "*" + x, true
		}
	case *ast.ArrayType:
		if elt, ok := GetFieldDeclTypeName(ftype.Elt); ok {
			return "[]" + elt, true
		}
	case *ast.SelectorExpr:
		if x, ok := GetFieldDeclTypeName(ftype.X); ok {
			if sel, ok := GetFieldDeclTypeName(ftype.Sel); ok {
				if len(x) != 0 {
					return x + "." + sel, true
				}
				return sel, true
			}
		}
	case *ast.IndexExpr:
		if x, ok := GetFieldDeclTypeName(ftype.X); ok {
			if index, ok := GetFieldDeclTypeName(ftype.Index); ok {
				return x + "[" + index + "]", true
			}
		}
	}

	return "", false
}

func FuncDeclRecvType(decl *ast.FuncDecl) (ast.Expr, bool) {
	if decl.Recv != nil && len(decl.Recv.List) == 1 {
		return decl.Recv.List[0].Type, true
	}

	return nil, false
}

func FuncTypeParamsSeq(decl *ast.FuncType) iter.Seq[*ast.Field] {
	if decl.Params != nil {
		return func(yield func(*ast.Field) bool) {
			for _, f := range decl.Params.List {
				if !yield(f) {
					return
				}
			}
		}
	}

	return func(yield func(*ast.Field) bool) {}
}

func FuncTypeResultsSeq(decl *ast.FuncType) iter.Seq[*ast.Field] {
	if decl.Results != nil {
		return func(yield func(*ast.Field) bool) {
			for _, f := range decl.Results.List {
				if !yield(f) {
					return
				}
			}
		}
	}

	return func(yield func(*ast.Field) bool) {}
}
