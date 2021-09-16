# go包demo与源码解析

## index
- [Context包源码分析,上下文](./context/README.md)
- [golang的基本类型源码分析](./types/README.md)
- [垃圾回收gc](./runtime/gc.md)
- [channel源码分析](./runtime/channel.md)
- [golang的调度模型，GMP](./runtime/GMP.md)
- [命令行以及flag的使用](./flag/README.md)
- [sync同步机制](./flag/README.md)
- [go语言时间包Time的用法](./time/README.md)
- [net包解析](./net/README.md)
- [反射的用法](./reflect/README.md)
- [error错误处理](./error/.)
- [单元测试与基准测试](./testing/README.md)

## 关于src的一些tips
### go目录
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
- [如何优雅的使用GDB调试Go](https://mp.weixin.qq.com/s/xfDydcpRCmX1dR5FybI0Rw)

### 关于源码中出现的参数`raceenabled`与`msanenabled`
- `raceenabled`参数代表是否启用数据竞争检测; 在`go build`或者`go run`中加入`-race`参数就代表该选项为`true`
- `msanenabled`参数: go1.6新增的参数,类似上面的`-race`,这个参数为`-msan`,并且仅在 linux/amd64上可用;作用是将调用插入到C/C++内存清理程序;这对于测试包含可疑 C 或 C++ 代码的程序很有用。在使用新的指针规则测试 cgo 代码时，您可能想尝试一下.

## 位移符号
y << x 代表将y的二进制左移x位, 比较常用的是以2为底的位移操作, 如:2的100次方`2 << 100`
```go
const (
        j = 1 << iota	// 1=1 => 左移0位, 仍然是1
        k = 3 << iota	// 3=11 => 左移1位, 110, 4+2=6
        l				// 3=11 => 左移2位, 1100, 8+4=12
        n				// 3=11 => 左移3位, 11000, 16+8=24
    )
    println(j, k, l, n)	// 1, 6, 12, 24
```

## 编码规范
1. 格式化规范：使用gofmt 或者 goimport；
2. 一行最长不要超过80个字符；
3. go vet 工具帮我们静态分析源码各种问题；
4. import规范：
   建议采用如下格式：
   ```go
   import (
       "encoding/json"
       "strings"

       "myproject/models"
       "myproject/controller"
       
       "github.com/astaxie/beego"
       "github.com/go-sql-driver/mysql"
   )
   ```
   有顺序地引入包， 不同类型使用换行符分离，第一种是标准库，第二种是项目包，第三种是第三方包。*并且不要使用相对路径*
5. 变量申明：使用驼峰命名，多个变量申明放在一起；函数外部的全局变量申明必须使用var，不然容易踩到作用域的坑；
6. 闭包循环的问题：
    ```go
    for i := 0; i <= 5; i++ {
        // 错误的写法，打印出来的都是6
        go func() {
            fmt.Println(i)
        }()

        // 正确的写法
        go func(i int) {
            fmt.Println(i)
        }(i)
    }
    time.Sleep(time.Second)
    ```
7. 不要随意使用panic，而是使用`log.Fatal`来记录错误
8. 使用`{}`来更好地区分代码块(go使用`{}`来划分变量的作用范围)



## reference


## todo
1. flag
3. strconv
4. strings
5. database
6. net
8. os
9. path
11. sort
12. io
13. bufio
