package handler

import (
	"context"
	"log"
	"user/domain/model"
	"user/domain/service"
	"user/proto"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u *User) Register(ctx context.Context, userRegisterRequest *user.UserRegisterRequest,
	userRegisterResponse *user.UserRegisterResponse) error {

	userRegister := &model.User{
		ID:           0,
		UserName:     userRegisterRequest.UserName,
		HashPassword: userRegisterRequest.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		log.Println("add failed")
		return err
	}

	userRegisterResponse.Msg = "新增成功"
	log.Println("success")
	return nil
}

func (u *User) Login(ctx context.Context, in *user.UserLoginRequest, out *user.UserLoginResponse) error {
	log.Println("Login UserName: ", in.UserName, " pwd:", in.Pwd)
	isOk, err := u.UserDataService.CheckPwd(in.UserName, in.Pwd)
	if err != nil {
		log.Println("check failed")
		return err
	}
	out.IsSuccess = isOk
	log.Println("success")
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, in *user.UserInfoRequest, out *user.UserInfoResponse) error {
	log.Println("GetUserInfo UserName: ", in.UserName)
	userInfo, err := u.UserDataService.FindUserByName(in.UserName)
	if err != nil {
		log.Println(" find failed")
		return err
	}
	out.UserId = userInfo.ID
	out.UserName = userInfo.UserName
	log.Printf("success: %+v ||| out: %+v \n", userInfo, out)
	return nil
}
