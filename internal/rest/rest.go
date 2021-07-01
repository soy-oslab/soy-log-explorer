package rest

import (
	"bytes"
	"encoding/gob"
	"github.com/go-resty/resty/v2"
	"github.com/soyoslab/soy_log_explorer/internal/global"
	"github.com/soyoslab/soy_log_generator/pkg/compressor"
)

func getDecompData(data []byte) []byte {
	c := &compressor.GzipComp{}
	return c.Decompress(data)
}

func getBytes(v interface{}) []byte {
	buf, ok := v.([]byte)
	if !ok {
		return nil
	}
	return buf
}

func castDocs(data []byte) global.ESdocs {
	var docs global.ESdocs

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	dec.Decode(&docs)

	return docs
}

// ESPush push the documents to elasticsearch
func ESPush(v ...interface{}) {
	docs := castDocs(getDecompData(getBytes(v[0])))
	resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(docs.Docs).
		Post("http://localhost:9200/" + docs.Index + "/_doc")
}
