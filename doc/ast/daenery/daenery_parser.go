package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type DaenerysParser struct {
	name string
	path string
	dir  map[string]*DirParser
}

type DirParser struct {
	name  string
	path  string
	dir   map[string]*DirParser
	files map[string]*ast.Package
}

func NewDaenerysParser(path string) *DaenerysParser {
	// ==== 解析路径 拿到当前服务名
	splits := strings.Split(path, "/")
	if len(splits) == 0 {
		panic("invalid path")
	}
	name := splits[len(splits)-1]
	if len(name) == 0 {
		name = splits[len(splits)-2]
	}

	// ==== 检查传入的路径是否符合框架结构
	dir, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if !dir.IsDir() {
		panic("this path is a file, not dir")
	}
	info, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	checkDir := map[string]struct{}{}
	for _, s := range checkDirMap {
		checkDir[s] = struct{}{}
	}

	dirs := map[string]*DirParser{}
	for _, fileInfo := range info {
		if fileInfo.IsDir() {
			if _, ok := checkDir[fileInfo.Name()]; ok {

				// 遍历当前dir下所有的文件，将所有包写入到map中
				packages := make(map[string]*ast.Package)
				filepath.Walk(filepath.Join(path, fileInfo.Name()), func(path string, info fs.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if info.IsDir() {
						itemPacks, err := parser.ParseDir(token.NewFileSet(), path, dirFilter, parser.AllErrors)
						if err != nil {
							panic(err)
						}
						for s, a := range itemPacks {
							//fmt.Printf("name = %s pack = %+v \n ", s, a)
							packages[s] = a
						}
					}
					return nil
				})

				dirs[fileInfo.Name()] = &DirParser{
					name:  fileInfo.Name(),
					path:  filepath.Join(path, fileInfo.Name()),
					files: packages,
				}
				delete(checkDir, fileInfo.Name())
			}
		}
	}

	if len(checkDir) != 0 {
		panic(fmt.Errorf("dir not match: %+v \n", checkDir))
	}

	return &DaenerysParser{name: name, path: path, dir: dirs}
}

func dirFilter(file fs.FileInfo) bool {
	return true
}
