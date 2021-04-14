package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/go-kit/user/common"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/go-kit/user/dao"

	"github.com/jinzhu/gorm"
)

// 用户信息 data transfer object,  不带password字段
type UserInfoDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// 注册用户展示 view object
type RegisterUserVO struct {
	Username string
	Password string
	Email    string
}

var (
	ErrUserExisted = errors.New("user is existed")
	ErrPassword    = errors.New("email and password are not match")
	ErrRegistering = errors.New("email is registering")
)

// 用户抽象接口 实现登录和注册两个方法
type UserService interface {
	// 登录接口
	Login(ctx context.Context, email, password string) (*UserInfoDTO, error)
	// 注册接口
	Register(ctx context.Context, vo *RegisterUserVO) (*UserInfoDTO, error)
}

// UserService接口的具体实现
type UserServiceImpl struct {
	userDAO dao.UserDAO
}

// 初始化UserService
func MakeUserServiceImpl(userDAO dao.UserDAO) UserService {
	return &UserServiceImpl{
		userDAO: userDAO,
	}
}

// 登录逻辑
func (userService *UserServiceImpl) Login(
	ctx context.Context, email, password string) (*UserInfoDTO, error) {
	user, err := userService.userDAO.SelectByEmail(email)
	if err == nil {
		if user.Password == password {
			return &UserInfoDTO{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			}, nil
		} else {
			return nil, ErrPassword
		}
	} else {
		log.Printf("err : %s", err)
	}
	return nil, err
}

// 注册逻辑
func (userService UserServiceImpl) Register(
	ctx context.Context, vo *RegisterUserVO) (*UserInfoDTO, error) {

	lock := common.GetRedisLock(vo.Email, time.Duration(5)*time.Second)
	err := lock.Lock()
	if err != nil {
		log.Printf("err : %s", err)
		return nil, ErrRegistering
	}
	defer lock.Unlock()

	existUser, err := userService.userDAO.SelectByEmail(vo.Email)

	if (err == nil && existUser == nil) || err == gorm.ErrRecordNotFound {
		newUser := &dao.UserEntity{
			Username: vo.Username,
			Password: vo.Password,
			Email:    vo.Email,
		}
		err = userService.userDAO.Save(newUser)
		if err == nil {
			return &UserInfoDTO{
				ID:       newUser.ID,
				Username: newUser.Username,
				Email:    newUser.Email,
			}, nil
		}
	}
	if err == nil {
		err = ErrUserExisted
	}
	return nil, err
}
