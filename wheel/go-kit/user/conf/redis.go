package conf

type Redis struct {
	Host   string
	Port   string
	Passwd string
}

func InitLocalRedis() Redis {
	return Redis{
		Host:   "127.0.0.1",
		Port:   "6379",
		Passwd: "",
	}
}
