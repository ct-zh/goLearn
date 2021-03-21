package services

import (
	"fmt"
	"go_iris/common"
	"go_iris/repositories"
	"testing"
)

func getService(t *testing.T) IProductService {
	db, err := common.NewMysqlConn()
	if err != nil {
		t.Error(err)
	}

	repository := repositories.NewProductManage("product", db)
	return NewProductService(repository)
}

func TestNewProductService(t *testing.T) {
	service := getService(t)
	fmt.Printf("%+v", service)
}

func TestProductService_DeleteById(t *testing.T) {
	// todo
}

func TestProductService_GetAll(t *testing.T) {
	service := getService(t)
	products, err := service.GetAll()
	if err != nil {
		t.Error(err)
	}
	for k, i := range products {
		fmt.Printf("Key: %d Value: %+v\n", k, i)
	}
}

func TestProductService_GetById(t *testing.T) {
	// todo
}

func TestProductService_Insert(t *testing.T) {
	// todo
}

func TestProductService_Update(t *testing.T) {
	// todo
}
