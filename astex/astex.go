package astex

import (
	"fmt"
	"go/ast"
)

func ExprGetTypeName(fieldType ast.Expr) string {
	switch ftype := fieldType.(type) {
	case *ast.Ident:
		return ftype.Name
	case *ast.StarExpr:
		return ExprGetTypeName(ftype.X)
	case *ast.ArrayType:
		return ExprGetTypeName(ftype.Elt)
	case *ast.SelectorExpr:
		x := ExprGetTypeName(ftype.X)
		sel := ExprGetTypeName(ftype.Sel)
		if len(x) != 0 {
			return x + "." + sel
		}
		return sel
	case *ast.IndexExpr:
		x := ExprGetTypeName(ftype.X)
		index := ExprGetTypeName(ftype.Index)
		return x + "[" + index + "]"
	}

	return ""
}

func GetFieldDeclTypeName(fieldType ast.Expr) (string, error) {
	switch ftype := fieldType.(type) {
	case *ast.Ident:
		return ftype.Name, nil
	case *ast.StarExpr:
		return "*" + ExprGetTypeName(ftype.X), nil
	case *ast.ArrayType:
		return "[]" + ExprGetTypeName(ftype.Elt), nil
	case *ast.SelectorExpr:
		x := ExprGetTypeName(ftype.X)
		sel := ExprGetTypeName(ftype.Sel)
		if len(x) != 0 {
			return x + "." + sel, nil
		}
		return sel, nil
	case *ast.IndexExpr:
		x := ExprGetTypeName(ftype.X)
		index := ExprGetTypeName(ftype.Index)
		return x + "[" + index + "]", nil
	}

	return "", fmt.Errorf("unknown parameter type")
}

func FuncDeclRecvType(decl *ast.FuncDecl) (ast.Expr, error) {
	if decl.Recv != nil && len(decl.Recv.List) == 1 {
		return decl.Recv.List[0].Type, nil
	}

	return nil, fmt.Errorf("recv type not found")
}

func FuncDeclParams(decl *ast.FuncDecl) ([]*ast.Field, error) {
	if decl.Type.Params != nil {
		return decl.Type.Params.List, nil
	}

	return []*ast.Field{}, nil
}
