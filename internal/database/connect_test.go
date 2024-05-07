package database

import (
	"testing"
)

func TestConnect(t *testing.T) {

	service := New(GetConfig())

	if service == nil || service.Client == nil {
		t.Errorf("service is nil")
	}
}
