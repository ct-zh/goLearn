package gomock_demo

// go mock demo 测试打桩
// 如下面的GetFromDb方法，需要传入一个db，
// 对于单元测试，我们可以对db打桩，传入一个mock db,从而解开对db的依赖
//
// from: https://geektutu.com/post/quick-gomock.html
// 安装： go get -u github.com/golang/mock/gomock
// go get -u github.com/golang/mock/mockgen
//
// 命令： mockgen -source=t1.go -destination=t1_mock.go -package=main
// 生成对应的mock文件

type DB interface {
	Get(key string) (int, error)
}

func GetFromDb(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}
	return -1
}
