package main

import "fmt"

// golang大坑: 两个为nil的变量未必相等

// main函数里,为什么Generate()与Test()返回的数据都是类型为*main.MagicError
// 值为nil, 为何Generate就能等于nil,而Test不等于nil呢?
//
// 因为Test()方法返回的error是interface类型;
// interface底层,无论是iface还是eface,都存在type与data,
// 类型为interface的变量判断是否等于nil,必须type和data都为nil
func main() {
	fmt.Printf("%T %+v \n", Test(), Test())
	fmt.Printf("%T %+v \n", Generate(), Generate())
	fmt.Println(Generate() == nil)
	fmt.Println(Test() == Generate())
	fmt.Println(Test() == nil)
	// 上面代码中`Generate()`返回的是nil指针;
	// `Test()`返回的是interface,因为type不等于nil,所以Test返回的值不等于nil;

	// 这里提一种写法: `(*interface{})(nil)`,
	// 意思是将nil转换成interface类型的指针;
	// 得到的结果仅仅是空接口类型指针指向无效的地址,
	// 这样写的作用是强调val虽然是无效的数据,但是它是有类型`*interface{}`的;
	a := (*interface{})(nil)
	fmt.Println(a == nil) // true

	b := (*MagicError)(nil)
	fmt.Println(b == nil) // true

	var c interface{}
	c = b                 // fmt.Println(c) => nil
	fmt.Println(c == nil) // false

	// 当b转换为interface类型后,就不等于nil了
}

type MagicError struct{}

func (MagicError) Error() string {
	return "[Magic]"
}
func Generate() *MagicError {
	return nil
}

func Test() error {
	return Generate()
}

// 上面Test是有问题的写法，生产环境应该这样抛出err:
func Test2() error {
	err := Generate()
	if err != nil {
		return err
	}
	return nil
}
