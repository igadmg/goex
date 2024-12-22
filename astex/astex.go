package astex

import (
	"fmt"
	"go/ast"
)

func ExprGetFullTypeName(fieldType ast.Expr) string {
	switch ftype := fieldType.(type) {
	case *ast.Ident:
		return ftype.Name
	case *ast.StarExpr:
		return ExprGetFullTypeName(ftype.X)
	case *ast.ArrayType:
		return ExprGetFullTypeName(ftype.Elt)
	case *ast.SelectorExpr:
		x := ExprGetFullTypeName(ftype.X)
		sel := ExprGetFullTypeName(ftype.Sel)
		if len(x) != 0 {
			return x + "." + sel
		}
		return sel
	case *ast.IndexExpr:
		x := ExprGetFullTypeName(ftype.X)
		index := ExprGetFullTypeName(ftype.Index)
		return x + "[" + index + "]"
	}

	return ""
}

func ExprGetCallTypeName(fieldType ast.Expr) string {
	switch ftype := fieldType.(type) {
	case *ast.Ident:
		return ftype.Name
	case *ast.StarExpr:
		return ExprGetFullTypeName(ftype.X)
	case *ast.ArrayType:
		return ExprGetFullTypeName(ftype.Elt)
	case *ast.SelectorExpr:
		sel := ExprGetCallTypeName(ftype.Sel)
		return sel
	case *ast.IndexExpr:
		x := ExprGetCallTypeName(ftype.X)
		return x
	}

	return ""
}

func GetFieldDeclTypeName(fieldType ast.Expr) (string, error) {
	switch ftype := fieldType.(type) {
	case *ast.Ident:
		return ftype.Name, nil
	case *ast.StarExpr:
		x, err := GetFieldDeclTypeName(ftype.X)
		if err != nil {
			return "", err
		}
		return "*" + x, nil
	case *ast.ArrayType:
		elt, err := GetFieldDeclTypeName(ftype.Elt)
		if err != nil {
			return "", err
		}
		return "[]" + elt, nil
	case *ast.SelectorExpr:
		x, err := GetFieldDeclTypeName(ftype.X)
		if err != nil {
			return "", err
		}
		sel, err := GetFieldDeclTypeName(ftype.Sel)
		if err != nil {
			return "", err
		}
		if len(x) != 0 {
			return x + "." + sel, nil
		}
		return sel, nil
	case *ast.IndexExpr:
		x, err := GetFieldDeclTypeName(ftype.X)
		if err != nil {
			return "", err
		}
		index, err := GetFieldDeclTypeName(ftype.Index)
		if err != nil {
			return "", err
		}
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
