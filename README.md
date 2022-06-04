# golang编程笔记
> 笨鸟先飞.✊✊✊

## index

|目录|注释|
|---|---|
|[doc](./doc/README.md)🐣🐣🐣|基础语法与demo|
|[leetcode](./leetcode/readme.md)🚗🚗🚗|golang版本刷题整理|
|[src](./src/README.md)✈️✈️✈️|开发场景demo|
|[wheel](./wheel/)|一些脚手架/插件/扩展的demo，类似src|

<img src="https://tip.golang.org/lib/godoc/images/footer-gopher.jpg">

- [golang面试常见问题](https://github.com/ct-zh/interview/tree/master/go)

## 第三方文章摘
- [你所应该知道的A/B测试基础](http://blog.leapoahead.com/2015/08/27/introduction-to-ab-testing/)
- [让代码审查扮演更好的角色](http://blog.leapoahead.com/2016/10/04/code-review-one-step-further/)
- [“函数是一等公民”背后的含义](http://blog.leapoahead.com/2015/09/19/function-as-first-class-citizen/)
- [从开源项目中获得的docker经验](http://blog.leapoahead.com/2015/10/07/docker-lessons-learned-md/)

### golang基础
- [for 和 range 的性能比较](https://geektutu.com/post/hpg-range.html)：range在迭代大value的数据时性能很差，因为range给的item是value的拷贝；
- [切片(slice)性能及陷阱](https://geektutu.com/post/hpg-slice.html)：大切片使用copy替代re-slice防止不被gc；
- [Reflect反射的性能](https://geektutu.com/post/hpg-reflect.html): 避免使用反射，比如官方json包就是用的反射，可以用[easyjson](https://github.com/mailru/easyjson)代替
- [字符串拼接性能](https://geektutu.com/post/hpg-string-concat.html): 尽量使用`strings.Builder`来进行字符串拼接;
- [Bilibili 毛剑：Go 业务基础库之 Error ](https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247487124&idx=1&sn=0f6141c2ccd9a0abc4baf26e04f0fd4c&source=41#wechat_redirect)
- [毛剑：Bilibili 的 Go 服务实践（上篇）](https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247487505&idx=1&sn=c9de6535528d2102bee364937201f6e6&source=41#wechat_redirect)
- [毛剑：Bilibili 的 Go 服务实践（下篇）](https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247487504&idx=1&sn=9b8663676ee689e0fcd4b990ecf99f3d&source=41#wechat_redirect)
- [Gopher China 2019 讲师专访-bilibili架构师毛剑 ](https://www.sohu.com/a/303913388_657921)
- [Gopher China 2019](https://www.bilibili.com/video/BV1c4411g77Y?p=5)
- [Gopher China 2019 PPT](https://github.com/gopherchina/conference/blob/master/README.md)

### 并发编程
- [读写锁与互斥锁的性能比较](https://geektutu.com/post/hpg-mutex.html): 读写锁的读写性能明显比互斥锁更好；*这篇文章里面有一段snyc的源码注释:互斥锁如何实现公平*，不清楚的话看下面这篇；
- [sync.mutex 源代码分析](https://colobu.com/2018/12/18/dive-into-sync-mutex/)
- [如何退出协程 goroutine (超时场景)](https://geektutu.com/post/hpg-timeout-goroutine.html): 子协程在读channel需要考虑退出的问题
- [如何退出协程 goroutine (其他场景)](https://geektutu.com/post/hpg-exit-goroutine.html)：子协程在读channel需要考虑退出的问题
- [控制协程的并发数量](https://geektutu.com/post/hpg-concurrency-control.html): 1. 使用buffer channel；2. 使用协程池；
- [sync.Once提升性能](https://geektutu.com/post/hpg-sync-once.html): 提到了最常用的字段作为结构体的第一个字段(热路径,不需要计算偏移);
    
    > 结构体第一个字段的地址和结构体的指针是相同的，如果是第一个字段，直接对结构体的指针解引用即可。如果是其他的字段，除了结构体指针外，还需要计算与第一个值的偏移(calculate offset)。在机器码中，偏移量是随指令传递的附加值，CPU 需要做一次偏移值与指针的加法运算，才能获取要访问的值的地址。因为，访问第一个字段的机器代码更紧凑，速度更快。


### debug、性能分析与编译
- [benchmark基准测试](https://geektutu.com/post/hpg-benchmark.html)
- [pprof 性能分析](https://geektutu.com/post/hpg-pprof.html)
- [减小 Go 代码编译后的二进制体积](https://geektutu.com/post/hpg-reduce-size.html)
- [死码消除与debug模式](https://geektutu.com/post/hpg-dead-code-elimination.html)


## 架构
- [bilibili技术总监毛剑：B站高可用架构实践](https://zhuanlan.zhihu.com/p/139258985)
