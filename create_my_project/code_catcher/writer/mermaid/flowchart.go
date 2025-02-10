package mermaid

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ct-zh/goLearn/create_my_project/code_catcher/types"
)

// FlowchartWriter Mermaid流程图生成器
type FlowchartWriter struct {
	sources map[string]*types.Source
}

// LayeredEntityReferences 分层的实体引用关系
type LayeredEntityReferences struct {
	EntityName string   // 当前实体名称
	Type       string   // 实体类型
	Callers    []string // 调用者
	Callees    []string // 被调用者
}

// NewFlowchartWriter 创建新的流程图生成器
func NewFlowchartWriter(sources map[string]*types.Source) *FlowchartWriter {
	return &FlowchartWriter{
		sources: sources,
	}
}

// GeneratePackageDependencies 生成包依赖关系图
func (w *FlowchartWriter) GeneratePackageDependencies() string {
	var sb strings.Builder
	sb.WriteString("graph TD\n")

	// 记录已处理的包
	processedPkgs := make(map[string]bool)
	// 记录包之间的依赖关系
	dependencies := make(map[string]map[string]bool)

	// 遍历所有源文件
	for _, source := range w.sources {
		pkg := source.PackageName
		if !processedPkgs[pkg] {
			processedPkgs[pkg] = true
			sb.WriteString(fmt.Sprintf("    %s[%s]\n", sanitizeID(pkg), pkg))
		}

		// 处理导入的包
		for _, imp := range source.Imports {
			// 去除引号和版本信息
			imp = cleanImportPath(imp)
			if !processedPkgs[imp] {
				processedPkgs[imp] = true
				sb.WriteString(fmt.Sprintf("    %s[%s]\n", sanitizeID(imp), imp))
			}

			// 记录依赖关系
			if dependencies[pkg] == nil {
				dependencies[pkg] = make(map[string]bool)
			}
			dependencies[pkg][imp] = true
		}
	}

	// 添加依赖关系连线
	for pkg, deps := range dependencies {
		for dep := range deps {
			sb.WriteString(fmt.Sprintf("    %s --> %s\n", sanitizeID(pkg), sanitizeID(dep)))
		}
	}

	return sb.String()
}

// GenerateEntityReferences 生成实体引用关系图
func (w *FlowchartWriter) GenerateEntityReferences() string {
	var sb strings.Builder
	sb.WriteString("graph TD\n")

	// 记录已处理的实体
	processedEntities := make(map[string]bool)
	// 记录实体之间的引用关系
	references := make(map[string]map[string]string)

	// 遍历所有源文件
	for _, source := range w.sources {
		// 处理结构体及其方法
		for _, st := range source.Structs {
			structID := fmt.Sprintf("struct_%s", sanitizeID(st.Name))
			if !processedEntities[structID] {
				processedEntities[structID] = true
				sb.WriteString(fmt.Sprintf("    %s[%s]\n", structID, st.Name))
			}

			// 处理方法
			for _, method := range st.Methods {
				methodID := fmt.Sprintf("method_%s_%s", sanitizeID(st.Name), sanitizeID(method.Name))
				if !processedEntities[methodID] {
					processedEntities[methodID] = true
					sb.WriteString(fmt.Sprintf("    %s[%s.%s]\n", methodID, st.Name, method.Name))
				}
				// 添加结构体到方法的连线
				sb.WriteString(fmt.Sprintf("    %s --> %s\n", structID, methodID))

				// 处理方法的引用
				for _, ref := range method.References {
					callerID := fmt.Sprintf("caller_%s", sanitizeID(ref.Caller.Name))
					if !processedEntities[callerID] {
						processedEntities[callerID] = true
						sb.WriteString(fmt.Sprintf("    %s[%s]\n", callerID, ref.Caller.Name))
					}
					if references[callerID] == nil {
						references[callerID] = make(map[string]string)
					}
					references[callerID][methodID] = ref.Context

					// 反向添加调用关系
					sb.WriteString(fmt.Sprintf("    %s -->|%s| %s\n", callerID, ref.Context, methodID))
				}
			}
		}

		// 处理函数
		for _, fn := range source.Functions {
			fnID := fmt.Sprintf("func_%s", sanitizeID(fn.Name))
			if !processedEntities[fnID] {
				processedEntities[fnID] = true
				sb.WriteString(fmt.Sprintf("    %s[%s]\n", fnID, fn.Name))
			}

			// 处理函数的引用
			for _, ref := range fn.References {
				callerID := fmt.Sprintf("caller_%s", sanitizeID(ref.Caller.Name))
				if !processedEntities[callerID] {
					processedEntities[callerID] = true
					sb.WriteString(fmt.Sprintf("    %s[%s]\n", callerID, ref.Caller.Name))
				}
				if references[callerID] == nil {
					references[callerID] = make(map[string]string)
				}
				references[callerID][fnID] = ref.Context
			}
		}
	}

	// 添加引用关系连线
	for caller, refs := range references {
		for target, context := range refs {
			sb.WriteString(fmt.Sprintf("    %s -->|%s| %s\n", caller, context, target))
		}
	}

	return sb.String()
}

// GenerateInterfaceImplementations 生成接口实现关系图
func (w *FlowchartWriter) GenerateInterfaceImplementations() string {
	var sb strings.Builder
	sb.WriteString("graph TD\n")

	// 记录已处理的实体
	processedEntities := make(map[string]bool)

	// 遍历所有源文件
	for _, source := range w.sources {
		// 处理接口
		for _, iface := range source.Interfaces {
			ifaceID := fmt.Sprintf("interface_%s", sanitizeID(iface.Name))
			if !processedEntities[ifaceID] {
				processedEntities[ifaceID] = true
				sb.WriteString(fmt.Sprintf("    %s[(%s)]\n", ifaceID, iface.Name))
			}

			// 查找实现此接口的结构体
			for _, src := range w.sources {
				for _, st := range src.Structs {
					if containsInterface(st.Implements, iface.Name) {
						structID := fmt.Sprintf("struct_%s", sanitizeID(st.Name))
						if !processedEntities[structID] {
							processedEntities[structID] = true
							sb.WriteString(fmt.Sprintf("    %s[%s]\n", structID, st.Name))
						}
						sb.WriteString(fmt.Sprintf("    %s -.->|implements| %s\n", structID, ifaceID))
					}
				}
			}
		}
	}

	return sb.String()
}

// sanitizeID 清理ID中的特殊字符
func sanitizeID(id string) string {
	return strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
			return r
		}
		return '_'
	}, id)
}

// cleanImportPath 清理导入路径
func cleanImportPath(path string) string {
	// 去除版本信息（如果有）
	if idx := strings.Index(path, "@"); idx != -1 {
		path = path[:idx]
	}
	// 获取最后一个包名
	return filepath.Base(path)
}

// containsInterface 检查接口是否在实现列表中
func containsInterface(implements []string, ifaceName string) bool {
	for _, impl := range implements {
		if impl == ifaceName {
			return true
		}
	}
	return false
}

// collectAllEntities 收集所有实体信息
func (w *FlowchartWriter) collectAllEntities() []LayeredEntityReferences {
	entities := make(map[string]*LayeredEntityReferences)

	// 遍历所有源文件
	for _, source := range w.sources {
		// 处理结构体
		for _, st := range source.Structs {
			entityName := st.Name
			if _, exists := entities[entityName]; !exists {
				entities[entityName] = &LayeredEntityReferences{
					EntityName: entityName,
					Type:       "struct",
				}
			}

			// 处理方法
			for _, method := range st.Methods {
				methodName := fmt.Sprintf("%s.%s", st.Name, method.Name)
				if _, exists := entities[methodName]; !exists {
					entities[methodName] = &LayeredEntityReferences{
						EntityName: methodName,
						Type:       "method",
					}
				}

				// 添加调用关系（排除自引用）
				for _, ref := range method.References {
					callerName := ref.Caller.Name
					if callerName != methodName { // 排除自引用
						entities[methodName].Callers = appendUnique(entities[methodName].Callers, callerName)
						if callerEntity, exists := entities[callerName]; exists {
							callerEntity.Callees = appendUnique(callerEntity.Callees, methodName)
						}
					}
				}
			}
		}

		// 处理函数
		for _, fn := range source.Functions {
			if _, exists := entities[fn.Name]; !exists {
				entities[fn.Name] = &LayeredEntityReferences{
					EntityName: fn.Name,
					Type:       "function",
				}
			}

			// 添加调用关系（排除自引用）
			for _, ref := range fn.References {
				callerName := ref.Caller.Name
				if callerName != fn.Name { // 排除自引用
					entities[fn.Name].Callers = appendUnique(entities[fn.Name].Callers, callerName)
					if callerEntity, exists := entities[callerName]; exists {
						callerEntity.Callees = appendUnique(callerEntity.Callees, fn.Name)
					}
				}
			}
		}
	}

	// 转换为切片
	result := make([]LayeredEntityReferences, 0, len(entities))
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result
}

// hasExternalReferences 检查实体是否有外部引用
func (w *FlowchartWriter) hasExternalReferences(entity LayeredEntityReferences) bool {
	return len(entity.Callers) > 0 || len(entity.Callees) > 0
}

// GenerateLayeredMarkdown 生成分层的markdown文档
func (w *FlowchartWriter) GenerateLayeredMarkdown(tree *types.CallTree) string {
	var sb strings.Builder
	sb.WriteString("# 代码分析报告\n\n")

	// 从根节点开始递归生成文档
	w.generateNodeMarkdown(&sb, tree.Root, "", make(map[string]bool), 0)

	return sb.String()
}

// generateNodeMarkdown 递归生成节点的markdown文档
func (w *FlowchartWriter) generateNodeMarkdown(sb *strings.Builder, node *types.CallNode, prefix string, rendered map[string]bool, depth int) {
	if node == nil || node.Entity == nil {
		return
	}

	// 递归处理子节点
	for _, child := range node.Children {
		if !rendered[child.Entity.Name] {
			rendered[child.Entity.Name] = true

			// 根据深度生成标题级别
			titleLevel := depth + 2 // 从h2开始
			if titleLevel > 6 {
				titleLevel = 6
			}
			title := strings.Repeat("#", titleLevel)

			sb.WriteString(fmt.Sprintf("%s %s (%s)\n\n", title, child.Entity.Name, child.Entity.Type))

			// 生成调用关系图
			if len(child.Children) > 0 {
				sb.WriteString(fmt.Sprintf("%s 调用关系\n\n", strings.Repeat("#", titleLevel+1)))
				sb.WriteString("```mermaid\n")
				sb.WriteString(w.generateNodeGraph(child))
				sb.WriteString("```\n\n")
			}

			// 生成实体详细信息
			switch child.Entity.Type {
			case types.EntityTypeFunction:
				if fn, ok := child.Entity.Data.(types.Function); ok {
					w.writeFunctionDetails(sb, &fn)
				}
			case types.EntityTypeMethod:
				if method, ok := child.Entity.Data.(types.Method); ok {
					w.writeMethodDetails(sb, &method)
				}
			case types.EntityTypeStruct:
				if st, ok := child.Entity.Data.(types.Struct); ok {
					w.writeStructDetails(sb, &st)
				}
			}

			sb.WriteString("\n---\n\n")

			// 递归处理子节点
			w.generateNodeMarkdown(sb, child, prefix+"  ", rendered, depth+1)
		}
	}
}

// generateNodeGraph 生成节点的调用关系图
func (w *FlowchartWriter) generateNodeGraph(node *types.CallNode) string {
	var sb strings.Builder
	sb.WriteString("graph TD\n")

	// 添加当前节点
	sb.WriteString(fmt.Sprintf("    %s%s\n",
		sanitizeID(node.Entity.Name),
		w.getNodeStyle(string(node.Entity.Type))))

	// 添加子节点和连接
	for _, child := range node.Children {
		sb.WriteString(fmt.Sprintf("    %s%s\n",
			sanitizeID(child.Entity.Name),
			w.getNodeStyle(string(child.Entity.Type))))
		sb.WriteString(fmt.Sprintf("    %s --> %s\n",
			sanitizeID(node.Entity.Name),
			sanitizeID(child.Entity.Name)))
	}

	return sb.String()
}

// writeFunctionDetails 写入函数详细信息
func (w *FlowchartWriter) writeFunctionDetails(sb *strings.Builder, fn *types.Function) {
	sb.WriteString("#### 函数签名\n\n")
	sb.WriteString("```go\n")
	sb.WriteString(fmt.Sprintf("func %s(", fn.Name))
	w.writeParams(sb, fn.Params)
	sb.WriteString(")")
	if len(fn.Results) > 0 {
		sb.WriteString(" (")
		w.writeParams(sb, fn.Results)
		sb.WriteString(")")
	}
	sb.WriteString("\n```\n\n")
}

// writeMethodDetails 写入方法详细信息
func (w *FlowchartWriter) writeMethodDetails(sb *strings.Builder, method *types.Method) {
	sb.WriteString("#### 方法签名\n\n")
	sb.WriteString("```go\n")
	sb.WriteString(fmt.Sprintf("func (receiver) %s(", method.Name))
	w.writeParams(sb, method.Params)
	sb.WriteString(")")
	if len(method.Results) > 0 {
		sb.WriteString(" (")
		w.writeParams(sb, method.Results)
		sb.WriteString(")")
	}
	sb.WriteString("\n```\n\n")
}

// writeStructDetails 写入结构体详细信息
func (w *FlowchartWriter) writeStructDetails(sb *strings.Builder, st *types.Struct) {
	sb.WriteString("#### 结构体定义\n\n")
	sb.WriteString("```go\n")
	sb.WriteString(fmt.Sprintf("type %s struct {\n", st.Name))
	for _, field := range st.Fields {
		sb.WriteString(fmt.Sprintf("    %s %s\n", field.Name, field.Type))
	}
	sb.WriteString("}\n```\n\n")
}

// writeParams 写入参数列表
func (w *FlowchartWriter) writeParams(sb *strings.Builder, params []types.Field) {
	for i, param := range params {
		if i > 0 {
			sb.WriteString(", ")
		}
		if param.Name != "" {
			sb.WriteString(fmt.Sprintf("%s %s", param.Name, param.Type))
		} else {
			sb.WriteString(param.Type)
		}
	}
}

// getNodeStyle 根据实体类型返回节点样式
func (w *FlowchartWriter) getNodeStyle(entityType string) string {
	switch entityType {
	case "struct":
		return "[" + entityType + "]"
	case "method":
		return "(" + entityType + ")"
	case "function":
		return "{" + entityType + "}"
	default:
		return "[" + entityType + "]"
	}
}

// appendUnique 添加唯一元素到切片
func appendUnique(slice []string, element string) []string {
	for _, e := range slice {
		if e == element {
			return slice
		}
	}
	return append(slice, element)
}
