# chapter2

> 对比ppt一起看

## section01 - 开源框架对比 (主要是看一下开源代码的实现)
httprouter 必看的项目, 是大多数 web框架的爸爸;
fasthttp 也可以看一下 ( [fasehttp快在哪里](https://xargin.com/why-fasthttp-is-fast-and-the-cost-of-it/) )

- chi框架 - router分组、middleware/Chain 稍微可以研究一下
- gin框架 - binding将decoder与validator合二为一, 可以看一下(用到了反射); 同样可以看一下validator; 
- echo 中规中矩吧
- beego mvc框架, 曾经的元老

> what is middleware ? 你现在能说明白吗 ? 

### 微服务框架
微服务比web框架需要更多的组件, 如:

- Config：配置管理组件
- Logger：遵守第三⽅⽇志收集规范的⽇志组件; 一个公司内部的日志格式必须统一;
- Metrics：使框架能够与 Prometheus 等监控系统集成的 metrics 组件
- Tracing：遵守 OpenTelemetry 的 tracing 组件
- Registry：服务发现组件
- MQ：可以切换不同队列实现的 mq 组件; 应该提供一套统一的api; 
- 依赖注⼊：wire，dig 等组件

框架如:
- go micro 开源框架;
- go zero
- yoyoggo
- dubbo go
- kratos

> 滴滴曾经 consul 出现过广播风暴 导致异常; 后续换成了自研的服务发现; 大多数服务发现都是自研的, 如阿里开源的nacos
> 服务发现应该是一个AP还是CP的问题
> 如果服务发现挂了, 服务会找不到其他服务

更复杂的可以看一下 ERDA

### 框架的评判

设计框架需要考虑的问题:
- 自动化
    - 代码自动生成/接入sdk自动生成
    - 自动发布/自动生成接口文档
    - cli工具
- 平台化
    - IDL服务定义, 统一的平台管理
    - 接口文档可检索
    - 服务部署流水化
- 集成化
- 组件化
- 插件化
- 通用化(不是公司框架考虑的问题)

> 再次推 Google SRE

### web框架观点
依赖树的问题: 同一个模块中A SDK与B SDK都依赖于 另一个模块, 如配置中心; 当配置中心代码从v1升级到v2时, 为了稳定性可能A SDK先升级, B SDK未必会同时升级. 此时A SDK依赖配置中心V2, B SDK依赖配置中心V1, 这种问题可能在编译时通过, 在运行时出现问题;


## section02 - 框架底层
- middleware
    - > 参考代码: https://github.com/gin-gonic/contrib/blob/master/secure/example/example.go
    - 基本思路是，把功能 性(业务代码)和⾮功能 性(⾮业务代码)分离
    - 不同语言/架构有不同的称呼: 责任链模式/拦截器/装饰器/代理/filter/洋葱, 建议统一为middleware吧

- 路由
    - > 参考项目: https://github.com/julienschmidt/httprouter/tree/master
    - 通常做文本匹配用的比较多的是字典树trie, 但是字典树每个子节点都是一个字, 比较浪费空间; 当前route用的比较多的是优化版的压缩前缀树(radix tree), 单个节点不再是一个字母, 而是路由路径的一段; 
        - 每一个http method都对应一棵radix tree;
        - 路由冲突

- validator
    - > 源码研究: https://github.com/go-playground/validator

- request bind
    - > bind  json https://github.com/gin-gonic/gin/blob/master/binding/json.go
    - 解析不同的请求格式, 返回对应的请求数据;
    - 将其与用户的结构体绑定

- sql bind
    - 参考项目:gorm、sqlx;
    - 解决新手总是忘记close的问题;
    - 给sql的占位符与变量绑定, 并且处理变量防止注入
    - 将结果与传入的结构体做绑定;

- context
    - 按照公司实际的需求来实现, 都是些逻辑性的东西

- log

## section3 - 业务分层,领域驱动设计
- 最早期为mvc架构, model/view/controller; 后续view层随着前后端分离不再需要, 服务端只存在model与controller, 几乎所有业务逻辑都写在controller里面;
- 之后很多人再次细分了后端架构, 大多数还是三层:
    - model -> dao, 该层model声明所有对象; dao层则是这些对象的控制方法(data access object), 如查询数据库、查询缓存等;
    - controller, 有的也叫做handler, 对外部请求的校验、协议处理等等, 将其数据转化成model对应的结构体传入logic; 同时也根据协议之类的处理返回的数据;
    - logic, 有的也叫做buz、domain object、service等等, 业务逻辑主体;
- 尽管controller层分走了与外部请求交互的相关逻辑, dao层分走了与外部数据源交互的相关逻辑, 但是随着业务的发展,controller层的业务逻辑仍然可能变的非常重;社区的人于是提出了两种模式:
    - 贫血模式: 业务逻辑都在logic层上
    - 充血模式: 要让domain object 也就是 entity 有更多的逻辑; 通过聚合来组合entity逻辑; (这里的domain object就是model层的所有对象模型)

- DDD社区推荐使用REPO的模式, 给各个domain object定义对应的repo; 例如一个用户的REPO
    ```go
    type CustomerRepo interface{
      func changeUserName(firstName, lastName string) error
      func disconnectHomeTel() error
      func refreshExtraInfo() error
    }
    ```
    
    相关repo的代码实现都要很简单, 主流程根据不同用例选择执行不同的函数, 而不是用数据建模, 把所有逻辑都放在一个函数里面;

-  SOLID  [面向对象的五个原则]([SOLID (物件導向設計) - 維基百科，自由的百科全書 (wikipedia.org)](https://zh.wikipedia.org/zh-tw/SOLID_(面向对象设计)) 即 单一功能、开闭(对扩展开放,对修改封闭)、里氏替换、接口隔离与依赖反转;

  - 特地提一下go语言中的D = dip 依赖反转, 在代码中体现应该是如下所示:

    ```go
    // 普通的业务逻辑
    // service层 user.go
    // 依赖关系是: controller -> logic -> dao
    func (s *Service) GetUserInfo(uid int64) (model.UserInfo, error) {
      m := model.UserInfo{}
      m.BaseInfo = s.dao.GetBaseUserInfo(uid)
      m.ExtraInfo = s.dao.GetExtra(uid)
      // ....
    }
    
    // 修改为依赖反转, 所有依赖关系收束在logic里
    // controller -> logic <- dao
    
    type UserInfoRepo interface{	// dao层依赖该repo, 应该实现该repo的所有方法
      GetBaseUserInfo(uid int64) (baseInfo model.BaseInfo, err error)
      GetExtra(uid int64) (extra model.Extra, err error)
      // .....
    }
    
    func (s *Service) GetUserInfo(uid int64) (model.UserInfo, error) {
      repo := NewUserInfoRepo(s.dao)		// 该处可以传dao, 从数据库/缓存中获取数据; 也可能传的是client, 从其他服务中获取数据; 但不管如何, 该处解除了对dao层的依赖, 只要是任何实现了repo接口的数据源都可以传入;
      m.BaseInfo = repo.GetBaseUserInfo(uid)
      m.ExtraInfo = repo.GetExtra(uid)
    }
    ```

### 整洁架构

整个整洁架构绝大多数都是根据DIP 依赖注入/控制反转来实现的, 参考: [eminetto/clean-architecture-go-v2: Clean Architecture sample (github.com)](https://github.com/eminetto/clean-architecture-go-v2)

- 不与框架绑定 - 业务层不应该有任何框架相关的代码; 

  ```go
  // 比如框架处理请求/发送请求等逻辑, 应该在业务外层进行处理, 
  // 业务层如果对这些存在依赖, 应该将依赖的行为抽象出来;
  
  // 例如 logic层 依赖框架进行client发送http请求
  func (s *Service) getUserInfo(uid int64) (model.UserInfo, error) {
    commInfo, err := user_center.NewClient().GetUserCommInfo(uid) 
    // ....
  }
  
  // 控制反转
  type UserClient interface{
    GetUserCommInfo(int64) (UserCommInfo, error)
  }
  func (s *Service) getUserInfo(uid int64) (model.UserInfo, error) {
    userClient := NewUserClient(s)
    commInfo, err := userClient.GetUserCommInfo(uid) 
    // ....
  }
  ```

- 可测试 - 通过依赖注入, 定义好外部环境的行为, 从而在测试时可以进行替换;

  ```go
  // 下面的代码依赖dao, 在单元测试的时候还需要传入dao实例
  func (s *Service) GetUserInfo(uid int64) (model.UserInfo, error) {
    m := model.UserInfo{}
    m.BaseInfo = s.dao.GetBaseUserInfo(uid)
    m.ExtraInfo = s.dao.GetExtra(uid)
    // ....
  }
  
  // 将dao方法抽象出来
  type UserInfoRepo interface{
    GetBaseUserInfo(uid int64) (baseInfo model.BaseInfo, err error)
    GetExtra(uid int64) (extra model.Extra, err error)
  }
  
  // 依赖注入repo参数, 只要repo参数实现了所需的接口就行了, 这里可以传dao也可以传一些方便测试的实例
  func (s *Service) GetUserInfo(repo UserInfoRepo, uid int64) (model.UserInfo, error) {
    m.BaseInfo = repo.GetBaseUserInfo(uid)
    m.ExtraInfo = repo.GetExtra(uid)
  }
  ```

- 不与ui绑定
- 不与数据库绑定
- 不与外部代理绑定 ( 公司里面基本不可能实现 )
  - 没有配置化的agent ( consul、config等等 )
  - 没有service mesh
  - 没有metrics采集模块


可以看一下这个文章, 感觉讲得更清楚: https://mp.weixin.qq.com/s/30iZP_93oSDo8844KDN3IA


### DDD

基本名词

- Value Object: 不可变的、可比较的值;

- Entity: 与`Value Object`最大的区别是多了一个ID字段, 有id就说明是可变数据支持增删查改; 基本上可以相当于数据库一条记录/mq一条消息/缓存的一个key;

- Aggregate: 

  - value object 与 Entity的聚合内容, 与entity有一定相似性, 也可能有id, 也是可变的;  

    ```go
    type UserInfo struct {	// 用户信息聚合 aggregate
      	Uid int64					// aggregate也可能有id值
      	Base BaseUserInfo // 这部分对应一个entity, 从外部查询到用户的基础信息
      	AuthLevel int8    // value object, 这里应该代表用户的权限等级
    }
    ```

  - 一定会有一个对应的repo;

  - 聚合需要保证内容的一致性; 

    - 例如商品购买操作, 可能聚合了  优惠券消耗、订单下单、下单活动完成等一系列操作, 这次操作需要保证一致性,即要么全部成功要么全部失败; 这些内容在商品的repo里维护; 
    - 上述商品操作只是一个例子, 实际优惠券、订单、活动应该分布在多个实体中;

- Aggregate Root:

  - 聚合根也是聚合, 但是对外暴露

- Repo Pattern:



### 插件化架构


## section4 - 优雅代码

- 常见坑
  - 返回值为interface与nil做比较

    ```go
    func fn(arg int) interface{}	{ // 函数返回interface{}, 如果没显式return nil, 那后续不能做判断result != nil
      var result interface{}
      // ...
      return result
    }	
    
    if result := fn(1); result != nil {
    }
    ```
  
  - 不处理err
  - 闭包捕获循环变量
  
    -  ` for k, v range data ` , k、v 传入协程中处理 
    - 1.22处理了该问题
  - receiver 是值而不是指针, 方法内对receiver的操作是无效的
  
    -   `func (s Service) Join(xx int)` 该函数内部对Service修改是无效的
  - 参考effective-go 和 50 shades
  
- linter

  - 对于常见的坑应该用linter来避免; 保证代码的下限, 而不是上限; 项目里有新手时能保证不会犯一些低级傻逼的错误;
  - go vet
  - [Codecov Free Trial - Codecov](https://about.codecov.io/codecov-free-trial/?utm_source=google&utm_medium=cpc&utm_campaign=Google_Search_NB_General_EMEA_SignUp_Alpha&utm_content=g&utm_term=code coverage&gclid=Cj0KCQjw4MSzBhC8ARIsAPFOuyUF906oPnoHwTZD_8tF3iUMe19hyJ3PYd2KL7ueqo2QQEd2cJHldSUaAnC9EALw_wcB&utm_id={20652195004}&gad_source=1) 检查你的开源项目pr是否存在问题, 单元测试覆盖率是否足够
  - golangci-lint
  - review-dog

- 优雅代码

  - 参考《重构》中间的坏代码: 命名不合适、
  - 活文档: https://xargin.com/about-living-doc/
  - [免费在线学习代码重构和设计模式 (refactoringguru.cn)](https://refactoringguru.cn/)

- 写测试 

  - https://segment.com/blog/5-advanced-testing-techniques-in-go/
  - 使用testify https://www.jetbrains.com/help/go/using-the-testify-toolkit.html#generate_test_for_package
  - 流量回放, 早期的需要修改runtime, 现在可以试试service mesh adapt

- 错误注入原理与实现

  - 参考: [pingcap/failpoint: An implementation of failpoints for Golang. (github.com)](https://github.com/pingcap/failpoint) 侵入代码注入错误, 编译时增加参数来执行这段代码

- code review

  - 推荐平台: Gerrit 


## section5 - 中台

- 中台发展
- Go语言在中台中的应用 / 一个数据中台迭代过程
  - jpath: 直接拿到json中特定的字段
- 问题



## section6 性能调优

- 非程序代码消耗的时间

  - 硬件的响应延迟;

  - 网络协议占用的时间; 了解相关协议的底层流程:http2、http3、rpc、Thrift;

- 代码耗时

  - 逃逸分析; 能知道变量是分配在堆上还是栈上; 预估gc阶段的次数与耗时影响;
  - benchmark; benchmark可以输出profile ( leetcode必备 )
  - 零内存分配 ( zero garbage / allocation ): 可以使benchmark 的 alloc 为0; 一般是使用sync.Pool来重用变量, 完全消灭堆分配;
    - 方法论: benchmark 跑 memory profile, 找到其中内存分配的部分, 将对应变量用sync.Pool重用;
    - 就是所谓的池化;
  - cache line 加载问题: 
    - cpu cache一次会将多个字节的内容从内存加载到cpu cache line中;
    - 命令: `getconf LEVEL1_DCACHE_LINESIZE` 可以得到cache line1的大小;  一般服务器cpu cache L1大小都为64字节;
    - 问题: 并发执行两个程序, 这两个程序对应结构体字段在内存上相邻, 导致被加载到同一个cache line上面; 当核A修改其中一个结构体, 核B并发修改另一个结构体, 会导致两个核反复操作同一个cache line, cpu内部会有消息传递, 造成类似广播风暴的效果, 降低并发程序性能.
    - 解决办法: 增加数据填充字段; 参考可见信号量`sema`结构体的实现;
  - 二维数组横遍历与竖着遍历, 横遍历效率更高
  - 局部性原理: 时间局部性与空间局部性; 

- 生产环境优化

  - 压测定位问题; api压测、全链路压测; 

    - 关注指标: 发压端QPS、服务返回错误数、P99

  - 优先选择应用层逻辑优化, 其次再往底层优化;

  - 基本流程与常见手段

    1. 首先查看依赖的上游服务是否有问题;

       - 让上游赶紧解决！
       - 是否能做缓存?
       - 超时/熔断

    2. CPU占用, 看cpu profile 火焰图, 优化占用cpu过多的逻辑;

       - 更换json库、使⽤⼆进制编码⽅式替代 JSON 编码;
       - 同物理节点通信，使⽤共享内存 IPC，直接⼲掉序列化开销;
       - MD5 算 hash 成本太⾼ -> 使⽤ cityhash，murmurhash

    3. 内存OOM

       - prometheus的rss确定是我们的进程内存占用高, 还是其他程序内存占用高但是OOM killer可能杀死的是我们的进程; 
       - goroutine数量是不是特别多? goroutine栈占用内存多少? stack
       - 查看heap profile中的inuse -> 定时类脚本需要看alloc
       - 堆对象使用过多
         - sync.Pool 对象复用
         - 为不同⼤⼩的对象提供不同⼤⼩ level 的 sync.Pool

    4. goroutine数量过多: 

       - profile 里可以看到goroutine在具体做什么; 

       - 连接数是不是太多了? 大多数连接可能都卡在read buffer上; 
       - goroutine  pool, 控制最大的goroutine数量;

    5. gc频率、gc延迟: gctrace日志( 注意会影响硬盘 )、stw时间

       - offheap: 用到了cgo  [glycerine/offheap: an off-heap hash-table in Go. Used to be called go-offheap-hashtable, but we shortened it. (github.com)](https://github.com/glycerine/offheap)
       - map中有大量指针, 会对gc扫描造成压力; ( 分代gc可以解决这个问题 )  ->   减少堆上对象的分配
         - sync.Pool 池化
         - Map ->  slice
         - 指针 -> 非指针对象
         - 多个小对象 -> 合并为一个大对象
       - 降低gc频率:  调大gogc( 触发gc的内存值 ) / 程序启动阶段make一个很大的slice, 1GB; ( 只适合内存不紧张的服务 )

  - 锁瓶颈的优化方案:

    - 缩小临界区
    - 降低锁粒度: 全局锁改对象锁、连接锁; 连接锁改请求锁、文件锁、多个文件锁; 
    - 同步改异步
    - 使用双buffer完全消除读锁阻塞
      - 全量更新:  在一个新的数据里面写, 然后使用atomic操作替换掉旧数据;
      - 增量更新:  先拷贝原数据, 更新key, 然后使用atomic操作替换; 特别注意多进程更新同一个数据不同key可能存在问题, 如 进程A开始更新key1, 同时进程B开始更新key2, 进程A更新完毕替换新数据, 进程B更新完毕替换新数据, 则最后key1是进程B读取时的老数据, 而不是进程A更新后的数据.

  - 问题例子:

    - 直接调用exec库, fd数量过多直接挂掉;
    - 长函数不要 lock + defer unlock; 一个是中间panic会导致死锁; 另一个是临界区执行时间过长整个系统吞吐下降; 缩小锁的颗粒度;
    - 日志、通信、上报等操作; 
      - 如日志写的比较大, 对一个文件fd会有锁, 影响程序速度;  
      - metrics  udp大量上报,底部也是有锁, 一般可能会卡在几千QPS上;

- Continuous Profiling

  - 在问题出现的时候尽量保留现场;
  - [mosn/holmes: self-aware Golang profile dumper (github.com)](https://github.com/mosn/holmes)

- 从fasthttp 与evio 见go语言常见 优化策略 
- 整体架构单元化: 将整个系统缩成微小的单元, 这个单元拥有整个系统所有功能, 但是它只负责对某个面提供服务
  - 例如按物理地址做划分, 该单元多少台机器只对某个市提供服务;
  - [什么是单元化原理,如何实现单元化_金融分布式架构(SOFAStack)-阿里云帮助中心 (aliyun.com)](https://help.aliyun.com/document_detail/159741.html)

​	


# xargin chapter3

## 服务拆分





## CI CD

- 开关配置
  - 功能开关;
  - A/B实验开关;
  - 灰度开关; 按城市、人群特征、手机号、随机百分比等比例开启功能;
  - 权限开关;
- 开关的生命周期








​    





