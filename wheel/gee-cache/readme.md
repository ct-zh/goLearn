# 分布式缓存GeeCache

> 来源于极客兔兔的[七天从零实现分布式缓存](https://geektutu.com/post/geecache-day4.html)
> github: [7days-golang](https://github.com/geektutu/7days-golang)

1. 实现了基于*LRU算法*的cache结构；
2. 实现了并发安全 => 加锁；
3. 提供http服务；
4. hash计算使用一致性hash实现;
5. 实现简单的singleflight缓冲器，在同一时间内一个key只会发起一次请求;
6. 实现了中间层api与peers分布式节点;
7. 采用protobuf;

## LRU算法实现
使用的链表实现:
- 一共有两个结构体, cache和cacheNode; cache由map与list实现,节点都是保存的cacheNode的指针,节约内存;
- lru算法使用list实现,每次新增时将节点追加到list尾部,更新时则遍历list拿到节点移动到尾部;
- 如果因为设置内存达到上限或者其他原因触发删除操作,则移除list头部节点,同时从map里面删除对应的数据;

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

## singleflight缓解击穿问题
这里模拟的是singleflight的实现, 存在两个结构体Group与call, Group保存了当前所有call的map,和一个锁; call保存了请求信息与waitGroup:
- 请求缓存时多套一层Do函数;
- Do函数的内容是: 先加Group锁,将请求写入map,发起请求,同时call.waitGroup+1;释放锁;
- 其他请求过来加锁, 发现map里面已经有call了,则解锁并waitGroup.wait等待;
- 先前的请求拿到数据,写一份在call里面,waitGroup.Done,返回;
- 其他请求wait解锁,从call里面拿数据,返回;

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