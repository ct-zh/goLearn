package handler

import (
	"context"
	"errors"
	"github.com/micro/go-micro/v2/logger"
	"hot/domain/model"
	"hot/domain/service"
	"hot/proto"
)

type Account struct {
	Srv service.IAccount
}

var TypeMap = map[hot.Type]model.Type{
	hot.Type_App: model.App,
	hot.Type_Wx:  model.Wx,
}

func (a *Account) GetAccountInfo(_ context.Context, req *hot.AccountRequest, res *hot.AccountResponse) error {
	acc := req.Account
	if len(acc) <= 0 || len(acc) >= 50 {
		return errors.New("account invalid")
	}

	info, err := a.Srv.GetByAccount(acc)
	if err != nil {
		return err
	}

	for k, v := range TypeMap {
		if info.ClientType == v {
			res.ClientType = k
		}
	}

	res.Account = info.Account
	res.SecurityCode = info.SecurityCode
	res.CreateTime = info.CreateAt.Format("2006-01-02 15:04:05")
	logger.Infof("info: %+v res: %+v", info, res)

	return nil
}

func (a *Account) SetAccountType(_ context.Context, req *hot.AccountTypeRequest, res *hot.AccountTypeResponse) error {
	var t model.Type
	t, ok := TypeMap[req.Type]
	if !ok {
		return errors.New("invalid")
	}

	logger.Infof("req.Type: %+v t: %+v", req.Type, t)

	acc := req.Account
	if len(acc) <= 0 || len(acc) >= 50 {
		return errors.New("account invalid")
	}

	info, err := a.Srv.GetByAccount(acc)
	if err != nil {
		return err
	}

	info.ClientType = t
	if err := a.Srv.UpdateTypeById(info); err != nil {
		return err
	}

	res.IsSuccess = true
	return nil
}
