# go cmd

## go tools
### go vet 与 go tool vet



### go get
1. 下载项目依赖`go get ./...`
2. 拉取最新的版本(优先择取 tag) `go get golang.org/x/text@latest`
3. 拉取 master 分支的最新 commit `go get golang.org/x/text@master`
4. 拉取 tag 为 v0.3.2 的 commit `go get golang.org/x/text@v0.3.2`
5. 更新 `go get -u`



### go mod
#### 开启 go module
系统环境变量修改, proxy国内建议使用阿里的代理:
```
GO111MODULE=on
GOPROXY=https://goproxy.cn,https://goproxy.io,direct
```
#### 用法
1. 初始化一个moudle，模块名为你项目名 `go mod init 模块名`
2. 下载modules到本地cache `go mod download`
	> 目前所有模块版本数据均缓存在 $GOPATH/pkg/mod和 ​$GOPATH/pkg/sum 下
3. 编辑go.mod文件 选项有-json、-require和-exclude，可以使用帮助go help mod edit
```
go mod edit
```
4. 以文本模式打印模块需求图 `go mod graph`
5. 删除错误或者不使用的modules  `go mod tidy`
6. 生成vendor目录 `go mod vendor`
7. 验证依赖是否正确 `go mod verify`
8. 查找依赖 `go mod why`
9. 替代只能翻墙下载的库
```
go mod edit -replace=golang.org/x/crypto@v0.0.0=github.com/golang/crypto@latest
go mod edit -replace=golang.org/x/sys@v0.0.0=github.com/golang/sys@latest
```
10. 清理moudle 缓存 `go clean -modcache`
11. 查看可下载版本 `go list -m -versions github.com/gogf/gf`



### godoc文档
命令行:`godoc -http=:12333`,访问本地12333端口,可以看到本地的go文档.


### go tool trace



### go tool pprof


