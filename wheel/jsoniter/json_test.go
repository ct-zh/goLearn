package jsoniter

import (
	"testing"

	orgJson "encoding/json"
)

func TestNewEncoder(t *testing.T) {
	str := `[{"content":"内容","source":"yinp","pass_count":1,"reject_count":2}]`
	bodyStr := []byte(str)

	model := []struct {
		Content     string `json:"content,omitempty"`
		Source      string `json:"source,omitempty"`
		PassCount   int    `json:"pass_count,omitempty"`
		RejectCount int    `json:"reject_count,omitempty"`
	}{}

	err := NewEncoder().Decode(bodyStr, &model)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("req: %+v", model)
}

type data struct {
	Content     string `json:"content,omitempty"`
	Source      string `json:"source,omitempty"`
	PassCount   int    `json:"pass_count,omitempty"`
	RejectCount int    `json:"reject_count,omitempty"`
}

// goos: darwin
//goarch: amd64
//pkg: github.com/ct-zh/goLearn/wheel/jsoniter
//cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
//BenchmarkJsoniter
//BenchmarkJsoniter-8   	 1402230	       756.1 ns/op
//BenchmarkJson
//BenchmarkJson-8       	 3959030	       310.0 ns/op
func BenchmarkJsoniter(b *testing.B) {
	u := &data{
		Content:     "内容",
		Source:      "yinp",
		PassCount:   1,
		RejectCount: 2,
	}
	for i := 0; i < b.N; i++ {
		NewEncoder().Encode(u)
	}
}

func BenchmarkJson(b *testing.B) {
	u := &data{
		Content:     "内容",
		Source:      "yinp",
		PassCount:   1,
		RejectCount: 2,
	}
	for i := 0; i < b.N; i++ {
		orgJson.Marshal(u)
	}
}
