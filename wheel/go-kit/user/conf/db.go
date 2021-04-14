package conf

type Db struct {
	Host   string
	Port   string
	User   string
	Passwd string
	DbName string
}

func InitLocalDb() Db {
	return Db{
		Host:   "localhost",
		Port:   "3306",
		User:   "root",
		Passwd: "root",
		DbName: "testdb",
	}
}
