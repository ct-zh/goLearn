package config_1

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"golang.org/x/sync/singleflight"
)

type client struct {
	history sync.Map // 上一次配置

	db    *gorm.DB
	redis *redis.Pool
}

const timeout = time.Second

// ErrNotFound 未找到配置
var (
	ErrNotFound = errors.New("config not found")
	ErrTimeout  = errors.New("get config timeout")
)

func InitCfg(db *gorm.DB, redis *redis.Pool) *client {
	return &client{db: db, redis: redis}
}

func (c *client) Get(ctx context.Context, val string, force ...bool) (string, error) {
	var (
		g      = singleflight.Group{} // 增加缓冲器 防止缓存击穿
		timer  = time.NewTimer(timeout)
		result singleflight.Result
	)

	fn := c.findFn(val)
	ch := g.DoChan(val, fn)

	select {
	case result = <-ch:
	case <-timer.C:
		result = singleflight.Result{Val: nil, Err: ErrTimeout}
	}

	if result.Err != nil {
		if len(force) == 1 && force[0] == true {
			return "", result.Err
		}
		if history, ok := c.GetHistory(val); ok {
			return history, nil
		}
		return "", result.Err
	}

	c.history.Store(val, result)

	return result.Val.(string), nil
}

func (c *client) findFn(val string) func() (interface{}, error) {
	return func() (interface{}, error) {
		// 从缓存读
		key := fmt.Sprintf("config:%s", val)
		cache, err := c.redis.Get().Do("GET", key)
		if err != nil {
			return nil, err
		}

		if cache.(string) != "" { // 有缓存，直接返回
			return cache.(string), nil
		}

		ret := new(KeyValStruct)
		err = c.db.Table(ret.getKeyValTable()).
			Where("`key` = ? AND status = 1", val).
			First(&ret).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return "", ErrNotFound
			}
			return nil, err
		}

		_, _ = c.redis.Get().Do("SET", key, ret.Value, "EX", 60*5, "NX")

		return ret.Value, nil
	}
}

func (c *client) GetHistory(key interface{}) (res string, ok bool) {
	if result, ok := c.history.Load(key); ok {
		return result.(string), true
	} else {
		return "", false
	}
}

func (c *client) SetHistory(defaultKV map[string]string) {
	for k, v := range defaultKV {
		c.history.Store(k, v)
	}
}
