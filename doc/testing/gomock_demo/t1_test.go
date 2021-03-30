package gomock_demo

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDb(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	if v := GetFromDb(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
