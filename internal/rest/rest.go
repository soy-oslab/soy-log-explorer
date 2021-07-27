package rest

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/soyoslab/soy_log_explorer/internal/compressor"
	"github.com/soyoslab/soy_log_explorer/pkg/esdocs"
)

func getESdocs(v interface{}) (esdocs.ESdocs, error) {
	docs, ok := v.(esdocs.ESdocs)
	if !ok {
		log.Println("[Erorr] Can't get ESdocs")
		return docs, errors.New("Can't type conversion from interface{} to ESdocs")
	}
	return docs, nil
}

func getBytes(v interface{}) []byte {
	buf, ok := v.(string)
	if !ok {
		log.Println("[Erorr] Can't get bytes")
		return nil
	}
	return []byte(buf)
}

// ESPush push the documents to elasticsearch
func ESPush(docs esdocs.ESdocs) {
	// ramdisk

	docsIndex := strings.ToLower(docs.Index)
	var logarr []map[string]string

	err := json.Unmarshal([]byte(docs.Docs), &logarr)
	if err != nil {
		log.Println("[Erorr] Can't convert to json")
		return
	}

	for log := range logarr {
		resty.New().R().
			SetHeader("Content-Type", "application/json").
			SetBody(logarr[log]).
			Post("http://localhost:9200/" + docsIndex + "/_doc")
	}
}

// ESPushHot push the documents to elasticsearch for Hot data
func ESPushHot(v ...interface{}) {
	docs, err := getESdocs(v[0])
	if err != nil {
		log.Println("[Erorr] Can't decode data")
		return
	}
	ESPush(docs)
}

// ESPushCold push the documents to elasticsearch for Cold data
func ESPushCold(v ...interface{}) {
	docs, err := compressor.DocsDecompress(getBytes(v[0]))
	if err != nil {
		log.Println("[Erorr] Can't decode data")
		return
	}
	ESPush(docs)
}
