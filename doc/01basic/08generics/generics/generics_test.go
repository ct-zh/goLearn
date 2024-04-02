package generics

import "testing"

func TestAdd(t *testing.T) {
	// 两个int64类型
	var a, b int64 = 1, 2
	c := Add(a, b)
	t.Logf("%d + %d = %d [%T]", a, b, c, c)

	// 两个float64类型
	var d, e float64 = 1.1, 2.2
	f := Add(d, e)
	t.Logf("%f + %f = %f [%T]", d, e, f, f)

	// 一个int64类型，一个float64类型 不行 会报错
	//g := Add(d, a)
	//t.Logf("%d + %f = %v [%T]", a, d, g, g)
}

func TestMyList_Append(t *testing.T) {
	mylist := &MyList[int]{}
	mylist.Append(1)
	t.Logf("list: %+v", mylist.data)

	myList2 := &MyList[string]{}
	myList2.Append("1")
	t.Logf("list: %+v", myList2.data)
}
