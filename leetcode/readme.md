# go语言刷leetcode

|目录                |   注释   |
|---                |---|
|basic              |算法基础知识|
|sort               |排序|
|array              |数组|
|linkedList         |链表|
|string             |字符串|
|queue              |队列|
|tree               |树|
|hashTable          |哈希表|
|backtracking       |回溯算法|
|greedyAlgorithm    |贪心算法|
|dynamicProgramming |动态规划|

## 基础内容索引index
0. [算法题解笔记](./basic/README.md)

1. 排序

   O(n^2)的算法
    - [选择排序](./basic/sort/selectionSort/selectionSort.go)
    - [插入排序、插入排序的改进](./basic/sort/insertionSort/insertionSort.go)

   O(nlogn)算法
    - [自顶向下、自底向上归并排序](./basic/sort/mergeSort/mergeSort.go)
    - [快速排序,双路、三路快排](./basic/sort/quickSort/quickSort.go)

   补充
    - [各种排序的性能比较](./basic/sort/main.go)

2. 链表
    - [简单链表实现](./basic/linkedList/simple/linkedList.go)

3. 堆

    - [堆的实现思路](./basic/heap/heap.md)
    - [最大堆](./basic/heap/maxHeap.go)
    - [最小堆](./basic/heap/minHeap.go)
    - [最小索引堆](./basic/heap/indexMinHeap.go)

4. 树

    - [树的实现思路](basic/tree/README.md)
    - [bs二分查找法](./basic/tree/bs.go)
    - [二分查找树](./basic/tree/bst.go)

5. 并查集
    - [并查集基础、实现并查集并优化](./basic/unionFind/unionFind.md)

6. 图论
    - [图的实现](./basic/graph/graph/graph.go)
    - [连通图](./basic/graph/graph/component.go)
    - [稠密图](./basic/graph/graph/denseGraph.go)
    - [稀疏图](./basic/graph/graph/spareGraph.go)
    - [寻路](./basic/graph/graph/path.go)
    - [带权图与带权图的边](./basic/graph/weightGraph/weightGraph.go)
    - [带权稠密图 - 邻接矩阵](./basic/graph/weightGraph/denseWeight.go)
    - [带权稀疏图 - 邻接表](./basic/graph/weightGraph/spareWeight.go)


7. 最小生成树
    - [最小生成树](./basic/graph/mst/readme.md)
    - [prim算法](./basic/graph/mst/primMst.go)
    - [lazy prim](./basic/graph/mst/lazyPrimMST.go)
    - [kruskal算法](./basic/graph/mst/kruskalMST.go)

8. 最短路径问题
    - [最短路径问题、松弛操作](./basic/graph/shortest/readme.md)
    - [Dijkstra算法](./basic/graph/shortest/dijkstra.go)
    - [Bellman-Ford算法](./basic/graph/shortest/bellmanFord.go)

## 学习路线

1. 查找基础;
2. 排序基础;
3. 数组: 数据结构的封装、数组的扩容和缩容、均摊时间复杂度分析
4. 栈与队列: 循环队列的实现、双端队列的实现;
5. 链表:
6. 归并排序: 理解分治算法、快速求解逆序对个数
7. 快排、希尔排序
8. 二分查找法
9. 二分搜索树: DFS 和 BFS、 集合和映射
10. 堆和优先队列: heapify, 堆排序
11. AVL树和红黑树: 左旋转右旋转维护平衡、2-3树和红黑树的等价性、
12. 线段树、并查集和字典树: 线段树和区间查询问题、并查集解决集合查询问题、路径压缩、字典树和模式匹配基础
13. 哈希表和SQRT分解:
14. 非比较排序: 计数排序法、桶排序、基数排序法、LSD 和 MSD、
15. 模式匹配算法: 深入理解 KMP 算法、KMP 中的 LPS 数组、状态机、状态机的基础应用、RK 算法
16. B类树和数据库、Knuth 洗牌算法、蓄水池抽样

## reference

- [玩转算法](https://github.com/liuyubobobo/Play-with-Algorithms)







