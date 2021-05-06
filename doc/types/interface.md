## interface
接口有两种底层结构,一种是空接口:eface;一种是有方法的接口:iface;

> eface源码见src/runtime/runtime2.go
eface的定义是:
```go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```

> iface源码见src/runtime/runtime2.go
iface的定义是:
```go
type iface struct {
	tab  *itab
	data unsafe.Pointer
}
```
eface有`_type`字段;而iface有内容更丰富的`itab`字段,其中就包含了`_type`字段;

注意,`[]interface{}`类型是数组,`*interface{}`类型是指针,这两个的类型都不是接口,没有接口相关特性,这很容易搞混(坑).见下面例子:

#### []interface
> 原文见: https://github.com/golang/go/wiki/InterfaceSlice
以下语句将会报错:
```go
dataSlice := []int{1, 2, 3, 4, 5}
var iFaceSlice []interface{}
iFaceSlice = dataSlice
```
interface不应该什么类型都能表示吗?为什么这里无法将dataSlice赋给interface切片呢?

因为`[]interface{}`的类型不是interface,而是slice;在slice中,每个interface类型占两个字,而int类型只占一个字,它们的底层结构是不相同的;

正确的写法应该是:
```go
// method 1, 直接赋给interface
var iFaceSlice interface{}
iFaceSlice = dataSlice

// method 2, for循环赋值
iFaces := make([]interface{}, len(dataSlice))
for k, v := range dataSlice {
    iFaces[k] = v
}
```

#### 复制接口内容
对于结构体,一般有两种初始化方法:`a := foo{}`和`b := &foo{}`

对于方法`a := foo{}`,如果我们想复制接口内容,直接赋值就可以了: `c = a`

但是对于`b := &foo{}`,由于是引用类型,复制的是指针地址,所以不能直接赋值;

```go
type User interface {
	Name() string
	SetName(name string)
}

type Admin struct {
	name string
}

func (a *Admin) Name() string {
	return a.name
}

func (a *Admin) SetName(name string) {
	a.name = name
}
```
见上面这个例子, 因为Admin结构体的方法需要`*Admin`才能调用,(*存在指针方法的结构体,只能初始化成引用类型*),所以会出现这种情况:
```go
var user1 User
user1 = &Admin{name:"user1"}
fmt.Printf("User1's name: %s\n", user1.Name())
// User1's name: user1

var user2 User
user2 = user1
user2.SetName("user2")
fmt.Printf("User1's name: %s\n", user1.Name())
// User1's name: user2
fmt.Printf("User2's name: %s\n", user2.Name())
// User2's name: user2
```

那么如何*值复制*user1呢?

##### 方法一,解引用
```go
var user3 User
// 先转换user1的类型为Admin
padmin := user1.(*Admin)
// 再取出user1底层的数据(解引用)
admin := *padmin
// 将其再赋给*Admin
user3 = &admin

user3.SetName("user3")
fmt.Printf("User3's name: %s\n", user3.Name())
// User3's name: user3
fmt.Printf("User1's name: %s\n", user1.Name())
// User1's name: user1
```

##### 方法二,reflect
```go
var user5 User

// 如果user1是指针类型
if reflect.TypeOf(user1).Kind() == reflect.Ptr {
	user5 = reflect.New(reflect.ValueOf(user1).Elem().Type()).Interface().(User)
} else {
	// 如果user1是值类型
	user5 = reflect.New(reflect.TypeOf(user1)).Elem().Interface().(User)
}
user5.SetName("uaaaa")
fmt.Printf("User5's name: %s\n", user5.Name())
// User5's name uaaaa
fmt.Printf("User1's name: %s\n", user1.Name())
// User1's name: user1
```