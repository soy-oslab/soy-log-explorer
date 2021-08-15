package rest

import (
	"testing"

	"github.com/soyoslab/soy_log_explorer/internal/compressor"
	"github.com/soyoslab/soy_log_explorer/pkg/esdocs"
)

func TestESPush(t *testing.T) {
	docs := esdocs.ESdocs{Index: "my_index", Docs: `[{"name":"Gomez"}]`}

	b, err := compressor.DocsCompress(docs)
	if err != nil {
		t.Error(err)
	}

	ESPush(docs)
	ESPushCold(string(b))
}

func TestESPushFailed(t *testing.T) {
	esdocs := esdocs.ESdocs{}
	esdocs.Docs = "{;l;23;"
	ESPush(esdocs)
	ESPushCold(esdocs)
}

func TestGetEsdocs(t *testing.T) {
	getESdocs(esdocs.ESdocs{})
	getESdocs("")
}
