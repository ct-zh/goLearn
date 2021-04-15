package model

import "time"

type Type int

type FullTableScanTest struct {
	Id           uint64    `gorm:"primary_key;not_null;auto_increment;Column=id" json:"id"`
	Account      string    `gorm:"not_null;Column=account" json:"account"`
	ClientType   Type      `gorm:"not_null;Column=client_type" json:"client_type"`
	SecurityCode string    `gorm:"not_null;Column=security_code" json:"security_code"`
	CreateAt     time.Time `gorm:"not_null;Column=create_at" json:"create_at"`
}

const (
	Wx Type = iota + 1
	App
)
