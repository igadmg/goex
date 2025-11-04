package astex

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExprGetFullTypeName(t *testing.T) {
	tests := []struct {
		expr     string
		expected string
	}{
		{"int", "int"},
		{"*int", "int"},
		{"[]int", "int"},
		{"pkg.Type", "pkg.Type"},
		{"pkg.Type[TypeParam]", "pkg.Type[TypeParam]"},
	}

	for _, test := range tests {
		expr, err := parser.ParseExpr(test.expr)
		assert.NoError(t, err)

		result, ok := ExprGetFullTypeName(expr)
		assert.Equal(t, ok, true)
		assert.Equal(t, test.expected, result)
	}
}

func TestGetFieldDeclTypeName(t *testing.T) {
	tests := []struct {
		expr     string
		expected string
	}{
		{"int", "int"},
		{"*int", "*int"},
		{"[]int", "[]int"},
		{"pkg.Type", "pkg.Type"},
		{"pkg.Type[TypeParam]", "pkg.Type[TypeParam]"},
	}

	for _, test := range tests {
		expr, err := parser.ParseExpr(test.expr)
		assert.NoError(t, err)

		result, ok := GetFieldDeclTypeName(expr)
		assert.Equal(t, ok, true)
		assert.Equal(t, test.expected, result)
	}
}

func TestFuncDeclRecvType(t *testing.T) {
	src := `
		package main
		func (r *Receiver) Method() {}
	`
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, 0)
	assert.NoError(t, err)

	for _, decl := range node.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			recvType, ok := FuncDeclRecvType(funcDecl)
			assert.Equal(t, ok, true)
			assert.NotNil(t, recvType)
		}
	}
}

func TestFuncDeclParams(t *testing.T) {
	src := `
		package main
		func Method(param1 int, param2 string) {}
	`
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, 0)
	assert.NoError(t, err)

	for _, decl := range node.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			params := FuncTypeParamsSeq(funcDecl.Type)
			assert.Equal(t, ok, true)
			assert.Len(t, params, 2)
		}
	}
}
