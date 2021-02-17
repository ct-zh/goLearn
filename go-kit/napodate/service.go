package napodate

import (
	"context"
	"time"
)

// 在 Go kit 中，必须将服务建模为接口
// 如下所述，我们将需要三个端点，这些端点将被映射到此接口。
type Service interface {
	Status(ctx context.Context) (string, error)
	Get(ctx context.Context) (string, error)
	Validate(ctx context.Context, date string) (bool, error)
}

type dateService struct {
}

// Service 的构造函数
func NewService() Service {
	return dateService{}
}

// Status接口说明服务健康
func (d dateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Get接口返回今天的日期
func (d dateService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format("02/01/2006"), nil
}

// Validate判断date是否为今天的日期
func (d dateService) Validate(ctx context.Context, date string) (bool, error) {
	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		return false, err
	}
	now := time.Now().Format("02/01/2006")
	if now == date {
		return true, nil
	} else {
		return false, nil
	}
}
