package datamodels

type Order struct {
	ID          int64  `json:"id" sql:"id" data:"id"`
	UserId      int64  `json:"userId" sql:"userId" data:"userId"`
	ProductId   int64  `json:"productId" sql:"productId" data:"productId"`
	OrderStatus string `json:"orderStatus" sql:"orderStatus" data:"orderStatus"`
}

const (
	OrderWait = iota
	OrderSuccess
	OrderFailed
)
