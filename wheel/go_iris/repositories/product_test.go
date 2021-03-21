package repositories

import (
	"database/sql"
	"fmt"
	"go_iris/common"
	"go_iris/datamodels"
	"reflect"
	"testing"
)

func getRepository(t *testing.T) IProduct {
	db, err := common.NewMysqlConn()
	if err != nil {
		t.Error(err)
	}

	return NewProductManage(db)
}

func TestFuncTest(t *testing.T) {
	repository := getRepository(t)
	result := repository.TestFunc()
	if result != "This is test func" {
		t.Errorf("Error output: %s, the true result is 'This is test func'", result)
	}
}

func TestProductManage_Insert(t *testing.T) {
	type fields struct {
		table  string
		dbConn *sql.DB
	}
	type args struct {
		product *datamodels.Product
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantProductId int64
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductManage{
				table:  tt.fields.table,
				dbConn: tt.fields.dbConn,
			}
			gotProductId, err := p.Insert(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotProductId != tt.wantProductId {
				t.Errorf("Insert() gotProductId = %v, want %v", gotProductId, tt.wantProductId)
			}
		})
	}
}

func TestProductManage_Delete(t *testing.T) {
	type fields struct {
		table  string
		dbConn *sql.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductManage{
				table:  tt.fields.table,
				dbConn: tt.fields.dbConn,
			}
			if got := p.Delete(tt.args.id); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductManage_Update(t *testing.T) {
	type fields struct {
		table  string
		dbConn *sql.DB
	}
	type args struct {
		product *datamodels.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductManage{
				table:  tt.fields.table,
				dbConn: tt.fields.dbConn,
			}
			if err := p.Update(tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductManage_SelectByKey(t *testing.T) {
	type fields struct {
		table  string
		dbConn *sql.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantProduct *datamodels.Product
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductManage{
				table:  tt.fields.table,
				dbConn: tt.fields.dbConn,
			}
			gotProduct, err := p.SelectByKey(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectByKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProduct, tt.wantProduct) {
				t.Errorf("SelectByKey() gotProduct = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}

func TestProductManage_SelectAll(t *testing.T) {
	repository := getRepository(t)
	result, err := repository.SelectAll()
	if err != nil {
		t.Error(err)
	}

	for _, i := range result {
		fmt.Println(i)
	}
}
