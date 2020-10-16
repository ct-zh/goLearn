基于iris框架的并发抢购程序

> <a href="https://github.com/kataras/iris"> iris github</a>

1. restful api
2. 使用rabbitMQ队列
3. 前端静态化


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
3. 后台接口用API格式，前端商品页面展示

