package rest

import (
	"github.com/go-resty/resty/v2"
)

// ESPush push the documents to elasticsearch
func ESPush(data string, index string, serv string) error {
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(serv + "/" + index + "/_doc")
	return err
}
