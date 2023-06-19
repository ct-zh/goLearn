package any_parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type AnyObject struct {
	ImportPath string
	MyType     ast.Node
	Methods    *ast.FieldList
}

func FindObject(path string) (map[string]*AnyObject, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, isGoFile, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	objects := make(map[string]*AnyObject)
	for _, pkg := range pkgs {

		fmt.Printf("\n pkg imports: %+v\n", pkg.Imports)
		for importPath, file := range pkg.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				switch node := n.(type) {
				case *ast.TypeSpec:
					typeName := node.Name.Name
					if len(node.Name.Name) == 0 {
						return true
					}
					if node.Type == nil {
						return true
					}
					methods := extractMethods(node)
					if len(methods.List) > 0 {
						object := &AnyObject{
							ImportPath: importPath,
							MyType:     node.Type,
							Methods:    methods,
						}
						objects[typeName] = object
					}
				}
				return true
			})
		}
	}

	return objects, nil
}

func extractMethods(typeSpec *ast.TypeSpec) *ast.FieldList {
	switch typ := typeSpec.Type.(type) {
	case *ast.InterfaceType:
		return typ.Methods
	case *ast.StructType:
		return typ.Fields
	default:
		return nil
	}
}

type AnyFunc struct {
	Name string
	Func *ast.FuncDecl
}
