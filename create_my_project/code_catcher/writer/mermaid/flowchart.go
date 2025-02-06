package mermaid

import (
	"fmt"
	"path/filepath"
	"strings"

	codecatcher "github.com/ct-zh/goLearn/create_my_project/code_catcher"
)

// FlowchartWriter Mermaid流程图生成器
type FlowchartWriter struct {
	sources map[string]*codecatcher.Source
}

// NewFlowchartWriter 创建新的流程图生成器
func NewFlowchartWriter(sources map[string]*codecatcher.Source) *FlowchartWriter {
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
