package port

import (
	"testing"
)

func TestHotPush(t *testing.T) {
	args := `{"message":"abc"}`
	var reply string
	new(Rpush).HotPush(nil, &args, &reply)
	if reply != args {
		t.Error("reply error")
	}
}

func TestColdPush(t *testing.T) {
	args := `{"message":"abc"}`
	var reply string
	new(Rpush).ColdPush(nil, &args, &reply)
	if reply != args {
		t.Error("reply error")
	}
}
