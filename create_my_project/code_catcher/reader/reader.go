package reader

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ct-zh/goLearn/create_my_project/code_catcher/types"
)

// Reader 代码读取器
type Reader struct {
	rootPath string
	sources  map[string]*types.Source // key: 文件路径
	fset     *token.FileSet           // 用于获取位置信息
}

// NewReader 创建新的Reader实例
func NewReader(rootPath string) *Reader {
	return &Reader{
		rootPath: rootPath,
		sources:  make(map[string]*types.Source),
		fset:     token.NewFileSet(),
	}
}

// ReadProject 读取整个项目
func (r *Reader) ReadProject() error {
	// 第一遍：解析所有文件的基本结构
	err := filepath.Walk(r.rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			if err := r.ReadGoFile(path); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// 第二遍：分析引用关系
	return filepath.Walk(r.rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			if err := r.analyzeReferences(path); err != nil {
				return err
			}
		}
		return nil
	})
}

// ReadGoFile 读取单个Go文件
func (r *Reader) ReadGoFile(filePath string) error {
	// 读取文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// 解析Go代码
	file, err := parser.ParseFile(r.fset, filePath, content, parser.ParseComments)
	if err != nil {
		return err
	}

	// 创建Source对象
	source := &types.Source{
		FilePath:    filePath,
		PackageName: file.Name.Name,
		CodeText:    content,
	}

	// 解析imports
	for _, imp := range file.Imports {
		importPath := strings.Trim(imp.Path.Value, "\"")
		source.Imports = append(source.Imports, importPath)
	}

	// 遍历AST
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			// 处理接口定义
			if t, ok := x.Type.(*ast.InterfaceType); ok {
				iface := r.parseInterface(x.Name.Name, t)
				source.Interfaces = append(source.Interfaces, iface)
			}
			// 处理结构体定义
			if t, ok := x.Type.(*ast.StructType); ok {
				st := r.parseStruct(x.Name.Name, t)
				source.Structs = append(source.Structs, st)
			}
		case *ast.FuncDecl:
			// 处理函数定义
			if x.Recv == nil { // 只处理非方法的函数
				fn := r.parseFunction(x)
				source.Functions = append(source.Functions, fn)
			} else {
				// 处理方法，将其添加到对应的结构体中
				r.parseMethod(x, source)
			}
		}
		return true
	})

	r.sources[filePath] = source
	return nil
}

// analyzeReferences 分析文件中的引用关系
func (r *Reader) analyzeReferences(filePath string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	file, err := parser.ParseFile(r.fset, filePath, content, parser.ParseComments)
	if err != nil {
		return err
	}

	var currentEntity *types.Entity // 当前正在处理的实体

	// 遍历AST查找引用
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			// 更新当前实体为函数或方法
			if x.Recv == nil {
				// 函数
				currentEntity = &types.Entity{
					Type:     types.EntityTypeFunction,
					Name:     x.Name.Name,
					FilePath: filePath,
				}
			} else {
				// 方法
				currentEntity = &types.Entity{
					Type:     types.EntityTypeMethod,
					Name:     x.Name.Name,
					FilePath: filePath,
				}
			}
		case *ast.SelectorExpr:
			// 处理选择器表达式（如：a.b）
			if currentEntity != nil {
				r.handleSelectorExpr(x, filePath, currentEntity)
			}
		case *ast.Ident:
			// 处理标识符
			if currentEntity != nil {
				r.handleIdentifier(x, filePath, currentEntity)
			}
		case *ast.CallExpr:
			// 处理函数调用
			if currentEntity != nil {
				r.handleFunctionCall(x, filePath, currentEntity)
			}
		case *ast.TypeAssertExpr:
			// 处理类型断言
			if currentEntity != nil {
				r.handleTypeAssertion(x, filePath, currentEntity)
			}
		case *ast.CompositeLit:
			// 处理复合字面量
			if currentEntity != nil {
				r.handleCompositeLit(x, filePath, currentEntity)
			}
		}
		return true
	})

	return nil
}

// handleSelectorExpr 处理选择器表达式
func (r *Reader) handleSelectorExpr(expr *ast.SelectorExpr, filePath string, caller *types.Entity) {
	// 如果父节点是标识符，可能是字段访问
	if _, ok := expr.X.(*ast.Ident); ok {
		pos := r.fset.Position(expr.Pos())
		ref := types.Reference{
			FilePath:   filePath,
			LineNumber: pos.Line,
			Context:    "field access",
			Caller:     caller,
		}

		// 遍历所有源文件查找匹配的结构体字段
		for _, source := range r.sources {
			for i := range source.Structs {
				// 检查字段
				for j := range source.Structs[i].Fields {
					if source.Structs[i].Fields[j].Name == expr.Sel.Name {
						r.addReference(source.Structs[i].Fields[j].References, ref)
					}
				}
			}
		}
	}
}

// handleIdentifier 处理标识符
func (r *Reader) handleIdentifier(ident *ast.Ident, filePath string, caller *types.Entity) {
	if ident.Obj == nil {
		return
	}

	pos := r.fset.Position(ident.Pos())
	ref := types.Reference{
		FilePath:   filePath,
		LineNumber: pos.Line,
		Context:    "identifier",
		Caller:     caller,
	}

	// 遍历所有源文件查找匹配的类型、函数等
	for _, source := range r.sources {
		// 检查结构体
		for i := range source.Structs {
			if source.Structs[i].Name == ident.Name {
				r.addReference(source.Structs[i].References, ref)
			}
		}
		// 检查接口
		for i := range source.Interfaces {
			if source.Interfaces[i].Name == ident.Name {
				r.addReference(source.Interfaces[i].References, ref)
			}
		}
		// 检查函数
		for i := range source.Functions {
			if source.Functions[i].Name == ident.Name {
				r.addReference(source.Functions[i].References, ref)
			}
		}
	}
}

// handleFunctionCall 处理函数调用
func (r *Reader) handleFunctionCall(call *ast.CallExpr, filePath string, caller *types.Entity) {
	pos := r.fset.Position(call.Pos())
	ref := types.Reference{
		FilePath:   filePath,
		LineNumber: pos.Line,
		Context:    "function call",
		Caller:     caller,
	}

	// 处理函数调用
	switch fun := call.Fun.(type) {
	case *ast.Ident:
		// 直接函数调用
		for _, source := range r.sources {
			for i := range source.Functions {
				if source.Functions[i].Name == fun.Name {
					r.addReference(source.Functions[i].References, ref)
				}
			}
		}
	case *ast.SelectorExpr:
		// 方法调用
		for _, source := range r.sources {
			for i := range source.Structs {
				for j := range source.Structs[i].Methods {
					if source.Structs[i].Methods[j].Name == fun.Sel.Name {
						r.addReference(source.Structs[i].Methods[j].References, ref)
					}
				}
			}
		}
	}
}

// handleTypeAssertion 处理类型断言
func (r *Reader) handleTypeAssertion(assert *ast.TypeAssertExpr, filePath string, caller *types.Entity) {
	pos := r.fset.Position(assert.Pos())
	ref := types.Reference{
		FilePath:   filePath,
		LineNumber: pos.Line,
		Context:    "type assertion",
		Caller:     caller,
	}

	if ident, ok := assert.Type.(*ast.Ident); ok {
		// 遍历所有源文件查找匹配的类型
		for _, source := range r.sources {
			// 检查结构体
			for i := range source.Structs {
				if source.Structs[i].Name == ident.Name {
					r.addReference(source.Structs[i].References, ref)
				}
			}
			// 检查接口
			for i := range source.Interfaces {
				if source.Interfaces[i].Name == ident.Name {
					r.addReference(source.Interfaces[i].References, ref)
				}
			}
		}
	}
}

// handleCompositeLit 处理复合字面量
func (r *Reader) handleCompositeLit(lit *ast.CompositeLit, filePath string, caller *types.Entity) {
	pos := r.fset.Position(lit.Pos())
	ref := types.Reference{
		FilePath:   filePath,
		LineNumber: pos.Line,
		Context:    "composite literal",
		Caller:     caller,
	}

	if ident, ok := lit.Type.(*ast.Ident); ok {
		// 遍历所有源文件查找匹配的类型
		for _, source := range r.sources {
			// 检查结构体
			for i := range source.Structs {
				if source.Structs[i].Name == ident.Name {
					r.addReference(source.Structs[i].References, ref)
				}
			}
		}
	}
}

// parseInterface 解析接口定义
func (r *Reader) parseInterface(name string, iface *ast.InterfaceType) types.Interface {
	result := types.Interface{
		Name:       name,
		References: make(map[types.ReferenceKey]types.Reference),
	}

	for _, method := range iface.Methods.List {
		if ft, ok := method.Type.(*ast.FuncType); ok {
			result.Methods = append(result.Methods, r.parseMethodType(method.Names[0].Name, ft))
		}
	}

	return result
}

// parseStruct 解析结构体定义
func (r *Reader) parseStruct(name string, st *ast.StructType) types.Struct {
	result := types.Struct{
		Name:       name,
		References: make(map[types.ReferenceKey]types.Reference),
	}

	for _, field := range st.Fields.List {
		if len(field.Names) > 0 {
			result.Fields = append(result.Fields, types.Field{
				Name:       field.Names[0].Name,
				Type:       r.typeToString(field.Type),
				References: make(map[types.ReferenceKey]types.Reference),
			})
		}
	}

	return result
}

// parseFunction 解析函数定义
func (r *Reader) parseFunction(fn *ast.FuncDecl) types.Function {
	return types.Function{
		Name:       fn.Name.Name,
		Params:     r.parseFieldList(fn.Type.Params),
		Results:    r.parseFieldList(fn.Type.Results),
		IsExported: ast.IsExported(fn.Name.Name),
		References: make(map[types.ReferenceKey]types.Reference),
	}
}

// parseMethod 解析方法并添加到对应的结构体中
func (r *Reader) parseMethod(fn *ast.FuncDecl, source *types.Source) {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		return
	}

	recvType := r.typeToString(fn.Recv.List[0].Type)
	method := r.parseMethodType(fn.Name.Name, fn.Type)

	// 找到对应的结构体并添加方法
	for i, st := range source.Structs {
		if st.Name == strings.TrimPrefix(recvType, "*") {
			source.Structs[i].Methods = append(source.Structs[i].Methods, method)
			break
		}
	}
}

// parseMethodType 解析方法类型
func (r *Reader) parseMethodType(name string, ft *ast.FuncType) types.Method {
	return types.Method{
		Name:       name,
		Params:     r.parseFieldList(ft.Params),
		Results:    r.parseFieldList(ft.Results),
		References: make(map[types.ReferenceKey]types.Reference),
	}
}

// parseFieldList 解析字段列表
func (r *Reader) parseFieldList(fields *ast.FieldList) []types.Field {
	var result []types.Field
	if fields == nil {
		return result
	}

	for _, field := range fields.List {
		fieldType := r.typeToString(field.Type)
		if len(field.Names) > 0 {
			for _, name := range field.Names {
				result = append(result, types.Field{
					Name:       name.Name,
					Type:       fieldType,
					References: make(map[types.ReferenceKey]types.Reference),
				})
			}
		} else {
			// 匿名字段或者返回值没有命名
			result = append(result, types.Field{
				Type:       fieldType,
				References: make(map[types.ReferenceKey]types.Reference),
			})
		}
	}
	return result
}

// typeToString 将AST类型转换为字符串
func (r *Reader) typeToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + r.typeToString(t.X)
	case *ast.SelectorExpr:
		return r.typeToString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + r.typeToString(t.Elt)
	case *ast.MapType:
		return "map[" + r.typeToString(t.Key) + "]" + r.typeToString(t.Value)
	case *ast.InterfaceType:
		return "interface{}"
	default:
		return ""
	}
}

// addReference 添加引用信息
func (r *Reader) addReference(refs map[types.ReferenceKey]types.Reference, ref types.Reference) {
	key := types.ReferenceKey{
		FilePath:   ref.FilePath,
		LineNumber: ref.LineNumber,
		CallerName: ref.Caller.Name,
	}
	refs[key] = ref
}

// GetSources 获取所有解析的源文件
func (r *Reader) GetSources() map[string]*types.Source {
	return r.sources
}
