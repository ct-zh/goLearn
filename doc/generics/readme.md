# 泛型

从历史上看，C++、D 乃至 Rust 等系统语言一直采用单态化方法实现泛型。然而，Go 1.18 的泛型实现并不完全依靠单态化 (Monomorphization)，而是采用了一种被称为"GCShape stenciling with Dictionaries"的部分单态化技术。这种方法的好处是可以大幅减少代码量，但在特定情况下，会导致代码速度变慢。
Ian Lance Taylor 表示，Go 的通用开发准则有要求：开发者应通过编写代码而不是定义类型来编写 Go 程序。当涉及到泛型时，如果通过定义类型参数约束来编写程序，那一开始就走错了路。正解应该是从编写函数开始，当明确了类型参数的作用后，再添加类型参数就很容易了。
接着，Ian 列举了 4 种类型参数能有效发挥作用的情况：
1. 使用语言定义的特殊容器类型；
2. 通用数据结构；
3. 类型参数首选是函数，而非方法的情况；
4. 不同类型需要实现通用方法；
 
同时也提醒了不适合使用类型参数的情况：
1. 不要使用类型参数替换接口类型 (Interface Type)；
2. 如果方法实现不同，不要使用类型参数；
3. 在适当的地方使用反射 (reflection)；

最后，Ian 给出了简要的泛型使用方针，当开发者发现自己多次编写完全相同的代码，而这些副本之间的唯一区别仅在于使用了不同类型，这时候便可以考虑使用类型参数。换句话说，即开发者应避免使用类型参数，直到发现自己要多次编写完全相同的代码。

入门泛型使用可参考代码[generics.go](./generics/generics.go)