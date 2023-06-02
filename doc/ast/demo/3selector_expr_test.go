package demo

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// *ast.selectorExpr 选择器表达式
// 它由一个包名或类型名称作为前缀，紧跟着一个点号（.），再跟着标识符名称。
// 例如，fmt.Println中的fmt就是一个选择器表达式，表示访问fmt包中的标识符Println。
// *ast.SelectorExpr表示了一个选择器表达式的AST节点。它包含了两个重要的字段：
// X字段是选择器表达式的前缀部分，表示包名或类型名称的表达式。它可以是另一个选择器表达式、标识符表达式等。
// Sel字段是选择器表达式的标识符部分，表示访问的标识符名称。

func Test_SelectorExpr(t *testing.T) {
	// 要解析的Go源代码
	src := `
	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		fmt.Println("Hello, world!")
		time.Sleep(time.Second)
	}
	`

	// 解析源代码为AST
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	// 遍历AST查找选择器表达式
	ast.Inspect(node, func(n ast.Node) bool {
		if selExpr, ok := n.(*ast.SelectorExpr); ok {
			fmt.Printf("选择器表达式: %s\n", fset.Position(selExpr.Pos()))
			fmt.Printf("前缀部分: %s\n", selExpr.X)
			fmt.Printf("标识符部分: %s\n", selExpr.Sel)
		}
		return true
	})
}

func Test_StarExpr(t *testing.T) {
	// 要解析的Go源代码
	src := `
	package main
	
	import "fmt"
	
	func main() {
		var ptr *int
		fmt.Println(ptr)
	}
	`

	// 解析源代码为AST
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	// 遍历AST查找指针类型
	ast.Inspect(node, func(n ast.Node) bool {
		if starExpr, ok := n.(*ast.StarExpr); ok {
			typeExpr := starExpr.X
			fmt.Printf("指针类型: %s\n", typeExpr)
		}
		return true
	})
}
