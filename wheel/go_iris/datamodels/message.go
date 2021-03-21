package datamodels

type Message struct {
	ProductId int64
	UserId    int64
}

func NewMessage(userId int64, productId int64) *Message {
	return &Message{
		ProductId: productId,
		UserId:    userId,
	}
}
