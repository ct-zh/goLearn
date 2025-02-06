package codecatcher

import (
	"github.com/ct-zh/goLearn/create_my_project/code_catcher/reader"
	"github.com/ct-zh/goLearn/create_my_project/code_catcher/types"
)

// FileSource 分析指定路径下的源代码
func FileSource(rootPath string) (*types.Source, error) {
	r := reader.NewReader(rootPath)
	if err := r.ReadProject(); err != nil {
		return nil, err
	}
	sources := r.GetSources()

	// 如果只有一个文件，直接返回
	if len(sources) == 1 {
		for _, source := range sources {
			return source, nil
		}
	}

	// 如果有多个文件，合并结果
	result := &types.Source{
		FilePath: rootPath,
	}
	for _, source := range sources {
		result.Interfaces = append(result.Interfaces, source.Interfaces...)
		result.Structs = append(result.Structs, source.Structs...)
		result.Functions = append(result.Functions, source.Functions...)
	}
	return result, nil
}
