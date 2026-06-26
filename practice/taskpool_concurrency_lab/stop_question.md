# stop question

定义三个逻辑：
    - worker: 执行task的协程；
    - stop: 代表外部调用Stop方法，停止继续提交任务，并尽量让worker执行完剩余的tasks；
    - submit：提交task的方法

定义两个channel：
    - tasks：类型是buf channel；submit是生产者，worker是消费者；
    - done：结束标记，在stop处进行close；submit不再提交数据；


## 版本一，不关闭tasks，保证worker全部执行完毕

- submit
    - 一个select
        - case 1: 判断 done close
        - case 2: 往tasks里写数据
        - case 3: ctx超时

- stop
    - 协程1
        - close done
        - waitgroup 等待worker结束
        - close 本地chan
    - 主逻辑：一个select
        - case1: 监听本地chan是否close
        - case2: ctx超时

- worker
    - select
        - case 1: 判断 done close
            - 

