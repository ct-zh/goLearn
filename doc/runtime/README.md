# go调度

## source
- [gc](./gc.md)
- [G、M、P](./GMP.md)
- [Channel](./channel.md)

## Q
- [创建channel](./channel/basic/main.go)
- [select配合channel](./channel/select/select.go)
- [如何优雅关闭channel、channel的关闭原则](./channel/closeChan/closeChan.go)
- [使用go tool trace分析调度流程](./trace/trace1/main.go)
- [go routine的超时问题与处理方法](./g/timeout/timeout_test.go)


## channel
[源码分析见](channel.md)

### channel经验谈
|操作|空值|非空已关闭|非空未关闭|
|---|---|---|---|
|关闭|panic|panic|成功关闭|
|发送|永久阻塞|panic|阻塞或者成功发送|
|接收|永久阻塞|永不阻塞(不停获得空值)|阻塞或者成功接收|

1. 往满的buffer channel里写数据： 会阻塞在写的位置；

2. 读空的buffer channel： 会阻塞;
3. close掉还有数据的buffer channel：能把写进去的数据读出来，读完之后再读的就是chan类型的空值；写则会panic;
4. 读空的buffer channel，其他协程把buffer channel给close了： 会无限读chan类型的空值；
5. 根据2、3、4, 建议不要轻易close buffer channel，最佳实践是增加done channel来通知子进程可以结束了；
6. 也可以使用`case val, notClose := <-ch`的方式，判断notClose是否为false;

7. channel建议配合select使用`time.After`或者`time.Ticker`来设置超时时间;


