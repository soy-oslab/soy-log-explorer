package rest

import (
	"testing"
)

func TestESPush(t *testing.T) {
	err := ESPush(`{"name":"James"}`, "my_index", "http://localhost:9200")
	if err != nil {
		t.Error("REST api error")
	}
}
