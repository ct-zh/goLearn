package repository

import (
	"testing"
	"user/conf"
	"user/domain/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

func newgorm() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", conf.NewDbArgs().DSN)
	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	return db, nil
}

func TestUserRepository_InitTable(t *testing.T) {
	type fields struct {
		mysqlDb *gorm.DB
	}
	db, err := newgorm()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "test1",
			fields: fields{mysqlDb: db},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepository{
				mysqlDb: tt.fields.mysqlDb,
			}
			err := u.InitTable()
			if err != nil {
				t.Errorf("InitTable() error = %v", err)
			} else {
				t.Log("InitTable() success")
			}
		})
	}
}

func TestUserRepository_CreateUser(t *testing.T) {
	type fields struct {
		mysqlDb *gorm.DB
	}
	type args struct {
		user *model.User
	}
	db, err := newgorm()
	if err != nil {
		t.Fatalf("db init error: %s", err.Error())
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test1",
			fields: fields{mysqlDb: db},
			args: args{user: &model.User{
				UserName:     "caoting",
				HashPassword: "123456",
			}},
		},
		{
			name:   "test2",
			fields: fields{mysqlDb: db},
			args: args{user: &model.User{
				UserName:     "dioBrando",
				HashPassword: "438438",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepository{
				mysqlDb: tt.fields.mysqlDb,
			}
			pw, err := bcrypt.GenerateFromPassword([]byte(tt.args.user.HashPassword), bcrypt.DefaultCost)
			if err != nil {
				t.Error(err)
			} else {
				tt.args.user.HashPassword = string(pw)
				got, err := u.CreateUser(tt.args.user)
				if err != nil {
					t.Error(err)
				}
				t.Logf("Create success: %+v \n", got)
			}
		})
	}
}

func TestUserRepository_DeleteUserByID(t *testing.T) {
	type fields struct {
		mysqlDb *gorm.DB
	}
	type args struct {
		userID int64
	}
	db, err := newgorm()
	if err != nil {
		t.Fatalf("db init error: %s", err.Error())
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test1",
			fields: fields{mysqlDb: db},
			args:   args{userID: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepository{
				mysqlDb: tt.fields.mysqlDb,
			}
			if err := u.DeleteUserByID(tt.args.userID); err != nil {
				t.Errorf("DeleteUserByID() error = %v", err)
			}
		})
	}
}

func TestUserRepository_FindAll(t *testing.T) {
	type fields struct {
		mysqlDb *gorm.DB
	}
	db, err := newgorm()
	if err != nil {
		t.Fatalf("db init error: %s", err.Error())
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			fields: fields{mysqlDb: db},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepository{
				mysqlDb: tt.fields.mysqlDb,
			}
			gotUserAll, err := u.FindAll()
			if err != nil {
				t.Error(err)
			} else {
				t.Logf("select Users: \n %+v\n", gotUserAll)
			}
		})
	}
}

func TestUserRepository_FindUserByID(t *testing.T) {
	type fields struct {
		mysqlDb *gorm.DB
	}
	type args struct {
		userID int64
	}
	db, err := newgorm()
	if err != nil {
		t.Fatalf("db init error: %s", err.Error())
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "",
			fields: fields{db},
			args:   args{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepository{
				mysqlDb: tt.fields.mysqlDb,
			}
			gotUser, err := u.FindUserByID(tt.args.userID)
			if err != nil {
				t.Error(err)
			} else {
				t.Log(gotUser)
			}
		})
	}
}

func TestUserRepository_FindUserByName(t *testing.T) {
	type fields struct {
		mysqlDb *gorm.DB
	}
	type args struct {
		name string
	}
	db, err := newgorm()
	if err != nil {
		t.Fatalf("db init error: %s", err.Error())
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "",
			fields: fields{db},
			args:   args{name: "caoting"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepository{
				mysqlDb: tt.fields.mysqlDb,
			}
			gotUser, err := u.FindUserByName(tt.args.name)
			if err != nil {
				t.Error(err)
			} else {
				t.Log(gotUser)
			}
		})
	}
}

func TestUserRepository_UpdateUser(t *testing.T) {
	type fields struct {
		mysqlDb *gorm.DB
	}
	type args struct {
		user *model.User
	}
	db, err := newgorm()
	if err != nil {
		t.Fatalf("db init error: %s", err.Error())
	}

	user, err := NewUserRepository(db).FindUserByName("caoting")
	if err != nil {
		t.Errorf("UpdateUser: find user failed")
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "",
			fields: fields{db},
			args:   args{user},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepository{
				mysqlDb: tt.fields.mysqlDb,
			}
			err := u.UpdateUser(tt.args.user)
			if err != nil {
				t.Error(err)
			} else {
				t.Log("update success")
			}
		})
	}
}
