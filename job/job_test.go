package job

import (
	"testing"
)

func TestRun(t *testing.T) {
	ch := make(chan int, 10)
	ch <-1
	ch <-2
	if len(ch) == 2 {
		t.Log("len == 2")
	} else {
		t.Error("len != 2")
	}
}