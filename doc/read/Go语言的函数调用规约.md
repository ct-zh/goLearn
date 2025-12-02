
## 什么是函数调用规约（Calling Convention）

简单来说， 函数调用规约是**“调用者 (Caller)” 和 “被调用者 (Callee)” 之间达成的一种 契约 或 协议**。

它规定了函数调用时，数据如何在计算机底层（寄存器、栈）之间进行传递。就像两个人打电话，需要约定好谁先说话，用什么语言，怎么结束一样，CPU 执行函数调用也需要这一套规则。

具体来说，它规定了几个核心问题：**参数怎么传？返回值怎么传？谁来清理栈？以及寄存器的保护责任。**

### 其他语言有函数调用规约吗？

当然有！函数调用规约不是Go语言独有的， 所有编译型语言都需要定义调用规约。 而且不同的操作系统、架构和语言，可能会使用不同的规约。

下面是几个常见语言的调用规约：
- C/C++ (x86_64 / AMD64)，C 语言是系统编程的基石，它的规约通常就是操作系统的标准规约。
- C (x86 / 32-bit)，32 位时代，规约非常混乱
- Java (JVM) JVM 是一个 基于栈的虚拟机 (Stack-based VM)。它的指令（如 iadd , invokevirtual ）都是在操作数栈上工作的，没有“寄存器”的概念。
	- JIT 编译后 : 当 HotSpot 虚拟机将 Java 字节码编译成本地机器码时，它会遵循宿主机操作系统的调用规约（如 System V AMD64），或者使用自己内部优化的规约来执行本地代码。
- Python：它是解释型语言，它的“调用规约”是在 Python 虚拟机的 C 语言实现层面的（PyObject 指针的传递）。本身不直接操作 CPU 寄存器或栈，而是操作 Python 的栈帧对象。

C语言的规约是非常经典的规约， 后续分析调用规约时，我们将一起分析C语言的规约，与Go进行比较。

那么问题来了，我们怎么分析调用规约呢？它是一个文档吗？还是一段代码？我们首先要知道这个规约定义在哪里，才能分析它的具体内容。

### 函数调用规约的出处

#### C语言的调用规约

C语言的函数调用规约并不是写在某一行代码里的，而是定义在编译器、操作系统标准文档以及 CPU 架构手册中的。它们是 约定 ，必须由编译器（生成代码的一方）和操作系统/库（运行代码的一方）共同遵守。

这些规约通常由操作系统厂商或 CPU 架构设计者发布，是该平台上的“法律”。

- System V AMD64 ABI (Linux, macOS, BSD, Solaris) :
	- 出处 : [System V Application Binary Interface AMD64 Architecture Processor Supplement](https://refspecs.linuxbase.org/elf/x86_64-abi-0.99.pdf)
	- 这是一份非常详细的 PDF 文档，由 Intel、AMD 和操作系统厂商共同维护。它规定了寄存器用法、栈布局、异常处理等所有细节。GCC 和 Clang 编译器严格遵循此文档。

- Windows x64 ABI :
	- 出处 : [Microsoft Learn 文档 - x64 calling convention](https://learn.microsoft.com/en-us/cpp/build/x64-calling-convention?view=msvc-170)
	- 微软定义了自己的标准，所有在 Windows 上运行的 x64 程序（VC++, MinGW 等）都必须遵守。

- ARM64 (AArch64) AAPCS64 :
	- 出处 : [ARM 官方文档 - Procedure Call Standard for the Arm 64-bit Architecture](https://github.com/ARM-software/abi-aa/blob/main/aapcs64/aapcs64.rst)
	- 规定了 ARM 架构下的 R0-R31 寄存器用法。


#### Go语言的函数调用规约

Go语言因为是一个**自举**的语言，拥有自己的编译器（gc）和运行时，因此它可以自己定义一套规则。Go 语言的规约定义在 Go 源码 和 设计文档 中。

- Go ABI 设计文档 :
	- 这是最权威的定义： [Go Register-based Calling Convention Proposal](https://github.com/golang/proposal/blob/master/design/40724-register-calling.md)
	- 这份设计文档详细解释了为什么要切换到寄存器 ABI，以及具体的寄存器映射规则。

- Go 编译器源码 :
	- 规约最终是落实到代码里的。在 Go 编译器源码中，我们可以找到具体的实现。
	- 文件位置 : src/cmd/compile/internal/abi/abiutils.go (定义了寄存器分配逻辑)
	- 文件位置 : src/cmd/compile/internal/ssagen/ssa.go (将 SSA 中间代码转换为具体的机器码，处理参数传递)

Go1.17 之前，



### 参数怎么传？


### 返回值怎么传？


### 谁来清理栈？



### 寄存器的保护责任



