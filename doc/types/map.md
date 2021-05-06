# map
## 前言
1. map是hash table实现,无序;
2. map*不是线程安全的*; (golang 1.9 在sync包里实现了并发安全的map)
3. hash冲突常用*线性探测*或者*拉链法*
   开放定址（线性探测）和拉链的优缺点
    - 拉链法比线性探测处理简单
    - 线性探测查找是会被拉链法会更消耗时间
    - 线性探测会更加容易导致扩容，而拉链不会
    - 拉链存储了指针，所以空间上会比线性探测占用多一点
    - 拉链是动态申请存储空间的，所以更适合链长不确定的

## 数据结构





## reference

- [go里面的哈希表](https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/)