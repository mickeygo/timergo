package sqlite

import (
	"testing"
)

// TestGetDB
func TestGetDB(t *testing.T) {
	db, err := Open()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}
	
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Log("ping ok")
	} else {
		t.Error("ping error")
	}
}