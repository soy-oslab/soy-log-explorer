package port

import (
	"github.com/soyoslab/soy_log_explorer/internal/global"
	"testing"
)

func TestHotPush(t *testing.T) {
	docs := global.ESdocs{Index: "my_index", Docs: `{"name":"Hotname"}`}
	var reply string

	new(Rpush).HotPush(nil, &docs, &reply)

	if reply != "" {
		t.Error("reply error")
	}
}

func TestColdPush(t *testing.T) {
	docs := global.ESdocs{Index: "my_index", Docs: `{"name":"Coldname"}`}
	var reply string

	new(Rpush).ColdPush(nil, &docs, &reply)

	if reply != "" {
		t.Error("reply error")
	}
}
