# 分布式缓存GeeCache

> 来源于极客兔兔的[七天从零实现分布式缓存](https://geektutu.com/post/geecache-day4.html)

1. 实现了基于*LRU算法*的cache结构；
2. 实现了并发安全；
3. 提供http服务；
4. hash计算使用一致性hash实现;
5. 实现简单的singleflight缓冲器，在同一时间内一个key只会发起一次请求;


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


## 缓冲雪崩、击穿、穿透的解决办法
### 雪崩
雪崩是在某一时间缓存的大量热点key过期，导致大量请求走到DB；

解决办法是缓存的ttl增加随机数；

### 击穿
击穿指某个热点key过期，然后收到大量请求走到DB；

解决办法： 1. 未命中缓存，先加分布式锁，再读DB设置缓存，再释放锁；2. 本文实现的缓冲器singleFlight

### 穿透
穿透指大量查询不存在的数据，不会命中缓存，大量查询查到DB

解决办法是1. 缓存这个key，值设置为nil; 2. 分布式布隆过滤器;