package main

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig()
	t.Logf("key = %s", cfg.OpenKey)
}
