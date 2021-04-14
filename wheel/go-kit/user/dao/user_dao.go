package dao

import (
	"fmt"

	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/go-kit/user/types/time"
)

// 用户表
type UserEntity struct {
	ID        int64     `gorm: "primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	Username  string    `json: "username"`
	Password  string    `json: "password"`
	Email     string    `json: "email"`
	CreatedAt time.Time `json: "created_at"`
}

// 返回表名
func (UserEntity) TableName() string {
	return "user"
}

// User接口 Data Access Object
type UserDAO interface {
	SelectByEmail(email string) (*UserEntity, error)
	Save(user *UserEntity) error
}

// User结构体, UserDAO接口的实现
type UserDAOImpl struct {
}

// 通过email获取用户实体
func (userDAO *UserDAOImpl) SelectByEmail(email string) (*UserEntity, error) {
	user := &UserEntity{}
	err := db.Where("email = ?", email).First(user).Error
	fmt.Println(user)
	return user, err
}

// 保存用户
func (userDAO *UserDAOImpl) Save(user *UserEntity) error {
	return db.Create(user).Error
}
