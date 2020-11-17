## 实现思路
### 堆
1. createByArr构造函数初始化堆，同时完成shiftDown操作
2. 增加insert方法，完成shiftUp操作
3. 依靠shiftDown方法，增加extract(推出元素)方法
4. 完成getTop方法

### 索引堆
1. 实现一个最大堆/最小堆
2. 增加indexes数组
3. 所有取值代码`data[i]`修改为`data[indexes[i]]`，insert方法增加index参数
4. 将所有交换代码`swap(data[i], data[j])`修改成`swap(index[i],index[j])`
5. 增加方法extractIndex，返回的是最值的索引
6. 增加getTopIndex方法
7. 增加reverse数组
8. 所有发生indexes操作的地方加上反向的reverse操作
9. 增加Change方法（将索引为index的值修改为v,同时对index做shiftUp与shiftDown）
10. 增加Contain方法(查看索引index的位置是否存在元素,使用reverse实现)

## 数据结构
### 堆
```go
package heap
type heap struct {
    data        []int       // 堆的数据 从1开始索引
    count       int         // 数据总数
    capacity    int         // 堆容量，这个参数可以没有
    shiftUp     func(k int) // 对新加入的子节点做交换操作，使其所在的树满足最小堆的定义
    shiftDown   func(k int) // 对根节点做交换操作，使其所在的树满足最小堆的定义
    Insert      func(v int) // 添加元素，做shiftUp操作
    ExtractMin  func() int  // 取出最小值
    GetMin      func() int  // 获取最小值
    Size        func() int
    IsEmpty     func() bool
}
```

### 索引堆
```go
package indexHeap
type indexHeap struct{
	data                []int
    indexes             []int
    reverse             []int
    count               int
    capacity            int
    shiftUp             func(k int)
    shiftDown           func(k int)
    Insert              func(index int, v int)
    ExtractMin          func() int
    GetMin              func() int
    Size                func() int
    IsEmpty             func() int                  // 这个方法之前都是普通堆的方法
    getVal              func(k int) int             // 在堆内通过索引k获取对应的值
    ExtractMinIndex     func() int                  // 使用indexes实现，获取最小值的索引
    GetMinIndex         func() int                  // 获取最小值的索引
    GetItem             func(index int) int         // 同getVal方法，不过这个是对外的，所以索引index需要+1
    Change              func(index int, v int)      // 通过索引index改变其值为v，修改完之后需要reverse完成shiftDown和shiftUp操作
    Contain             func(index int) bool        // 判断索引index处是否存在值，使用reverse实现
}
```

## 理论
1. 根节点大于等于所有子节点的二叉树，就被称为最大堆;同理，根节点小于等于所有子节点的二叉树，就被称为最小堆
2. 堆排序的流程是：每次从堆中获取到根节点，然后将堆的根节点与最后一个节点交换，删除最后一个节点（即之前的根节点），同时对堆做heapify操作

4. 索引堆：数据存储在data中，建立indexes数组，将data的索引写入indexes，所有堆的操作都在indexes里面进行

6. 索引堆的注意事项：索引堆是从1开始索引的，但是对外面用户来说是从0开始索引的；所以所有index相关的操作，从外面传进来需要+1，返回到外面需要-1；
 