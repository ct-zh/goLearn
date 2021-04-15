package repository

import (
	"hot/common"
	"hot/domain/model"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGet(t *testing.T) {
	db, err := common.GetTestDb()
	if err != nil {
		t.Fatal(err)
	}

	a := NewAccount(db)
	if data, err := a.Get(1); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("get data: %+v", data)
	}
}

func TestGetByAccount(t *testing.T) {
	db, err := common.GetTestDb()
	if err != nil {
		t.Fatal(err)
	}

	a := NewAccount(db)
	if data, err := a.GetByAccount("test10"); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("get data: %+v", data)
	}
}

func TestUpdateTypeById(t *testing.T) {
	db, err := common.GetTestDb()
	if err != nil {
		t.Fatal(err)
	}

	a := NewAccount(db)
	if data, err := a.Get(1); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("get data: %+v", data)
		oldType := data.ClientType
		var newType model.Type
		switch oldType {
		case model.Wx:
			newType = model.App
		case model.App:
			newType = model.Wx
		default:
			t.Fatal("error type")
		}

		data.ClientType = newType
		err = a.UpdateTypeById(data)
		if err != nil {
			t.Fatal(err)
		}

		if data2, err := a.Get(1); err != nil {
			t.Fatal(err)
		} else {
			data2.ClientType = oldType
			err = a.UpdateTypeById(data2)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
