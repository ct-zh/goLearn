package datamodels

type Product struct {
	Id           int64  `json:"id" sql:"id" data:"id"`
	ProductName  string `json:"productName" sql:"productName" data:"productName"`
	ProductNum   int64  `json:"productNum" sql:"productNum" data:"productNum"`
	ProductImage string `json:"productImage" sql:"productImage" data:"ProductImage"`
	ProductUrl   string `json:"productUrl" sql:"productUrl" data:"productUrl"`
}
