package types

// Reference 引用信息
type Reference struct {
	FilePath   string  // 引用所在的文件
	LineNumber int     // 引用所在的行号
	Context    string  // 引用上下文（如方法名、函数名等）
	Caller     *Entity // 调用方实体
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
