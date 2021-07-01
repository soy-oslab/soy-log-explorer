package port

import (
	"testing"

	"github.com/soyoslab/soy_log_explorer/internal/compressor"
	"github.com/soyoslab/soy_log_explorer/internal/global"
)

func TestHotPush(t *testing.T) {
	docs := global.ESdocs{Index: "my_index", Docs: `{"name":"Hotname"}`}
	var reply string

	b, err := compressor.DocsCompress(docs)
	if err != nil {
		t.Error(err)
	}

	new(Rpush).HotPush(nil, &b, &reply)

	if reply != "" {
		t.Error("reply error")
	}
}

func TestColdPush(t *testing.T) {
	docs := global.ESdocs{Index: "my_index", Docs: `{"name":"Coldname"}`}
	var reply string

	b, err := compressor.DocsCompress(docs)
	if err != nil {
		t.Error(err)
	}

	new(Rpush).ColdPush(nil, &b, &reply)

	if reply != "" {
		t.Error("reply error")
	}
}
