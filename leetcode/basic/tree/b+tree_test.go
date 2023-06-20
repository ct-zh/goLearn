package tree

import (
	"testing"
)

func TestTree(t *testing.T) {
	tree := &BPTree{}
	tree.Initialize()

	i := 1
	for i < 100 {
		_, result := tree.Insert(i, i*10)
		t.Log(i)
		if result == false {
			print("数据已存在")
		}
		i++
	}

	tree.Remove(7)
	tree.Remove(6)
	tree.Remove(5)
	resultData, succ := tree.FindData(5)
	if succ == true {
		t.Logf("%+v\n", resultData)
	}

	t.Logf("%+v", tree.root)

	//i2 := 0
	//for i2 < tree.root.Children[1].KeyNum {
	//	t.Log(tree.root.Children[1].leafNode.data[i])
	//	i++
	//}

}
