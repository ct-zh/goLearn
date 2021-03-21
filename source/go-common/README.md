# bilibili go-common源码分析

## 技术栈
0. go vendor
1. 使用go-convery进行单元测试
2. 使用Bazel构建;(BUILD文件) (rider是什么鬼)
3. `github.com/BurntSushi/toml` toml配置文件
4. `protoc-gen-gogo`    proto文件自动生成rpc代码
5. zookeeper

> https://github.com/go-kratos/kratos/blob/main/README_zh.md

## library文件夹
### cache
- memcache
- redis: redigo

### conf
- dsn
- env
- paladin: 配置中心客户端

### container
- pool: 通用连接池
- queue: 队列? (代码很短)

### database
- bfs: 分布式文件系统,毛剑自己开源了: https://github.com/Terry-Mao/bfs  [bfs:支撑Bilibili的小文件存储系统](https://www.toutiao.com/i6272104949560115714/)
- elastic: 
- hbase.v2: (Hadoop)在hbase基础上加入了链路追踪和统计 https://github.com/tsuna/gohbase
- orm: github.com/jinzhu/gorm
- sql: MySQL数据库驱动，进行封装加入了链路追踪和统计。
- tidb: TiDB数据库驱动 对mysql驱动进行封装

### ecode
error code管理

### exp
- feature go-common 里的 Feature 管理工具。用于灰度测试一些基础库的功能。

### log
基于uber的zap封装的日志框架

### naming
服务发现、服务注册相关的SDK集合: 初期为zk,后面换成了自研的discovery
- livezk: zookeeper客户端 github.com/samuel/go-zookeeper/zk
- discovery: discovery的客户端SDK，包括了服务发现和服务注册功能: https://github.com/bilibili/discovery

### net
- http: blademaster来自 bilibili 主站技术部的 http 框架，融合主站技术部的核心科技，带来如飞一般的体验。(基于gin)
- ip
- metadata: 用于储存各种元信息
- netutil/breaker: Hystrix熔断器
- rpc: context,liverpc,interceptor,warden
- trace: 追踪

#### warden
来自 bilibili 主站技术部的 RPC 框架，融合主站技术部的核心科技，带来如飞一般的体验。
(基于grpc封装?)

### os
- signal

### queue
- databus: 消息队列 基于kafka封装的databus

### rate
- limit/bench/stress 框架压测专用
- tcp vegas

### stat
stat 统计库，包含Counter、Summary等
- prometheus
- sys: 获取Linux平台下的系统信息，包括cpu主频、cpu使用率等

### sync
- errgroup: 提供带recover的errgroup，err中包含详细堆栈信息
- errgroupv2:
- pipeline: 提供内存批量聚合工具 内部区分压测流量 
- pipeline/fanout: 代替library/cache

### syscall

### text
- translate: https://github.com/BYVoid/OpenCC/

### time
Kratos 的时间模块，主要用于mysql时间戳转换、配置文件读取并转换、Context超时时间比较

### xstr


## app文件夹
### admin
运营管理服务
- bbq
- ep
- live
- main
- openplatform


### common
- openplatform: 开放平台公共模块
- live


### infra
各种服务
- canal: db
- config: 配置中心服务端，提供配置文件的管理和拉取
- databus: databus是一个通过使用redis协议来简化kafka的消费方/生产方的一个中间件
- discovery: 服务注册发现
- notify: 消息中间件，消息注册监听推送事件通知

### interface
Gateway&Interface对外网关服务

### job
后台异步服务job

### service
rpc service

### tool
自动生成代码
- bgr: golang syntax 检查规则解释器
- bmproto: 根据protobuf文件，生成grpc和blademaster框架http代码及文档
- cache: 缓存代码生成
- ci: 
- creator: 生成各种代码
- gdoc: 自动化doc
- gengo: k8s代码生成库? https://github.com/kubernetes/gengo
- gorpc: 根据service 方法生成rpc client 以及rpc model层 代码
- grpc-http-proxy: http调试工具
- kratos
- liverpc
- mkprow
- owner
- protoc-gen-bm: app/tool/bmproto的旧版本
- saga: 1.提供大仓库pkg依赖关系DAG 2.提供gitlab MR自动构建、测试、覆盖率、代码静态检查
- warden: warden proto 自动生成工具



# 2.7 内容
## 某个服务的目录结构, 
按照`app/admin/main/vip/`:
    1. cmd: 入口文件,命令行,用于启动http服务
    2. conf: type struct 配置文件, 用到了`toml`
    3. dao: 数据层 redis, memcache访问方法, 还有一些RPC调用也放在这里面
    4. http: http服务,还有对应的接口  主要是提供协议转换, 聚合. 逻辑还是再service层做
    5. model: 模型层
    6. service: 业务层 对于后端服务来说, 该目录提供服务的实现, 对于http服务, 该目录提供http服务的实现

    流程大概是: 入口文件cmd,启动http,注册服务service到http.svc,注册路由之类的工作;在注册路由里面写入了http层的一些方法;当请求过来,http根据路由调用对应的方法,这些方法通过http.svc调用service的方法拿到具体的业务数据再返回给前端.

    所有的服务均遵守该目录结构. model层放VO, DO等, dao层用于数据层封装, 隔离本服务的领域逻辑与外部数据. http层提供协议转换. service实现具体逻辑.
    比较像Java开发的模式, 可能在公司很多人不是很喜欢这样复杂的目录, 喜欢什么都放在一个目录下.
    不过这样的分目录是一种比较好的实践. 各层分的清清楚楚, 一个服务从1个接口到10个接口, 都比较清晰. 对于服务改动来说,也比较好聚焦于某一层.

或者按照`app/service/main/vip`:

    1. api: 
    2. cmd: 入口文件
    3. conf: 配置文件
    4. dao: 数据层
    5. http: 
    6. model: 模型
    7. rpc: rpc服务
    8. server: grpc服务
    9. service
    10. verify

