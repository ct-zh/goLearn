package simple

import (
	"fmt"
	"sync"
	"time"
)

// Pool是一个可以分别存取的临时对象的集合。
//
// Pool中保存的任何item都可能随时不做通告的释放掉。(pool里的内容随时会被GC)
// 如果Pool持有该对象的唯一引用，这个item就可能被回收。
//
// Pool可以安全的被多个线程同时使用。
//
// Pool的目的是缓存申请但未使用的item用于之后的重用，以减轻GC的压力。
// 也就是说，让创建高效而线程安全的空闲列表更容易。但Pool并不适用于所有空闲列表。
//
// Pool的合理用法是用于管理一组静静的被多个独立并发线程共享并可能重用的临时item。
// Pool提供了让多个线程分摊内存申请消耗的方法。
//
// Pool的一个好例子在fmt包里。该Pool维护一个动态大小的临时输出缓存仓库。
// 该仓库会在过载（许多线程活跃的打印时）增大，在沉寂时缩小。
//
// 另一方面，管理着短寿命对象的空闲列表不适合使用Pool，
// 因为这种情况下内存申请消耗不能很好的分配。这时应该由这些对象自己实现空闲列表。

// Get()
// Get方法从池中选择任意一个item，删除其在池中的引用计数，并提供给调用者。
// Get方法也可能选择无视内存池，将其当作空的。
// 调用者不应认为Get的返回这和传递给Put的值之间有任何关系。
// 假使Get方法没有取得item：如p.New非nil，Get返回调用p.New的结果；否则返回nil。

// 看不懂？ 看看下面这个例子

func poolFn() {
	pool := &sync.Pool{
		New: func() interface{} {
			b := make([]byte, 5)
			return &b
		},
	}

	// 第一个get，因为pool是空的，会调用new
	res := pool.Get().(*[]byte)
	fmt.Printf("%+v \n", res)

	(*res)[0] = 'a'
	(*res)[1] = 'b'
	(*res)[2] = 'c'

	pool.Put(res)

	time.Sleep(time.Second)

	// 第二次拿到数据,获得的byte切片是刚刚PUT进去的

	// ATTENTION： 如果触发过GC，
	// 如果PUT进去的切片只有pool的唯一引用，则会被GC；
	// 因此，pool不能保证拿到的item就是PUT进去的item
	// runtime.GC()

	res2 := pool.Get().(*[]byte)
	fmt.Printf("%+v \n", res2)
}
