# assembly go

## todolist

### 一、 完善现有内容

1.  **完成 “delve实践” 章节**：
    *   **实践案例：** 详细编写 “验证空结构体不占用任何内存空间” 的步骤。
    *   **操作步骤：**
        1.  编写 Go 代码，定义一个空结构体 `EmptyStruct` 和一个包含该空结构体的 `TestStruct`。
        2.  使用 `dlv debug` 启动调试。
        3.  使用 `b main.main` 和 `c` 进入主函数。
        4.  使用 `n` 逐行执行，实例化结构体。
        5.  使用 `whatis` 和 `print` 查看变量类型和值。
        6.  关键是使用 `print unsafe.Sizeof(variable)` 来打印不同变量所占用的内存大小，直观地证明空结构体大小为 0。

### 二、 增加 Go 核心特性实践

1.  **调试 Goroutine：**
    *   **目标：** 理解 Go 如何调度 Goroutine。
    *   **实践：**
        1.  编写一个启动多个 Goroutine 的程序。
        2.  使用 `goroutines` (`grs`) 命令查看所有 Goroutine 的状态（如 `running`, `runnable`, `waiting`）。
        3.  使用 `goroutine <id>` (`gr <id>`) 命令切换到特定的 Goroutine 上下文。
        4.  在不同 Goroutine 中使用 `stack` (`bt`) 命令查看各自独立的调用堆栈，体会其轻量级的本质。

2.  **剖析 Channel 通信：**
    *   **目标：** 观察 Goroutine 如何通过 Channel 阻塞和交换数据。
    *   **实践：**
        1.  创建一个带缓冲和不带缓冲的 Channel 示例。
        2.  在发送方和接收方设置断点。
        3.  当程序阻塞时，使用 `goroutines` 查看 Goroutine 的等待状态（`chan receive`/`chan send`）。
        4.  使用 `print ch` 查看 Channel 的内部状态，如 `qcount` (队列中元素数量), `dataqsiz` (队列大小), `closed` (是否关闭)等。

3.  **探究 Interface 内部结构：**
    *   **目标：** 理解 Go 接口（`eface` 和 `iface`）是如何实现动态类型的。
    *   **实践：**
        1.  创建一个接口变量，并为其赋予不同类型的值。
        2.  使用 `print myInterfaceVar` 查看 `delve` 为我们格式化后的结果，它会清晰地展示接口指向的具体类型和值。
        3.  可以进一步结合 `examinemem` (`x`) 和指针操作，查看接口变量在内存中的两个字：类型信息指针和数据指针。

4.  **解析 Slice 和 Map 的内存布局：**
    *   **目标：** 理解 Slice 和 Map 的底层数据结构。
    *   **实践：**
        *   **Slice:**
            1.  创建一个 Slice。
            2.  使用 `print mySlice` 查看它的三个核心部分：指向底层数组的指针 (`array`), 长度 (`len`), 和容量 (`cap`)。
            3.  通过 `append` 操作，观察当容量不足时，底层数组指针的变化，印证其扩容机制。
        *   **Map:**
            1.  创建一个 Map 并添加一些键值对。
            2.  使用 `print myMap` 查看其内部结构（它是一个指向 `hmap` 结构体的指针）。
            3.  观察 Map 在动态增长时内部的变化。

### 三、 深入汇编与内存

1.  **汇编级调试：**
    *   **目标：** 将 Go 代码与最终生成的汇编指令对应起来。
    *   **实践：**
        1.  使用 `disassemble` (`disass`) 命令查看当前函数的汇编代码。
        2.  使用 `step-instruction` (`si`) 命令进行单条汇编指令的步进。
        3.  使用 `regs` 命令查看 CPU 寄存器的状态，特别是 `PC` (程序计数器), `SP` (栈指针), `BP` (基址指针) 的变化，理解函数调用的过程。

2.  **观察变量的栈与堆分配：**
    *   **目标：** 验证 Go 的逃逸分析。
    *   **实践：**
        1.  编写一个函数，其中一个变量分配在栈上，另一个因为被外部引用而“逃逸”到堆上。
        2.  （可选）先用 `go build -gcflags="-m"` 编译，查看编译器的逃逸分析报告。
        3.  在 `delve` 中，分别使用 `print &stackVar` 和 `print &heapVar` 获取两个变量的内存地址。
        4.  通过地址值可以明显区分栈（通常是高位地址）和堆（通常是低位地址），从而直观地看到逃逸分析的结果。

### 四、 探索 Delve 高级功能

1.  **条件断点：**
    *   **目标：** 在循环或频繁调用的函数中，只在特定条件下中断。
    *   **实践：** 在一个 for 循环中，使用 `condition <breakpoint_id> i == 5` 命令，让程序只在 `i` 等于 5 时停下来。

2.  **远程调试与 Headless 模式：**
    *   **目标：** 介绍如何在开发环境（如 VS Code）中连接到服务器或 Docker 容器中的 `delve` 实例。
    *   **实践：** 简要介绍 `dlv debug --headless --listen=:2345` 命令，并说明如何在 VS Code 的 `launch.json` 中配置 `request: "attach"` 和 `mode: "remote"` 来连接调试。

3.  **事后调试 (Post-mortem Debugging)：**
    *   **目标：** 介绍如何分析程序崩溃后产生的核心转储文件 (core dump)。
    *   **实践：** 简要介绍 `dlv core <executable_file> <core_dump_file>` 命令的用途，这是排查生产环境中复杂问题的利器。

### 五、 汇编知识学习路线

1.  **理解基本概念：**
    *   **寄存器 (Register):** 了解什么是寄存器，以及通用寄存器（如 `RAX`, `RBX`, `RCX`, `RDX`）、指针寄存器（`RSP`, `RBP`）和指令指针寄存器（`RIP`）在 x86-64 架构下的基本用途。
    *   **内存地址 (Memory Address):** 理解程序如何读取和写入内存。
    *   **指令 (Instruction):** 了解基本的汇编指令，如 `MOV` (数据移动), `ADD`/`SUB` (算术运算), `JMP` (无条件跳转), `CMP` (比较), `JNE`/`JE` (条件跳转) 等。
    *   **栈 (Stack):** 理解栈的 LIFO (后进先出) 原理，以及 `PUSH` 和 `POP` 指令如何操作栈。理解函数调用时，栈是如何用于传递参数、存储返回地址和保存局部变量的。

2.  **学习 Go 汇编 (Plan 9 汇编)：**
    *   **语法差异：** 了解 Go 汇编与常见的 AT&T 或 Intel 汇编在语法上的不同，例如源操作数和目标操作数的位置是相反的 (`MOVQ src, dst`)。
    *   **伪寄存器：** 学习 Go 汇编中的四个重要伪寄存器：`FP` (帧指针，用于访问函数参数和局部变量), `PC` (程序计数器，但主要用于定义符号), `SB` (静态基指针，用于引用全局符号), `SP` (栈指针，用于引用栈上变量)。
    *   **函数定义：** 学习 `TEXT` 指令如何定义一个函数，包括函数名、标志和帧大小。例如 `TEXT ·myFunction(SB), NOSPLIT, $0-16`。

3.  **实践与分析：**
    *   **简单函数分析：** 编写一个简单的 Go 函数（如 `func add(a, b int) int { return a + b }`）。
    *   **生成汇编代码：** 使用 `go build -gcflags="-S"` 或 `go tool compile -S` 命令来查看编译器生成的汇编代码。
    *   **代码对读：** 逐行将 Go 代码与汇编指令对应起来，理解参数是如何通过栈传递的，计算是如何在寄存器中完成的，以及返回值是如何被传回的。
    *   **使用 Delve 验证：** 在 `delve` 中使用 `disassemble` 命令查看汇编，并用 `step-instruction` 和 `regs` 来单步跟踪执行流程和寄存器变化，加深理解。

作为求助欲望旺盛的Go语言开发者, 难免对Go语言底层实现逻辑兴致盎然. 尤其Go底层自举实现, 面对世界顶级的Go语言大师提供的黄金屋, 相较于其他语言, 我们更容易享受到这款饕餮盛宴. 

可能想要了解底层的一些关键实现, 懂Go语言的基本语法就可以窥见一二. 但是想要将其流程如探囊取物般完整走下去, 可能我们还需要一些操作系统和汇编的基础知识. 下面我将以一名小白的视角, 试图去了解Go语言的运行全貌.

“工欲善其事, 必先利其器”, 对于小白来说, 可能挑选一个好用方便的工具, 比其他方法更容易提高效率. 于是我去问了一下chatGPT, 它给我推荐了两款工具: `delve`和`gdb`. 那我们先从delve开始, 揭开语法背后的秘密~

## delve

### 安装与启动

首先安装docker;之后找一个空目录,新建文件`Dockerfile`, 将下面内容粘贴到文件中:
```dockerfile
FROM centos
RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-Linux-* \
&& sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-Linux-* \
&& yum install golang -y \
&& yum install dlv -y \
&& yum install binutils -y \
&& yum install vim -y \
&& yum install gdb -y
```

在dockerfile的目录下执行`docker build -t dlv .`命令, 开始构建镜像.

> 为什么要使用docker环境 ?  因为不同系统环境规范不同, go在不同的环境下有不同的编译方案; 大部分情况下我们的二进制程序会运行在linux环境下,因此建议统一使用linux环境下的编译结果进行调试.
> 如果你的环境本来就是linux系统, 可以参考这篇官方文档进行操作: [如何使用dlv调试go语言程序](https://github.com/go-delve/delve/blob/master/README.md)

进入需要调试的目录，假设需要调试的文件为test.go, 执行命令`docker run --rm -it -v $(pwd):/app dlv dlv debug /app/test.go`开始dlv调试。下面是命令详解:
- `--rm`选项指示docker在容器退出后自动删除容器;

- `-it`选项指示docker在启动容器时分配一个伪tty;

- `-v $(pwd):/app` 选项将当前目录映射到容器中的/app目录;

- 第一个dlv是上面`docker build -t dlv .`中构建的dlv docker image; 

- 第二个dlv代表在容器中执行dlv命令，命令为`dlv debug /app/test.go` 调试test.go程序

### 开始尝试调试程序

举个例子:

```go
func main() {
    var nums1 []interface{}
	nums2 := []int{1, 2, 3}
	num3 := append(nums1, nums2)
	fmt.Println(len(num3))
}
```

首先执行命令`dlv debug .`进入调试模式，键入`list`查看当前程序的执行位置：

```shell
     3:	// license that can be found in the LICENSE file.
     4:
     5:	#include "textflag.h"
     6:
     7:	TEXT _rt0_amd64_linux(SB),NOSPLIT,$-8
=>   8:		JMP	_rt0_amd64(SB)
     9:
    10:	TEXT _rt0_amd64_linux_lib(SB),NOSPLIT,$0
    11:		JMP	_rt0_amd64_lib(SB)
```

当前程序似乎还在程序入口的位置，并未开始真正执行用户代码。我们还是先从用户代码看起吧, 可以在main函数开头打断点, 使用命令:`break main.main`或者`b main.main`，然后输入命令`continue`或者`c`使程序运行到这里:

```shell
(dlv) b main.main
Breakpoint 1 set at 0x4b66b8 for main.main() .app/test.go:5
(dlv) c
> main.main() .app/test.go:5 (hits goroutine(1):1 total:1) (PC: 0x4b66b8)
     1:	package main
     2:
     3:	import "fmt"
     4:
=>   5:	func main() {
     6:		var nums1 []interface{}
     7:		nums2 := []int{1, 2, 3}
     8:		num3 := append(nums1, nums2)
     9:		fmt.Println(len(num3))
    10:	}
```

假设我们想查看`nums1`变量的具体内容，可以先给第二行打断点：`b main.main:2`，然后用`c`执行到第二行。可以使用一系列命令来查看`nums1`变量的具体内容:

```shell
> main.main() .app/test.go:7 (hits goroutine(1):1 total:1) (PC: 0x4b66e6)
     2:
     3:	import "fmt"
     4:
     5:	func main() {
     6:		var nums1 []interface{}
=>   7:		nums2 := []int{1, 2, 3}
     8:		num3 := append(nums1, nums2)
     9:		fmt.Println(len(num3))
    10:	}
(dlv) locals
nums1 = []interface {} len: 0, cap: 0, nil
(dlv) whatis nums1
[]interface {}
(dlv) print nums1
[]interface {} len: 0, cap: 0, nil
(dlv) print &nums1
(*[]interface {})(0xc000191f30)
```

- locals命令可以看到当前函数堆栈的信息, 上面可以看到 `num1`变量是`[]interface{}`类型, len=0, cap=0, 实际内容是nil.
- `whatis nums1`获取nums1变量的类型
- `print nums1`获取该变量内部数据
- `print &nums1`来获取该变量的内存地址

调试完毕后可以使用`exit`退出，或者使用`restart`重开程序。

我们跟随上述例子, 第一次使用delve工具来调试了一个Go语言程序, 下面我们将列出delve的一些常用的命令, 并结合一些例子, 具体讲解如何使用delve, 查看Go语言底层的运行逻辑.

### Delve命令

####  启动调试

- `dlv debug [package]`编译当前目录下的 "main "软件包，并开始调试;
    - 要向程序传递flag，请使用`--`分隔它们： `dlv debug github.com/me/foo/cmd/foo -- -arg1 `

- `dlv test [package]`可以在单元测试的上下文中开始新的调试会话;

- `dlv exec [binary]`执行预编译二进制文件并开始调试会话。

    - 该命令将导致 Delve 执行二进制文件，并立即附加到该二进制文件以开始新的调试会话。请注意，如果二进制文件在编译时未禁用优化，则可能难以正确调试。请考虑在 Go 1.10 或更高版本中使用 -gcflags="all=-N -l "编译调试二进制文件，在 Go 早期版本中使用 -gcflags="-N -l "编译调试二进制文件。

- `dlv attach <pid>`附加到已运行的进程并开始调试

- `dlv core <exe> <core>` 检查核心转储coredump（仅支持 linux 和 windows 的core dump）。
    - core 命令将打开指定的 core 文件和相关的可执行文件，让你检查获取 core dump 时的进程状态。目前支持 linux/amd64 和 linux/arm64 core 文件、windows/amd64 minidump 以及 Delve 的 "dump "命令生成的 core 文件。
    
- `dlv replay <rr trace>` 重放[Mozilla rr](https://github.com/mozilla/rr)的trace

- `dlv trace [package]` 跟踪程序执行。
    - 跟踪子命令将在与所提供的正则表达式匹配的每个函数上设置跟踪点，并在跟踪点被击中时输出信息。如果您不想开始整个调试会话，而只是想知道进程正在执行哪些函数，那么这条命令就非常有用。跟踪子命令的输出会打印到 stderr，因此如果只想查看跟踪操作的输出，可以重定向 stdout。

- 调试参数headless `dlv --headless <command>` 使用debug、test、exec、attach、core和replay命令时可以附加此参数，会启动一个backend server提供给前端。例如可以在`VS Code`上安装dlv工具，使用该命令放开后vs code即可连接到dlv，然后使用VS Code来调试。

    - 也可使用`dlv connect <addr>`来连接一个放开的headless调试程序

- `help`
调试模式下使用该命令会打印出所有调试模式下可使用的命令

#### 调试命令

##### 调试命令的位置参数
dlv命令如果存在程序位置参数，一般遵循以下规则：

- *<address> 指定内存地址的位置。可指定为十进制、十六进制或八进制数

- <filename>:<line> 指定文件名中的行。只要表达式不含糊，文件名可以是文件的部分路径，甚至只是文件的基本名称。

- <line> 指定当前文件的行数

- +<offset> 指定当前行之后的行偏移量

- -<offset> 指定当前行之前的行偏移行数

- <function>[:<line>] 指定函数内部的行。函数的完整语法是 <package>.(<接收器类型>).<函数名>，但唯一需要的元素是函数名，只要表达式不含糊，其他元素都可以省略。在初始函数（例如：main.init）上设置断点时，应使用 <filename>:<line> 语法在正确的位置断开正确的初始函数。

- /<regex>/ 指定与 regex 匹配的所有函数的位置


##### 运行程序
```
call ------------------------ 恢复进程，注入函数调用（实验！！）
continue (alias: c) --------- 运行到断点或程序终止
next (alias: n) ------------- Step over to next source line.
rebuild --------------------- 重新生成目标可执行文件并重新启动它。如果可执行文件不是通过dlv生成的，则此操作不起作用
restart (alias: r) ---------- Restart process.
step (alias: s) ------------- 分步执行程序
step-instruction (alias: si)  分步调试汇编指令
stepout (alias: so) --------- Step out of the current function.
```

##### 操作断点
```
break (alias: b) ------- 设置断点。
breakpoints (alias: bp)  打印出活动断点的信息
clear ------------------ Deletes breakpoint.
clearall --------------- Deletes multiple breakpoints.
condition (alias: cond)  设置断点条件
on --------------------- 在遇到断点时执行命令
trace (alias: t) ------- 设置跟踪点
```


##### 查看程序变量和内存
```
args ----------------- 打印函数参数
display -------------- 每次程序停止时打印表达式的值
examinemem (alias: x)  检查内存
locals --------------- 打印局部变量
print (alias: p) ----- 对表达式求值
regs ----------------- 打印CPU寄存器的内容
set ------------------ 更改变量的值
vars ----------------- 打印包变量
whatis --------------- 打印表达式的类型
```

##### 列出并切换线程和程序
```
goroutine (alias: gr) -- 显示或更改当前goroutine
goroutines (alias: grs)  列出程序goroutines
thread (alias: tr) ----- 切换到指定的线程
threads ---------------- 打印出每个跟踪线程的信息
```

##### 查看调用堆栈和选择帧
```
deferred --------- 在延迟调用的上下文中执行命令
down ------------- 向下移动当前帧
frame ------------ 设置当前帧，或在其他帧上执行命令
stack (alias: bt)  打印堆栈跟踪
up --------------- 向上移动当前帧
```

##### 其他
```
config --------------------- Changes configuration parameters.
disassemble (alias: disass)  反汇编程序.
edit (alias: ed) ----------- 打开$DELVE_EDITOR或$EDITOR中的位置
exit (alias: quit | q) ----- Exit the debugger.
funcs ---------------------- 打印函数列表
help (alias: h) ------------ Prints the help message.
libraries ------------------ 列出加载的动态库
list (alias: ls | l) ------- 显示源代码
source --------------------- 执行包含dlv命令列表的文件
sources -------------------- 打印源文件列表
types ---------------------- 打印类型列表
```



### delve实践

#### 验证空结构体不占用任何内存空间

```