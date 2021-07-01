package rest

import (
	"bytes"
	"encoding/gob"
	"github.com/soyoslab/soy_log_explorer/internal/global"
	"github.com/soyoslab/soy_log_generator/pkg/compressor"
	"testing"
)

func TestESPush(t *testing.T) {
	docs := global.ESdocs{Index: "my_index", Docs: `{"name":"Gomez"}`}
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(docs)
	if err != nil {
		t.Error("cannot cast from ESdocs to []byte", err)
	}

	c := &compressor.GzipComp{}
	data := c.Compress(buf.Bytes())

	ESPush(data)
}
