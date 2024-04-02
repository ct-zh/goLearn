package any_parser

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
)

var TypeErr = errors.New("path type error")

type AnyDir struct {
	path string

	dir  map[string]*AnyDir
	file map[string]*AnyFile

	object map[string]*AnyObject
	funcs  map[string]*AnyFunc
}

func NewAnyDir(path string) (*AnyDir, error) {
	dir, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !dir.IsDir() {
		return nil, TypeErr
	}
	info, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	anyDir := &AnyDir{
		path:   path,
		dir:    make(map[string]*AnyDir),
		file:   make(map[string]*AnyFile),
		object: make(map[string]*AnyObject),
		funcs:  make(map[string]*AnyFunc),
	}
	for _, fileInfo := range info {
		if fileInfo.IsDir() {
			newPath := path + "/" + fileInfo.Name()
			itemDir, err := NewAnyDir(newPath)
			if err != nil {
				return nil, err
			}
			anyDir.dir[fileInfo.Name()] = itemDir
		} else if isGoFile(fileInfo) {
			itemFile, err := NewAnyFile(path + "/" + fileInfo.Name())
			if err != nil {
				return nil, err
			}
			anyDir.file[fileInfo.Name()] = itemFile
			for name, fn := range itemFile.Funcs {
				anyDir.funcs[name] = fn
			}

			itemObj, err := FindObject(path + "/" + fileInfo.Name())
			if err != nil {
				return nil, err
			}
			for s, object := range itemObj {
				anyDir.object[s] = object
			}
		}
	}

	return anyDir, nil
}

type AnyFile struct {
	Name   string
	MyType *ast.File
	Funcs  map[string]*AnyFunc
}

func NewAnyFile(path string) (*AnyFile, error) {
	// 初始化 AnyFile 对象
	anyFile := &AnyFile{
		Name:  filepath.Base(path),
		Funcs: make(map[string]*AnyFunc),
	}

	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	anyFile.MyType = astFile

	for _, decl := range astFile.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			funcName := funcDecl.Name.Name
			anyFunc := &AnyFunc{
				Name: funcName,
				Func: funcDecl,
			}
			anyFile.Funcs[funcName] = anyFunc
		}
	}

	return anyFile, nil
}
