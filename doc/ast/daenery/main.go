package main

import (
	"go/ast"
)

var checkDirMap = []string{
	"api",
	"app",
	"conf",
	"dao",
	"manager",
	"model",
	"server",
	"service",
}

func main() {
	cfg := NewConfig()

	daeParser := NewDaenerysParser(cfg.Path)

	parseRouter(daeParser)

	return
}

// 1. 目标1 解析router
func parseRouter(daeParser *DaenerysParser) {

	server, ok := daeParser.dir["server"]
	if !ok {
		panic("server dir not exist")
	}

	httpFiles, ok := server.files["http"]
	if !ok {
		panic("http file not exist")
	}

	// 获取一个函数，他的名称是initRoute, 并且只有一个参数, 参数包前缀是httpserver 包名是Server
	initRouteFn := findFuncFilterParams(httpFiles, "initRoute", []fileParams{
		{"httpserver", "Server"},
	})

	//fmt.Printf("initRouteFn = %+v \n", initRouteFn)

	// httpserver.Server这个变量名称
	httpServerValName := initRouteFn.Type.Params.List[0].Names[0].Name

	router := &Router{
		Items: make([]*RouterItem, 0),
	}

	groupMap := make(map[string]*RouterGroup)

	// 获取initRoute函数里注册router的列表
	for _, stmt := range initRouteFn.Body.List {
		switch itemStmt := stmt.(type) {
		case *ast.AssignStmt:
			if len(itemStmt.Lhs) != 1 || len(itemStmt.Rhs) < 1 {
				continue
			}
			call, ok := itemStmt.Rhs[0].(*ast.CallExpr)
			if !ok {
				panic("not call expr")
			}
			if call.Fun.(*ast.SelectorExpr).X.(*ast.Ident).Name != httpServerValName {
				continue
			}
			if HttpserverRequestType(call.Fun.(*ast.SelectorExpr).Sel.Name) != Group {
				continue
			}
			if len(call.Args) < 1 {
				continue
			}
			callParams1, ok := call.Args[0].(*ast.BasicLit)
			if !ok {
				continue
			}

			// 返回变量的名称
			valName := itemStmt.Lhs[0].(*ast.Ident).Name

			groupMap[valName] = &RouterGroup{
				Name:   valName,
				Prefix: callParams1.Value,
			}
		case *ast.ExprStmt:
			// 获取函数调用
			call, ok := itemStmt.X.(*ast.CallExpr)
			if !ok {
				panic("not call expr")
			}

			// 要保证call方法是 httpserver.Server这个变量调用的方法
			// 或者是group的分组方法
			var group *RouterGroup
			valName := call.Fun.(*ast.SelectorExpr).X.(*ast.Ident).Name
			if valName != httpServerValName {
				if itemGroup, ok := groupMap[valName]; ok {
					group = itemGroup
				} else {
					continue
				}
			}

			// 这个方法至少有两个参数：
			if len(call.Args) < 2 {
				continue
			}

			// call.Args 这是 s.GET("/ping", httpserver.Ping) 这个方法的参数切片
			// 第一个参数返回的是 *ast.BasicLit
			// 后面则是注册指定的函数，需要到这个包的函数列表里面去找
			// 如果有多个函数，这里只能默认最后一个函数是实际执行service的函数
			callParams1, ok := call.Args[0].(*ast.BasicLit)
			if !ok {
				continue
			}
			callParamsLast, ok := call.Args[len(call.Args)-1].(*ast.Ident)
			if !ok {
				continue
			}

			routerItem := &RouterItem{
				Method:   HttpserverRequestType(call.Fun.(*ast.SelectorExpr).Sel.Name),
				Uri:      callParams1.Value,
				FuncName: callParamsLast.Name,
				Group:    group,
			}

			//printType(routerItem)

			router.Items = append(router.Items, routerItem)
		case *ast.BlockStmt:
			for _, blockItem := range itemStmt.List {
				printType(blockItem)
			}
		default:
			continue
		}

	}

	// 获取initRoute里的所有函数调用
	// 解析 s.ANY、s.POST 、 s.GET、s.GROUP
	// 目标：拿到uri
	//ast.Inspect(initRouteFn, func(node ast.Node) bool {
	//	if node == nil {
	//		return true
	//	}
	//	switch n := node.(type) {
	//	case *ast.ExprStmt:
	//		fmt.Printf("n = [%T]%+v x=%+v \n", n, n, n.X)
	//
	//	}
	//
	//	return true
	//})

}
