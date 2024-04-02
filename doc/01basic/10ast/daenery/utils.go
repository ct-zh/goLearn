package main

import (
	"fmt"
	"go/ast"
)

type fileParams struct {
	prefix string
	name   string
}

func printType(a any) {
	fmt.Printf("[type=%T] %+v \n", a, a)
}

// 从package中找到对应名称的函数
func findFunc(pack *ast.Package, name string) *ast.FuncDecl {
	var nodeFn *ast.FuncDecl
	ast.Inspect(pack, func(node ast.Node) bool {
		if fn, ok := node.(*ast.FuncDecl); ok {
			if fn.Name.Name == name {
				nodeFn = fn
				return false
			}
		}
		return true
	})
	return nodeFn
}

// 从package中找到对应名称的函数，比较参数数量与类型是否一致
func findFuncFilterParams(pack *ast.Package, name string, paramsFilter []fileParams) *ast.FuncDecl {
	var nodeFn *ast.FuncDecl
	ast.Inspect(pack, func(node ast.Node) bool {
		if fn, ok := node.(*ast.FuncDecl); ok {
			if fn.Name.Name != name {
				return true
			}
			if len(fn.Type.Params.List) != len(paramsFilter) {
				return true
			}
			for i, params := range fn.Type.Params.List {
				//paramsType := params.Type
				//fmt.Printf("paramsType = [%T]%+v \n", params.Type, params.Type)
				switch paramsType := params.Type.(type) {
				case *ast.SelectorExpr: //
					selectorPrefix := fmt.Sprintf("%s", paramsType.X)
					selectorName := fmt.Sprintf("%s", paramsType.Sel)
					//fmt.Printf("selectorPrefix=%s selectorName=%s \n", selectorPrefix, selectorName)
					if paramsFilter[i].prefix != selectorPrefix || paramsFilter[i].name != selectorName {
						return true
					}
				case *ast.StarExpr:
					if selectorExpr, ok := paramsType.X.(*ast.SelectorExpr); ok {
						selectorPrefix := fmt.Sprintf("%s", selectorExpr.X)
						selectorName := fmt.Sprintf("%s", selectorExpr.Sel)
						//fmt.Printf("selectorPrefix=%s selectorName=%s \n", selectorPrefix, selectorName)
						if paramsFilter[i].prefix != selectorPrefix || paramsFilter[i].name != selectorName {
							return true
						}
					} else {
						return true
					}
				case *ast.Ident:
					selectorName := fmt.Sprintf("%s", paramsType.Name)
					//fmt.Printf("selectorName=%s paramsName=%s %t \n", selectorName, paramsFilter[i].name, selectorName == paramsFilter[i].name)
					if paramsFilter[i].name != selectorName {
						return true
					}
				}
			}
			nodeFn = fn
			return false
		}
		return true
	})
	return nodeFn
}
