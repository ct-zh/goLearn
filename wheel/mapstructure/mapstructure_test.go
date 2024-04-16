package mapstructure

import "testing"

func TestParser(t *testing.T) {
	type TestStruct struct {
		Name  string
		Age   int
		Email string
	}

	var originData = map[string]interface{}{
		"name":  "张三",
		"age":   30,
		"email": "zhangsan@example.com",
	}

	data := &TestStruct{}
	err := Parser(data, originData)
	if err != nil {
		t.Fatal(err)
	}
	if data.Name != originData["name"] {
		t.Fatal("Name mismatch")
	}
	if data.Age != originData["age"] {
		t.Fatal("Age mismatch")
	}
	if data.Email != originData["email"] {
		t.Fatal("Email mismatch")
	}
	t.Log("success!")
}
