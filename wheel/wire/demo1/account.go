package main

import "errors"

type Account struct {
	Name     string
	Password string
	User     *User
}

func NewAccount(name string, password string, user *User) (*Account, error) {
	if user == nil {
		return nil, errors.New("user can not be nil")
	}
	return &Account{
		Name:     name,
		Password: password,
		User:     user,
	}, nil
}
