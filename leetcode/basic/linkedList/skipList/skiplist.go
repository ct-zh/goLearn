package skipList

import (
	"math/rand"
	"time"
)

// 跳跃表
// 双向链表的优化版，将前进指针修改为跳跃数组,通过维护若干条捷径(下面skip.Node.l参数)加快链表的查找效率;
// 因为会经常增删结点，所以捷径l的生成长度全靠随机(抛硬币，正就增加一层，负就不增加；幂次定律,层数越高，概率越小)
//
// 注意跳跃表有一个头结点，用于保存level的头部，score、data以及bw这几个值都不使用；
// 头结点不参与len的计算，只用来查找;
//
// 查找：从头结点level的最高层开始查找，往下走，类似二分搜索；
// 新增：找到插入位置后，跟链表一样对前后结点进行操作，然后再抛硬币生成对应长度的level；
// 删除：和链表一样
//
// 跳跃表查找的时间复杂度是平均O(logN)，最差O(N)
// 对比平衡二叉树的优点是：维持结构平衡的成本比较低，完全依靠随机(不需要像红黑树那样旋转)
// 缺点：空间复杂度接近O(2N),都是额外存储捷径l的空间；
//
// 跳跃表的优化：
// 空间上的优化：将底层链表保存在磁盘上，选择若干高层L写入内存中用数组排列，数组结点是指向磁盘偏移地址的指针；
// 性能优化：
// 1. 全部加载在内存里，或者用LRU算法把经常访问的结点加载在内存里；
// 2. 层级不再使用随机算法，而是具体业务具体分析，比如score代表时间时，可以每隔一段时间增加一个层级，这样基本接近于二分查找;
// 3. 如果保存在磁盘上，可以考虑一个结点划分一个页，类似于b+tree的叶子结点；
//

type skipList struct {
	header   *skipNode // 头结点指针
	tail     *skipNode // 尾结点指针
	maxLevel int       // 最大层级
	len      int       // 结点数量(不包括头结点)
}

type skipNode struct {
	bw    *skipNode   // 后退指针
	score float64     // 结点分数
	data  interface{} // 结点数据
	l     []level     // 层级
}

type level struct {
	span    int       // 跨度
	forward *skipNode // 前进指针
}

// 抛硬币
func coin() bool {
	rand.Seed(time.Now().Unix())
	return rand.Intn(2) == 1
}

// create new skip list
func New() *skipList {
	return &skipList{
		header: &skipNode{l: make([]level, 1)}, // 生成空的头结点
	}
}

// 释放跳跃表空间
func (s *skipList) Free() {

}

func (s *skipList) Add(score float64, data interface{}) error {
	panic("")
}

func (s *skipList) Search(score float64) ([]interface{}, error) {
	panic("")
}

// 获取指定对象所在排位
func (s *skipList) GetRank(score float64, data interface{}) int {
	panic("")
}

func (s *skipList) Remove(score float64, data interface{}) error {
	panic("")
}
