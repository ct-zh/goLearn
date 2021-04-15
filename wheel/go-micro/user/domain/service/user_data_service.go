package service

import (
	"errors"

	"user/domain/model"
	"user/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) (err error)
	FindUserByName(string) (*model.User, error)
	CheckPwd(userName string, pwd string) (isOk bool, err error)
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

//创建实例
func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed string) (is_ok bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword))
	if err != nil {
		return false, errors.New("密码错误")
	}
	return true, nil
}

func (u *UserDataService) AddUser(user *model.User) (userID int64, err error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

func (u *UserDataService) DeleteUser(userID int64) error {
	return u.UserRepository.DeleteUserByID(userID)
}

func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	if isChangePwd {
		pwdBytes, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdBytes)
	}
	return u.UserRepository.UpdateUser(user)
}

func (u *UserDataService) FindUserByName(name string) (user *model.User, err error) {
	// check name
	if len(name) == 0 {
		return user, errors.New("name非法")
	}
	return u.UserRepository.FindUserByName(name)
}

func (u *UserDataService) CheckPwd(userName string, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}

	return ValidatePassword(pwd, user.HashPassword)
}
