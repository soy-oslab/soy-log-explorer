package rest

import (
	"github.com/go-resty/resty/v2"
	"github.com/soyoslab/soy_log_explorer/internal/compressor"
)

func getBytes(v interface{}) []byte {
	buf, ok := v.(string)
	if !ok {
		return nil
	}
	return []byte(buf)
}

// ESPush push the documents to elasticsearch
func ESPush(v ...interface{}) {
	docs, err := compressor.DocsDecompress(getBytes(v[0]))
	if err != nil {
		return
	}
	// ramdisk
	resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(docs.Docs).
		Post("http://localhost:9200/" + docs.Index + "/_doc")
}
