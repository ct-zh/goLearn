package tree

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRbNode_getBrother(t *testing.T) {
	// 新建一棵二叉树
	tree := NewRBTree()
	tree.Insert(10, nil)
	tree.Insert(1, nil)
	tree.Insert(15, nil)
	tree.Insert(12, nil)
	tree.Insert(18, nil)

	//tree.Print()
	//    10
	//  1	15
	//	   12 18

	// 获取兄弟结点
	if tree.root.getBrother() != nil {
		t.Fatal("error: root应该没有兄弟结点")
	}
	if tree.root.left.getBrother().key != 15 {
		t.Fatal("error: 结点1的兄弟结点应该是15")
	}
	if tree.root.right.left.getBrother().key != 18 {
		t.Fatal("error: 结点12的兄弟结点应该是18")
	}
}

func TestRbNode_exchange(t *testing.T) {
	Convey("exchange方法测试 ", t, func() {
		Convey("交换叶子结点", func() {
			// 交换4和66的位置
			//      55
			//   10    77
			//  4    66
			tree := NewRBTree()
			tree.Insert(55, nil)
			tree.Insert(10, nil)
			tree.Insert(77, nil)
			tree.Insert(4, nil)
			tree.Insert(66, nil)
			//tree.Print()

			tree.exchange(tree.root.right.left, tree.root.left.left)
			// 结果应该是
			//      55
			//   10    77
			//  66    4
			//tree.Print()

			So(tree.root.left.left, ShouldNotBeNil)
			So(tree.root.left.left.key, ShouldEqual, 66)
			So(tree.root.right.left, ShouldNotBeNil)
			So(tree.root.right.left.key, ShouldEqual, 4)
		})

		Convey("交换根结点", func() {
			// 交换根结点的情况
			// 交换55和66的位置
			//      55
			//   10    77
			//  4    66
			tree2 := NewRBTree()
			tree2.Insert(55, nil)
			tree2.Insert(10, nil)
			tree2.Insert(77, nil)
			tree2.Insert(4, nil)
			tree2.Insert(66, nil)
			//tree2.Print()

			tree2.exchange(tree2.root, tree2.root.right.left)
			//tree2.Print()

			So(tree2.root, ShouldNotBeNil)
			So(tree2.root.key, ShouldEqual, 66)
			So(tree2.root.right.left, ShouldNotBeNil)
			So(tree2.root.right.left.key, ShouldEqual, 55)
		})
	})
}

func TestRbNode_leftRotation(t *testing.T) {
	// 新建一棵二叉树
	tree := NewRBTree()
	tree.Insert(10, nil)
	tree.Insert(1, nil)
	tree.Insert(15, nil)
	tree.Insert(12, nil)
	tree.Insert(18, nil)

	//tree.Print()
	//    10
	//  1	15
	//	   12 18

	// 左旋之后应该是
	// 		 15
	// 	  10    18
	//	 1 12

	tree.leftRotation(tree.root)
	tree.Print()

	tree2 := NewRBTree()
	tree2.Insert(50, nil)
	tree2.Insert(70, nil)
	tree2.Insert(30, nil)
	tree2.Insert(10, nil)
	tree2.Insert(40, nil)
	tree2.Insert(35, nil)
	tree2.Insert(45, nil)

	// 			50
	//       30    70
	//    10   40
	//		  35 45

	// 对30左旋,变成:
	// 			50
	//       40    70
	//    30   45
	//	10 35

	tree2.leftRotation(tree2.root.left)
	tree2.Print()
}

func TestRbNode_rightRotation(t *testing.T) {
	// 新建一棵二叉树
	tree := NewRBTree()
	tree.Insert(15, nil)
	tree.Insert(10, nil)
	tree.Insert(20, nil)
	tree.Insert(8, nil)
	tree.Insert(13, nil)

	//tree.Print()
	//    15
	//  10	20
	// 8 13
	// 右旋之后应该是
	// 	   10
	// 	 8   15
	//	 	13 20

	tree.rightRotation(tree.root)
	tree.Print()

	fmt.Println("======================")

	tree2 := NewRBTree()
	tree2.Insert(50, nil)
	tree2.Insert(70, nil)
	tree2.Insert(30, nil)
	tree2.Insert(55, nil)
	tree2.Insert(80, nil)
	tree2.Insert(52, nil)
	tree2.Insert(60, nil)

	// 			50
	//     30       70
	//    		 55   80
	//         52 60

	// 对70右旋,变成:
	// 			50
	//       30    55
	//           52   70
	//               60 80

	tree2.rightRotation(tree2.root.right)
	tree2.Print()
}

func TestCreateByMap(t *testing.T) {

}

func TestRedBlackTree_Size(t *testing.T) {

}

func TestRedBlackTree_IsEmpty(t *testing.T) {

}

func TestRedBlackTree_Contain(t *testing.T) {

}

func TestRedBlackTree_Search(t *testing.T) {
}

func TestRedBlackTree_Insert(t *testing.T) {
	tree := &redBlackTree{}
	tree.Insert(85, nil)
	tree.Insert(80, nil)
	tree.Insert(89, nil)
	tree.Insert(82, nil)
	tree.Insert(75, nil)
	tree.Insert(68, nil)
	tree.Insert(77, nil)
	tree.Insert(66, nil)
	tree.Print()
}

func TestRedBlackTree_Print(t *testing.T) {
	m := make(map[int]interface{})
	for i := 1; i < 10; i++ {
		m[i] = "aaa" + strconv.Itoa(i)
	}
	tree := CreateByMap(m)
	tree.Print()
}

func TestRedBlackTree_Remove(t *testing.T) {
	tree := &redBlackTree{}
	tree.Insert(10, nil)
	tree.Insert(1, nil)
	tree.Insert(15, nil)
	tree.Insert(12, nil)
	tree.Insert(18, nil)

	//    10
	//  1	15
	//	   12 18
	tree.Print()
	//
	fmt.Println("============")
	//// 删除结点15
	tree.Remove(15)
	tree.Print()

}
