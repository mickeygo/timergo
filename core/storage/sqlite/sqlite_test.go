package sqlite

import (
	"testing"
)

func TestGetEngine(t *testing.T) {
	engine, err := GetEngine()

	if err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}

	if err := engine.Ping(); err != nil {
		t.Log("ping ok")
	} else {
		t.Error("ping error")
	}
}