package services

import (
	"errors"
	"go_iris/datamodels"
	repositories2 "go_iris/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	IsPwdSuccess(userName string, pwd string) (user *datamodels.User, ok bool)
	AddUser(user *datamodels.User) (userId int64, err error)
}

type UserService struct {
	Repository repositories2.IUser
}

func (u UserService) IsPwdSuccess(userName string, pwd string) (user *datamodels.User, ok bool) {
	user, err := u.Repository.Select(userName)
	if err != nil {
		return nil, false
	}

	if isOk, _ := validatePwd(pwd, user.Password); !isOk {
		return nil, false
	}

	return user, true
}

func validatePwd(pwd string, hash string) (ok bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(hash)); err != nil {
		return false, errors.New("error password")
	}
	return true, nil
}

func (u UserService) AddUser(user *datamodels.User) (userId int64, err error) {
	pwdByte, errPwd := generatePwd(user.Password)
	if errPwd != nil {
		return userId, errPwd
	}
	user.Password = string(pwdByte)
	return u.Repository.Insert(user)
}

func generatePwd(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

func NewUserService(repository repositories2.IUser) IUserService {
	return &UserService{
		Repository: repository,
	}
}
