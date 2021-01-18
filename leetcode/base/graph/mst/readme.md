# 最小生成树
最小生成树(Minimum Spanning Tree, MST),
一个有 n 个结点的连通图的生成树是原图的极小连通子图，且包含原图中的所有 n 个结点，
并且有保持图连通的最少的边。最小生成树可以用kruskal（克鲁斯卡尔）算法或prim（普里姆）算法求出。

> 简化版： 只针对带权图连通图,存在一棵树,连接了图里的所有节点,所有边的权值相加是最小的

# Prim算法
1. 切分 cut: 把图的所有结点分为两个部分，称为一个切分
2. 横切边 Crossing Edge: 如果一个边的两个端点，属于切分不同的两边，这个边称为横切边 Crossing Edge
3. 切分定理 cut property: 给定任意的切分，横切边中权值最小的边必然属于最小生成树

## lazy prim
lazy prim需要用到最小堆结构，见[最小堆](./minHeap.go)

代码见[lazyPrimMST](./lazyPrimMST.go)

## prim
代码见[prim](./primMst.go),需要用到[最小索引堆](./indexMinHeap.go)

# kruskal算法
[克鲁斯卡尔算法代码](./kruskalMST.go),需要用到[并查集](./unionFind.go)


