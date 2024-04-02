# go语言基础与原理解析

## 目录index
- [basic - go语言基础入门](./01basic/)
- [types - go语言基础类型原理分析](./types/)
- [runtime - go语言运行时逻辑研究](./runtime/)
- [sync - 并发编程](./sync/)
- [testing - 单元测试最佳实践](./testing/)
- [net - go语言中的网络](./net/)
- [read - 相关读物](./read/)
- [question - 存在疑问的短代码集合](./question/)

## go语言基础
最基本的入门教程可以参考：[菜鸟编程](https://www.runoob.com/go/go-tutorial.html)

在日常使用中可以参考官方文档：[Packages](https://pkg.go.dev/)或者中文翻译的[标准库文档](https://studygolang.com/pkgdoc)

开始动手编程了，建议先读一遍[EffectiveGo](https://go.dev/doc/effective_go)或者[我的翻译版本](./read/effective-go.md)，以及[The Go Programming Language Specification](https://go.dev/ref/spec)或者[翻译版本go语言编程规范](./read/go_spec.md)

## go语言源码分析
### go源码目录
目录中分为api、doc、include、lib、src、misc、test这7个原始目录，编译后还会生成bin、pkg，2个目录
- api：对应工具"go tool api"相应源码 src/cmd/api;
- doc：文档目录，如果要编译，或者了解一些版本信息，可以看一下这个目录中html内容;
- include：依赖C头文件, go1.5实现自举以后, 不需要这个目录了;
- lib：依赖库文件;
- src：源代码;
- misc：一些杂项脚本都在这里面;
- test：测试;
- bin：go、godoc、gofmt等等;
- pkg：生成对应系统动态连接库，以及对应系统的工具命令都在该目录，如cgo、api、yacc等等;

### 源码地址
- [在github上](https://github.com/golang/go/tree/go1.14.15/src) ,通过tag找到对应的版本;
- [golang.org](https://golang.org/doc/faq#history) golang文档;
- [review](https://go-review.googlesource.com/c/go/+/36476)

### 源码调试
GDB和Dlv都可用于调试Go语言程序。它们都支持设置断点、查看变量值、单步执行程序等功能。区别是dlv是go语言的专用调试器，比较好简单上手，而GDB则是通用调试工具。关于使用dlv debug go程序，可以[参考这一篇](./01basic/源码调试.md)。

