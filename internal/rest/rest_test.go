package rest

import (
	"github.com/soyoslab/soy_log_explorer/internal/global"
	"testing"
)

func TestESPush(t *testing.T) {
	docs := global.ESdocs{Index: "my_index", Docs: `{"name":"James"}`}
	ESPush(docs)
}
