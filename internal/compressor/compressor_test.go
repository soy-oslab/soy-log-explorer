package compressor

import (
	"testing"

	"github.com/soyoslab/soy_log_explorer/pkg/esdocs"
)

func TestCompress(t *testing.T) {
	docs := esdocs.ESdocs{Index: "my_index", Docs: `{"name":"hong"}`}

	b, err := DocsCompress(docs)
	if err != nil {
		t.Error(err)
	}

	ret, err := DocsDecompress(b)
	if err != nil {
		t.Error(err)
	}

	if docs.Index != ret.Index {
		t.Error("Index is different!")
	}

	if docs.Docs != ret.Docs {
		t.Error("Docs is different!")
	}
}
