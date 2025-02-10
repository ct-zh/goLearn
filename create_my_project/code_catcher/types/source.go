package types

import (
	"fmt"
	"sort"
	"strings"
)

// Reference 引用信息
type Reference struct {
	FilePath   string  // 引用所在的文件
	LineNumber int     // 引用所在的行号
	Context    string  // 引用上下文（如方法名、函数名等）
	Caller     *Entity // 调用方实体
}

// CallNode 调用树节点
type CallNode struct {
	Entity   *Entity              // 当前实体
	Children map[string]*CallNode // 子节点，key为实体名称
	Rendered bool                 // 是否已经被渲染
}

// String 返回节点的字符串表示
func (n *CallNode) String(level int) string {
	if n == nil || n.Entity == nil {
		return ""
	}

	// 创建缩进
	indent := strings.Repeat("  ", level)

	// 构建当前节点的字符串
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s%s (%s)\n", indent, n.Entity.Name, n.Entity.Type))

	// 递归处理子节点
	childNames := make([]string, 0, len(n.Children))
	for name := range n.Children {
		childNames = append(childNames, name)
	}
	// 按名称排序，使输出稳定
	sort.Strings(childNames)

	for _, name := range childNames {
		sb.WriteString(n.Children[name].String(level + 1))
	}

	return sb.String()
}

// CallTree 调用树
type CallTree struct {
	Root      *CallNode            // 根节点
	NodeIndex map[string]*CallNode // 用于快速查找节点，key为实体名称
}

// String 返回整个调用树的字符串表示
func (t *CallTree) String() string {
	if t == nil || t.Root == nil {
		return "空调用树"
	}

	var sb strings.Builder
	sb.WriteString("调用树结构：\n")

	// 获取所有入口点（根节点的直接子节点）
	entryPoints := make([]string, 0, len(t.Root.Children))
	for name := range t.Root.Children {
		entryPoints = append(entryPoints, name)
	}
	// 按名称排序，使输出稳定
	sort.Strings(entryPoints)

	// 打印每个入口点及其子树
	for _, name := range entryPoints {
		sb.WriteString(t.Root.Children[name].String(0))
	}

	// 打印节点索引统计
	sb.WriteString(fmt.Sprintf("\n总节点数：%d\n", len(t.NodeIndex)))
	return sb.String()
}

// ReferenceKey 用于标识唯一的引用
type ReferenceKey struct {
	FilePath   string // 引用所在的文件
	LineNumber int    // 引用所在的行号
	CallerName string // 调用方名称
}

// EntityType 实体类型
type EntityType int

const (
	EntityTypeInterface EntityType = iota
	EntityTypeStruct
	EntityTypeField
	EntityTypeMethod
	EntityTypeFunction
)

// String 返回实体类型的字符串表示
func (t EntityType) String() string {
	switch t {
	case EntityTypeInterface:
		return "interface"
	case EntityTypeStruct:
		return "struct"
	case EntityTypeField:
		return "field"
	case EntityTypeMethod:
		return "method"
	case EntityTypeFunction:
		return "function"
	default:
		return "unknown"
	}
}

// Entity 统一的实体结构
type Entity struct {
	Type     EntityType // 实体类型
	Name     string     // 实体名称
	FilePath string     // 实体所在文件
	Data     any        // 实体具体数据（Interface/Struct/Field/Method/Function）
}

// Source 源代码分析结果
type Source struct {
	FilePath    string      // 文件路径
	PackageName string      // 包名
	Imports     []string    // 导入的包
	Interfaces  []Interface // 接口定义
	Structs     []Struct    // 结构体定义
	Functions   []Function  // 函数定义
	CodeText    []byte      // 具体代码内容
}

// Interface 接口定义
type Interface struct {
	Name       string                     // 接口名称
	Methods    []Method                   // 接口方法
	References map[ReferenceKey]Reference // 被引用信息
}

// Struct 结构体定义
type Struct struct {
	Name       string                     // 结构体名称
	Fields     []Field                    // 结构体字段
	Methods    []Method                   // 结构体方法
	Implements []string                   // 实现的接口
	References map[ReferenceKey]Reference // 被引用信息
}

// Field 字段定义
type Field struct {
	Name       string                     // 字段名
	Type       string                     // 字段类型
	References map[ReferenceKey]Reference // 被引用信息
}

// Method 方法定义
type Method struct {
	Name       string                     // 方法名
	Params     []Field                    // 参数
	Results    []Field                    // 返回值
	References map[ReferenceKey]Reference // 被引用信息
}

// Function 函数定义
type Function struct {
	Name       string                     // 函数名
	Params     []Field                    // 参数
	Results    []Field                    // 返回值
	IsExported bool                       // 是否导出
	References map[ReferenceKey]Reference // 被引用信息
}
