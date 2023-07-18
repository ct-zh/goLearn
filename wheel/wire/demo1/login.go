package main

type LoginAction struct {
	Account    Account
	ResultCode int
	ResultMsg  string
	Err        error
}
