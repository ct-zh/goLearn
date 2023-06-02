package demo

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func Test_ParseField(t *testing.T) {
	// 要解析的Go源代码
	src := `
	package main
	
	type Person struct {
		Name string
		Age  int
	}
	
	func SayHello(name string) {
		fmt.Println("Hello, " + name)
	}
	`

	// 解析源代码为AST
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	// 遍历AST查找结构体字段和函数的参数/结果
	ast.Inspect(node, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.Field:
			// 结构体字段
			fmt.Printf("结构体字段: %s\n", node.Names[0].Name)
			fmt.Printf("字段类型: %s\n", node.Type)
			fmt.Printf("\n")
		case *ast.FieldList:
			// 函数的参数/结果列表
			for _, field := range node.List {
				fmt.Printf("参数/结果名称: %s\n", field.Names[0].Name)
				fmt.Printf("参数/结果类型: %s\n", field.Type)
				fmt.Printf("\n")
			}
		}
		return true
	})
}
