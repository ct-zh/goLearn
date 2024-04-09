package easyjson

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUser(t *testing.T) {
	u := &User{
		Name:   "郑元畅",
		Age:    19,
		Gender: Man,
	}
	t.Logf("%+v", u)
	json, err := u.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", string(json))

	u2 := &User{}
	err = u2.UnmarshalJSON(json)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", u2)

	if !reflect.DeepEqual(u, u2) {
		t.Fatalf("两个结构体不相等， u： %+v u2: %+v ", u, u2)
	}
}

// BenchmarkEasyJson-8   	 8223699	       141 ns/op	     128 B/op	       1 allocs/op
// BenchmarkGoJson-8     	 1820175	       650 ns/op	     176 B/op	       2 allocs/op

func BenchmarkEasyJson(b *testing.B) {
	u := &User{
		Name:   "郑元畅",
		Age:    19,
		Gender: Man,
	}
	for i := 0; i < b.N; i++ {
		u.MarshalJSON()
	}
}

func BenchmarkGoJson(b *testing.B) {
	u := &User{
		Name:   "郑元畅",
		Age:    19,
		Gender: Man,
	}
	for i := 0; i < b.N; i++ {
		json.Marshal(u)
	}
}
