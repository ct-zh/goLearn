# 分布式缓存GeeCache

> 来源于极客兔兔的[七天从零实现分布式缓存](https://geektutu.com/post/geecache-day4.html)

1. 实现了基于*LRU算法*的cache结构；
2. 实现了并发安全；
3. 提供http服务；
4. hash计算使用一致性hash实现



## 一致性hash的实现
实现的是简单版，使用crc32当作hash函数，计算出来的hash值是int型，不考虑hash冲突等问题；
> crc32返回的是uint32
- 定义一个int切片当作hash环，使用map记录hash与key的映射关系（虚拟节点）
- 对key进行hash，虚拟节点则是以在key前面加入编号的规则进行hash
- 对int切片进行排序
- 将hash数据以append的方式写入int切片，hash与key的映射写入map
- 任意传入一个查找str（如ip地址，uid等）
- 对str进行hash得到某个值，在切片找到第一个大于等于strhash的hash值
- 从map处使用hash取出映射，即可得到对应的key；在分布式系统中key一般代表服务器编号；
- 也就是说只要str不变化，服务器处于上线状态，str每次获取的都是同一台服务器
- 服务器下线，int切片对应的数据删除，str继续找到下一个服务器，不影响用户体验


