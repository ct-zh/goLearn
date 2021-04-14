package service

import (
	"context"
	"testing"

	"github.com/LannisterAlwaysPaysHisDebts/lesson38/common/redis"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module2/user/conf"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module2/user/dao"
)

func TestUserServiceImpl_Register(t *testing.T) {
	err := dao.InitMysql(conf.InitLocalDb())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	redisConf := conf.InitLocalRedis()
	err = redis.InitRedis(redisConf.Host, redisConf.Port, redisConf.Passwd)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	userServ := &UserServiceImpl{
		userDAO: &dao.UserDAOImpl{},
	}

	user, err := userServ.Register(context.Background(),
		&RegisterUserVO{
			Username: "aoho",
			Password: "aoho",
			Email:    "aoho@mail.com",
		})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("user id is %d", user.ID)
}

func TestUserServiceImpl_Login(t *testing.T) {
	err := dao.InitMysql(conf.InitLocalDb())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	redisConf := conf.InitLocalRedis()
	err = redis.InitRedis(redisConf.Host, redisConf.Port, redisConf.Passwd)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	userServ := &UserServiceImpl{
		userDAO: &dao.UserDAOImpl{},
	}
	user, err := userServ.Login(context.Background(), "aoho@mail.com", "aoho")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("user id is %d", user.ID)
}
