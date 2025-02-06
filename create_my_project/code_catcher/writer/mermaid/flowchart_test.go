package mermaid

import (
	"strings"
	"testing"

	codecatcher "github.com/ct-zh/goLearn/create_my_project/code_catcher"
	"github.com/stretchr/testify/assert"
)

// createTestSource 创建测试用的Source对象
func createTestSource(pkg string, imports []string) *codecatcher.Source {
	return &codecatcher.Source{
		PackageName: pkg,
		Imports:     imports,
	}
}

func TestGeneratePackageDependencies(t *testing.T) {
	// 创建测试数据
	sources := map[string]*codecatcher.Source{
		"main.go": createTestSource("main", []string{
			"fmt",
			"strings",
		}),
		"helper.go": createTestSource("main", []string{
			"encoding/json",
			"fmt",
		}),
	}

	writer := NewFlowchartWriter(sources)
	result := writer.GeneratePackageDependencies()

	// 验证结果
	assert.Contains(t, result, "graph TD")
	assert.Contains(t, result, "main[main]")
	assert.Contains(t, result, "fmt[fmt]")
	assert.Contains(t, result, "strings[strings]")
	assert.Contains(t, result, "json[json]")
	assert.Contains(t, result, "main --> fmt")
	assert.Contains(t, result, "main --> strings")
	assert.Contains(t, result, "main --> json")
}

func TestGenerateEntityReferences(t *testing.T) {
	// 创建测试数据
	source := &codecatcher.Source{
		PackageName: "test",
		Structs: []codecatcher.Struct{
			{
				Name: "TestStruct",
				Methods: []codecatcher.Method{
					{
						Name: "TestMethod",
						References: map[codecatcher.ReferenceKey]codecatcher.Reference{
							{FilePath: "main.go", LineNumber: 10, CallerName: "main"}: {
								FilePath:   "main.go",
								LineNumber: 10,
								Context:    "function call",
								Caller: &codecatcher.Entity{
									Type: codecatcher.EntityTypeFunction,
									Name: "main",
								},
							},
						},
					},
				},
			},
		},
		Functions: []codecatcher.Function{
			{
				Name: "HelperFunc",
				References: map[codecatcher.ReferenceKey]codecatcher.Reference{
					{FilePath: "main.go", LineNumber: 15, CallerName: "TestMethod"}: {
						FilePath:   "main.go",
						LineNumber: 15,
						Context:    "function call",
						Caller: &codecatcher.Entity{
							Type: codecatcher.EntityTypeMethod,
							Name: "TestMethod",
						},
					},
				},
			},
		},
	}

	sources := map[string]*codecatcher.Source{
		"test.go": source,
	}

	writer := NewFlowchartWriter(sources)
	result := writer.GenerateEntityReferences()

	// 验证结果
	assert.Contains(t, result, "graph TD")
	assert.Contains(t, result, "struct_TestStruct[TestStruct]")
	assert.Contains(t, result, "method_TestStruct_TestMethod[TestStruct.TestMethod]")
	assert.Contains(t, result, "func_HelperFunc[HelperFunc]")
	assert.Contains(t, result, "caller_main[main]")
	assert.Contains(t, result, "-->|function call|")
}

func TestGenerateInterfaceImplementations(t *testing.T) {
	// 创建测试数据
	source := &codecatcher.Source{
		PackageName: "test",
		Interfaces: []codecatcher.Interface{
			{
				Name: "TestInterface",
				Methods: []codecatcher.Method{
					{Name: "TestMethod"},
				},
			},
		},
		Structs: []codecatcher.Struct{
			{
				Name:       "TestStruct",
				Implements: []string{"TestInterface"},
				Methods: []codecatcher.Method{
					{Name: "TestMethod"},
				},
			},
		},
	}

	sources := map[string]*codecatcher.Source{
		"test.go": source,
	}

	writer := NewFlowchartWriter(sources)
	result := writer.GenerateInterfaceImplementations()

	// 验证结果
	assert.Contains(t, result, "graph TD")
	assert.Contains(t, result, "interface_TestInterface[(TestInterface)]")
	assert.Contains(t, result, "struct_TestStruct[TestStruct]")
	assert.Contains(t, result, "-.->|implements|")
}

func TestSanitizeID(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"normal", "normal"},
		{"with space", "with_space"},
		{"with-dash", "with_dash"},
		{"with.dot", "with_dot"},
		{"with/slash", "with_slash"},
		{"with$special@chars", "with_special_chars"},
	}

	for _, test := range tests {
		result := sanitizeID(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestCleanImportPath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"github.com/user/repo", "repo"},
		{"github.com/user/repo@v1.0.0", "repo"},
		{"encoding/json", "json"},
		{"fmt", "fmt"},
	}

	for _, test := range tests {
		result := cleanImportPath(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestContainsInterface(t *testing.T) {
	implements := []string{"Interface1", "Interface2", "Interface3"}

	// 测试存在的接口
	assert.True(t, containsInterface(implements, "Interface1"))
	assert.True(t, containsInterface(implements, "Interface2"))
	assert.True(t, containsInterface(implements, "Interface3"))

	// 测试不存在的接口
	assert.False(t, containsInterface(implements, "Interface4"))
	assert.False(t, containsInterface(implements, ""))
}

func TestFlowchartWriterWithEmptySources(t *testing.T) {
	writer := NewFlowchartWriter(make(map[string]*codecatcher.Source))

	// 测试包依赖关系图
	pkgResult := writer.GeneratePackageDependencies()
	assert.Equal(t, "graph TD\n", pkgResult)

	// 测试实体引用关系图
	entityResult := writer.GenerateEntityReferences()
	assert.Equal(t, "graph TD\n", entityResult)

	// 测试接口实现关系图
	interfaceResult := writer.GenerateInterfaceImplementations()
	assert.Equal(t, "graph TD\n", interfaceResult)
}

func TestFlowchartWriterWithNilSources(t *testing.T) {
	writer := NewFlowchartWriter(nil)

	// 测试包依赖关系图
	pkgResult := writer.GeneratePackageDependencies()
	assert.Equal(t, "graph TD\n", pkgResult)

	// 测试实体引用关系图
	entityResult := writer.GenerateEntityReferences()
	assert.Equal(t, "graph TD\n", entityResult)

	// 测试接口实现关系图
	interfaceResult := writer.GenerateInterfaceImplementations()
	assert.Equal(t, "graph TD\n", interfaceResult)
}

func TestMermaidSyntax(t *testing.T) {
	// 创建一个包含所有类型节点的复杂测试用例
	source := &codecatcher.Source{
		PackageName: "test",
		Imports:     []string{"fmt", "strings"},
		Interfaces: []codecatcher.Interface{
			{
				Name: "TestInterface",
				Methods: []codecatcher.Method{
					{Name: "TestMethod"},
				},
			},
		},
		Structs: []codecatcher.Struct{
			{
				Name:       "TestStruct",
				Implements: []string{"TestInterface"},
				Methods: []codecatcher.Method{
					{
						Name: "TestMethod",
						References: map[codecatcher.ReferenceKey]codecatcher.Reference{
							{FilePath: "test.go", LineNumber: 10, CallerName: "main"}: {
								Context: "function call",
								Caller: &codecatcher.Entity{
									Name: "main",
								},
							},
						},
					},
				},
			},
		},
	}

	sources := map[string]*codecatcher.Source{
		"test.go": source,
	}

	writer := NewFlowchartWriter(sources)

	// 测试所有生成的图
	graphs := []string{
		writer.GeneratePackageDependencies(),
		writer.GenerateEntityReferences(),
		writer.GenerateInterfaceImplementations(),
	}

	for _, graph := range graphs {
		// 验证基本语法
		assert.True(t, strings.HasPrefix(graph, "graph TD\n"))

		// 验证节点定义格式
		lines := strings.Split(graph, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || line == "graph TD" {
				continue
			}

			// 检查节点定义（例如：node[label]）
			if strings.Contains(line, "[") && !strings.Contains(line, "-->") {
				assert.Regexp(t, `^\s*[\w_]+\[[^\]]+\]\s*$`, line, "节点定义格式错误: %s", line)
			}

			// 检查连接定义
			if strings.Contains(line, "-->") {
				if strings.Contains(line, "|") {
					// 带标签的连接（例如：A -->|label| B）
					assert.Regexp(t, `^\s*[\w_]+\s*(-->|-.->)\s*\|[^|]+\|\s*[\w_]+\s*$`, line, "带标签的连接格式错误: %s", line)
				} else {
					// 普通连接（例如：A --> B）
					assert.Regexp(t, `^\s*[\w_]+\s*(-->|-.->)\s*[\w_]+\s*$`, line, "普通连接格式错误: %s", line)
				}
			}
		}
	}
}
