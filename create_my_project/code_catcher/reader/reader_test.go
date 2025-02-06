package reader

import (
	"os"
	"path/filepath"
	"testing"

	codecatcher "github.com/ct-zh/goLearn/create_my_project/code_catcher"
	"github.com/stretchr/testify/assert"
)

// 创建测试用的临时Go文件
func createTestFile(t *testing.T, dir, filename, content string) string {
	path := filepath.Join(dir, filename)
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatalf("无法创建测试文件: %v", err)
	}
	return path
}

func TestReader_ReadProject(t *testing.T) {
	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "reader_test")
	if err != nil {
		t.Fatalf("无法创建临时目录: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建测试文件
	mainPath := createTestFile(t, tmpDir, "main.go", `
package main

import "fmt"

type TestInterface interface {
	TestMethod() string
}

type TestStruct struct {
	Field1 string
	Field2 int
}

func (t *TestStruct) TestMethod() string {
	return t.Field1
}

func main() {
	ts := &TestStruct{Field1: "test"}
	fmt.Println(ts.TestMethod())
}
`)

	// 创建Reader实例
	reader := NewReader(tmpDir)

	// 测试ReadProject
	err = reader.ReadProject()
	assert.NoError(t, err)

	// 获取解析结果
	sources := reader.GetSources()
	mainFile := sources[mainPath]
	assert.NotNil(t, mainFile)

	// 测试包名解析
	assert.Equal(t, "main", mainFile.PackageName)

	// 测试接口解析
	assert.Equal(t, 1, len(mainFile.Interfaces))
	assert.Equal(t, "TestInterface", mainFile.Interfaces[0].Name)
	assert.Equal(t, 1, len(mainFile.Interfaces[0].Methods))
	assert.Equal(t, "TestMethod", mainFile.Interfaces[0].Methods[0].Name)

	// 测试结构体解析
	assert.Equal(t, 1, len(mainFile.Structs))
	assert.Equal(t, "TestStruct", mainFile.Structs[0].Name)
	assert.Equal(t, 2, len(mainFile.Structs[0].Fields))
	assert.Equal(t, "Field1", mainFile.Structs[0].Fields[0].Name)
	assert.Equal(t, "string", mainFile.Structs[0].Fields[0].Type)
	assert.Equal(t, "Field2", mainFile.Structs[0].Fields[1].Name)
	assert.Equal(t, "int", mainFile.Structs[0].Fields[1].Type)

	// 测试方法解析
	assert.Equal(t, 1, len(mainFile.Structs[0].Methods))
	assert.Equal(t, "TestMethod", mainFile.Structs[0].Methods[0].Name)

	// 测试引用分析
	// TestMethod被main函数调用
	methodRefs := mainFile.Structs[0].Methods[0].References
	assert.Equal(t, 1, len(methodRefs))

	// 获取引用信息
	var ref codecatcher.Reference
	for _, r := range methodRefs {
		ref = r
		break
	}

	assert.Equal(t, "function call", ref.Context)
	assert.Equal(t, "main", ref.Caller.Name)
}

func TestReader_ReadGoFile(t *testing.T) {
	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "reader_test")
	if err != nil {
		t.Fatalf("无法创建临时目录: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建测试文件
	filePath := createTestFile(t, tmpDir, "test.go", `
package test

type TestStruct struct {
	Field string
}

func (t *TestStruct) Method() {}
`)

	// 创建Reader实例
	reader := NewReader(tmpDir)

	// 测试ReadGoFile
	err = reader.ReadGoFile(filePath)
	assert.NoError(t, err)

	// 获取解析结果
	sources := reader.GetSources()
	assert.Equal(t, 1, len(sources))

	source := sources[filePath]
	assert.NotNil(t, source)
	assert.Equal(t, "test", source.PackageName)
	assert.Equal(t, 1, len(source.Structs))
	assert.Equal(t, "TestStruct", source.Structs[0].Name)
	assert.Equal(t, 1, len(source.Structs[0].Fields))
	assert.Equal(t, 1, len(source.Structs[0].Methods))
}

func TestReader_ComplexCode(t *testing.T) {
	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "reader_test")
	if err != nil {
		t.Fatalf("无法创建临时目录: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建多个相互依赖的测试文件
	interfacePath := createTestFile(t, tmpDir, "interface.go", `
package test

type Handler interface {
	Handle(data string) error
}
`)

	implPath := createTestFile(t, tmpDir, "impl.go", `
package test

import "fmt"

type MyHandler struct {
	prefix string
}

func (h *MyHandler) Handle(data string) error {
	fmt.Println(h.prefix + data)
	return nil
}

func NewHandler(prefix string) Handler {
	return &MyHandler{prefix: prefix}
}
`)

	userPath := createTestFile(t, tmpDir, "user.go", `
package test

func UseHandler(h Handler) {
	h.Handle("test")
}
`)

	// 创建Reader实例并读取项目
	reader := NewReader(tmpDir)
	err = reader.ReadProject()
	assert.NoError(t, err)

	sources := reader.GetSources()
	assert.Equal(t, 3, len(sources))

	// 验证接口实现关系
	interfaceFile := sources[interfacePath]
	implFile := sources[implPath]
	userFile := sources[userPath]

	assert.NotNil(t, interfaceFile)
	assert.NotNil(t, implFile)
	assert.NotNil(t, userFile)

	// 验证接口定义
	assert.Equal(t, 1, len(interfaceFile.Interfaces))
	handler := interfaceFile.Interfaces[0]
	assert.Equal(t, "Handler", handler.Name)

	// 验证结构体实现
	assert.Equal(t, 1, len(implFile.Structs))
	myHandler := implFile.Structs[0]
	assert.Equal(t, "MyHandler", myHandler.Name)

	// 验证方法引用
	found := false
	for _, method := range myHandler.Methods {
		if method.Name == "Handle" {
			for _, ref := range method.References {
				if ref.Caller.Name == "UseHandler" {
					found = true
					break
				}
			}
		}
	}
	assert.True(t, found, "应该找到UseHandler中对Handle方法的调用")
}

func TestReader_ParseErrors(t *testing.T) {
	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "reader_test")
	if err != nil {
		t.Fatalf("无法创建临时目录: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建语法错误的Go文件
	createTestFile(t, tmpDir, "error.go", `
package test

this is not valid go code
`)

	reader := NewReader(tmpDir)
	err = reader.ReadProject()
	assert.Error(t, err, "应该返回语法错误")
}

func TestReader_EmptyProject(t *testing.T) {
	// 创建空的临时目录
	tmpDir, err := os.MkdirTemp("", "reader_test")
	if err != nil {
		t.Fatalf("无法创建临时目录: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	reader := NewReader(tmpDir)
	err = reader.ReadProject()
	assert.NoError(t, err)

	sources := reader.GetSources()
	assert.Equal(t, 0, len(sources), "空项目不应该有任何源文件")
}
