
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

- Go ABI内部规范，这是最权威的定义 [Go internal ABI specification Go 内部 ABI 规范](https://github.com/golang/go/blob/master/src/cmd/compile/abi-internal.md)

- 基于寄存器的调用约定提案: [Go Register-based Calling Convention Proposal](https://github.com/golang/proposal/blob/master/design/40724-register-calling.md)
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

go1.17+为基于寄存器的函数调用规约，对于**整数/指针参数 (Integer Arguments)**，go使用一组固定的整数寄存器来传递整数、指针、布尔值等。

在AMD64 (x86-64) 架构中，整数、指针参数的寄存器顺序是：RAX、RBX、RCX、RDI、RSI、R8、R9、R10、R11（[参考文档](https://github.com/golang/go/blob/master/src/cmd/compile/abi-internal.md#architecture-specifics)）。

你的日常开发平台可能是在M芯片的MacOs上，它是ARM64架构的。在这种情况下使用的寄存器与数量都和AMD64不同。在ARM64架构上，go的整数、指针参数依次使用 R0 到 R15 (共 16 个通用寄存器)。([参考文档](https://github.com/golang/go/blob/master/src/cmd/compile/abi-internal.md#architecture-specifics))

在AMD64架构下，**浮点数参数 (Floating Point Arguments)**使用 XMM 寄存器传递：X0 ~ X14。而在ARM64架构下，浮点参数 ：依次使用 F0 到 F15。

**栈溢出 (Stack Spill)**
如果参数过多，超过了可用寄存器的数量，剩余的参数会依然通过**栈**传递。


### 是否可以验证呢？

当然可以！眼见为实。我们可以通过汇编语言来查看go函数具体调用的寄存器。我们拿如下代码举例验证：

```go
func add(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11 int) (b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11 int) {
	return a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11
}

func main() {
	// 使用 16 进制方便在汇编/调试器中观察
	// 预期寄存器分配 (AMD64):
	// 0x11 -> RAX
	// 0x22 -> RBX
	// 0x33 -> RCX
	// 0x44 -> RDI
	// 0x55 -> RSI
	// 0x66 -> R8
	// 0x77 -> R9
	// 0x88 -> R10
	// 0x99 -> R11
	// 0xAA -> Stack (栈)
	// 0xBB -> Stack (栈)
	a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11 := add(
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB,
	)
	println(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}
```

我们有两种方式查看汇编代码，第一种是直接使用`go tool compile`命令：
```shell
# -S: 输出汇编代码
# -N: 禁用优化 (防止代码被优化掉，导致你看不到参数传递)
# -l: 禁用内联 (强制发生函数调用)
go tool compile -S -N -l main.go > assembly.s
```

然后打开生成的 assembly.s 文件（或者直接在终端看输出），搜索 main 函数中调用 add 的部分。你应该能看到类似这样的指令序列（对应 AMD64 架构）：

```
MOVQ    $17, AX    // 0x11 -> RAX (Arg 0)
MOVQ    $34, BX    // 0x22 -> RBX (Arg 1)
MOVQ    $51, CX    // 0x33 -> RCX (Arg 2)
MOVQ    $68, DI    // 0x44 -> RDI (Arg 3)
MOVQ    $85, SI    // 0x55 -> RSI (Arg 4)
MOVQ    $102, R8   // 0x66 -> R8  (Arg 5)
MOVQ    $119, R9   // 0x77 -> R9  (Arg 6)
MOVQ    $136, R10  // 0x88 -> R10 (Arg 7)
MOVQ    $153, R11  // 0x99 -> R11 (Arg 8)
// 寄存器用完了，剩下的通过栈传递
MOVQ    $170, (SP) // 0xAA -> Stack (Arg 9)
MOVQ    $187, 8(SP)// 0xBB -> Stack (Arg 10)
CALL    "".add(SB)
```

如果是ARM64架构，可能是这样：
```
0x0020 00032 	MOVD	$17, R0
0x0024 00036 	MOVD	$34, R1
0x0028 00040 	MOVD	$51, R2
0x002c 00044 	MOVD	$68, R3
0x0030 00048 	MOVD	$85, R4
0x0034 00052 	MOVD	$102, R5
0x0038 00056 	MOVD	$119, R6
0x003c 00060 	MOVD	$136, R7
0x0040 00064 	MOVD	$153, R8
0x0044 00068 	MOVD	$170, R9
0x0048 00072 	MOVD	$187, R10 // ARM64架构有16个通用寄存器，在这个例子中不需要栈
0x004c 00076 	PCDATA	$1, $0
0x004c 00076 	CALL	main.add(SB)
```

第二种方法发是使用dlv debug，可以动态观察到寄存器的变化：
```shell
# 启动调试
dlv debug main.go
```

在dlv的交互界面进行如下操作：
```shell
# 设置断点并运行
break main.main
continue

# 使用 next 命令单步执行，直到 add(...) 这一行
next

# 当执行流即将进入 add 函数时（或者刚好进入 add 函数的第一行），输入
regs
# 你会看到具体的寄存器调用
```

通过以上两种方法，我们可以清晰明显地看到示例代码中`add`函数的参数是如何传递的。

### c语言函数参数怎么传？

C语言在AMD64 ABI下，整数或指针类型的参数传递通常使用这几个寄存器：RDI, RSI, RDX, RCX, R8, R9。如果参数超过 6 个，剩下的通过 栈 (Stack) 传递。

浮点数参数 ：前 8 个浮点数使用 XMM0 - XMM7 寄存器。

在 ARM64 的C语言调用规约中，前8个整数参数都是通过寄存器 R0 - R7 (代码中的 w0 - w6 ) 传递的。这比 x86_64 (只用 6 个寄存器) 能通过寄存器传递更多的参数，效率可能略高一点。

#### 官方文档与出处
这些规约定义在 System V Application Binary Interface 中。

- 文档名称 : System V Application Binary Interface - AMD64 Architecture Processor Supplement
- 关键章节 : "3.2 Function Calling Sequence"
- 在线地址 : refspecs.linuxbase.org/elf/x86_64-abi-0.99.pdf

#### 验证实例
我们可以创建一个简单的 C 文件，然后使用编译器输出汇编代码来验证上述寄存器：

```c
// test_c_call.c
// 定义一个接受 6 个以上参数的函数，确保存入寄存器和栈
int my_c_function(int a, int b, int c, int d, int e, int f, int g) {
    return a + b + c + d + e + f + g;
}

int main() {
    // 调用函数，传入特定数值方便在汇编中查找
    // 0x11 -> RDI
    // 0x22 -> RSI
    // 0x33 -> RDX
    // 0x44 -> RCX
    // 0x55 -> R8
    // 0x66 -> R9
    // 0x77 -> Stack
    my_c_function(0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77);
    return 0;
}
```

在终端中运行以下命令（使用 clang 或 gcc ）：
```shell
# -S: 生成汇编代码
# -O0: 关闭优化 (防止代码被优化掉，保持清晰的参数传递过程)
clang -S -O0 test_c_call.c -o test_c_call.s
```

打开生成的 .s 文件，找到 main 函数调用 my_c_function 的部分。你会看到类似下面的指令（AT&T 语法）：
```Assembly
	.globl	_main
_main:
    ...
    ; 准备参数
	movl	$119, (%rsp)    ; 0x77 (第7个参数) 放进栈顶 (Stack)
	movl	$17, %edi       ; 0x11 -> EDI/RDI (第1个参数)
	movl	$34, %esi       ; 0x22 -> ESI/RSI (第2个参数)
	movl	$51, %edx       ; 0x33 -> EDX/RDX (第3个参数)
	movl	$68, %ecx       ; 0x44 -> ECX/RCX (第4个参数)
	movr	$85, %r8d       ; 0x55 -> R8D/R8  (第5个参数)
	movr	$102, %r9d      ; 0x66 -> R9D/R9  (第6个参数)
	callq	_my_c_function
    ...
```

如果你的电脑是ARM64架构的，可能看到的是下面的指令：
```
; 准备参数 (Parameter Passing)
mov 	 w0, #17    ; 参数 1 (0x11) -> w0 (对应 x0 寄存器的低32位)
mov 	 w1, #34    ; 参数 2 (0x22) -> w1
mov 	 w2, #51    ; 参数 3 (0x33) -> w2
mov 	 w3, #68    ; 参数 4 (0x44) -> w3
mov 	 w4, #85    ; 参数 5 (0x55) -> w4
mov 	 w5, #102   ; 参数 6 (0x66) -> w5
mov 	 w6, #119   ; 参数 7 (0x77) -> w6

; 函数调用 (Function Call)
bl 	 _my_c_function ; BL = Branch with Link (跳转并保存返回地址)
```

从上面的汇编输出中，我们可以非常清楚地看到：

1. **前 6 个整数参数分别进入 RDI、RSI、RDX、RCX、R8、R9**  
	这与 AMD64 System V ABI 的官方规定完全一致。
2. **第 7 个参数开始放入栈 (Stack)**  
	在 x86-64 上，超过寄存器数量的参数全部会通过栈传递。
3. 整个参数装载顺序是由编译器保证的，开发者在 C 代码中看不见，但汇编层面完全透明。

到这里，我们已经亲眼验证了：
-  C 语言在 x86-64 上使用固定寄存器传递前 6 个整数参数  
- 超过的参数会退回栈  
- 调用规约不是文档上的概念，而是可以直接从机器码中观察到的“事实”

这样一来，我们就真正理解了 C 语言的 “参数怎么传”。

#### 为什么要验证 C 的方式？
之所以花时间验证 C 的参数传递，是因为：

**Go 的调用规约（尤其是 Go1.17+ 的寄存器 ABI）在设计上，明显参考了 C 的做法，但又做了很多优化。**

例如：
- C 只支持一个返回值 → Go 支持多返回值
- C 只有 6 个整数寄存器用来传参 → Go 使用更多寄存器
- C 的寄存器区分 caller-saved / callee-saved → Go 几乎全部 caller-saved
- C 遇到结构体返回需要隐藏指针 → Go 直接用多个寄存器返回

因此理解 C，是理解 Go 的最佳切入口。

接下来我们就继续回答第二个核心问题：


## 函数调用规约之 返回值怎么传？

在上一节中我们分析了“参数是如何传进函数的”。  
函数调用的另一端，是调用结束后——**返回值是怎么传回来的？**

Go 在返回值的处理上非常有特点，因为 Go 原生支持：
- 多返回值
- 多种类型组合的返回值（指针、结构体、接口等）
- 编译器自动分配返回值空间，无需开发者操心

因此在 ABI 层面，Go 设计了一套相对统一、可扩展的“返回值传递规则”。


### 1. Go 的返回值传递方式与参数非常相似

Go 1.17+ 的寄存器 ABI 采用一个非常重要的原则：
>传参数使用哪组寄存器，返回值就使用同一组寄存器。

例如在 AMD64 架构下：
- 第 1 个返回值 → RAX
- 第 2 个返回值 → RBX
- 第 3 个返回值 → RCX
- …依次类推

浮点数返回值则同样使用 XMM0、XMM1、XMM2… 等浮点寄存器。

这意味着，对于一个像：`func f() (int, int, int)` 这样的函数，Go 编译器会直接使用多个寄存器把返回值“平铺”传递给调用者，无需任何额外拷贝。

### 2. 返回值太多怎么办？依然退回栈上传递

返回值数量多到寄存器无法容纳时，Go 会自动回退到“栈传递”。
不过与 C 不同的是，这些栈空间仍然是由 **调用者 (Caller)** 提前分配的，Go 的 ABI 在这点上保持了高度一致性（Caller 清理、Caller 分配布局）。

### 3. 为什么 Go 的返回值设计比 C 简单？

与 C 语言相比：
- C 函数只有一个返回值 → 用 RAX 或 XMM0 即可
- 如果 C 想返回多个值，需要手动 **结构体返回** 或 **调用者传入指针**
- 而 Go 完全由编译器帮你处理，只需要定义多个返回值即可

换句话说，C 返回多个值只能是“曲线救国”，  而Go 返回多个值是“直通罗马”。

这使得 Go 的返回值 ABI 虽然看似复杂，但实际使用时对开发者完全透明。


### 验证实例

参考函数调用章节的方法，我们同样可以编写两个例子来快速观察返回值寄存器的具体调用情况：

```go
func multi() (int, int, int) {
	return 1, 2, 3
}

func main() {
	a, b, c := multi()
	_ = a; _ = b; _ = c
}
```

使用上一节介绍的方式查看汇编：`go tool compile -S -N -l go_return.go`

你会看到类似下面的返回值分布（AMD64）：
- 第 1 个返回值 → **RAX**
- 第 2 个返回值 → **RBX**
- 第 3 个返回值 → **RCX**

ARM64 下则会是：
- R0、R1、R2

对比而言，C 语言本身只支持一个返回值:
```c
int one() {
    return 123;
}

int main() {
    int v = one();
    return 0;
}
```

查看汇编：`clang -S -O0 c_return.c -o c_return.s`，我们会看到：

- 返回值永远在 **RAX**（AMD64）
- ARM64 则固定在 **X0**



## 函数调用规约之 谁来清理栈？

在参数传递和返回值传递都弄清楚之后，我们还需要回答一个很现实的问题：

> 函数调用完毕后，**栈上的空间究竟由谁来负责清理？**  
> 是被调用者（Callee）清理，还是调用者（Caller）清理？

这个问题是函数调用规约中的关键点，因为它会影响：
- 栈指针（SP）如何移动
- 函数调用指令序列的长度
- 支持不定参数（variadic）时的灵活性
- 整个语言的 ABI 设计复杂度

不同语言的选择不一样，而 Go 和 C 在这点上恰好做出了相同的决定。

###  Go 与 C 都使用：**Caller 清理栈（Caller-Clean）**

无论是在 C 的 System V ABI，还是 Go 的寄存器 ABI 中，结论都是：

> **调用者负责恢复栈指针（SP），也就是 Caller 清理栈。**

也就是说：
- 调用者在调用前会预留参数或返回值空间
- 被调用函数不负责清理这块空间
- 函数返回后，调用者恢复 SP
- 整个调用链保持简单统一

这种方案有几个明显优势：
1. **支持可变参数（variadic）非常灵活**
2. **函数内部更简单，无需知道调用者传了多少额外东西**
3. **寄存器 ABI 更容易实现**
4. **Go 的栈收缩/扩容（stack growth）更容易处理**

### 验证实例

这里我们给出一个最小例子，你可以用上一节的方法观察汇编:
```go
func foo(a, b int) int { return a + b }

func main() {
	_ = foo(1, 2)
}
```

查看汇编:`go tool compile -S -N -l main.go`，我们可以看到，在main函数（调用者Caller）中:

```assembly
// 1. 调用前：main 扩展栈空间SP (为参数/返回值预留空间)
0x000a 00010 (main.go:5)    SUBQ    $16, SP        
// main 负责申请 16 字节栈空间

// 2. 参数传递 (Go 1.17+ 优先使用寄存器)
0x000e 00014 (main.go:6)    MOVL    $1, AX         // 参数 a 放入 AX 寄存器
0x0013 00019 (main.go:6)    MOVL    $2, BX         // 参数 b 放入 BX 寄存器

// 3. 执行调用
0x0018 00024 (main.go:6)    CALL    main.foo(SB)


// 4. 调用后：main 负责恢复栈空间 (Caller-Clean)
0x001d 00029 (main.go:7)    ADDQ    $16, SP        
// main 负责清理这 16 字节
0x0021 00033 (main.go:7)    POPQ    BP
0x0022 00034 (main.go:7)    RET
```

而在`foo`函数中(被调用者Callee):

```assembly
TEXT    main.foo(SB), NOSPLIT|ABIInternal, $8-16
    0x0000 00000 (main.go:3)    SUBQ    $8, SP         // foo 只申请自己的局部栈帧
    ...
    //  使用 Caller 提供的空间 (Spilling)
    0x000e 00014 (main.go:3)    MOVQ    AX, main.a+16(SP) // 将寄存器 AX 中的参数存入 main 分配的栈空间
    0x0013 00019 (main.go:3)    MOVQ    BX, main.b+24(SP) // 将寄存器 BX 中的参数存入 main 分配的栈空间
    ...
    0x001f 00031 (main.go:3)    ADDQ    $8, SP         // foo 只清理自己的 8 字节
    0x0023 00035 (main.go:3)    POPQ    BP
    0x0024 00036 (main.go:3)    RET
```
这段汇编印证了上文中的描述：
1. 调用前 ( main ): SUBQ $16, SP —— Caller ( main ) 负责分配参数所需的栈空间（即使参数通过寄存器传递，栈上依然预留了 Spill 空间）。
2. 函数内部 ( foo ): foo 函数内部只处理自己的局部变量空间（ SUBQ $8, SP 和 ADDQ $8, SP ），完全不触碰参数栈空间的清理。
3. 调用后 ( main ): ADDQ $16, SP —— Caller ( main ) 在 CALL 返回后，负责回收之前分配的 16 字节空间。
这就是标准的 Caller-Clean 模式： 谁分配（参数空间），谁清理 。


### C语言的验证实例

**让我们继续看看C语言：**

```c
int foo(int a, int b) { return a + b; }

int main() {
    foo(1, 2);
    return 0;
}
```

查看汇编：`clang -S -O0 c_stack.c -o c_stack.s`，我们会看到 main 函数 (调用者 Caller)：
```assembly
_main:
    pushq   %rbp
    movq    %rsp, %rbp
    
    // 1. 调用前：main 调整栈指针 (Caller-Clean 准备)
    subq    $16, %rsp      // 分配 16 字节栈空间（用于局部变量或对齐）
    
    movl    $0, -4(%rbp)   // main 的返回值 0 暂存
    
    // 2. 参数传递 (System V AMD64 ABI 优先使用寄存器)
    movl    $1, %edi       // 第一个参数 a -> 寄存器 EDI
    movl    $2, %esi       // 第二个参数 b -> 寄存器 ESI
    
    // 3. 执行调用
    callq   _foo
    
    // 4. 调用后：main 恢复栈指针 (Caller-Clean)
    xorl    %eax, %eax     // 设置 main 返回值为 0
    addq    $16, %rsp      // 恢复栈指针，清理之前分配的 16 字节
    popq    %rbp
    retq
```

foo 函数 (被调用者 Callee):
```assembly
_foo:
    pushq   %rbp
    movq    %rsp, %rbp
    
    // 5. 内部操作：只处理自己的栈帧
    movl    %edi, -4(%rbp) // 将寄存器传来的参数 a 存入自己的栈帧
    movl    %esi, -8(%rbp) // 将寄存器传来的参数 b 存入自己的栈帧
    
    movl    -4(%rbp), %eax // 取出 a
    addl    -8(%rbp), %eax // a + b，结果放入 EAX (返回值寄存器)
    
    // 6. 返回：不负责清理 Caller 分配的参数空间（如果有的话）
    popq    %rbp
    retq
```
这个 C 语言的例子同样展示了 Caller-Clean 的特征，但细节与 Go 略有不同（主要受 System V AMD64 ABI 影响）：

1. 栈调整 (`main`) : `subq $16, %rsp`和`addq $16, %rsp`。`main`负责分配和回收栈空间。虽然这里主要是为了对齐（16字节对齐）或局部变量，而非直接用于传参（因为参数1和2放进了寄存器），但原则一致： 栈的变动由*Caller*负责复原 。
2. 被调用者 (`foo`) : 内部非常干净，只保存自己的帧指针，计算完直接返回。它完全不知道也不关心`main`是否在栈上给它留了空间。
3. 对比*Go*:
   - *Go*: 即使使用寄存器传参，Caller(`main`) 依然会在栈上预留 "Spill slots"（溢出槽），以防`Callee`需要把寄存器里的参数刷回内存（例如进行栈扩容或调试时）。
   - *C* (Clang -O0) : 在这个简单例子中，参数直接进寄存器，Caller 分配的栈空间主要是为了满足 ABI 的 16 字节对齐要求，或者预留给局部变量。如果参数超过 6 个，多出的参数才会真正压栈，那时候 Caller-Clean 的特征会更加明显（Caller`push`参数，Call 完后`add sp`）。
总结：无论是 Go 还是 C（在 x86-64 System V ABI 下）， Caller 负责维护由它发起的函数调用所引起的栈指针变化 ，这就是 Caller-Clean。



## 函数调用规约之 寄存器的保护责任

在参数传递、返回值传递、栈空间管理都讲清楚之后，我们还剩下一个非常重要的问题：

> 当一个函数被调用时，它可以随意使用哪些寄存器？
> 哪些寄存器必须被保护？

这在函数调用规约里叫做：

> 寄存器保护责任（Register Saving Convention）

它决定了：
- 函数调用时，寄存器会不会被破坏
- 谁需要保存寄存器内容（Caller 还是 Callee）
- ABI 的性能与灵活性
- 是否容易支持多返回值与 GC 扫描

不同语言的策略不同，而 Go 与 C 的选择几乎走向了两个极端。

### C语言：Caller-Saved + Callee-Saved 混合模型

在 C（System V ABI）中，寄存器被分成两类：

| 类型               | 责任方       | 代表寄存器                           |
| ---------------- | --------- | ------------------------------- |
| **Caller-Saved** | 调用者保存     | RAX, RCX, RDX, RSI, RDI, R8–R11 |
| **Callee-Saved** | 被调用者保存+恢复 | RBX, RBP, R12–R15               |

含义很简单：

- 如果调用者希望调用后某些寄存器内容仍然保持不变
	→ 调用前需要自己保存（push 到栈）

- 如果被调用者使用了某些“必须保持稳定”的寄存器
	→ 使用前必须保存，用完后恢复

这种模式历史悠久，灵活性很好，但也增加了 ABI 的复杂度。

### Go语言：几乎全部 Caller-Saved

Go 的寄存器 ABI 采取了一个非常强势的策略：

> 除了少数特殊寄存器外，全部寄存器都是 Caller-Saved。
> Callee 可以自由使用寄存器，而不需要保存或恢复。

也就是说：
1. Go 函数“想用多少寄存器都随便用”
2. 使用前无需保存旧值
3. 返回前也无需恢复
4. 所有破坏风险由 Caller 自己负责

为什么 Go 要这么“激进”？

原因非常重要：

### Go Caller-Saved设计的原因

#### 1. GC（垃圾回收）更容易实现

Go 的 GC 需要在安全点扫描栈、识别指针位置。

如果寄存器被 Callee-Saved 频繁压栈/恢复，
GC 需要处理更多“不确定性”，复杂度翻倍。

全部 Caller-Saved 意味着：
- 某一时刻“哪些寄存器里存了指针”只需看 Caller 的 Stack Map
- 被调用代码可以随意覆盖寄存器
- 更容易做抢占（preemption）与栈扩容


#### 2. 简化栈扩容（stack growth）

Goroutine 的栈会动态扩容、收缩，而这要求 ABI：
- 栈布局必须容易重新定位
- 不能把寄存器状态复杂地压栈、恢复

Caller-Saved 刚好非常匹配 Go 的栈模型。

#### 3. 配合多返回值的寄存器平铺

C 要保护寄存器，否则多返回值会极其困难。

Go 则可以随意占用：
- RAX、RBX、RCX…
- X0、X1、X2…

做多返回值“平铺”传输，Callee不需要保存任何寄存器，多返回值也天然高效。

#### 精简示例：Go 如何“随意破坏寄存器”？

```go
func use() {
	// 随意计算，破坏寄存器
	_ = 1 + 2 + 3
}

func main() {
	x := 10
	use()
	_ = x
}
```

查看汇编：`go tool compile -S -N -l go_reg.go`

`main.use`函数 (被调用者 Callee):
```assembly
TEXT    main.use(SB), NOSPLIT|NOFRAME|ABIInternal, $0-0
    // 这里没有任何保存寄存器的代码 (如 PUSH BX, PUSH R12 等)
    // 也没有分配栈帧 (NOFRAME)
    RET
```
现象 : use 函数非常干净，因为它只是做了一些编译器优化后直接能算出来的常量计算（实际上 _ = 1 + 2 + 3 被编译器直接优化掉了，甚至没生成计算指令，直接 RET ）。但关键点在于： 它没有义务保存任何通用寄存器 。如果它真的使用了寄存器（例如复杂的逻辑），它会直接覆盖使用，而不需要先保存旧值再恢复。

`main.main`函数 (调用者 Caller):
```assembly
TEXT    main.main(SB), ABIInternal, $16-0
    ...
    // 1. 保存现场 (Spilling)
    SUBQ    $8, SP                 // main 分配 8 字节栈空间
    ...
    MOVQ    $10, main.x(SP)        // 这里的 x=10 是直接存放在栈上的！
                                   // 即使 x 是局部变量，main 也没有把它一直放在寄存器里跨越函数调用
    
    // 2. 调用函数
    CALL    main.use(SB)           // 调用 use()
                                   // 在这期间，use() 可以随意破坏任何通用寄存器
    
    // 3. 恢复现场 (Restoring)
    // 实际上这里没有显式的 POP 操作，因为 x 本来就在栈上 (main.x(SP))
    // 如果 main 后续要用 x，它会从栈上 main.x(SP) 重新读取
    
    ADDQ    $8, SP                 // 恢复栈指针
    POPQ    BP
    RET
```

你会看到：
- Callee (`use`) 无负担:进入`use()`后，不需要保存和恢复任何通用寄存器，它拥有对寄存器的完全支配权。
- Caller (`main`) 负责保护 : main 函数知道调用 use 可能会破坏寄存器。因此， main 必须确保它关心的变量（如 x ）在调用前后是安全的。
	- 在这个例子中，Go 编译器直接把`x`分配在了栈上 (`main.x(SP)`)。
	- 这意味着`x`天然就“保存”在了内存中。
	- 调用`use`后，如果`main`还要用`x`，它会从栈上读，而不是指望某个寄存器里还存着`x`。

这就是 Caller-Saved 策略的体现。谁（调用者）想保留数据，谁就自己负责存到栈上（Spill），别指望被调用者帮你守着寄存器。


### 精简示例：C 确实会主动保存 Callee-Saved 寄存器

```c
// c_reg.c
int use() {
    return 1 + 2 + 3;
}

int main() {
    int x = use();
    return x;
}
```

查看汇编：`clang -S -O0 c_reg.c -o c_reg.s`你会看到类似：

```assembbly
pushq   %rbx   ; 保存 RBX
...     ; 使用 RBX
popq    %rbx   ; 恢复 RBX
```

这说明，C 函数使用 RBX，必须按照顺序: 保存 → 使用 → 恢复
- 这是 Callee-Saved 的严格要求
- 与 Go 的“自由使用寄存器”形成鲜明对比

### 小结
Go 与 C 最大的 ABI 差异之一，就是寄存器保护责任。
- C：Caller-Saved + Callee-Saved 混合
- Go：几乎全部 Caller-Saved

这种设计让 Go 在以下方面高度优化：
- 垃圾回收（GC）
- 栈扩容（stack growth）
- 抢占（preemption）
- 多返回值传递
- 运行时实现简化

可以说，Go 的寄存器策略完全为“运行时 + GC”而生。


