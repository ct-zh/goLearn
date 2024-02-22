# 字符串操作与变量转换

## 前言
关于 Go 语言字符串的使用，我们需要了解标准库`strconv`和标准库`strings`的使用方式，它们分别用于字符串类型转换和字符串操作。

## 字符串操作
这里给出常用的字符串操作函数，基本基于标准库`strings`

- **字符串查找** `Contains(s, substr string) bool` 
  - 如： `strings.Contains("seafood", "foo")`返回true

- 字符串中是否包含指定字符串中任意字符 `ContainsAny(s, chars string) bool` 
  - 如 `strings.ContainsAny("fail", "ui")`返回`true`

- *删除指定字符* `Trim(s, cutset string) string`
  - 如 `strings.Trim("¡¡¡Hello, Gophers!!!", "!¡")` 返回`Hello, Gophers`


- 转换大小写 `ToUpper(s string) string` 与 `ToLower(string) string`
- **拆分** `Split(s, sep string) []string`  将字符串 `s` 以字符串 `sep` 为分隔符，拆分为字符串切片

- **拼接** `Join(elems []string, sep string) string`

## 字符串转换

使用标准库 `strconv`

- **字符串 to 布尔** `ParseBool(str string) (bool, error)`
  - 该函数接收参数的值是有限制的，除了 1、t、T、TRUE、true、True、0、f、F、FALSE、false、False 之外，其它任何值都会返回 error
- **布尔 to 字符串** `FormatBool(b bool) string`

- **字符串 to 浮点型** `ParseFloat(s string, bitSize int) (float64, error)`
  - bitSize可以为32、64
  - 该函数接收参数可以识别值为 `NaN`、`Inf`（有符号 `+Inf` 或 `-Inf`），并且忽略它们的大小写。
- **浮点型 to 字符串** `FormatFloat(f float64, fmt byte, prec, bitSize int) string`
  - 第一个参数是需要转换的浮点数；第二个参数是进制；第三个参数是精度，第四个参数是转换后值的取值范围
  - 第二个参数 `b` 代表二进制指数；`e` 或 `E` 代表十进制指数；`f` 代表无进制指数；`g` 或 `G` 代表指数大时 为 `e`，反之为 `f`；`x` 或 `X` 代表十六进制分数和二进制指数
  - 第三个参数，精度 prec 控制由 'e'，'E'，'f'，'g'，'G'，'x' 和 'X' 格式打印的位数(不包括指数)。对于 'e'，'E'，'f'，'x' 和 'X'，它是小数点后的位数。对于 'g' 和 'G'，它是有效数字的最大数目(去掉后面的零)。特殊精度 -1 使用所需的最小位数，以便 ParseFloat 精确返回 `f`
- **字符串 to 整型** `ParseInt(s string, base int, bitSize int) (i int64, err error)`
  - 第一个入参为字符串类型的数值，可以 "+" 或 "-" 符号开头
  - 第二个参数指定进制，它的值如果是 `0`，进制则以第一个参数符号后的前缀决定，例如："0b" 为 2，"0" 或 "0o" 为 8，"0x" 为 16，否则为 10；
  - 第三个参数指定返回结果必须符合整数类型的取值范围，它的值为 0、8、16、32 和 64，分别代表 `int`、`int8`、`int16`、`int32` 和 `int64`。
  - 实际项目开发中，十进制使用的最多，所以标准库 `strconv` 提供了函数 `func Atoi(s string) (int, error)`，它的功能类似 `ParseInt(s, 10, 0)`，需要注意的是，它的返回值类型是 `int`（需要注意取值范围），而不是 `int64`
- **整型 to 字符串** `FormatInt(i int64, base int) string`
  - 第二个参数的取值范围 `2 <= base <= 36`
  - 在实际项目开发中，十进制使用的最多，所以标准库 `strconv` 提供了函数 `func Itoa(i int) string`，它的功能类似 `FormatInt(int64(i), 10)`，需要注意的是，该函数入参的类型是 `int`。

> 对于类型转换，也可以使用第三方库如[GitHub - spf13/cast: safe and easy casting from one type to another in Go](https://github.com/spf13/cast)  可以避免一些panic



## 高效拼接字符串

-  使用`+`，如`str := "a" + "b" + "c"`
  - 每次拼接，都涉及内存拷贝，需要分配一块新内存，并且该方式也仅适用于字符串类型的变量
- 使用`strings.Join`
- 使用 `fmt.Sprint` 涉及到类型转换，性能一般
- 使用 `bytes.Buffer`
  - 底层涉及到`string`与`[]byte`的转换，另外WriteString 方法使用的 buffer 太长，会导致 panic，只适合少量字符变量和字符串变量进行字符串拼接的场景
- 使用`strings.Builder` 
  - 使用 `unsafe.Pointer` 优化了 `string` 和 `[]byte` 之间的转换，所以，在大量字符串拼接的场景，推荐使用该种方式
