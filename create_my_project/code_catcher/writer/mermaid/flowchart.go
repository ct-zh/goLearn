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
func (w *FlowchartWriter) GenerateLayeredMarkdown() string {
	var sb strings.Builder
	sb.WriteString("# 代码分析报告\n\n")

	// 收集所有实体
	entities := w.collectAllEntities()

	// 按类型分组
	structMethods := make(map[string][]LayeredEntityReferences)
	var structs, functions []LayeredEntityReferences

	for _, entity := range entities {
		if entity.Type == "method" {
			if parts := strings.Split(entity.EntityName, "."); len(parts) == 2 {
				structName := parts[0]
				structMethods[structName] = append(structMethods[structName], entity)
			}
		} else if entity.Type == "struct" {
			structs = append(structs, entity)
		} else if entity.Type == "function" {
			functions = append(functions, entity)
		}
	}

	// 生成结构体及其方法的文档
	for _, st := range structs {
		sb.WriteString(fmt.Sprintf("## %s (%s)\n\n", st.EntityName, st.Type))

		// 只有在有外部引用时才生成调用关系图
		if w.hasExternalReferences(st) {
			context := w.getEntityContext(st, 5)
			sb.WriteString("### 调用关系\n\n")
			sb.WriteString("```mermaid\n")
			sb.WriteString(w.generateContextGraph(context))
			sb.WriteString("```\n\n")
		}

		w.writeEntityReferences(&sb, st)

		// 生成该结构体的方法文档
		if methods := structMethods[st.EntityName]; len(methods) > 0 {
			for _, method := range methods {
				sb.WriteString(fmt.Sprintf("### %s\n\n", method.EntityName))

				// 只有在有外部引用时才生成调用关系图
				if w.hasExternalReferences(method) {
					context := w.getEntityContext(method, 5)
					sb.WriteString("#### 调用关系\n\n")
					sb.WriteString("```mermaid\n")
					sb.WriteString(w.generateContextGraph(context))
					sb.WriteString("```\n\n")
				}

				w.writeEntityReferences(&sb, method)
			}
		}

		sb.WriteString("---\n\n")
	}

	// 生成独立函数的文档
	if len(functions) > 0 {
		sb.WriteString("## 独立函数\n\n")
		for _, fn := range functions {
			sb.WriteString(fmt.Sprintf("### %s (%s)\n\n", fn.EntityName, fn.Type))

			// 只有在有外部引用时才生成调用关系图
			if w.hasExternalReferences(fn) {
				context := w.getEntityContext(fn, 5)
				sb.WriteString("#### 调用关系\n\n")
				sb.WriteString("```mermaid\n")
				sb.WriteString(w.generateContextGraph(context))
				sb.WriteString("```\n\n")
			}

			w.writeEntityReferences(&sb, fn)
			sb.WriteString("---\n\n")
		}
	}

	return sb.String()
}

// writeEntityReferences 写入实体的引用信息
func (w *FlowchartWriter) writeEntityReferences(sb *strings.Builder, entity LayeredEntityReferences) {
	if len(entity.Callers) > 0 {
		sb.WriteString("#### 调用者\n\n")
		for _, caller := range entity.Callers {
			sb.WriteString(fmt.Sprintf("- %s\n", caller))
		}
		sb.WriteString("\n")
	}

	if len(entity.Callees) > 0 {
		sb.WriteString("#### 被调用者\n\n")
		for _, callee := range entity.Callees {
			sb.WriteString(fmt.Sprintf("- %s\n", callee))
		}
		sb.WriteString("\n")
	}
}

// getEntityContext 获取实体的上下文（上下n层关系）
func (w *FlowchartWriter) getEntityContext(entity LayeredEntityReferences, depth int) map[string]LayeredEntityReferences {
	context := make(map[string]LayeredEntityReferences)
	context[entity.EntityName] = entity

	// 向上遍历调用者
	w.traverseUp(entity.EntityName, entity.Callers, context, depth)
	// 向下遍历被调用者
	w.traverseDown(entity.EntityName, entity.Callees, context, depth)

	return context
}

// traverseUp 向上遍历调用关系
func (w *FlowchartWriter) traverseUp(currentName string, callers []string, context map[string]LayeredEntityReferences, remainingDepth int) {
	if remainingDepth <= 0 || len(callers) == 0 {
		return
	}

	for _, caller := range callers {
		if _, exists := context[caller]; exists {
			continue
		}

		// 获取调用者信息
		for _, entity := range w.collectAllEntities() {
			if entity.EntityName == caller {
				context[caller] = entity
				w.traverseUp(caller, entity.Callers, context, remainingDepth-1)
				break
			}
		}
	}
}

// traverseDown 向下遍历调用关系
func (w *FlowchartWriter) traverseDown(currentName string, callees []string, context map[string]LayeredEntityReferences, remainingDepth int) {
	if remainingDepth <= 0 || len(callees) == 0 {
		return
	}

	for _, callee := range callees {
		if _, exists := context[callee]; exists {
			continue
		}

		// 获取被调用者信息
		for _, entity := range w.collectAllEntities() {
			if entity.EntityName == callee {
				context[callee] = entity
				w.traverseDown(callee, entity.Callees, context, remainingDepth-1)
				break
			}
		}
	}
}

// generateContextGraph 生成上下文关系图
func (w *FlowchartWriter) generateContextGraph(context map[string]LayeredEntityReferences) string {
	var sb strings.Builder
	sb.WriteString("graph TD\n")

	// 添加节点
	processedNodes := make(map[string]bool)
	for name, entity := range context {
		if !processedNodes[name] {
			processedNodes[name] = true
			nodeStyle := w.getNodeStyle(entity.Type)
			sb.WriteString(fmt.Sprintf("    %s%s\n", sanitizeID(name), nodeStyle))
		}
	}

	// 添加连接
	processedEdges := make(map[string]bool)
	for name, entity := range context {
		for _, callee := range entity.Callees {
			if _, exists := context[callee]; exists {
				edgeKey := fmt.Sprintf("%s->%s", name, callee)
				if !processedEdges[edgeKey] {
					processedEdges[edgeKey] = true
					sb.WriteString(fmt.Sprintf("    %s --> %s\n", sanitizeID(name), sanitizeID(callee)))
				}
			}
		}
	}

	return sb.String()
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
