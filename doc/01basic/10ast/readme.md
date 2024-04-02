# ast


## ast.Package
拿到路径下所有go文件packages: `itemPacks, err := parser.ParseDir(token.NewFileSet(), path, dirFilter, parser.AllErrors)`
    
可以获得*ast.Package 变量

有了*ast.Package变量，使用`ast.Inspect()` 深度优先遍历，获得`ast.Node`类型变量；

> [示例 - Test_AstParse](./demo/1ast_parser_test.go)

### ast.Node
`ast.Node`可能的值常见有：
- `*ast.File`：代表整个源文件。
- `*ast.Package`：代表一个包的信息。
- `*ast.Ident`：代表标识符，如变量名、函数名等。
- `*ast.FuncDecl`：代表函数或方法的声明。
- `*ast.TypeSpec`：代表类型定义，如结构体、接口、别名等。
- `*ast.GenDecl`：代表通用的声明，如import、const、var等。
- `*ast.AssignStmt`：代表赋值语句。
- `*ast.ExprStmt`：代表表达式语句，如函数调用、方法调用等。
- `*ast.CallExpr`：代表函数或方法的调用表达式。
- `*ast.ReturnStmt`：代表函数或方法的返回语句。
- `*ast.IfStmt`：代表if语句。
- `*ast.ForStmt`：代表for循环语句。
- `*ast.RangeStmt`：代表range循环语句。
- `*ast.ImportSpec`：代表导入语句的规范。
- `*ast.StructType`：代表结构体类型。
- `*ast.InterfaceType`：代表接口类型。

还可能是其他值，见ast包。不过这些值都实现了`ast.Node`的两个方法`Pos()`和`End()`，用于标识源代码的起始位置和结束位置。

> [示例 - 见Test_AstNode](./demo/1ast_parser_test.go)

## 代码解析
```go
package main  // *ast.Ident main

type Person struct {    // *ast.TypeSpec(Name:Person)|*ast.Ident Person|*ast.StructType
    // *ast.FieldList 结构体的字段列表(通过for range node.List:可以获取到字段Name和Age)
    Name string     // *ast.Field(Names:Name,Type:string)|*ast.Ident Name|*ast.Ident string
    Age  int        // *ast.Field(Names:Age,Type:int)|*ast.Ident Age|*ast.Ident int
}

func SayHello(name string) (err error) {    // *ast.FuncDecl （Name:SayHello）|*ast.Ident SayHello|*ast.FuncType
    // *ast.FieldList (通过for range node.List:可以获取到字段name)
    // *ast.Field Names:[name] Type:string | *ast.Ident name| *ast.Ident string
    // *ast.FieldList (字段 err)
    // *ast.Field Names:[err] Type:error | *ast.Ident err | *ast.Ident error


    fmt.Println("Hello, " + name)   // *ast.SelectorExpr &{X:fmt Sel:Println}| *ast.Ident fmt| *ast.Ident Println
    // *ast.BinaryExpr &{X:0xc0001501e0 OpPos:136 Op:+ Y:name}
    // *ast.BasicLit &{ValuePos:126 Kind:STRING Value:"Hello, "}
    // *ast.Ident name
}

```


## ast.Field && ast.FieldList
表示结构体字段或函数参数/结果

具有以下字段
- `Names`字段是一个标识符列表，表示字段或参数/结果的名称。对于匿名字段或只有一个字段/参数/结果的情况，可以忽略该字段。
- `Type`字段是表示字段或参数/结果类型的表达式，可以是标识符、指针、数组、映射等类型表达式。


## ast.FuncDecl
表示一个方法或者函数的声明。

- fn.Name.Name可以获取函数的名称
- fn.Type.Params.List可以获取参数列表
- 参数列表可能有多种类型：
    - `*ast.SelectorExpr`类型名称表达式，例如：`user model.User`，返回的是`model.User`
    - `*ast.StarExpr`指针名称表达式。例如：`user *model.Use`，返回的是`*model.User`
    - `*ast.Ident`可能是基础类型，如`name string`



