package dao

import (
	"testing"

	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module2/user/conf"
)

func TestUserDAOImpl_Save(t *testing.T) {
	userDAO := &UserDAOImpl{}

	err := InitMysql(conf.InitLocalDb())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// 初始化用户实例
	user := &UserEntity{
		Username: "aoho",
		Password: "aoho",
		Email:    "aoho@mail.com",
	}

	// 保存
	err = userDAO.Save(user)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("new User ID is %d", user.ID)
}

func TestUserDAOImpl_SelectByEmail(t *testing.T) {
	userDAO := &UserDAOImpl{}

	err := InitMysql(conf.InitLocalDb())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	user, err := userDAO.SelectByEmail("aoho@mail.com")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("result uesrname is %s", user.Username)
}
