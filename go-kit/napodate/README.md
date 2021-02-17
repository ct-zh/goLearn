> source code is forked by https://github.com/napolux/go-kit-microservice-example-tutorial-99999
> content is from https://learnku.com/go/t/38417

mux做路由；

1. main： 映射端点，启动http服务；
2. endpoint: 公开端点逻辑，返回数据；
3. transport: 约束请求与响应的数据类型;注册请求与响应数据的编码器；
4. service: 具体实现逻辑;
5. server: 注册路由与中间件；