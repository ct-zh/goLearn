基于iris框架的并发抢购程序

> <a href="https://github.com/kataras/iris"> iris github</a>

## 优化思想：
> 筛选有效流量，异步处理数据

## 要点解析
1. 用户请求静态资源挂CDN，包括动态页面静态化挂CDN
2. SLB 用户请求 流量负载均衡
3. 用户请求 的 分布式安全验证，流量拦截
4. 商品数量控制，防止超卖，增加性能
5. web服务 => rabbitmq，防止爆库


## todo
1. repository的newInstance逻辑，db参数需要支持多种数据库
2. common的代码
    ```
    // 控制器参数
    dec := common.NewDecoder(&common.DecoderOptions{TagName: "imooc"})
    if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
       //  
    }
    ```
3. 后台接口加上权限验证

