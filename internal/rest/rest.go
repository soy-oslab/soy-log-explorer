package rest

import (
	"github.com/go-resty/resty/v2"
	"github.com/soyoslab/soy_log_explorer/internal/global"
)

func castDocs(v interface{}) global.ESdocs {
	docs := v.(global.ESdocs)
	return docs
}

// ESPush push the documents to elasticsearch
func ESPush(v ...interface{}) {
	docs := castDocs(v[0])
	index := docs.Index
	data := docs.Docs
	serv := "http://localhost:9200"

	resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(serv + "/" + index + "/_doc")
}
