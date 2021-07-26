package rest

import (
	"testing"

	"github.com/soyoslab/soy_log_explorer/internal/compressor"
	"github.com/soyoslab/soy_log_explorer/pkg/esdocs"
)

func TestESPush(t *testing.T) {
	docs := esdocs.ESdocs{Index: "my_index", Docs: `{"name":"Gomez"}`}

	b, err := compressor.DocsCompress(docs)
	if err != nil {
		t.Error(err)
	}

	ESPushHot(docs)
	ESPushCold(string(b))
}
