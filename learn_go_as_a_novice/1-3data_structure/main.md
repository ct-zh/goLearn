# go内置的数据结构

## 内置数据结构一览
大致先看一个go语言的内置数据结构，列表如下，参考go语言版本为1.13：

| package   | data strcuture                                         |
| :-------- | :------------------------------------------------------ |
| runtime   | channel, timer, semaphore, map, iface, eface, slice, string |
| sync      | mutex, cond, pool, once, map, waitgroup                |
| container | heap, list, ring                                       |
| netpoll   | netpoll related                                        |
| memory    | allocation related, gc related                         |
| os        | os related                                             |
| context   | context                                                |


我们常接触的一般是runtime里面相关的数据结构，其中semaphore是信号量，iface与eface是interface底层相关的数据结构。


## channel


