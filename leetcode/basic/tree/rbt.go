package tree

import "fmt"

// 红黑树
// 红黑树，Red-Black Tree 「RBT」是一个自平衡(不是绝对的平衡)的二叉查找树(BST)，树上的每个节点都遵循下面的规则:
// ps. 红黑树叶子结点会挂两个黑色的null空结点
// 1. 每个节点不是红色就是黑色
// 2. 树的根始终是黑色的, NULL结点也是黑色的(黑土地孕育黑树根， )
// 3. 红色结点的子结点一定是黑色（也就是说按照父子子的结构，只可能出现： 红黑黑、黑红黑、黑黑红、黑黑黑四种情况）
// *4. 从任意节点（包括根）到其任何后代NULL节点的每条路径都具有相同数量的黑色节点
// *从4也可以推出: 如果一个结点存在黑子结点，那么该结点肯定有两个子结点

// 红黑树两个操作:
// 1. recolor (重新标记黑色或红色)
// 2. rotation (旋转，这是树达到平衡的关键)
// rotation分为左旋和右旋：
// 左旋：以某个结点作为支点(旋转结点)，其{右子结点}变为{旋转结点}的父结点，{右子结点}的{左子结点}变为{旋转结点}的{右子结点}，左子结点保持不变。
// 右旋：以某个结点作为支点(旋转结点)，其{左子结点}变为{旋转结点}的父结点，{左子结点}的{右子结点}变为{旋转结点}的{左子结点}，右子结点保持不变。
// ps. (二叉树左边比根结点小、右边比根结点大，所以左边称小儿子、右边称大儿子)
// 大儿子变爸爸、大儿子把自己的小儿子过继给他爸，这称为左旋；
// 小儿子变爸爸、小儿子把自己的大儿子过继给他爸，这称为右旋；

// 红黑树总是通过旋转和变色达到自平衡

type Color int

const (
	RED Color = iota + 1
	BLACK
)

func GetColor(color Color) string {
	if color == RED {
		return "红"
	} else {
		return "黑"
	}
}

// 红黑树结构
// 不要给root直接赋值，请走Insert方法添加结点
type redBlackTree struct {
	root  *rbNode // 根节点
	count int     // 节点个数
}

// 新建一个空的红黑树
func NewRBTree() *redBlackTree {
	return &redBlackTree{}
}

// 结点,私有结构体，外界不需要知道具体实现
type rbNode struct {
	key   int         // 数据在二叉树里的值，用于在二叉树里排序
	value interface{} // 二叉树保存的数据
	color Color       // 结点颜色，见const

	// 位置参数
	parent *rbNode
	left   *rbNode
	right  *rbNode
}

// 获取某个结点的兄弟结点
func (n *rbNode) getBrother() *rbNode {
	if n.parent == nil {
		return nil
	}
	if n.parent.left == n {
		return n.parent.right
	}
	return n.parent.left
}

// 获取叔叔结点
func (n *rbNode) getUncle() *rbNode {
	if n.parent == nil {
		return nil
	}
	return n.parent.getBrother()
}

// 获取祖父结点
func (n *rbNode) getGrandFa() *rbNode {
	if n.parent == nil || n.parent.parent == nil {
		return nil
	}
	return n.parent.parent
}

// 结点n1与n2交换位置，直接替换kv值与color
func (r *redBlackTree) exchange(n1, n2 *rbNode) {
	n1.key, n1.value, n1.color, n2.key, n2.value, n2.color = n2.key, n2.value, n2.color, n1.key, n1.value, n1.color
}

// 请不要调用
func (r *redBlackTree) exchangeBug(n1, n2 *rbNode) {
	// 互相换子结点
	n1.left, n1.right, n2.left, n2.right = n2.left, n2.right, n1.left, n1.right
	// 互换子结点的父结点
	if n1.left != nil {
		n1.left.parent = n2
	}
	if n1.right != nil {
		n1.right.parent = n2
	}
	if n2.left != nil {
		n2.left.parent = n1
	}
	if n2.right != nil {
		n2.right.parent = n1
	}

	// 如果存在根结点
	if n1.parent == nil {
		r.root, n2.parent, n1.parent = n2, nil, n2.parent

		// 替换父节点的连接
		if n2.parent.left == n2 {
			n2.parent.left = n1
		} else if n2.parent.right == n2 {
			n2.parent.right = n1
		}
	} else if n2.parent == nil {
		r.root, n2.parent, n1.parent = n1, n1.parent, nil

		// 替换父节点的连接
		if n1.parent.left == n1 {
			n1.parent.left = n2
		} else if n1.parent.right == n1 {
			n1.parent.right = n2
		}
	}
	fmt.Printf("n1: %+v n2: %+v", n1, n2)
}

// 以n为根左旋
func (r *redBlackTree) leftRotation(n *rbNode) {
	if n.right == nil { // 左旋必须有右子结点
		return
	}

	// 1. 将根结点的右子结点设置为根结点；

	// 先解决父结点的问题
	if n.parent == nil { // n结点是root结点
		r.root = n.right
	} else if n.parent.left == n {
		n.parent.left = n.right
	} else {
		n.parent.right = n.right
	}
	// 2. 右子结点的左子结点设置为根结点的右子结点；
	RC := n.right.left
	n.parent, n.right, n.right.left = n.right, RC, n
}

// 以n为根右旋
func (r *redBlackTree) rightRotation(n *rbNode) {
	if n.left == nil { // 右旋必须要有左子结点
		return
	}

	// 1. 将根结点的左子结点设置为根结点；

	// 先解决父结点的问题
	if n.parent == nil { // n结点是root结点
		r.root = n.left
	} else if n.parent.left == n {
		n.parent.left = n.left
	} else {
		n.parent.right = n.left
	}
	// 2. 左子结点的右子结点设置为根结点的左子结点；
	RC := n.left.right
	n.parent, n.left, n.left.right = n.left, RC, n
}

// 通过map创建一个红黑树
// map key为排序的key； value为存储的值；
func CreateByMap(m map[int]interface{}) *redBlackTree {
	tree := &redBlackTree{}
	for key, value := range m {
		tree.Insert(key, value)
	}
	return tree
}

// 打印二叉树
// 打印思路：先遍历二叉树，拿到每个结点的值和所在的层级，然后根据所在层级循环打印；
func (r *redBlackTree) Print() {
	total := [][]*rbNode{{}}
	r.print(r.root, 0, &total) // root为0级

	for i := 0; i < len(total); i++ {
		for j := 0; j < len(total[i]); j++ {
			if total[i][j] == nil {
				fmt.Printf(" <nil> ")
			} else {
				fmt.Printf(" <key:%d color:%s> ", total[i][j].key, GetColor(total[i][j].color))
				//fmt.Printf("<%+v>", total[i][j])
			}
		}
		fmt.Println()
	}
}

func (r *redBlackTree) print(n *rbNode, level int, total *[][]*rbNode) {
	if len(*total) <= level {
		*total = append(*total, []*rbNode{})
	}

	if n == nil {
		(*total)[level] = append((*total)[level], nil)
		return
	}

	(*total)[level] = append((*total)[level], n)

	// 从左子树开始
	r.print(n.left, level+1, total)
	r.print(n.right, level+1, total)
}

// 获取红黑树大小
func (r *redBlackTree) Size() int {
	return r.count
}

// 判断红黑树是否为空
func (r *redBlackTree) IsEmpty() bool {
	return r.count == 0
}

// 获取最小值
func (r *redBlackTree) Min() (int, interface{}) {
	minNode := r.min(r.root)
	return minNode.key, minNode.value
}

func (r *redBlackTree) min(n *rbNode) *rbNode {
	if n.left == nil {
		return n
	}
	return r.min(n.left)
}

// 获取最大值
func (r *redBlackTree) Max() (int, interface{}) {
	max := r.max(r.root)
	return max.key, max.value
}

func (r *redBlackTree) max(n *rbNode) *rbNode {
	if n.right == nil {
		return n
	}
	return r.max(n.right)
}

// 从二分搜索树中删除最小值所在节点 问题：如果根节点是最小值就无法删除
func (r *redBlackTree) RemoveMin() {
	if r.root != nil {
		r.removeMin(r.root)
	}
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (r *redBlackTree) removeMin(n *rbNode) *rbNode {
	// 左节点是空，说明该节点已经是最小值了
	// 将右子树提取出来接到上面一层的左子树
	if n.left == nil {
		r.count--
		return n.right
	}
	n.left = r.removeMin(n.left)

	// 维护parent字段
	n.left.parent = n

	return n
}

// 红黑树的查找, 和 二叉搜索树一样
// 1. 从根结点开始查找，把根结点设置为当前结点；
// 2. 若当前结点为空，返回null；
// 3. 若当前结点不为空，用当前结点的key跟查找key作比较；
// 4. 若当前结点key等于查找key，那么该key就是查找目标，返回当前结点；
// 5. 若当前结点key大于查找key，把当前结点的左子结点设置为当前结点，重复步骤2；
// 6. 若当前结点key小于查找key，把当前结点的右子结点设置为当前结点，重复步骤2；

// 确定key是否在二叉树内
func (r *redBlackTree) Contain(key int) bool {
	_, ok := r.traverse(r.root, key)
	return ok
}

// 根据key获取value
func (r *redBlackTree) Search(key int) (interface{}, bool) {
	if node, isOk := r.traverse(r.root, key); isOk {
		return node.value, true
	}
	return nil, false
}

// 遍历树寻找key
func (r *redBlackTree) traverse(n *rbNode, key int) (*rbNode, bool) {
	if n == nil {
		return nil, false
	}

	if key == n.key {
		return n, true
	} else if key < n.key {
		return r.traverse(n.left, key)
	} else {
		return r.traverse(n.right, key)
	}
}

// 红黑树插入
// 插入操作包括两部分工作：一查找插入的位置；二插入后自平衡。查找插入的父结点很简单，跟查找操作区别不大：
//
//1. 从根结点开始查找；
//2. 若根结点为空，那么插入结点作为根结点，结束。
//3. 若根结点不为空，那么把根结点作为当前结点；
//4. 若当前结点为null，返回当前结点的父结点，结束。
//5. 若当前结点key等于查找key，那么该key所在结点就是插入结点，更新结点的值，结束。
//6. 若当前结点key大于查找key，把当前结点的左子结点设置为当前结点，重复步骤4；
//7. 若当前结点key小于查找key，把当前结点的右子结点设置为当前结点，重复步骤4；

// 新插入的结点始终为红色
// 红色在父结点（如果存在）为黑色结点时，红黑树的黑色平衡没被破坏，
// 不需要做自平衡操作。但如果插入结点是黑色，那么插入位置所在的子树黑色结点总是多1，必须做自平衡。

// 一共有四种情景：
// 1. 红黑树为空树: 把插入结点作为根结点，并把结点设置为黑色;
// 2. 插入结点的Key已存在: 只需要更新value;
// 3. 插入结点的父结点为黑结点:由于插入的结点是红色的，并不会影响红黑树的平衡，直接插入即可，无需做自平衡;
// 4. 插入结点的父结点P为红结点:那么该父结点不可能为根结点，所以插入结点总是存在祖父结点PP。
//		4.1 叔叔结点S(父节点的兄弟结点)存在并且为红结点: 1. 将P和S设置为黑色; 2. 将PP设置为红色; 3. 把PP设置为当前插入结点;
//		4.2 叔叔结点不存在或为黑结点，并且插入结点的父亲结点是祖父结点的左子结点:
// 			4.2.1 插入结点是其父结点P的左子结点:	1. 将P设为黑色; 2. 将PP设为红色; 3. 对PP进行右旋;
//			4.2.2 插入结点是其父结点的右子结点: 1. 对P进行左旋; 2. 把P设置为插入结点，得到情景4.2.1; 3. 进行情景4.2.1的处理;
//		4.3 叔叔结点不存在或为黑结点，并且插入结点的父亲结点是祖父结点的右子结点:(与4.2正好相反)
//			4.3.1 插入结点是其父结点的右子结点: 1. 将P设为黑色; 2. 将PP设为红色; 3. 对PP进行左旋;
//			4.3.2 插入结点是其父结点的左子结点: 1. 对P进行右旋; 2. 把P设置为插入结点，得到情景4.2.1; 3. 进行情景4.2.1的处理;
//

// 插入key和value, 根据key的值调整红黑树
func (r *redBlackTree) Insert(key int, value interface{}) {
	// step 1. 和二叉树逻辑一样，先根据key插入到指定位置；
	target, isNew := r.insert(r.root, key, value)
	if isNew {
		r.root = target
	}

	// step 2. 调整树的结构使其平衡;
	r.balance(target)
}

// 寻找插入key的位置
// @param n：当前判断的结点；key：插入的key值；value：插入的数据
// @return target：新插入的目标结点； isNew：对于当前结点n来说，target结点是不是新增加的结点；
func (r *redBlackTree) insert(n *rbNode, key int, value interface{}) (target *rbNode, isNew bool) {
	if n == nil {
		r.count++
		return &rbNode{key: key, value: value, color: RED}, true // 新增结点，返回true
	}

	if key == n.key { // 查找到相等的值，则直接更新
		n.value = value
		return n, false // 不是新增结点，只是更新， 返回false
	} else if key < n.key { // 往左子树找
		target, isNew = r.insert(n.left, key, value)
		if isNew { // 该结点的左子结点是新增的结点
			n.left = target
			n.left.parent = n
		}
		return target, false
	} else { // key > n.key, 往右子树找
		target, isNew = r.insert(n.right, key, value)
		if isNew { // 结点n的右子结点是新增的结点
			n.right = target
			n.right.parent = n
		}
		return target, false
	}
}

// 使二叉树重新平衡
func (r *redBlackTree) balance(target *rbNode) {
	// 1. 红黑树为空树: 把插入结点作为根结点，并把结点设置为黑色;
	if target == r.root {
		r.root.color = BLACK // 保证root为黑色
		return
	}
	if target.parent.color == BLACK {
		// 3. 插入结点的父结点为黑结点:由于插入的结点是红色的，并不会影响红黑树的平衡，直接插入即可，无需做自平衡;
		return
	}
	// 4. 插入结点的父结点为红结点

	// 4.1 叔叔结点S(父节点的兄弟结点)存在并且为红结点
	if target.getUncle() != nil && target.getUncle().color == RED {
		// 1. 将P和S设置为黑色; 2. 将PP设置为红色; 3. 把PP设置为当前插入结点;
		target.parent.color = BLACK
		target.getUncle().color = BLACK
		target.getGrandFa().color = RED
		r.balance(target.getGrandFa())
		// 自底向上重新平衡二叉树
		// ps.这也是唯一一种会增加红黑树黑色结点层数的插入情景。
	} else if target.getGrandFa().left == target.parent {
		// 4.2 叔叔结点不存在或为黑结点，并且插入结点的父亲结点是祖父结点的左子结点:
		if target.parent.left == target {
			// 4.2.1 插入结点是其父结点P的左子结点:	1. 将P设为黑色; 2. 将PP设为红色; 3. 对PP进行右旋;
			target.parent.color = BLACK
			target.getGrandFa().color = RED
			r.rightRotation(target.getGrandFa())
		} else {
			// 4.2.2 插入结点是其父结点的右子结点: 1. 对P进行左旋; 2. 把P设置为插入结点，得到情景4.2.1; 3. 进行情景4.2.1的处理;
			r.leftRotation(target.parent)

			target.getGrandFa().color = BLACK
			if target.getGrandFa().parent != nil {
				target.getGrandFa().parent.color = RED
				r.rightRotation(target.getGrandFa().parent)
			}
		}
	} else if target.getGrandFa().right == target.parent {
		// 4.3 叔叔结点不存在或为黑结点，并且插入结点的父亲结点是祖父结点的右子结点:
		if target.parent.right == target {
			// 4.3.1 插入结点是其父结点的右子结点: 1. 将P设为黑色; 2. 将PP设为红色; 3. 对PP进行左旋;
			target.parent.color = BLACK
			target.getGrandFa().color = RED
			r.leftRotation(target.getGrandFa())
		} else {
			// 4.3.2 插入结点是其父结点的左子结点: 1. 对P进行右旋; 2. 把P设置为插入结点，得到情景4.2.1; 3. 进行情景4.2.1的处理;
			r.rightRotation(target.parent)
			target = target.parent

			target.getGrandFa().color = BLACK
			if target.getGrandFa().parent != nil {
				target.getGrandFa().parent.color = RED
				r.leftRotation(target.getGrandFa().parent)
			}
		}
	}
}

// 删除结点
// 1. 若删除结点无子结点，直接删除
// 2. 若删除结点只有一个子结点，用子结点替换删除结点
// 3. 若删除结点有两个子结点，用后继结点(右子树里最小的结点)替换删除结点
// 与bst的实现方式一样

// 删除结点被替代后，在不考虑结点的键值的情况下，对于树来说，可以认为删除的是替代结点
// 基于此，上面所说的3种二叉树的删除情景可以相互转换并且最终都是转换为情景1
// 情景2：删除结点用其唯一的子结点替换，子结点替换为删除结点后，可以认为删除的是子结点，
// 若子结点又有两个子结点，那么相当于转换为情景3，一直自顶向下转换，总是能转换为情景1。（对于红黑树来说，根据性质4.1，只存在一个子结点的结点肯定在树末了）
// 情景3：删除结点用后继结点（肯定不存在左结点），如果后继结点有右子结点，那么相当于转换为情景2，否则转为为情景1。

// 红黑树删除
// R表示替代结点，P表示替代结点的父结点，S表示替代结点的兄弟结点，SL表示兄弟结点的左子结点，SR表示兄弟结点的右子结点。
// R是即将被替换到删除结点的位置的替代结点，在删除前，它还在原来所在位置参与树的子平衡，平衡后再替换到删除结点的位置，才算删除完成。
// 情景：
// 1. 替换结点是红色结点：颜色变为删除结点的颜色(因为红色不影响平衡)；
// 2. 替换结点是黑结点:
//		2.1 替换结点是其父结点的左子结点;
//			2.1.1 替换结点的兄弟结点是红结点: 1. 将S设为黑色; 2. 将P设为红色; 3. 对P进行左旋; 4. 进行情景2.1.2.3的处理;
//			2.1.2 替换结点的兄弟结点是黑结点:
//				2.1.2.1 替换结点的兄弟结点的右子结点是红结点，左子结点任意颜色: 1. 将S的颜色设为P的颜色;2. 将P设为黑色; 3. 将SR设为黑色;4.对P进行左旋;
//				2.1.2.2 替换结点的兄弟结点的右子结点为黑结点，左子结点为红结点: 1. 将S设为红色; 2. 将SL设为黑色; 3. 对S进行右旋; 4. 进行情景2.1.2.1;
// 				2.1.2.3 替换结点的兄弟结点的子结点都为黑结点: 1. 将S设为红色;2. 把P作为新的替换结点; 3. 重新进行删除结点情景处理;
//		2.2 替换结点是其父结点的右子结点(与上面对称)
//			2.2.1 替换结点的兄弟结点是红结点: 1. 将S设为黑色; 2. 将P设为红色; 3. 对P进行右旋; 4. 进行情景2.2.2.3的处理;
//			2.2.2 替换结点的兄弟结点是黑结点:
//				2.2.2.1 替换结点的兄弟结点的左子结点是红结点，右子结点任意颜色: 1. 将S的颜色设为P的颜色;2. 将P设为黑色; 3. 将SL设为黑色;4.对P进行右旋;
//				2.2.2.2 替换结点的兄弟结点的左子结点为黑结点，右子结点为红结点: 1. 将S设为红色; 2. 将SR设为黑色; 3. 对S进行左旋; 4. 进行情景2.2.2.1;
// 				2.2.2.3 替换结点的兄弟结点的子结点都为黑结点: 1. 将S设为红色;2. 把P作为新的替换结点; 3. 重新进行删除结点情景处理;
// 总结： 1. 自己能平衡就自己平衡(1); 2. 自己不能平衡就找兄弟结点帮忙移一个黑色过来(除1，2.1.2.3, 2.2.2.3)
// 		3. 帮不了就把问题往上抛

func (r *redBlackTree) Remove(key int) {
	// step1. 删除指定node
	var removeNode, replace *rbNode
	removeNode, isOk := r.traverse(r.root, key)
	if !isOk {
		return // 未找到
	}

	// 叶子结点,直接删除自身
	if removeNode.left == nil && removeNode.right == nil {
		if removeNode.parent.left == removeNode {
			removeNode.parent.left = nil
		} else {
			removeNode.parent.right = nil
		}
		return
	}

	// 只有一个子树，直接替换子树
	if removeNode.left == nil {
		replace = removeNode.right
	} else if removeNode.right == nil {
		replace = removeNode.left
	} else {
		// 待删除节点左右子树均不为空的情况
		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		replace = r.min(removeNode) // 找到右子树里最小的结点
	}
	if replace.left != nil || replace.right != nil {
		panic("replace node 选择错误, 请检查逻辑")
	}

	// step2. 开始重新平衡红黑树
	r.balanceDelete(removeNode, replace)

	// 注意： 交换的是替换结点和删除结点的颜色和值，也就是此时
	// replace -> 替换结点的地址(一个叶子结点)，值和颜色是删除结点;
	// remove -> 删除结点的地址， 替换结点的颜色和值
	r.exchange(replace, removeNode)

	// 删除结点
	if replace.parent.left == replace {
		replace.parent.left = nil
	} else {
		replace.parent.right = nil
	}
}

// 删除掉以node为根的二分搜索树中键值为key的节点, 递归算法
// 返回删除节点后新的二分搜索树的根, 与被删除结点的颜色
func (r *redBlackTree) remove(n *rbNode, key int) (replace *rbNode, deleteColor Color) {
	if n == nil {
		return nil, RED
	}
	if key < n.key {
		n.left, deleteColor = r.remove(n.left, key)
		n.left.parent = n // 维护结点的parent
		return n, deleteColor
	} else if key > n.key {
		n.right, deleteColor = r.remove(n.right, key)
		n.right.parent = n
		return n, deleteColor
	} else { // 删除逻辑
		// 待删除结点只有一个子结点的情况
		if n.left == nil {
			r.count--
			return n.right, n.color
		}
		if n.right == nil {
			r.count--
			return n.left, n.color
		}
		// 待删除节点左右子树均不为空的情况
		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := r.min(n.right) // 找到右子树里最小的结点，提取出来

		// 将右子树最小结点替换结点n
		successor.right = r.removeMin(n.right) // 这里有count--的逻辑
		successor.left = n.left
		return successor, n.color
	}
}

func (r *redBlackTree) balanceDelete(rm, rep *rbNode) {
	// 1. 替换结点是红色结点：颜色变为删除结点的颜色(因为红色不影响平衡)；
	if rep.color == RED {
		rep.color = rm.color
	} else { // 2. 替换结点是黑结点:
		// 2.1 替换结点是其父结点的左子结点;
		s := rep.getBrother()
		p := rep.parent

		if p.left == rep {
			// 2.1.1 替换结点的兄弟结点是红结点(因为替换结点是黑结点，说明肯定有兄弟结点)
			if s.color == RED {
				// 1. 将S设为黑色; 2. 将P设为红色; 3. 对P进行左旋; 4. 进行情景2.1.2.3的处理;
				s.color = BLACK
				p.color = RED
				r.leftRotation(p)
				r.balanceDelete(rm, rep)
			} else {
				// 2.1.2 替换结点的兄弟结点是黑结点:
				// (这里将nil子结点也算黑结点, 所以要注意nil判断)
				if s.right != nil && s.right.color == RED {
					// 2.1.2.1 替换结点的兄弟结点的右子结点是红结点，左子结点任意颜色
					// 1. 将S的颜色设为P的颜色;2. 将P设为黑色; 3. 将SR设为黑色;4.对P进行左旋;
					s.color = p.color
					p.color = BLACK
					s.right.color = BLACK
					r.leftRotation(p)
				} else if s.left != nil && s.left.color == RED &&
					(s.right == nil || s.right.color == BLACK) {
					// 2.1.2.2 替换结点的兄弟结点的右子结点为黑结点，左子结点为红结点:
					// 1. 将S设为红色; 2. 将SL设为黑色; 3. 对S进行右旋; 4. 进行情景2.1.2.1;
					s.color = RED
					s.left.color = BLACK
					r.rightRotation(s)
					r.balanceDelete(rm, rep)
				} else if (s.left == nil || s.left.color == BLACK) &&
					(s.right == nil || s.right.color == BLACK) {
					// 2.1.2.3 替换结点的兄弟结点的子结点都为黑结点:
					// 1. 将S设为红色;2. 把P作为新的替换结点; 3. 重新进行删除结点情景处理;
					s.color = RED
					r.balanceDelete(rm, p)
				}
			}
		} else {
			// 2.2 替换结点是其父结点的右子结点(与上面对称)

			// 2.2.1 替换结点的兄弟结点是红结点(因为替换结点是黑结点，说明肯定有兄弟结点)
			if s.color == RED {
				// 1. 将S设为黑色; 2. 将P设为红色; 3. 对P进行右旋; 4. 进行情景2.2.2.3的处理;
				s.color = BLACK
				p.color = RED
				r.rightRotation(p)
				r.balanceDelete(rm, rep)
			} else {
				// 2.2.2 替换结点的兄弟结点是黑结点:
				// (这里将nil子结点也算黑结点, 所以要注意nil判断)
				if s.left != nil && s.left.color == RED {
					// 2.2.2.1 替换结点的兄弟结点的左子结点是红结点，右子结点任意颜色
					// 1. 将S的颜色设为P的颜色;2. 将P设为黑色; 3. 将SL设为黑色;4.对P进行右旋;
					s.color = p.color
					p.color = BLACK
					s.left.color = BLACK
					r.leftRotation(p)
				} else if s.right != nil && s.right.color == RED &&
					(s.left == nil || s.left.color == BLACK) {
					// 2.2.2.2 替换结点的兄弟结点的左子结点为黑结点，右子结点为红结点:
					// 1. 将S设为红色; 2. 将SR设为黑色; 3. 对S进行左旋; 4. 进行情景2.2.2.1;
					s.color = RED
					s.right.color = BLACK
					r.leftRotation(s)
					r.balanceDelete(rm, rep)
				} else if (s.left == nil || s.left.color == BLACK) &&
					(s.right == nil || s.right.color == BLACK) {
					// 2.2.2.3 替换结点的兄弟结点的子结点都为黑结点:
					// 1. 将S设为红色;2. 把P作为新的替换结点; 3. 重新进行删除结点情景处理;
					s.color = RED
					r.balanceDelete(rm, p)
				}
			}
		}
	}
}
