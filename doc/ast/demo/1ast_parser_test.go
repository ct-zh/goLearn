package demo

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// ast解析文件或者解析文件夹
// parser.ParseFile
// parser.ParseDir
//

func Test_AstParse(t *testing.T) {
	// 要解析的Go源代码
	src := `
	package main

	import "fmt"

	func main() {
		fmt.Println("Hello, world!")
	}
	`

	// 解析源代码为AST
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	// 深度优先遍历AST
	ast.Inspect(node, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			fmt.Println("函数名称:", fn.Name.Name)
		}
		return true
	})
}

// ast.Node 有两个方法Pos和End，代表这个变量在源代码中的起始位置和结束位置
func Test_AstNode(t *testing.T) {
	// 要解析的Go源代码
	src := `
	package main
	
	import "fmt"
	
	func main() {
		fmt.Println("Hello, world!")
	}
	`

	// 解析源代码为AST
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	// 遍历AST查找函数调用表达式
	ast.Inspect(node, func(n ast.Node) bool {
		if callExpr, ok := n.(*ast.CallExpr); ok {
			pos := fset.Position(callExpr.Pos())
			end := fset.Position(callExpr.End())
			fmt.Printf("函数调用表达式: %+v\n", callExpr)
			fmt.Printf("起始位置: %s\n", pos)
			fmt.Printf("结束位置: %s\n", end)
		}
		return true
	})
}
