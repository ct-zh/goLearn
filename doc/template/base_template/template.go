package template

import (
	"bytes"
	"text/template"
)

// ExecuteTemplate 执行模板并返回结果
// tmpl: 模板字符串
// data: 要填充到模板中的数据
func ExecuteTemplate(tmpl string, data interface{}) (string, error) {
	// 创建模板实例
	t := template.Must(template.New("template").Parse(tmpl))

	// 创建一个buffer来存储结果
	var buf bytes.Buffer

	// 执行模板
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
