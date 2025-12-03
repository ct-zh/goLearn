
## 什么是函数调用规约（Calling Convention）

简单来说， 函数调用规约是**“调用者 (Caller)” 和 “被调用者 (Callee)” 之间达成的一种 契约 或 协议**。

它规定了函数调用时，数据如何在计算机底层（寄存器、栈）之间进行传递。就像两个人打电话，需要约定好谁先说话，用什么语言，怎么结束一样，CPU 执行函数调用也需要这一套规则。

具体来说，它规定了几个核心问题：**参数怎么传？返回值怎么传？谁来清理栈？以及寄存器的保护责任。**

语言往往需要对比才能体现它的特色。在正式开始讨论这几个核心问题之前，我们需要讨论一个话题，其他语言有函数调用规约吗？

### 其他语言有函数调用规约吗？

当然有！函数调用规约不是Go语言独有的， 所有编译型语言都需要定义调用规约。 而且不同的操作系统、架构和语言，可能会使用不同的规约。

下面是几个常见语言的调用规约：
- C/C++ (x86_64 / AMD64)，C 语言是系统编程的基石，它的规约通常就是操作系统的标准规约。
- C (x86 / 32-bit)，32 位时代，规约非常混乱
- Java (JVM) JVM 是一个 基于栈的虚拟机 (Stack-based VM)。它的指令（如 iadd , invokevirtual ）都是在操作数栈上工作的，没有“寄存器”的概念。
	- JIT 编译后 : 当 HotSpot 虚拟机将 Java 字节码编译成本地机器码时，它会遵循宿主机操作系统的调用规约（如 System V AMD64），或者使用自己内部优化的规约来执行本地代码。
- Python：它是解释型语言，它的“调用规约”是在 Python 虚拟机的 C 语言实现层面的（PyObject 指针的传递）。本身不直接操作 CPU 寄存器或栈，而是操作 Python 的栈帧对象。

C语言的函数调用规约是非常经典的规约， 后续分析调用规约时，我们将一起分析C语言的规约，与Go进行比较。

那么问题来了，我们怎么分析各个语言的函数调用规约呢？它是一个文档吗？还是一段代码？我们首先要知道这个规约定义在哪里，才能分析它的具体内容。

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

C语言的函数调用规约都是写在文档中，那么go语言的函数调用规约也是写在文档中吗？

#### Go语言的函数调用规约

Go语言因为是一个**自举**的语言，拥有自己的编译器（gc）和运行时，因此它可以自己定义一套规则。Go 语言的规约定义在 Go 源码 和 设计文档 中。

- Go ABI 设计文档 :
	- 这是最权威的定义： [Go Register-based Calling Convention Proposal](https://github.com/golang/proposal/blob/master/design/40724-register-calling.md)
	- 这份设计文档详细解释了为什么要切换到寄存器 ABI，以及具体的寄存器映射规则。

- Go 编译器源码 :
	- 作为自举的语言，go语言的函数调用规约最终是可以落实到代码里的。在 Go 编译器源码中，我们可以找到具体的实现。
	- 文件位置 : src/cmd/compile/internal/abi/abiutils.go (定义了寄存器分配逻辑)
	- 文件位置 : src/cmd/compile/internal/ssagen/ssa.go (将 SSA 中间代码转换为具体的机器码，处理参数传递)

- 历史背景
	- go语言其实维护了两套ABI（Application Binary Interface）。
	- 1.17版本以前，go语言为了跨平台性与设计简洁，使用的是基于栈（Stack-based ABI）的函数调用规约。
	- 1.17版本之后，go语言改为了基于寄存器（Register-based ABI）的函数调用规约。
	- ABI0 (Legacy, Stack-based) :
		- 所有 参数和返回值都通过 栈 传递。
		- 这是 Go 1.17 之前的标准。
		- 目前主要用于汇编实现的函数（ .s 文件），为了保持跨平台的兼容性和简单性。
	- ABIInternal (Modern, Register-based) :
		- 优先使用 寄存器 传递参数和返回值，不够用时才退化为栈。
		- 这是 Go 编译器内部使用的规约，也是现在 Go 代码编译后的默认行为。
		- 它不稳定，可能会随 Go 版本变化，但性能更高。

我们主要研究1.17+的标准，也就是使用寄存器传递参数与返回值。

前置工作已经准备好了，我们开始正式讨论函数调用规约的那几个核心问题：**参数怎么传？返回值怎么传？谁来清理栈？以及寄存器的保护责任。**

## 函数调用规约之 参数怎么传？

go1.17+为基于寄存器的函数调用规约，分为两种：

**整数/指针参数 (Integer Arguments)**

使用一组固定的整数寄存器来传递整数、指针、布尔值等。对于 AMD64 (x86-64) 架构，顺序如下：
- RAX (Arg 0)
- RBX (Arg 1)
- RCX (Arg 2)
- RDI (Arg 3)
- RSI (Arg 4)
- R8 (Arg 5)
- R9 (Arg 6)
- R10 (Arg 7)
- R11 (Arg 8)

go语言使用的寄存器与C语言不同，C语言在 Linux 下的 System V AMD64 ABI通常使用：RDI, RSI, RDX, RCX, R8, R9

**浮点数参数 (Floating Point Arguments)**
浮点数使用 XMM 寄存器传递：X0 ~ X14

**栈溢出 (Stack Spill)**
如果参数过多，超过了可用寄存器的数量，剩余的参数会依然通过 栈 传递。




## 函数调用规约之 返回值怎么传？

在 Go 1.17+ 的寄存器 ABI 中，返回值的传递方式与参数非常相似：

1.  **使用相同的寄存器序列**：
    返回值也使用 `RAX`, `RBX`, `RCX` ... 等整数寄存器和 `X0` ... 等浮点寄存器。
    - 第一个返回值存放在 `RAX`。
    - 第二个返回值存放在 `RBX`。
    - 以此类推。

2.  **多返回值支持**：
    Go 语言的一个显著特性是支持多返回值。在底层，这仅仅意味着使用更多的寄存器。例如，如果一个函数返回 `(int, error)`，那么 `int` 值可能在 `RAX` 中，而 `error` (本质是一个 interface，包含两个指针) 可能占用 `RBX` 和 `RCX`。

3.  **栈溢出**：
    同样，如果返回值过多，装不下所有寄存器，剩余的返回值会通过栈内存传递。这部分空间也是由**调用者 (Caller)** 预先分配好的。

这与 C 语言有很大不同。C 语言通常只通过 `RAX` 返回一个值。如果需要返回结构体或多个值，通常需要隐式地传递一个指针，或者由调用者分配内存。Go 的寄存器规约使得多返回值非常高效。






## 函数调用规约之 谁来清理栈？


答案是：**调用者 (Caller)**。

在 Go 语言的函数调用中：
1.  **Caller 分配空间**：在调用子函数之前，Caller 会调整栈指针（SP），为子函数的参数和返回值预留空间（如果寄存器放不下的话）。
2.  **Caller 清理空间**：函数调用结束后，Caller 负责回收这部分栈空间（通常是通过恢复 SP 或者仅仅是逻辑上回收）。

这种方式被称为 **Caller-Clean**。与之相对的是 Callee-Clean（如 Windows 的 stdcall），由被调用函数负责清理堆栈。Go 选择 Caller-Clean 主要是为了灵活性（支持变长参数等）和与 C 语言规约的某种兼容性（尽管现在有了自己的 ABI）。



## 函数调用规约之 寄存器的保护责任


这是一个非常重要的设计决策：**Go 的寄存器 ABI 中，绝大多数寄存器都是 Caller-Saved（调用者保存）的。**

这意味着：
- **被调用者 (Callee)**：可以随意使用任意寄存器（除了 SP、BP 和特殊的 g 寄存器），不需要在使用前保存它们的旧值，也不需要在返回前恢复它们。
- **调用者 (Caller)**：如果你在寄存器里存了重要的数据，并且接下来要调用另一个函数，那你必须自己把这些数据保存到栈上（Spill）。因为那个函数可能会把你的寄存器覆盖掉。

**对比 C 语言**：
在 C 语言的 System V AMD64 ABI 中，寄存器分为两类：
- Caller-Saved (如 RAX, RCX, RDX...)
- Callee-Saved (如 RBX, RBP, R12-R15...)：被调用者如果想用这些寄存器，必须先保存，用完再恢复。

**Go 为什么这么做？**
Go 这样做主要是为了配合 **垃圾回收 (GC)** 和 **栈扩容 (Stack Growth)**。
当发生函数调用时（也是一个可能的 GC 安全点），如果所有寄存器都是 Caller-Saved，那么关于“哪些寄存器里存了指针”的信息，完全由 Caller 的栈映射 (Stack Map) 决定。Callee 不需要维护复杂的寄存器保存/恢复逻辑，这简化了运行时（Runtime）扫描栈和移动栈的实现。
