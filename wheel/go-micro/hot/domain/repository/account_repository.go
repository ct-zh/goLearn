package repository

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"hot/domain/model"
)

const (
	RedisCache = "account:%s"
)

type IAccount interface {
	// 主键查询
	Get(id int) (*model.FullTableScanTest, error)
	// 获取账号信息
	GetByAccount(string) (*model.FullTableScanTest, error)
	// 更新类型
	UpdateTypeById(test *model.FullTableScanTest) error
}

type Account struct {
	db    *gorm.DB
	redis redis.Conn
}

func NewAccount(db *gorm.DB, redis redis.Conn) *Account {
	return &Account{db: db, redis: redis}
}

func (a *Account) UpdateTypeById(test *model.FullTableScanTest) error {
	return a.db.Model(test).Update(test).Error
}

func (a *Account) Get(id int) (*model.FullTableScanTest, error) {
	acc := &model.FullTableScanTest{}
	return acc, a.db.Where("id = ?", id).First(acc).Error
}

func (a *Account) GetByAccount(s string) (*model.FullTableScanTest, error) {
	cacheKey := fmt.Sprintf(RedisCache, s)

	reply, err := a.redis.Do("GET", cacheKey)
	if err != nil {
		return nil, err
	}

	acc := &model.FullTableScanTest{}
	if reply != nil {
		err = json.Unmarshal(reply.([]byte), acc)
		if err != nil {
			return nil, err
		}
		return acc, nil
	} else {
		defer func() {
			cache, err := json.Marshal(acc)
			if err != nil {
				return
			}
			_, _ = a.redis.Do("SET", cacheKey, cache)
		}()
		return acc, a.db.Where("account = ?", s).First(acc).Error
	}
}
