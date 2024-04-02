package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func Test_findFunc(t *testing.T) {
	pkgs, err := parser.ParseDir(token.NewFileSet(), "./", nil, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	var fn *ast.FuncDecl
	for _, pack := range pkgs {
		tmpFn := findFunc(pack, "findFunc")
		if tmpFn != nil {
			fn = tmpFn
			break
		}
	}
	if fn == nil {
		t.Fatal("not found")
	}
	t.Log("success")
}

func Test_findFuncFilterParams(t *testing.T) {
	pkgs, err := parser.ParseDir(token.NewFileSet(), "./", nil, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}

	params := []fileParams{
		{
			prefix: "ast",
			name:   "Package",
		},
		{
			prefix: "",
			name:   "string",
		},
	}
	var fn *ast.FuncDecl
	for _, pack := range pkgs {
		tmpFn := findFuncFilterParams(pack, "findFunc", params)
		if tmpFn != nil {
			fn = tmpFn
			break
		}
	}
	if fn == nil {
		t.Fatal("not found")
	}
	t.Log("success")
}
