package graph

type Weight interface{}

// 有权图的边的内容
type Edge struct {
	a      int // 边的两个顶点a, b
	b      int
	weight Weight
}

func (e *Edge) V() int {
	return e.a
}

func (e *Edge) W() int {
	return e.b
}

func (e *Edge) Wt() Weight {
	return e.weight
}

// 匹配 Item interface
func (e Edge) GetValue() interface{} {
	return e.weight
}

// 获取边里非节点x的节点
func (e *Edge) Other(x int) int {
	if e.a == x {
		return e.b
	} else if e.b == x {
		return e.a
	} else {
		panic("x 必须是 两个节点中的一个")
	}
}
