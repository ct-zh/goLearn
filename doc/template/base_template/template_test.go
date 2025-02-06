package template

import (
	"strings"
	"testing"
)

func TestExecuteTemplate(t *testing.T) {
	tests := []struct {
		name     string
		tmpl     string
		data     interface{}
		expected string
		wantErr  bool
	}{
		{
			name: "基础用例-结构体",
			tmpl: "Hello {{.Name}}! You are {{.Age}} years old.",
			data: struct {
				Name string
				Age  int
			}{
				Name: "张三",
				Age:  25,
			},
			expected: "Hello 张三! You are 25 years old.",
			wantErr:  false,
		},
		{
			name: "Map数据",
			tmpl: "{{.message}} from {{.author}}",
			data: map[string]string{
				"message": "你好",
				"author":  "李四",
			},
			expected: "你好 from 李四",
			wantErr:  false,
		},
		{
			name: "条件语句",
			tmpl: "{{if .IsAdmin}}管理员{{else}}普通用户{{end}}",
			data: struct {
				IsAdmin bool
			}{
				IsAdmin: true,
			},
			expected: "管理员",
			wantErr:  false,
		},
		{
			name: "循环语句",
			tmpl: "{{range .Items}}{{.}},{{end}}",
			data: struct {
				Items []string
			}{
				Items: []string{"苹果", "香蕉", "橙子"},
			},
			expected: "苹果,香蕉,橙子,",
			wantErr:  false,
		},
		{
			name:     "错误的模板语法",
			tmpl:     "Hello {{.Name}", // 缺少结束括号
			data:     struct{}{},
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ExecuteTemplate(tt.tmpl, tt.data)

			// 检查错误
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecuteTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 如果期望错误，就不需要检查结果
			if tt.wantErr {
				return
			}

			// 检查结果
			if strings.TrimSpace(result) != strings.TrimSpace(tt.expected) {
				t.Errorf("ExecuteTemplate() = %v, want %v", result, tt.expected)
			}
		})
	}
}
