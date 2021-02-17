package napodate

import (
	"context"
	"testing"
	"time"
)

func Test_dateService_Status(t *testing.T) {
	srv, ctx := setup()

	s, err := srv.Status(ctx)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if ok := s == "ok"; !ok {
		t.Errorf("expected service to be ok")
		t.FailNow()
	}

	t.Logf("status test success! \n")
}

func Test_dateService_Get(t *testing.T) {
	srv, ctx := setup()
	date, err := srv.Get(ctx)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Interface Get returned %s \n", date)
}

func Test_dateService_Validate(t *testing.T) {
	date := time.Now().Format("02/01/2006")
	t.Logf("validate date: %+v", date)
	srv, ctx := setup()
	result, err := srv.Validate(ctx, date)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Validate: result is %+v \n", result)
}

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}
