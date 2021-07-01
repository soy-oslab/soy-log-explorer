package rest

import (
	"testing"

	"github.com/soyoslab/soy_log_explorer/internal/compressor"
	"github.com/soyoslab/soy_log_explorer/internal/global"
)

func TestESPush(t *testing.T) {
	docs := global.ESdocs{Index: "my_index", Docs: `{"name":"Gomez"}`}

	b, err := compressor.DocsCompress(docs)
	if err != nil {
		t.Error(err)
	}

	ESPush(string(b))
}
