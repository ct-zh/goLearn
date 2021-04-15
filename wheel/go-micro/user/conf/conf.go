package conf

// 测试环境

type Mysql struct {
	DSN string
}

func NewDbArgs() Mysql {
	return Mysql{
		DSN: "micro:12345678@tcp(192.168.199.198:3306)/microSever?charset=utf8&parseTime=True&loc=Local",
	}
}
