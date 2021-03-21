package datamodels

type User struct {
	ID       int64  `json:"id" sql:"id" data:"id"`
	NickName string `json:"nickName" sql:"nickName" data:"nickName"`
	Account  string `json:"account" sql:"account" data:"account"`
	Password string `json:"password" sql:"password" data:"password"`
}
