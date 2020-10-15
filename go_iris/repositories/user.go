package repositories

import (
	"database/sql"
	"go_iris/common"
	"go_iris/datamodels"
)

type IUser interface {
	Conn() error
	Insert(user *datamodels.User) (int64, error)
	Select(userName string) (user *datamodels.User, err error)
}

type UserManage struct {
	table  string
	dbConn *sql.DB
}

func (u UserManage) Conn() error {
	if u.dbConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		u.dbConn = mysql
	}
	if u.table == "" {
		u.table = "user"
	}
	return nil
}

func (u UserManage) Insert(user *datamodels.User) (int64, error) {
	panic("implement me")
}

func (u UserManage) Select(userName string) (user *datamodels.User, err error) {
	panic("implement me")
}

func NewUserManage(mysql *sql.DB) IUser {
	return &UserManage{
		table:  "user",
		dbConn: mysql,
	}
}
