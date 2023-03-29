package main

import (
	"context"
	"testing"
)

func TestAskForOpenAI(t *testing.T) {
	cfg = NewConfig()

	content, err := AskForOpenAI(context.Background(), "", "hello world")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result=%s", content)
}
