package reflect

import "testing"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonMapDecode(t *testing.T) {
	mapA := map[string]string{
		"key1": `{"name": "John", "age": 30}`,
		"key2": `{"city": "New York", "country": "USA"}`,
	}

	mapB := make(map[string]UserInfo)

	if err := JsonMapDecode(mapA, &mapB); err != nil {
		t.Fatal(err)
	}

	t.Logf("mapB = %+v", mapB)
}
