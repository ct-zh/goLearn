# [effective-go](https://go.dev/doc/effective_go)
## 介绍
Go 是一种新语言。虽然它借鉴了现有语言的思想，但它具有不同寻常的特性，使得有效的 Go 程序在性质上有别于用它的同类语言编写的程序。将 C++ 或 Java 程序直接翻译成 Go 语言不太可能产生令人满意的结果--Java 程序是用 Java 编写的，而不是 Go。另一方面，从 Go 的角度思考问题，可能会产生一个成功但完全不同的程序。换句话说，要写好 Go 语言，了解它的特性(properties)和习语(idioms)很重要。同样重要的是，要了解 Go 编程的既定惯例，如命名、格式化、程序结构等，这样你编写的程序才能让其他 Go 程序员容易理解。

本文档提供了编写清晰、惯用的 Go 代码的技巧。它是对语言规范、Go 之旅和如何编写 Go 代码的补充，所有这些内容都应首先阅读。

2022 年 1 月添加的注释：本文档是为 Go 于 2009 年发布时编写的，此后未作重大更新。尽管由于 Go 语言本身的稳定性，它是了解如何使用该语言的很好指南，但它对库的介绍很少，对 Go 生态系统自撰写以来发生的重大变化，如构建系统、测试、模块和多态性，也未作任何介绍。由于已经发生了太多的变化，而且越来越多的文档、博客和书籍已经很好地描述了 Go 的现代用法，因此我们没有更新该书的计划。Effective-Go仍然有用，但读者应该明白它远非一本完整的指南。上下文请参见 [issue 28782](https://go.dev/issue/28782) 。

### Examples

Go Package sources 的目的不仅是作为核心库，也是作为如何使用该语言的示例。此外，许多软件包都包含可直接从 go.dev 网站运行的、独立的可执行示例，例如[这个示例](https://go.dev/pkg/strings/#example-Map)（如果需要，请单击 "示例 "打开）。如果您对如何处理某个问题或如何实现某个功能有疑问，库中的文档、代码和示例可以为您提供答案、思路和背景。

## 格式化Formatting

格式问题争议最大，但影响最小。人们可以适应不同的格式风格，但最好是不必适应，而且如果每个人都遵守相同的风格，就可以减少花在这个主题上的时间。问题是如何在没有长篇规范性风格指南的情况下实现这一目标。

在 Go 中，我们采取了一种不同寻常的方法，让机器来处理大部分格式问题。gofmt 程序（也可用 go fmt，在软件包级别而非源文件级别运行）读取 Go 程序，并以标准的缩进和垂直对齐方式输出源代码，保留注释，必要时重新格式化。如果你想知道如何处理某种新的布局情况，请运行 gofmt；如果得到的结果似乎不对，请重新整理你的程序（或提交有关 gofmt 的 bug），不要绕过它。

举例来说，您无需花时间将结构字段上的注释排成一行。Gofmt 会为你完成这项工作。给定声明

```go
type T struct {
    name string // name of the object
    value int // its value
}
```

gofmt 会将列排成一行：

```go
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

标准软件包中的所有 Go 代码都已使用 gofmt 格式化。

一些格式化细节仍然存在。例如：

-  缩进
  我们使用制表符缩进，gofmt 默认使用制表符。如果有必要，请使用空格。

- 行长
  Go 没有行长限制。不用担心代码过长会影响程序运行。如果觉得行太长，可以用额外的制表符包起来并缩进。

- 括号
  Go 需要的括号比 C 和 Java 少：控制结构（if、for、switch）的语法中没有括号。此外，运算符的优先级层次结构更短、更清晰，所以

  ```go
  x<<8 + y<<16
  ```

  的意思，这与其他语言不同。

## 注释Commentary

Go 提供 C 风格的 /* */ 块注释和 C++ 风格的 // 行注释。行注释是规范；块注释主要作为包注释出现，但在表达式中或禁用大量代码时非常有用。

在顶层声明之前出现的注释，中间没有换行符，被视为声明本身的文档。这些 "文档注释 "是特定 Go 软件包或命令的主要文档。有关文档注释的更多信息，请参阅 ["Go 文档注释"](https://go.dev/doc/comment)。

## 命名Names

在 Go 语言中，名称和其他语言一样重要。它们甚至有语义上的影响：名字在包外的可见性取决于它的第一个字符是否大写。因此，我们值得花一点时间来谈谈 Go 程序中的命名规则。

### Package names

导入软件包时，软件包名称将成为内容的访问器。如下：

```go
import "bytes"
```

导入软件包后可以使用 bytes.Buffer。如果每个使用软件包的人都能使用相同的名称来指代其内容，那将会很有帮助，这意味着好的软件包名称应该是简短、简洁、令人回味的。按照惯例，软件包的名称都是小写、单个单词；不需要下划线或混合大写。尽量简短，因为每个使用你的软件包的人都会输入这个名字。不要先验地担心重名问题。软件包名称只是导入时的默认名称；它不需要在所有源代码中都是唯一的，在极少数重名的情况下，导入软件包可以选择不同的名称在本地使用。无论如何，由于导入的文件名决定了使用的是哪个软件包，因此很少发生混淆。

另一个惯例是，软件包名称是其源目录的基名；src/encoding/base64 中的软件包被导入为 "encoding/base64"，但名称是 base64，而不是 encoding_base64，也不是 encodingBase64。

软件包的导入者将使用名称来引用其内容，因此软件包中的导出名称可以使用这一事实来避免重复。(不要使用` import .`符号，它可以简化必须在测试包之外运行的测试，但在其他情况下应避免使用）。例如，bufio 包中的缓冲阅读器类型被称为 Reader，而不是 BufReader，因为用户看到的是 bufio.Reader，这是一个简洁明了的名称。此外，由于导入实体总是使用其软件包名称，因此 bufio.Reader 不会与 io.Reader 冲突。类似地，在 Go 中，ring.Ring 的构造函数通常被称为 NewRing，但由于 Ring 是该包导出的唯一类型，而该包又被称为 ring，因此它只被称为 New，该包的客户端将其视为 ring.New。使用软件包结构可以帮助你选择好的名称。

另一个简短的例子是 once.Do；once.Do(setup) 读起来很好，写成 once.DoOrWaitUntilDone(setup) 也不会有什么改进。长名称并不会自动提高可读性。一个有用的文档注释往往比一个超长的名称更有价值。

### Getters

Go 并不自动支持 getter 和 setter。自己提供获取器和设置器并没有什么问题，而且这样做通常也很合适，但在获取器的名称中加入 Get 既不习惯，也没有必要。如果有一个字段叫 owner（小写，未导出），那么获取方法就应该叫 Owner（大写，导出），而不是 GetOwner。导出时使用大写名称提供了区分字段和方法的钩子。如果需要，设置函数可能会被称为 SetOwner。这两个名称在实际使用中都很好理解：

```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

### Interface names

按照惯例，单方法接口以方法名加上-er 后缀或类似的修饰来命名，以构造一个代理名词：如`Reader`、`Writer`、`Formatter`、`CloseNotifier` 等。

有许多这样的名称，尊重它们和它们捕获的函数名是很有成效的。`Read`、`Write`、`Close`、`Flush`、`String`等都有规范的签名和含义。为避免混淆，除非你的方法具有相同的签名和含义，否则不要使用这些名称。反之，如果您的类型实现了一个与知名类型上的方法具有相同含义的方法，则应赋予其相同的名称和签名；将字符串转换方法称为 String，而不是 ToString。

### MixedCaps

最后，Go 的惯例是使用 MixedCaps 或 mixedCaps 而不是下划线来书写多字名称。

## 分号Semicolons

与 C 语言一样，Go 的正式语法也使用分号来结束语句，但与 C 语言不同的是，这些分号不会出现在源代码中。取而代之的是，词法词典会使用一条简单的规则，在扫描时自动插入分号，因此输入文本中基本上没有分号。

规则是这样的，如果换行符前的最后一个标记是标识符（包括 int 和 float64 等字）、数字或字符串常量等基本字面量，或下列标记之一

```
break continue fallthrough return ++ -- ) }
```

词法词典总是在符号后插入分号。这可以概括为 "如果换行符出现在可以结束语句的标记之后，则插入分号"。

分号也可以在紧靠结尾括号之前省略，因此语句如

```
 go func() { for { dst <- <-src } }()
```

不需要分号。惯用的 Go 程序只有在 for 循环子句等地方才有分号，用于分隔初始化器、条件和继续元素。分号也是分隔一行中多条语句所必需的，如果你这样写代码的话。

分号插入规则的一个后果是，不能将控制结构（if、for、switch 或 select）的开头分号放在下一行。如果这样做，就会在括号前插入分号，这可能会造成不必要的影响。写法如下

```go
if i < f() {
    g()
}

// not like this

if i < f()  // wrong!
{           // wrong!
    g()
}
```

## Control stutures

Go的控制结构与 C 语言的控制结构有关，但在一些重要方面有所不同。Go 没有 do 或 while 循环，只有略微通用的 for；switch 更为灵活；if 和 switch 与 for 一样接受可选的初始化语句；break 和 continue 语句接受可选的标签来标识要中断或继续的内容；还有新的控制结构，包括类型转换和多路通信复用器 select。语法也略有不同：没有圆括号，主体必须始终以`{}`括号分隔。

### If
在 Go 中，一个简单的 if 是这样的

```go
if x > 0 {
    return y
}
```

强制性大括号鼓励在多行中编写简单的 if 语句。无论如何，这样做都是很好的风格，尤其是当正文包含控制语句（如 return 或 break）时。

由于 if 和 switch 接受初始化语句，因此经常可以看到用 if 和 switch 来设置局部变量。

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

在 Go 库中，你会发现当 if 语句没有流入下一条语句时，即主体以 break、continue、goto 或 return 结尾时，不必要的 else 会被省略。

```go
f, err := os.Open(name)
if err != nil {
    return err
}
codeUsing(f)
```

这是代码必须防范一系列错误情况的常见示例。如果成功的控制流沿页面向下运行，在出现错误时将其消除，那么代码的可读性就会很好。由于出错情况往往以返回语句结束，因此生成的代码不需要 else 语句。

```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```

### Redeclaration and reassignment

题外话:上一节的最后一个示例演示了 := 短声明形式如何工作的细节。调用 os.Open 的声明为

```go
f, err := os.Open(name)
```

该语句声明了两个变量 f 和 err。几行之后，对 f.Stat 的调用为

```go
d, err := f.Stat()
```

看起来好像声明了 d 和 err。但请注意，err 在两条语句中都出现了。这种重复是合法的：err 在第一条语句中声明，但只是在第二条语句中重新赋值。这意味着 f.Stat 的调用使用了上面声明的 Err 变量，只是赋予了它一个新值。

在 := 声明中，变量 v 即使已经声明过，也可以出现，前提是：

- 该声明与 v 的现有声明处于同一作用域（如果 v 已在外层作用域中声明，则该声明将创建一个新变量 §）、

- 初始化中的相应值可赋值给 v

- 至少还有一个变量是由声明创建的。

这个不寻常的属性纯粹是实用主义的体现，例如，它使得在很长的 if-else 链中使用一个 err 值变得容易。你会经常看到使用它。

§ 这里值得注意的是，在 Go 中，函数参数和返回值的作用域与函数体相同，尽管它们在词法上出现在包围函数体的大括号之外。

### For

Go 的 for 循环与 C 的相似，但并不相同。它统一了 for 和 while，没有 do-while。它有三种形式，其中只有一种有分号。

```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }

// 简短的声明可以方便地在循环中直接声明索引变量。
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}

// 如果要对array、slice、string或者map进行循环遍历，或从通道中读取数据，则range可以管理这个循环。
for key, value := range oldMap {
    newMap[key] = value
}

// 如果只需要范围内的第一个项目（key或index），放弃第二个项目：
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}

// 如果只需要范围中的第二个项目（值），则使用空白标识符（下划线）来舍弃第一个项目：
sum := 0
for _, value := range array {
    sum += value
}
// 空白标识符有很多用途，将在后文介绍。
```

对于字符串来说，range的作用更大，它可以通过解析 UTF-8 来分解出单个的 Unicode 代码点。错误的编码会消耗一个字节，并产生替换符 U+FFD。(符文名称（带相关内置类型）是 Go 对单个 Unicode 代码点的术语。详情请参见[语言规范](https://go.dev/ref/spec#Rune_literals)）。循环

```go
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}

// 输出

character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
```

最后，Go 没有逗号操作符，++ 和 -- 是语句而不是表达式。因此，如果要在 for 中运行多个变量，应使用并行赋值（尽管这排除了 ++ 和 --）。

```go
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

### Switch

Go 的 switch 比 C 的更通用。表达式不一定是常量，甚至也不一定是整数，从上到下对情况进行比较，直到找到匹配为止，如果switch没有表达式，则切换到 true。因此，将 if-else-if-else 链写成 switch 是可能的，也是习以为常的。

```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

Swtich不会自动fall through，但可以通过逗号分隔列出多个 case:

```go
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

虽然在 Go 中，break 语句并不像其他一些类 C 语言那样常见，但它可以用来提前终止switch。但有时，需要中断的是周围的循环，而不是switch，在 Go 中，可以通过在循环上添加一个标签并 "breaking "到该标签来实现。本例展示了这两种用法。

```go
Loop:
    for n := 0; n < len(src); n += size {
        switch {
        case src[n] < sizeOne:
            if validateOnly {
                break
            }
            size = 1
            update(src[n])

        case src[n] < sizeTwo:
            if n+1 >= len(src) {
                err = errShortInput
                break Loop
            }
            if validateOnly {
                break
            }
            size = 2
            update(src[n] + src[n+1]<<shift)
        }
    }
```

当然，continue 语句也接受一个可选标签，但它只适用于循环。

在本节的最后，下面是一个使用两个switch语句的bytes slices比较例子：

```go
// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b []byte) int {
    for i := 0; i < len(a) && i < len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1
        case a[i] < b[i]:
            return -1
        }
    }
    switch {
    case len(a) > len(b):
        return 1
    case len(a) < len(b):
        return -1
    }
    return 0
}
```

### Type switch类型断言

**类型断言（type switch）**可以用于检测接口变量的动态类型。它使用带有关键字`type`的类型断言语法。在类型断言中声明的变量，会在每个分支中拥有对应的具体类型。通常会复用变量名，在每个分支中实际上是声明了同名但不同类型的变量。

```go
var t interface{}
t = functionOfSomeType()  // 将某个函数的返回值赋值给接口变量 t

switch t := t.(type) {  // 使用类型断言判断 t 的实际类型
default:
    fmt.Printf("unexpected type %T\n", t)  // 未知类型
case bool:
    fmt.Printf("boolean %t\n", t)         // 布尔型
case int:
    fmt.Printf("integer %d\n", t)         // 整型
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t)  // 指向布尔型的指针
case *int:
    fmt.Printf("pointer to integer %d\n", *t)  // 指向整型的指针
}
```

## Functions

### 多返回值 multiple return values

Go 语言的一大特色是函数和方法可以返回多个值。这种特性可以用来改善 C 语言中一些笨拙的惯用法，例如带内错误返回（例如 EOF 时返回 -1）和通过地址传递的参数进行修改。

在 C 语言中，写操作的错误会通过负值的字节数来表示，错误代码则隐藏在一个临时变量中。在 Go 语言中，`Write` 函数可以返回写入的字节数和一个错误值：“是的，您写了一些字节，但由于设备已满，所以并非全部”。来自 `os` 包的用于文件的 `Write` 方法的函数签名为：

```go
func (file *File) Write(b []byte) (n int, err error)
```

正如文档所述，它返回写入的字节数 (`n`) 和一个非 `nil` 的错误值 (`err`)，当 `n` 不等于 `len(b)` 时表示发生了错误。这是一种常见的风格，有关更多示例请参阅错误处理部分。

类似地，这种方法避免了需要传递指向返回值的指针来模拟引用参数。下面是一个简单的函数，用于从字节切片中的指定位置获取一个数字，并返回该数字和下一个位置。

```go
func nextInt(b []byte, i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}
```

您可以像这样使用它来扫描输入切片 `b` 中的数字：

```go
for i := 0; i < len(b); {
    x, i = nextInt(b, i)
    fmt.Println(x)
}
```

### 命名结果参数
Go 语言允许为函数的返回结果（也称为返回值参数）命名，命名后的结果参数可以像普通变量一样使用。命名后，它们会在函数开始时初始化为各自类型的零值。如果函数执行了一个没有参数的 `return` 语句，那么将会使用结果参数当前的值作为返回的值。

使用命名结果参数不是强制的，但是可以使代码更简洁易懂。因为它们相当于文档注释，可以清楚地表明每个返回值的含义。例如，为 `nextInt` 函数的结果命名，就可以一目了然地看出返回的两个 `int` 值分别代表什么。

```go
func nextInt(b []byte, pos int) (value, nextPos int) {}
```

由于命名结果参数在初始化后会直接作为 `return` 语句的一部分，因此可以简化并提升代码的可读性。下面是 `io.ReadFull` 函数使用命名结果参数的一个良好示例：

```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return // 省略了 n 和 err 的参数
}
```

在这个例子中，`n` 和 `err` 被直接作为 `return` 语句的一部分返回，而不需要显式地写成 `return n, err`。

### Defer语句

Go 语言的 `defer` 语句可以安排一个函数调用（称为“延迟函数”）在执行 `defer` 语句的函数返回之前立即运行。这是一种不同寻常但有效的方式来处理某些情况，例如必须释放资源，无论函数采取何种路径返回。典型的例子是解锁互斥锁或关闭文件。

```go
// Contents 函数返回文件的内容作为字符串。
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close 将在函数返回时执行。

    // ... 读取文件内容 ...
}

```

延迟调用诸如 `Close` 的函数有两个优点。首先，它可以确保您永远不会忘记关闭文件，这是一个很容易犯的错误，尤其是在您稍后编辑函数以添加新的返回路径时。其次，这意味着关闭操作紧贴着打开操作，这比将它放在函数的末尾要清晰得多。

延迟函数的参数（包括方法的接收者）是在 `defer` 执行时计算的，而不是在函数调用时计算的。除了避免因函数执行过程中变量值的变化带来的困扰之外，这也意味着单个延迟调用位置可以延迟多个函数执行。下面是一个不切实际的例子：

```go
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```

延迟函数以后进先出的顺序执行，因此当函数返回时，这段代码将导致打印 4 3 2 1 0。一个更可信的例子是在程序中跟踪函数执行情况的简单方法。我们可以这样编写几个简单的跟踪例程：

```go
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// Use them like this:
func a() {
    trace("a")
    defer untrace("a")
    // do something....
}
```

我们可以通过利用延迟函数的参数在 `defer` 执行时计算这一特性来做得更好。跟踪例程可以设置传递给取消跟踪例程的参数。

```go
func trace(s string) string {
  fmt.Println("entering:", s)
  return s
}

func un(s string) {
  fmt.Println("leaving:", s)
}

func a() {
  defer un(trace("a"))  // 将要执行的函数名传递给 defer
  fmt.Println("in a")
}

func b() {
  defer un(trace("b"))
  fmt.Println("in b")
  a()
}

func main() {
  b()
}

```

对于习惯于其他语言的块级资源管理的程序员来说，`defer` 可能看起来有些奇怪，但它最有趣和最强大的应用恰恰来自于它不是基于块的，而是基于函数的。在关于 `panic` 和 `recover` 的章节中，我们将看到另一个它的强大功能的例子。



## 数据 Data

### 新的内存分配 Allocation with new

Go 语言拥有两种内存分配原语：内置函数 `new` 和 `make`。它们的功能不同，适用于不同的类型，这可能会让人困惑，但规则很简单。我们先来谈谈` new`。这是一个内置函数，用于分配内存。但与其他一些语言中的同名函数不同，它不会初始化内存，只会将其归零。也就是说，`new(T)` 会为一个类型为 T 的新项目分配归零的存储空间，并返回其地址（一个类型为 *T 的值）。用 Go 术语来说，它返回一个指向新分配的类型 T 的零值的指针。

由于 `new` 返回的内存已归零，因此在设计数据结构时，最好让每个类型的零值可以直接使用，而无需进一步初始化。这意味着数据结构的用户可以使用 `new` 创建一个数据结构，然后立即开始使用。例如，`bytes.Buffer` 的文档指出“`Buffer` 的零值是一个可以直接使用的空缓冲区”。类似地，`sync.Mutex` 没有显式的构造函数或 `Init` 方法。相反，`sync.Mutex` 的零值被定义为一个未锁定的互斥锁。

零值可用的特性是可传递的。考虑下面的类型声明：

```go
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
```

`SyncedBuffer` 类型的值也可以在分配或声明后立即使用。在下一个代码段中，p 和 v 都可以正常工作，无需进一步安排。

```go
p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
```

### 构造函数和复合字面量

有时，零值并不能满足需求，需要使用初始化的构造函数，例如下面来自 `os` 包的例子：

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

这段代码中存在很多重复的初始化操作。我们可以使用复合字面量来简化它。复合字面量是一种表达式，每次求值都会创建一个新的实例。

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

需要注意的是，与 C 语言不同，在 Go 语言中，返回局部变量的地址完全是合法的。因为函数返回后，该变量的存储空间仍然存在。事实上，每次求值时，取复合字面量的地址都会分配一个新的实例，因此我们可以将最后两行合并。

```go
return &File{fd, name, nil, 0}
```

复合字面量的字段按顺序排列，所有字段都必须存在。但是，通过将元素显式标记为 `field:value` 对，初始化器可以按任何顺序出现，缺失的元素将保留其各自的零值。因此我们可以写成：

```go
return &File{fd: fd, name: name}
```

作为一个特例，如果一个复合字面量完全没有字段，它将创建一个该类型的零值。`new(File)` 和 `&File{}` 是等价的。

复合字面量也可以用于创建数组、切片和映射，字段标签将根据需要作为索引或映射键。在这些示例中，初始化操作可以正常工作，即使 `Enone`、`Eio` 和 `Einval` 的值不明确，只要它们是不同的即可。

```go
a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
```



### 使用 make 进行内存分配

回到内存分配。内置函数 `make(T, args)` 的用途与 `new(T)` 不同。它只用于创建切片、映射和通道，并且它返回一个类型为 T 的已初始化（非零）值（而不是 *T）。之所以进行区分，是因为这三种类型在底层表示引用数据结构，这些数据结构在使用之前必须进行初始化。例如，切片是一个包含指向数据（位于数组内部）的指针、长度和容量的三项描述符，在这些项初始化之前，切片是 nil。对于切片、映射和通道，`make` 会初始化内部数据结构并准备该值供使用。例如：

```go
make([]int, 10, 100)
```

它会分配一个包含 100 个 int 的数组，然后创建一个指向数组前 10 个元素的切片结构，该结构的长度为 10，容量为 100。（创建切片时，可以省略容量；有关详细信息，请参阅有关切片的章节。）相比之下，`new([]int)` 返回一个指向新分配的、零值的切片结构的指针，即指向一个 nil 切片值的指针。

以下示例说明了 `new` 和 `make` 之间的区别：

```go
var p *[]int = new([]int)       // 分配切片结构；*p == nil；很少使用
var v  []int = make([]int, 100) // 切片 v 现在引用一个新的包含 100 个 int 的数组

// 非常复杂的写法：
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// 惯用写法：
v := make([]int, 100)
```

请记住，`make` 仅适用于maps、slices和channels，并且不返回指针。要获取显式指针，可以使用 `new` 分配内存或显式获取变量的地址。

### 数组

数组在规划内存的详细布局时很有用，有时可以帮助避免分配，但它们主要用作切片的构建块（下一节将介绍切片）。为了为该主题奠定基础，这里简要介绍一下数组。

Go 语言中数组的工作方式与 C 语言有本质的区别。在 Go 语言中：

- 数组是值。将一个数组赋值给另一个数组会复制所有元素。
- 具体来说，如果您将一个数组传递给函数，它将接收数组的副本，而不是指向数组的指针。
- 数组的大小是其类型的一部分。类型 `[10]int` 和 `[20]int` 是不同的类型。

值的属性既有用又昂贵；如果您想要类似 C 的行为和效率，可以传递指向数组的指针。

```go
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // 注意显式的地址运算符
```

但是即使是这种风格也不是典型的 Go 代码。还是使用切片更惯用。

### 切片

切片是对数组的封装，提供了更通用、更强大、更方便的数据序列接口。除了具有显式维度的项目（例如变换矩阵）之外，Go 语言中的大多数数组编程都使用切片而不是简单数组。

切片引用底层数组，如果您将一个切片赋值给另一个切片，那么它们都将引用同一个数组。如果函数接受一个切片参数，并且该函数修改了切片元素，那么调用者将可以看到这些更改，类似于传递指向底层数组的指针。因此，`Read` 函数可以接受一个切片参数而不是指针和计数；切片中的长度设置了要读取的数据量的上限。以下是来自 `os` 包的 `File` 类型 `Read` 方法的签名：

```go
func (f *File) Read(buf []byte) (n int, err error)
```

该方法返回读取的字节数和一个错误值（如果存在）。要将数据读入更大缓冲区 `buf` 的前 32 个字节，可以使用切片运算符（这里用作动词）对缓冲区进行切片。

```go
n, err := f.Read(buf[0:32])
```

这种切片操作很常见且高效。事实上，暂时抛开效率不谈，以下代码片段也可以读取缓冲区的前 32 个字节。

```go
var n int
var err error
for i := 0; i < 32; i++ {
    nbytes, e := f.Read(buf[i:i+1])  // 读取一个字节。
    n += nbytes
    if nbytes == 0 || e != nil {
        err = e
        break
    }
}
```

只要不超出底层数组的限制，就可以更改切片的长度；只需将其自身的一部分重新赋值给它即可。切片的容量可以通过内置函数 `cap` 获取，它报告切片可能拥有的最大长度。下面是一个将数据追加到切片上的函数。如果数据超过容量，则会重新分配切片。该函数会返回生成的切片。该函数利用了这样一个事实，即 `len` 和 `cap` 应用于空切片时是合法的，并且会返回 0。

```go
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // 需要重新分配
        // 为未来增长分配所需的两倍空间。
        newSlice := make([]byte, (l+len(data))*2)
        // copy 函数是预声明的，适用于任何切片类型。
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}
```

必须在函数后返回切片，因为虽然 `Append` 可以修改 `slice` 的元素，但是切片本身（包含指针、长度和容量的运行时数据结构）是按值传递的。

追加到切片的概念非常有用，以至于内置函数 `append` 也实现了该功能。但是，为了理解该函数的设计，我们还需要更多一些信息，因此我们稍后会再介绍它。

### 二维切片

Go 语言的数组和切片是一维的。要创建相当于二维数组或切片的东西，需要定义一个数组的数组或切片的切片，如下所示：

```go
type Transform [3][3]float64  // 一个 3x3 的数组，实际上是一个数组的数组。
type LinesOfText [][]byte     // 一个字节切片的切片。
```

因为切片是可变长度的，所以每个内部切片都可以有不同的长度。这在许多情况下很常见，例如我们的 `LinesOfText` 示例：每行都有独立的长度。

```go
text := LinesOfText{
    []byte("Now is the time"),
    []byte("for all good gophers"),
    []byte("to bring some fun to the party."),
}
```

有时需要分配一个二维切片，例如处理像素的扫描线时就可能会遇到这种情况。实现这一点有两种方法。一种是独立分配每个切片；另一种是分配一个单一的数组，然后将各个切片指向它。使用哪种方法取决于您的应用程序。如果切片可能会增长或缩小，则应独立分配它们以避免覆盖下一行；如果不是，则使用单个分配来构造对象可能会更有效。为了参考，下面是这两种方法的概述。首先，逐行分配：

```go
// 分配顶层切片。
picture := make([][]uint8, YSize) // 每单位 y 一个行。
// 循环遍历行，为每一行分配切片。
for i := range picture {
    picture[i] = make([]uint8, XSize)
}
```

然后，作为一个分配，切成单独的行：

```go
// 分配顶层切片，与之前相同。
picture := make([][]uint8, YSize) // 每单位 y 一个行。
// 分配一个大切片来容纳所有像素。
pixels := make([]uint8, XSize*YSize) // 即使 picture 是 [][]uint8，它的类型也是 []uint8。
// 循环遍历行，从剩余像素切片的前部切出每一行。
for i := range picture {
    picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
```

### 字典maps 

字典 (Maps) 是 Go 语言内置的一种方便强大的数据结构，它将一种类型的的值 (键) 与另一种类型的的值 (元素或值) 关联起来。键可以是任何定义了相等运算符的类型，例如整数、浮点数和复数、字符串、指针、接口（只要动态类型支持相等运算）、结构体和数组。切片不能用作字典键，因为它们没有定义相等运算。

类似于切片，字典也引用底层数据结构。如果您将一个字典传递给修改了字典内容的函数，那么调用者将可以看到这些更改。

字典可以使用熟悉的复合字面量语法来构建，键值对用冒号分隔，因此很容易在初始化过程中构建它们。

```go
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
```

就像数组和切片一样，分配和获取字典值的语法看起来非常相似，只不过索引不需要是整数。

```go
offset := timeZone["EST"]
```

尝试使用不在字典中的键来获取字典值会返回该字典条目类型的零值。例如，如果字典包含整数，则查找不存在的键将返回 0。可以使用值类型为 bool 的字典来实现集合。将字典条目设置为 true 以将值放入集合中，然后通过简单索引进行测试。

```go
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // 如果 person 不在字典中，则会返回 false
    fmt.Println(person, "was at the meeting")
}
```

有时您需要区分缺失的条目和零值。是存在 "UTC" 条目还是因为它根本不在字典中而为 0？您可以使用多重赋值的形式进行区分。

```go
var seconds int
var ok bool
seconds, ok = timeZone[tz]
```

出于显而易见的原因，这被称为“逗号 ok”习惯用法。在此示例中，如果存在 tz，则会适当地设置 seconds，并且 ok 为 true；如果不存在，则 seconds 将设置为零，并且 ok 将为 false。下面是一个函数，它将它与一个漂亮的错误报告结合在一起：

```go
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```

要测试字典中是否存在某个元素而不用担心实际值，您可以使用空标识符 (_) 代替用于值的常规变量。

```go
_, present := timeZone[tz]
```

要删除字典条目，请使用内置函数 delete，其参数是字典和要删除的键。即使键已经不存在于字典中，也可以安全地执行此操作。

```go
delete(timeZone, "PDT")  // 现在是标准时间了
```



### 打印 Printing

Go 语言的格式化输出使用类似于 C 语言的 printf 家族的风格，但更加丰富和通用。这些函数位于 fmt 包中，名称以大写字母开头： fmt.Printf、 fmt.Fprintf、 fmt.Sprintf 等。字符串函数（Sprintf 等）会返回一个字符串，而不是填充提供的缓冲区。

您不必提供格式化字符串。对于 Printf、Fprintf 和 Sprintf 中的每一个，都还有另一对函数，例如 Print 和 Println。这些函数不接受格式化字符串，而是为每个参数生成默认格式。 Println 版本还会在参数之间插入一个空格并在输出后追加一个换行符，而 Print 版本仅在两侧的操作数都不是字符串时才添加空格。在这个例子中，每一行都产生相同的输出。

```go
fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))
```

格式化打印函数 fmt.Fprint 等的第一个参数接受任何实现了 io.Writer 接口的对象；os.Stdout 和 os.Stderr 都是常见的实例。

这里开始与 C 语言分道扬镳。首先，诸如 %d 的数字格式不采用符号或大小的标志；而是由打印例程根据参数的类型来决定这些属性。

```go
var x uint64 = 1<<64 - 1
fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))

// 会输出：
18446744073709551615 ffffffffffffffff; -1 -1
```

如果您只想使用默认转换，例如将整数转换为十进制，可以使用万能格式 %v（表示“值”）；结果与 Print 和 Println 生成的完全一样。此外，该格式可以打印任何值，甚至是数组、切片、结构体和字典。下面是针对上一节定义的时区字典的打印语句。

```go
fmt.Printf("%v\n", timeZone)  // 或者直接 fmt.Println(timeZone)
// 会输出
map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]
```

对于字典，Printf 等函数会按键按字典序对输出进行排序。

打印结构体时，修改后的格式 %+v 会使用字段名注释结构体的字段，对于任何值，备用格式 %#v 会使用完整的 Go 语法打印该值。

```go
type T struct {
    a int
    b float64
    c string
}
t := &T{7, -2.35, "abc\tdef"}
fmt.Printf("%v\n", t)
fmt.Printf("%+v\n", t)
fmt.Printf("%#v\n", t)
fmt.Printf("%#v\n", timeZone)

// 输出
&{7 -2.35 abc   def}
&{a:7 b:-2.35 c:abc     def}
&main.T{a:7, b:-2.35, c:"abc\tdef"}
map[string]int{"CST":-21600, "EST":-18000, "MST":-25200, "PST":-28800, "UTC":0}
```


(注意引用符号。) 那种引用字符串的格式也可以通过对类型为 string 或 []byte 的值应用 %q 来实现。备用格式 %#q 在可能的情况下会改用反引号。 ( %q 格式也适用于整数和 rune 类型，会生成单引号的 rune 常量。) 此外， %x 也适用于字符串、字节数组和字节切片，以及整数，它会生成一个长十六进制字符串，格式中带有空格 (% x) 则会在字节之间加入空格。

另一个方便的格式是 %T，它会打印值的类型。

```go
fmt.Printf("%T\n", timeZone)
// 打印： map[string]int
```

如果您想控制自定义类型的默认格式，只需在类型上定义一个签名为 String() string 的方法即可。对于我们简单的类型 T，可以这样实现。

```go
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
fmt.Printf("%v\n", t)

// 以上代码会按照如下格式打印：
7/-2.35/"abc\tdef"
```

(如果您需要同时打印 T 类型的的值和指向 T 的指针，那么 String 方法的接收者必须是值类型；这个例子使用指针是因为对于结构体类型来说这样做更有效且更符合惯用法。有关更多信息，请参见下面关于指针接收者 vs. 值接收者的部分。)

我们的 String 方法能够调用 Sprintf，因为打印例程是完全可重入的，可以这样包装使用。但是，关于这种方法，有一个重要的细节需要理解：不要通过以一种方式调用 Sprintf 来构建 String 方法，这种方式会导致您的 String 方法无限递归。如果 Sprintf 调用尝试直接将接收者打印为字符串，然后又会再次调用该方法，就会发生这种情况。这是一个常见且易犯的错误，如下例所示。

```go
type MyString string

func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", m) // 错误：将永远递归下去。
}
```

修复也很简单：将参数转换为基本字符串类型，该类型没有该方法。

```go
type MyString string

func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", string(m)) // 可以：注意类型转换。
}
```

我们将在初始化部分看到另一种避免这种递归的技术。

另一种打印技术是将打印例程的参数直接传递给另一个这样的例程。Printf 的签名为其最终参数使用类型 ...interface{}，用于指定格式之后可以出现任意数量的任意类型参数。

```go
func Printf(format string, v ...interface{}) (n int, err error) {
```

在 Printf 函数内部，v 的行为就像一个 []interface{}类型的变量，但是如果它被传递给另一个可变参数函数，它就会表现得像一个常规的参数列表。下面是我们之前使用过的 log.Println 函数的实现。它直接将它的参数传递给 fmt.Sprintln 进行实际格式化。

```go
// Println prints to the standard logger in the manner of fmt.Println.
func Println(v ...interface{}) {
    std.Output(2, fmt.Sprintln(v...))  // Output 接受 (int, string) 参数
}
```

我们在嵌套调用 Sprintln 时，在 v 后面写 ... 是为了告诉编译器将 v 作为一个参数列表来对待；否则它只会将 v 传递作为一个切片参数。

关于打印还有更多内容，这里没有涵盖。有关详细信息，请参阅 fmt 包的 godoc 文档。

顺便说一句，... 参数可以是特定的类型，例如 ...int 用于选择一组整数中最小的 min 函数：

```go
func Min(a ...int) int {
    min := int(^uint(0) >> 1)  // 最大 int
    for _, i := range a {
        if i < min {
            min = i
        }
    }
    return min
}
```

### Append函数

现在我们拥有了解 append 内置函数设计所需的关键信息。append 的签名与我们上面自定义的 Append 函数不同。概略地来说，它像这样：

```go
func append(slice []T, elements ...T) []T
```

其中 T 是任何给定类型的占位符。您实际上无法在 Go 中编写一个由调用者决定类型 T 的函数。这就是 append 是内置的原因：它需要编译器的支持。

append 所做的就是将元素追加到切片的末尾并返回结果。之所以需要返回结果，是因为和我们手工编写的 Append 一样，底层数组可能会发生变化。这个简单的例子：

```go
x := []int{1,2,3}
x = append(x, 4, 5, 6)
fmt.Println(x)
```

会输出 [1 2 3 4 5 6]。因此，append 有点类似于 Printf，它可以收集任意数量的参数。

但是，如果我们想像我们的 Append 那样将一个切片追加到另一个切片呢？很简单：在调用位置使用 ...，就像我们上面调用 Output 一样。下面的代码片段会产生与上面相同的输出。

```go
x := []int{1,2,3}
y := []int{4,5,6}
x = append(x, y...)
fmt.Println(x)
```

如果没有那个 ..., 它将无法编译，因为类型会错误；y 不是 int 类型。



## 初始化Initialization

虽然乍一看 Go 语言的初始化与 C 或 C++ 的初始化并没有太大的区别，但 Go 语言的初始化更加强大。可以在初始化过程中构建复杂的结构，并且可以正确处理初始化对象之间的顺序问题，即使是在不同的包之间。

### 常量

Go 语言中的常量就是真正的常量。它们在编译时创建，即使在函数中作为局部变量定义也是如此，并且只能是数字、字符（rune）、字符串或布尔值。由于编译时限制，定义常量的表达式必须是常量表达式，即编译器可以计算的表达式。例如，1<<3 是一个常量表达式，而 math.Sin(math.Pi/4) 不是，因为 math.Sin 的函数调用需要在运行时进行。

在 Go 语言中，可以使用枚举器 iota 创建枚举常量。由于 iota 可以是表达式的一部分，并且表达式可以隐式重复，因此很容易构建复杂的数值集合。

```go
type ByteSize float64

const (
    _           = iota // 通过赋值给空标识符来忽略第一个值
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
```

将 String 方法附加到任何用户定义的类型，可以使任意值在打印时自动格式化。虽然这种技术最常应用于结构体，但对于标量类型（如 ByteSize 等浮点类型）也很有用。

```go
func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}
```

表达式 YB 打印为 1.00YB，而 ByteSize(1e13) 打印为 9.09TB。

这里使用 Sprintf 实现 ByteSize 的 String 方法是安全的（避免了无限期重复），不是因为转换，而是因为它使用 %f 调用 Sprintf，而 %f 并非字符串格式：Sprintf 只有在需要字符串时才会调用 String 方法，而 %f 需要的是浮点数值。



### 变量

变量可以像常量一样初始化，但是初始化器可以是一般运行时计算的表达式。

```go
var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)
```

### init 函数

最后，每个源文件可以定义自己的无参的 init 函数来设置所需的状态。（实际上每个文件可以有多个 init 函数。） “最后”的意思是最终：init 函数在包中所有变量声明计算完初始化器之后才被调用，而这些初始化器只有在所有导入的包都被初始化之后才会被计算。

除了无法表示为声明的初始化之外，init 函数的一个常见用法是在实际执行开始之前验证或修复程序状态的正确性。

```go
func init() {
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath 可以被命令行的 --gopath 标志覆盖。
    flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}
```

### 指针接收者 vs. 值接收者

正如我们在 ByteSize 中看到的那样，可以为任何命名类型（除了指针或接口）定义方法；接收者不必是结构体。

在上面有关切片的讨论中，我们写了一个 Append 函数。我们可以将它定义为切片上的方法。为此，我们首先声明一个命名类型，可以将方法绑定到该类型，然后使方法的接收者成为该类型的的值。

```go
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
    // 代码与上面定义的 Append 函数完全相同。
}
```

这仍然要求方法返回更新后的切片。我们可以通过将方法重新定义为使其接收指向 ByteSlice 的指针作为其接收者来消除这种笨拙性，这样方法可以覆盖调用者的切片。

```go
func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // 代码同上，无需返回。
    *p = slice
}
```

事实上，我们做得更好。如果我们将函数修改为看起来像标准的 Write 方法，如下所示，

```go
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // 与上面一样。
    *p = slice
    return len(data), nil
}
```

那么类型 *ByteSlice 满足标准接口 io.Writer，这很方便。例如，我们可以向其中打印内容。

```go
var b ByteSlice
fmt.Fprintf(&b, "This hour has %d days\n", 7)
```

我们传递 ByteSlice 的地址，因为只有 *ByteSlice 满足 io.Writer。关于接收者的指针 vs. 值的规则是，值方法可以对指针和值进行调用，但是指针方法只能对指针进行调用。

这条规则之所以存在，是因为指针方法可以修改接收者；在值上调用它们会导致方法接收值的副本，因此任何修改都将被丢弃。因此，语言不允许这种错误。不过，也有一个方便的例外。当值可以寻址时，语言会自动插入地址运算符来处理调用值上的指针方法的常见情况。在我们的例子中，变量 b 是可以寻址的，所以我们可以只用 b.Write 调用它的 Write 方法。编译器会把它重写为 (&b).Write。

顺便说一句，在字节切片上使用 Write 的思想是 bytes.Buffer 实现的核心。



## 接口和其他类型

### 接口

Go 语言中的接口提供了一种指定对象行为的方式：如果某个东西可以做 "this"，那么它就可以用于 "here"。我们已经看到了一些简单的例子；自定义打印器可以通过 String 方法实现，而 Fprintf 可以将输出生成到任何具有 Write 方法的东西。在 Go 代码中，只有一个或两个方法的接口很常见，并且通常根据方法命名，例如实现了 Write 的 io.Writer。

一个类型可以实现多个接口。例如，一个集合可以通过 sort 包中的例程进行排序，如果它实现了 sort.Interface，其中包含 Len()、Less(i, j int) bool 和 Swap(i, j int)，它还可以拥有一个自定义格式化器。在这个人为的例子中，Sequence 同时满足这两个条件。

```go
type Sequence []int

// sort.Interface 要求的方法。
func (s Sequence) Len() int {
    return len(s)
}
func (s Sequence) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Copy 返回 Sequence 的副本。
func (s Sequence) Copy() Sequence {
    copy := make(Sequence, 0, len(s))
    return append(copy, s...)
}

// 用于打印的方法 - 在打印之前对元素进行排序。
func (s Sequence) String() string {
    s = s.Copy() // 制作一个副本；不要覆盖参数。
    sort.Sort(s)
    str := "["
    for i, elem := range s { // 循环是 O(N²)，将在下一个例子中修复。
        if i > 0 {
            str += " "
        }
        str += fmt.Sprint(elem)
    }
    return str + "]"
}
```

### 类型转换

Sequence 的 String 方法重复了 Sprint 已经为切片完成的工作。（它还有 O(N²) 的复杂度，效率低下。）如果在调用 Sprint 之前将 Sequence 转换为普通的 []int，我们可以共享工作 (并加快速度)。

```go
func (s Sequence) String() string {
    s = s.Copy()
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}
```

这个方法是另一个从 String 方法安全调用 Sprintf 的转换技巧示例。因为这两种类型 (Sequence 和 []int) 在忽略类型名的情况下是相同的，所以它们之间进行转换是合法的。转换不会创建新值，它只是临时地将现有值视为具有新类型。（还有一些其他合法的转换，例如从整数到浮点型的转换，会创建一个新值。）

在 Go 程序中，将表达式的类型转换为访问一组不同的方法是一种惯用法。例如，我们可以使用现有的类型 sort.IntSlice 将整个示例简化为以下内容：

```go
type Sequence []int

// 用于打印的方法 - 在打印之前对元素进行排序
func (s Sequence) String() string {
    s = s.Copy()
    sort.IntSlice(s).Sort()
    return fmt.Sprint([]int(s))
}
```

现在，我们没有让 Sequence 实现多个接口（排序和打印），而是利用数据项可以转换为多种类型 (Sequence、sort.IntSlice 和 []int) 的能力，每种类型都完成一部分工作。这在实践中并不常见，但可能是有效的。

### 接口转换和类型断言

类型转换是一种转换形式：它们接受一个接口，并且在 switch 的每个 case 中，从某种意义上将其转换为该 case 的类型。下面是 fmt.Printf 下面的代码如何使用类型转换将值转换为字符串的简化版本。如果它已经是字符串，我们想要接口持有的实际字符串值，而如果它有一个 String 方法，我们想要调用该方法的结果。

```go
type Stringer interface {
    String() string
}

var value interface{} // 调用者提供的 Value。
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```

第一种情况会找到一个具体的值；第二种情况会将接口转换为另一个接口。以这种方式混合类型完全没问题。

如果我们只关心一种类型怎么办？如果我们知道该值包含一个字符串，我们只想提取它呢？使用单案例的类型转换也可以，但类型断言也可以。类型断言接受一个接口值，并从中提取一个指定显式类型的值。语法借用了打开类型转换的子句，但使用了显式类型而不是 type 关键字：

```go
value.(typeName)
```

结果是一个具有静态类型 typeName 的新值。该类型要么是接口持有的具体类型，要么是该值可以转换为的第二个接口类型。为了提取我们知道存在于 value 中的字符串，我们可以写成：

```go
str := value.(string)
```

但是，如果该值不包含字符串，程序将崩溃并出现运行时错误。为了防止这种情况，可以使用 "comma, ok" 惯用法来安全地测试该值是否为字符串：

```go
str, ok := value.(string)
if ok {
fmt.Printf("string value is: %q\n", str)
} else {
fmt.Printf("value is not a string\n")
}
```

如果类型断言失败，str 仍然存在并且是 string 类型，但它将具有零值，即空字符串。

作为功能演示，这里有一个 if-else 语句，等同于打开此部分的类型转换。

```go
if str, ok := value.(string); ok {
	return str
} else if str, ok := value.(Stringer); ok {
	return str.String()
}
```

### 通用性Generality

如果一个类型仅用于实现一个接口，并且除了该接口之外不会有导出的方法，那么就没有必要导出类型本身。只导出接口可以清楚地表明该值除了接口中描述的内容之外没有任何有趣的行为。它还避免了在每个常见方法的实例上重复文档的需要。

在这种情况下，构造函数应该返回接口值而不是实现类型。例如，在哈希库中，crc32.NewIEEE 和 adler32.New 都返回接口类型 hash.Hash32。在 Go 程序中将 CRC-32 算法替换为 Adler-32 只需要更改构造函数调用；代码的其余部分不受算法更改的影响。

类似的方法允许将各种加密包中的流式密码算法与它们链接在一起的块密码分开。crypto/cipher 包中的 Block 接口指定了块密码的行为，该密码提供单个数据块的加密。然后，通过与 bufio 包类似的方式，实现此接口的密码包可用于构建流式密码（由 Stream 接口表示），而无需了解块加密的细节。

crypto/cipher 的接口如下所示：

```go
type Block interface {
    BlockSize() int
    Encrypt(dst, src []byte)
    Decrypt(dst, src []byte)
}

type Stream interface {
    XORKeyStream(dst, src []byte)
}
```

下面是计数器模式 (CTR) 流的定义，它将块密码转换为流式密码；请注意块密码的细节被抽象掉了：

> // NewCTR 返回一个 Stream，它使用给定的 Block 在计数器模式下进行加密/解密。iv 的长度必须与 Block 的块大小相同。 func NewCTR(block Block, iv []byte) Stream
>
> NewCTR 不仅适用于一种特定的加密算法和数据源，而且适用于 Block 接口的任何实现和任何 Stream。因为它们返回接口值，所以用其他加密模式替换 CTR 加密是一个局部更改。必须编辑构造函数调用，但由于周围代码只需要将结果视为 Stream，因此它不会察觉到差异。

### 接口和方法

由于几乎任何东西都可以附加方法，因此几乎任何东西都可以满足一个接口。一个说明性的例子是在 http 包中，它定义了 Handler 接口。任何实现 Handler 的对象都可以处理 HTTP 请求。

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

ResponseWriter 本身也是一个接口，它提供访问向客户端返回响应所需的方法。这些方法包括标准的 Write 方法，因此 http.ResponseWriter 可以用于任何可以使用 io.Writer 的地方。Request 是一个包含来自客户端的请求的解析表示的结构体。

为了简洁起见，让我们忽略 POST 请求，假设 HTTP 请求总是 GET 请求；这种简化不会影响处理程序的设置方式。下面是一个简单的处理程序实现，用于计算页面被访问的次数。

```go
// 简单计数器服务器。
type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}
```

（延续我们的主题，请注意 Fprintf 如何打印到 http.ResponseWriter。）在真正的服务器中，需要保护 ctr.n 访问并发访问。有关建议，请参阅 sync 和 atomic 包。

为了参考，下面是如何将这样的服务器附加到 URL 树上的节点。

```go
import "net/http"
...
ctr := new(Counter)
http.Handle("/counter", ctr)
```

但是，为什么要把 Counter 做成一个结构体呢？只需要一个整数就够了。（接收者需要是一个指针，这样调用者才能看到增量。）

```go
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}
```

如果您的程序有一些内部状态需要被通知页面已被访问，请将channel绑定到网页。

```go
// 一个频道，在每次访问时发送通知。(可能希望频道被缓冲。)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprint(w, "notification sent")
}
```

最后，假设我们想要在 /args 上显示调用服务器二进制文件时使用的参数。编写一个函数来打印参数很容易。

```go
func ArgServer() {
	fmt.Println(os.Args)
}
```

我们如何将其转换成一个 HTTP 服务器呢？我们可以将 ArgServer 作为某种类型的函数来使用，忽略其值，但是有一种更简洁的方法。由于我们可以为除指针和接口之外的任何类型定义方法，因此我们可以为函数编写方法。http 包包含以下代码：

```go
// HandlerFunc 类型是一个适配器，允许将普通函数用作 HTTP 处理程序。如果 f 是具有适当签名的函数，
// 则 HandlerFunc(f) 是一个调用 f 的 Handler 对象。
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP 调用 f(w, req)。
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
	f(f, req)
}
```

HandlerFunc 是一个具有方法 ServeHTTP 的类型，因此该类型的的值可以用于处理 HTTP 请求。看看方法的实现：接收者是一个函数 f，方法调用 f。这看起来可能很奇怪，但它与接收者是频道并且方法在频道上发送数据没什么不同。

为了将 ArgServer 变成一个 HTTP 服务器，我们首先修改它以使其具有正确的签名。

```go
// 参数服务器。
func ArgServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, os.Args)
}
```

ArgServer 现在具有与 HandlerFunc 相同的签名，因此它可以转换为该类型以访问其方法，就像我们将 Sequence 转换为 IntSlice 以访问 IntSlice.Sort 一样。设置它的代码很简洁：

```go
http.Handle("/args", http.HandlerFunc(ArgServer))
```

当有人访问页面 /args 时，安装在该页面的处理程序的值为 ArgServer，类型为 HandlerFunc。HTTP 服务器将调用该类型的 ServeHTTP 方法，以 ArgServer 为接收者，后者将反过来调用 ArgServer（通过 HandlerFunc.ServeHTTP 内部 的 f(w, req) 调用）。然后将显示参数。

在本节中，我们从结构体、整数、频道和函数创建了一个 HTTP 服务器，这都是因为接口只是方法集，几乎可以为任何类型定义这些方法集。

## 空白标识符 The blank identifier

我们之前已经讨论了几次空白标识符，都是在 for 循环和 map 的上下文中。空白标识符可以被赋值或声明任何类型和值的变量，而该值会被丢弃。这有点像写入 Unix 的 /dev/null 文件：它表示一个只写值，用作占位符，在需要变量但实际值无关紧要的地方使用。除了我们已经看到的使用之外，它还有其他用途。

### 多重赋值中的空白标识符

在 for 循环中使用空白标识符是一种特殊情况，它属于一般情况：多重赋值。

如果一个赋值需要左侧有多个值，但是程序不会使用其中一个值，那么在赋值的左侧使用一个空白标识符就可以避免创建哑变量，并清楚地表明该值应该被丢弃。例如，当调用一个返回一个值和一个错误的函数时，但只有错误很重要，可以使用空白标识符来丢弃无关的返回值。

```go
if _, err := os.Stat(path); os.IsNotExist(err) {
    fmt.Printf("%s does not exist\n", path)
}
```

偶尔你会看到丢弃错误值以忽略错误的代码；这是一个非常糟糕的习惯。总是要检查错误返回值，它们是有原因提供的。

```go
// 很糟糕！如果 path 不存在，此代码将崩溃。
fi, _ := os.Stat(path)
if fi.IsDir() {
	fmt.Printf("%s is a directory\n", path)
}
```

### 未使用的导入和变量

导入包或声明变量而不使用它们会产生错误。未使用的导入会使程序臃肿并减慢编译速度，而初始化但未使用的变量至少会浪费计算资源，并且可能表示更大的错误。然而，当程序正处于积极开发阶段时，经常会出现未使用的导入和变量，为了让编译继续进行而删除它们可能会很烦人，但稍后又可能再次需要它们。空白标识符提供了一种解决方法。

这个半写的程序有两个未使用的导入 (fmt 和 io) 和一个未使用的变量 (fd)，因此它将无法编译，但是查看迄今为止的代码是否正确会很好。

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
}
```

为了消除有关未使用导入的panic，可以使用空白标识符来引用导入包中的符号。类似地，将未使用的变量 fd 赋值给空白标识符将消除未使用变量的错误。这个版本的程序可以编译。

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf // 用于调试；完成后删除。
var _ io.Reader    // 用于调试；完成后删除。

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}
```

根据惯例，为了便于查找和提醒稍后清理，用于消除导入错误的全局声明应该紧跟在导入之后并进行注释。

### 导入的副作用Import for side effect

前一个例子中像 fmt 或 io 这样的未使用导入最终应该被使用或删除：空白赋值将代码标识为正在进行中的工作。但有时仅为其副作用导入包而没有任何显式使用也是很有用的。例如，在 net/http/pprof 包的 init 函数期间，它会注册提供调试信息的 HTTP 处理程序。它具有导出的 API，但大多数客户端只需要处理程序注册并通过网页访问数据。要仅为其副作用而导入包，请将包重命名为空白标识符：

```go
import _ "net/http/pprof"
```

这种形式的导入清楚地表明该包是为了其副作用而导入的，因为该包没有其他可能的用途：在这个文件中，它没有名称。（如果它确实有名称，并且我们没有使用该名称，编译器将拒绝该程序。）

### 接口检查

正如我们在上面关于接口的讨论中看到的，类型不需要显式声明它实现了某个接口。相反，只需实现接口的方法，类型就可以实现该接口。在实践中，大多数接口转换都是静态的，因此会在编译时进行检查。例如，将 *os.File 传递给期望 io.Reader 的函数不会编译，除非 *os.File 实现 io.Reader 接口。

不过，有些接口检查确实发生在运行时。一个例子是 encoding/json 包，它定义了一个 Marshaler 接口。当 JSON 编码器接收到实现该接口的值时，编码器会调用值自身的 marshaling 方法将其转换为 JSON，而不是进行标准转换。编码器在运行时使用类型断言来检查此属性，例如：

```go
m, ok := val.(json.Marshaler)
```

如果只需要询问类型是否实现了接口，而不实际使用接口本身（例如作为错误检查的一部分），可以使用空白标识符来忽略类型断言的值：

```go
if _, ok := val.(json.Marshaler); ok {
    fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
}
```

这种情况出现的一个地方原因是，在实现类型的包中，需要保证它确实满足了接口的要求。如果一个类型（例如 json.RawMessage）需要自定义的 JSON 表示，它应该实现 json.Marshaler，但是没有静态转换会使编译器自动验证这一点。如果类型意外地无法满足接口，JSON 编码器仍然可以工作，但不会使用自定义实现。为了保证实现正确，可以在包中使用带有空白标识符的全局声明：

```go
var _ json.Marshaler = (*RawMessage)(nil)
```

在这个声明中，涉及将 *RawMessage 转换为 Marshaler 的赋值要求 *RawMessage 实现 Marshaler，并且该属性将在编译时进行检查。如果 json.Marshaler 接口发生变化，这个包将无法再编译，并且我们会注意到它需要更新。

此结构中空白标识符的存在表示声明仅用于类型检查，而不是创建变量。不过，不要对每个满足接口的类型都这样做。根据惯例，此类声明仅在代码中不存在静态转换时才使用，这种情况很少见。

## 嵌入 Embedding

Go 语言没有典型面向类型的子类继承概念，但是它可以通过在结构体或接口内嵌入类型来“借用”实现的一部分。

接口嵌入非常简单。我们之前提到了 io.Reader 和 io.Writer 接口，下面是它们的定义：

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

io 包还导出了几个其他接口，用于指定可以实现几种此类方法的对象。例如，io.ReadWriter 是一个包含 Read 和 Write 的接口。我们可以通过显式列出这两种方法来指定 io.ReadWriter，但是通过嵌入这两个接口来形成新的接口更简单、更直观，如下所示：

```go
// ReadWriter 是组合了 Reader 和 Writer 接口的接口。
type ReadWriter interface {
    Reader
    Writer
}
```

这正像它看起来一样：ReadWriter 可以像 Reader 那样做，也可以像 Writer 那样做；它是嵌入接口的联合。只有接口可以嵌入到接口中。

同样的基本概念也适用于结构体，但影响更深远。bufio 包具有两种结构体类型，bufio.Reader 和 bufio.Writer，它们当然都实现了来自 io 包的类似接口。bufio 还实现了缓冲的读写器，它通过使用嵌入将一个读取器和一个写入器组合成一个结构体来实现：它在结构体中列出类型，但不给它们字段名。

```go
// ReadWriter 存储指向 Reader 和 Writer 的指针。
// 它实现了 io.ReadWriter。
type ReadWriter struct {
    *Reader  // *bufio.Reader
    *Writer  // *bufio.Writer
}
```

嵌入的元素是结构体的指针，当然在使用之前必须初始化为指向有效的结构体。ReadWriter 结构体可以写成：

```go
type ReadWriter struct {
    reader *Reader
    writer *Writer
}
```

但是为了提升字段的方法并满足 io 接口，我们还需要提供转发方法，如下所示：

```go
func (rw *ReadWriter) Read(p []byte) (n int, err error) {
    return rw.reader.Read(p)
}
```

通过直接嵌入结构体，我们避免了这种繁琐的记录。嵌入类型的的方法会免费提供，这意味着 bufio.ReadWriter 不仅具有 bufio.Reader 和 bufio.Writer 的方法，而且还满足所有三个接口：io.Reader、io.Writer 和 io.ReadWriter。

嵌入与子类化的区别之处很重要。当我们嵌入一个类型时，该类型的函数变成了外部类型的函数，但是当它们被调用时，方法的接收者是内部类型，而不是外部类型。在我们的例子中，当调用 bufio.ReadWriter 的 Read 方法时，它与上面写出的转发方法具有完全相同的效果；接收者是 ReadWriter 的 reader 字段，而不是 ReadWriter 本身。

嵌入也可以很简单地带来便利。此示例展示了一个嵌入字段和一个常规命名字段。

```go
type Job struct {
    Command string
    *log.Logger
}
```

Job 类型现在具有 *log.Logger 的 Print、Printf、Println 等方法。当然，我们可以给 Logger 指定一个字段名，但这并不是必须的。现在，一旦初始化，我们就可以记录到 Job 中：

```go
job.Println("starting now...")
```

Logger 是 Job 结构体的常规字段，因此我们可以像往常一样在 Job 的构造函数中对其进行初始化，例如：

```go
func NewJob(command string, logger *log.Logger) *Job {
    return &Job{command, logger}
}
```

或者使用复合字面量：

```go
job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
```

如果我们需要直接引用嵌入字段，则字段的类型名（忽略包限定符）用作字段名，就像我们在 ReadWriter 结构体的 Read 方法中所做的那样。这里，如果我们需要访问 Job 变量 job 的 *log.Logger，我们可以写成 job.Logger，这在想要细化 Logger 的方法时很有用。

```go
func (job *Job) Printf(format string, args ...interface{}) {
    job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}
```

嵌入类型会引入名称冲突问题，但解决这些问题的规则很简单。首先，字段或方法 X 会隐藏类型更深层嵌套部分中的任何其他项目 X。如果 log.Logger 包含一个名为 Command 的字段或方法，那么 Job 的 Command 字段将支配它。

其次，如果相同的名字出现在相同的嵌套级别，则通常是错误的；如果 Job 结构体包含另一个名为 Logger 的字段或方法，则嵌入 log.Logger 会出错。但是，如果程序中从未在类型定义之外提到重复的名称，则也可以。此限定条件可以防止对外部嵌入的类型进行更改；如果从未使用过任何一个字段，那么即使添加了一个字段与另一个子类型中的另一个字段冲突也没有问题。

## 并发Concurrency

## Go 语言并发编程

## 通过通信共享

并发编程是一个庞大的主题，这里只介绍 Go 语言的一些特定要点。

在许多环境中，并发编程因实现正确访问共享变量所需的微妙之处而变得困难。Go 鼓励采用一种不同的方法，在这种方法中，共享值通过 channel 传递，并且实际上从未由单独的执行线程主动共享。任何给定时刻，只有一个 goroutine 可以访问该值。数据竞争根本不会发生，这是设计使然。为了鼓励这种思维方式，我们将其简化为一句口号：

> Do not communicate by sharing memory;instead, share memory by communicating.
>
> 不要通过共享内存进行通信；而是通过通信共享内存。

这种方法可能会被滥用。例如，引用计数最好通过将互斥锁放在整数变量周围来完成。但是，作为一种高级方法，使用 channel 来控制访问可以更轻松地编写清晰、正确的程序。

思考这个模型的一种方式是考虑一个典型的单线程程序，它在一个 CPU 上运行。它不需要同步原语。现在再运行另一个这样的实例；它也同样不需要同步。现在让这两个实例进行通信；如果通信是同步器，那么仍然不需要其他同步。例如，Unix 管道完美符合这个模型。虽然 Go 的并发方法起源于 Hoare 的 Communicating Sequential Processes (CSP)，但它也可以被视为 Unix 管道的类型安全泛化。

### 协程Goroutine

它们被称为 goroutine，因为现有术语（线程、协程、进程等）会带来不准确的含义。goroutine 具有一个简单的模型：它是一个与同一地址空间中的其他 goroutine 并发执行的函数。它非常轻量级，只需要分配一点点栈空间。而且它们的栈空间从很小开始，因此开销很低，并且会根据需要通过分配（和释放）堆存储器来增长。

Goroutine 会被多路复用(multiplexed)到多个操作系统线程上，因此如果一个 goroutine 阻塞（例如等待 I/O），其他 goroutine 仍会继续运行。它们的設計隐藏了线程创建和管理的许多复杂性。

在函数或方法调用之前加上 go 关键字可以在新的 goroutine 中运行该调用。当调用完成时，goroutine 会静默退出。（效果类似于 Unix shell 的 & 符号，用于在后台运行命令。）

```go
go list.Sort()  // 并发运行 list.Sort；不要等待它完成。
```

函数字面量在 goroutine 调用中非常方便。

```go
func Announce(message string, delay time.Duration) {
    go func() {  // 注意括号 - 必须调用该函数。
        time.Sleep(delay)
        fmt.Println(message)
    }()
}
```

在 Go 语言中，函数字面量(function literals)是闭包(closures)：其实现确保函数引用的变量只要处于活动状态就会一直存在。

这些例子并不是非常实用，因为函数没有办法通知完成。为此，我们需要使用 channel。

### 管道channels

与 map 一样，channel 也使用 make 分配， resulting value 充当底层数据结构的引用。如果提供了可选的整型参数，它会设置 channel 的缓冲区大小。默认值为零，表示非缓冲或同步 channel。

```go
ci := make(chan int)            // 无缓冲的整数 channel
cj := make(chan int, 0)         // 无缓冲的整数 channel
cs := make(chan *os.File, 100)  // 缓冲的指向文件的指针 channel
```

非缓冲 channel 将通信（值交换）与同步（保证两个计算（goroutine）处于已知状态）结合在一起。

使用 channel 有很多很好的习惯用法。这里有一个例子可以帮助我们入门。在前一部分，我们在后台启动了一个排序。channel 可以让启动 goroutine 等待排序完成。

```go
c := make(chan int)  // 分配一个 channel。
// 在 goroutine 中启动排序；完成后，通过 channel 发送信号。
go func() {
    list.Sort()
    c <- 1  // 发送信号；值无关紧要。
}()
doSomethingForAWhile()
<-c   // 等待排序完成；丢弃发送的值。
```

*接收者总是会阻塞，直到有数据可接收。如果 channel 是非缓冲的，则发送者会阻塞，直到接收者接收了该值*。如果 channel 有缓冲区，则发送者只会阻塞到该值被复制到缓冲区为止；如果缓冲区已满，则意味着要等到某个接收者检索到一个值才会继续。

缓冲 channel 可以用作信号量，例如限制吞吐量。在这个例子中，传入的请求被传递给 handle 函数，该函数将一个值发送到 channel 中，处理请求，然后从 channel 中接收一个值，为下一个使用者准备好“信号量”。channel 缓冲区的容量限制了可以同时调用 process 的次数。

```go
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    // 等待活动队列清空。
    process(r)  // 可能需要很长时间。
    <-sem       // 完成；启用下一个请求运行。
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req)  // 不要等待 handle 完成。
    }
}
```

一旦有 MaxOutstanding 个处理程序正在执行 process，任何额外的处理程序都将因尝试向已满的 channel 缓冲区发送数据而阻塞，直到现有处理程序之一完成并从缓冲区接收数据为止。

这个设计存在一个问题:Serve 为每个传入的请求都创建了一个新的 goroutine,尽管同一时刻只有 MaxOutstanding 个 goroutine 可以运行。结果是,如果请求来得太快,程序可能会无限消耗资源。我们可以通过修改 Serve 来限制 goroutine 的创建来解决这个问题。以下是一个明显的解决方案,但需要注意它存在一个我们后续会解决的 bug:

```go
func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func() {
            process(req) // 有 bug,详见下面的解释
            <-sem
        }()
    }
}
```

这个 bug 在于,在 Go 的 for 循环中,循环变量 req 在每次迭代时都会被重用,因此所有的 goroutine 都共享同一个 req 变量。这不是我们想要的。我们需要确保每个 goroutine 都有自己唯一的 req 变量。以下是一种解决方法,将 req 的值作为参数传递给 goroutine 中的闭包:

```go
func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func(req *Request) {
            process(req)
            <-sem
        }(req)
    }
}
```

将这个版本与前一个版本进行比较,可以看到闭包的声明和运行方式有所不同。另一种解决方法是简单地创建一个与循环变量同名的新变量,如下所示:

```go
func Serve(queue chan *Request) {
    for req := range queue {
        req := req // 为 goroutine 创建 req 的新实例
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}
```

在 Go 中,写`req := req`虽然看起来有些奇怪,但是这是一种常见且合法的做法。这样可以创建一个与循环变量同名但独立于每个 goroutine 的新变量。

回到编写服务器的一般问题,另一种管理资源的方法是启动一个固定数量的 handle goroutine,它们都从请求队列中读取请求。goroutine 的数量限制了同时调用 process 的数量。这个 Serve 函数还接受一个用于退出的信号通道。

```go
func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    // 启动处理器
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }

    <-quit // 等待退出信号
}
```

### channel 的 channel

Go 语言的一个重要特性是 channel 是一等值，可以像其他值一样进行分配和传递。这种特性的常见用法之一是实现安全并行的多路复用(demultiplexing)。

在前一部分的示例中，handle 是一个处理请求的理想化处理程序，但我们没有定义它处理的类型。如果该类型包含用于回复的 channel，则每个客户端都可以为答案提供自己的路径。下面是 Request 类型的示意图定义。

```go
type Request struct {
    args        []int
    f           func([]int) int
    resultChan  chan int
}
```

客户端提供一个函数及其参数，以及一个位于请求对象内部用于接收答案的 channel。

```go
func sum(a []int) (s int) {
    for _, v := range a {
        s += v
    }
    return
}

request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
// 发送请求
clientRequests <- request
// 等待响应
fmt.Printf("answer: %d\n", <-request.resultChan)
```

在服务器端，只有处理程序函数会发生变化。

```go
func handle(queue chan *Request) {
    for req := range queue {
        req.resultChan <- req.f(req.args)
    }
}
```

这显然还有很多需要改进的地方才能使其更加现实，但这段代码为限速、并行、非阻塞的 RPC 系统提供了一个框架，并且看不到任何互斥锁。

### 并行化

这些思想的另一个应用是跨多个 CPU 内核并行计算。如果计算可以分解成独立执行的单独部分，则可以并行化，并使用 channel 来表示每个部分完成时的信号。

假设我们对一个项目向量执行昂贵的操作，并且每个项目的操作值都是独立的，如下面的理想化示例所示。

```go
type Vector []float64

// 将操作应用于 v[i], v[i+1] ... 直到 v[n-1]。
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
    for ; i < n; i++ {
        v[i] += u.Op(v[i])
    }
    c <- 1    // 信号此部分完成
}
```

我们在循环中独立启动各个部分，每个 CPU 一个。它们可以按任何顺序完成，这没关系；我们只需在启动所有 goroutine 后通过清空 channel 来计算完成信号。

```go
const numCPU = 4 // CPU 核心数量
```

```go
func (v Vector) DoAll(u Vector) {
    c := make(chan int, numCPU)  // 可选缓冲，但合理。
    for i := 0; i < numCPU; i++ {
        go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
    }
    // 清空 channel。
    for i := 0; i < numCPU; i++ {
        <-c    // 等待一个任务完成
    }
    // 全部完成。
}
```

我们可以不创建 numCPU 的常量值，而是询问运行时哪个值合适。函数 runtime.NumCPU 返回机器中硬件 CPU 核心的数量，因此我们可以写成：

```go
var numCPU = runtime.NumCPU()
```

还有一个函数 runtime.GOMAXPROCS，它报告（或设置）用户指定的 Go 程序可以同时运行的内核数量。它默认为 runtime.NumCPU 的值，但可以通过设置同名 shell 环境变量或使用带正数的函数来覆盖。用零调用它只会查询值。因此，如果我们想要尊重用户的资源请求，我们应该写成：

```go
var numCPU = runtime.GOMAXPROCS(0)
```

一定要注意并发（将程序结构为独立执行的组件）和并行（在多个 CPU 上并行执行计算以提高效率）的概念区别。虽然 Go 的并发特性可以使一些问题易于结构化为并行计算，但 Go 是一门并发语言，而不是并行语言，并不是所有并行化问题都适合 Go 的模型。有关区别的讨论，请参阅[这篇博文](https://go.dev/blog/concurrency-is-not-parallelism)中引用的演讲。

### Go 语言 - 泄漏缓存区

并发编程的工具甚至可以使非并行化的想法更容易表达。这里是一个来自 RPC 包的抽象示例。客户端 goroutine 循环从某个源（例如网络）接收数据。为了避免分配和释放缓冲区，它保留了一个空闲列表，并使用缓冲的 channel 来表示它。如果 channel 为空，则分配一个新的缓冲区。一旦消息缓冲区准备就绪，它就会通过 serverChan 发送到服务器。

```go
var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
    for {
        var b *Buffer
        // 如果可用，则获取一个缓冲区；如果没有，则分配一个。
        select {
        case b = <-freeList:
            // 获取了一个；无需进一步操作。
        default:
            // 没有空闲的，因此分配一个新的。
            b = new(Buffer)
        }
        load(b)              // 从网络读取下一个消息。
        serverChan <- b      // 发送到服务器。
    }
}
```

服务器循环接收来自客户端的每个消息，对其进行处理，然后将缓冲区返回给空闲列表。

```go
func server() {
    for {
        b := <-serverChan    // 等待工作。
        process(b)
        // 如果有空间，则重用缓冲区。
        select {
        case freeList <- b:
            // 缓冲区在空闲列表中；无需进一步操作。
        default:
            // 空闲列表已满，只需继续即可。
        }
    }
}
```

客户端尝试从 freeList 检索缓冲区；如果没有可用缓冲区，则分配一个新的。服务器发送到 freeList 会将 b 放回空闲列表，除非列表已满，在这种情况下，缓冲区将被丢弃到内存中由垃圾回收器回收。（select 语句中的 default 子句在没有其他 case 准备就绪时执行，这意味着 select 语句永远不会阻塞。）这种实现仅使用几行代码就构建了一个泄漏桶式空闲列表，依靠缓冲的 channel 和垃圾回收器进行记录。

## Errors

库函数(Library routines)通常需要向调用者返回某种错误指示。如前所述,Go 的多值返回使得可以在返回正常值的同时返回一个详细的错误描述。使用这个功能来提供详细的错误信息是一个很好的做法。例如,如你将看到的,os.Open 不只是在失败时返回一个 nil 指针,它还会返回一个描述发生了什么问题的错误值。

按照惯例,错误有一个简单的内置接口类型 error:

```go
type error interface {
    Error() string
}
```

库的编写者可以实现这个接口,在内部使用一个更丰富的模型,不仅可以看到错误,还可以提供一些上下文信息。如前所述,除了通常的 *os.File 返回值,os.Open 还会返回一个错误值。如果文件成功打开,错误值将为 nil,但在出现问题时,它将包含一个 os.PathError:

```go
// PathError 记录了一个错误,以及导致该错误的操作和文件路径。
type PathError struct {
    Op   string // "open", "unlink" 等
    Path string // 相关文件
    Err  error  // 系统调用返回的错误
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```

PathError 的 Error 方法生成一个类似这样的字符串:

```
open /etc/passwx: no such file or directory
```

这样的错误信息,包含了问题文件名、操作和触发的操作系统错误,即使在远离引起错误的调用处打印,也是很有用的,比简单的"no such file or directory"更具信息性。

当可行时,错误字符串应该标明它们的来源,比如有一个前缀来标识生成错误的操作或包。例如,在 image 包中,由于未知格式导致解码错误的字符串表示为"image: unknown format"。

关心具体错误细节的调用者可以使用类型转换或类型断言来查找特定的错误,并提取详细信息。对于 PathError,这可能包括检查内部的 Err 字段,以获取可恢复的失败。

```go
for try := 0; try < 2; try++ {
    file, err = os.Create(filename)
    if err == nil {
        return
    }
    if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
        deleteTempFiles() // 回收一些空间
        continue
    }
    return
}
```

这里的第二个 if 语句是另一个类型断言。如果失败,ok 将为 false,e 将为 nil。如果成功,ok 将为 true,这意味着错误的类型是 *os.PathError,此时 e 也是这种类型,我们可以对其进行更深入的检查。

### Panic

通常向调用者报告错误的方式是将 `error` 作为额外的返回值。著名的 `Read` 方法就是一个很好的例子,它返回一个字节计数和一个 `error`。但是如果错误是不可恢复的呢?有时程序根本无法继续运行下去。

为此,Go 提供了一个内置函数 `panic`,它会在运行时创建一个错误,从而停止程序的执行(但请参见下一节)。该函数接受一个任意类型的参数(通常是一个字符串)作为程序崩溃时要打印的信息。它也是一种表示发生了不可能事件的方式,比如退出一个无限循环。

```go
// 使用牛顿法实现立方根的一个玩具版本。
func CubeRoot(x float64) float64 {
    z := x/3   // 任意初始值
    for i := 0; i < 1e6; i++ {
        prevz := z
        z -= (z*z*z-x) / (3*z*z)
        if veryClose(z, prevz) {
            return z
        }
    }
    // 经过一百万次迭代仍未收敛,说明出现了问题。
    panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}
```

这只是一个例子,但是真正的库函数应该避免使用 `panic`。如果问题可以被掩盖或者绕过,那么总是让程序继续运行要比直接终止整个程序要好。一个可能的反例是在初始化期间:如果库确实无法设置自己,那么使用 `panic` 可能是合理的。

```go
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```

### Recover

当 `panic` 被调用时,包括因为下标越界或类型断言失败而隐式引发的运行时错误,它会立即停止当前函数的执行,并开始解开该 goroutine 的调用栈,在此过程中会运行任何延迟执行的函数。如果这种解开过程一直到达 goroutine 栈的顶部,程序就会崩溃。但是,可以使用内置函数 `recover` 来重新获得对 goroutine 的控制并恢复正常执行。

对 `recover` 的调用会停止解开过程,并返回传递给 `panic` 的参数。由于在解开过程中只有延迟执行的函数会运行,因此 `recover` 只有在延迟函数中才有用。

`recover` 的一个应用场景是在服务器内关闭失败的 goroutine,而不会杀死其他正在执行的 goroutine。

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

在这个例子中,如果 `do(work)` 引发了 panic,结果会被记录下来,而 goroutine 会干净地退出,不会影响其他 goroutine。在延迟的函数闭包中不需要做任何其他事情,调用 `recover` 就可以完全处理这种情况。

因为除非直接从延迟函数中调用,否则 `recover` 总是返回 `nil`,所以延迟的代码可以调用自身使用 `panic` 和 `recover` 的库函数而不会失败。例如,`safelyDo` 中的延迟函数可能会在调用 `recover` 之前调用一个日志函数,而这个日志代码会在 panic 状态下照常运行。

有了这种恢复模式,`do` 函数(及其调用的任何内容)都可以通过调用 `panic` 来干净地摆脱任何糟糕的情况。我们可以利用这个想法来简化复杂软件中的错误处理。让我们来看一个理想化的 `regexp` 包,它通过调用 `panic` 来报告解析错误,使用一个局部的错误类型。下面是 `Error` 的定义、`error` 方法和 `Compile` 函数。

```go
// Error is the type of a parse error; it satisfies the error interface.
type Error string
func (e Error) Error() string {
    return string(e)
}

// error is a method of *Regexp that reports parsing errors by
// panicking with an Error.
func (regexp *Regexp) error(err string) {
    panic(Error(err))
}

// Compile returns a parsed representation of the regular expression.
func Compile(str string) (regexp *Regexp, err error) {
    regexp = new(Regexp)
    // doParse will panic if there is a parse error.
    defer func() {
        if e := recover(); e != nil {
            regexp = nil    // Clear return value.
            err = e.(Error) // Will re-panic if not a parse error.
        }
    }()
    return regexp.doParse(str), nil
}
```



当 doParse panic 时，恢复块会将返回值设置为 nil - deferred 函数可以修改命名返回值。然后，它将在分配给 err 的过程中检查问题是否是解析错误，通过断言它具有本地类型 Error 来进行检查。如果它不是，类型断言将失败，导致运行时错误，该错误会继续堆栈展开，就好像没有任何中断一样。此检查意味着，即使我们使用 panic 和 recover 来处理解析错误，如果发生某些意外情况（例如索引超出范围），代码也会失败。

有了错误处理机制，error 方法（因为它是一个绑定到类型的函数，因此它与内置的 error 类型同名是完全可以的，甚至可以说是自然的）可以轻松地报告解析错误，而无需担心手动展开解析堆栈：

```go
if pos == 0 {
    re.error("'*' illegal at start of expression")
}
```

这种模式虽然有用，但应该只在包内使用。Parse 函数会将其内部的 panic 调用转换为 error 值；它不会将 panic 暴露给其客户端。这是一个很好的规则。

顺便说一下，这种重新 panic 的惯用法会改变实际发生错误时的 panic 值。但是，崩溃报告中会同时显示原始错误和新错误，因此问题的根源仍然可见。因此，这种简单的重新 panic 方法通常就足够了 - 毕竟它是一个崩溃。但是，如果您只想显示原始值，可以编写更多代码来过滤意外问题并使用原始错误重新 panic。这留给读者作为练习。



## 网络服务器Web sserver

让我们用一个完整的Go程序,一个网络服务器来结束。这个程序实际上是一种网络重定向服务器。Google在`chart.apis.google.com`提供了一项自动将数据格式化为图表和图形的服务。但是这个服务很难交互使用,因为你需要将数据作为查询参数放在URL中。这个程序为一种数据提供了一个更好的接口:给定一小段文本,它会调用图表服务器生成一个二维码,一个编码这段文本的方格矩阵。这个图像可以用你手机的相机抓取,并解释为一个URL,从而避免在手机小键盘上输入URL。

下面是完整的程序与解释：

```go
package main

import (
    "flag"
    "html/template"
    "log"
    "net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
    flag.Parse()
    http.Handle("/", http.HandlerFunc(QR))
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func QR(w http.ResponseWriter, req *http.Request) {
    templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`
```

到 `main` 函数为止的部分应该很容易理解。唯一的一个标志设置了我们服务器的默认HTTP端口。模板变量 `templ` 是关键所在。它构建了一个HTML模板,将由服务器执行以显示页面;我们稍后会详细介绍。

`main` 函数解析标志,并使用我们之前讨论过的机制将函数 `QR` 绑定到服务器的根路径。然后调用 `http.ListenAndServe` 启动服务器,它会一直阻塞直到服务器运行结束。

`QR` 只是接收包含表单数据的请求,并在名为 `s` 的表单值上执行模板。

`html/template` 包非常强大;这个程序只是触及了它的部分功能。本质上,它通过替换传递给 `templ.Execute` 的数据项派生的元素,实时重写一段HTML文本。在模板文本(`templateStr`)中,用双大括号括起来的部分表示模板操作。从 `{{if .}}` 到 `{{end}}` 的部分只有在当前数据项(称为 `.`)的值非空时才会执行。也就是说,当字符串为空时,这部分模板将被抑制。

两个 `{{.}}` 片段用于在网页上显示传递给模板的数据-查询字符串。HTML模板包会自动提供适当的转义,使文本安全显示。

模板字符串的其余部分只是在页面加载时显示的HTML。如果这个解释还不够详细,请参考[模板包的文档](https://go.dev/pkg/html/template/)以获得更深入的讨论。

总之,这就是一个使用几行代码加上一些数据驱动的HTML文本来实现的有用的Web服务器。Go足够强大,可以在几行代码中实现很多功能。















