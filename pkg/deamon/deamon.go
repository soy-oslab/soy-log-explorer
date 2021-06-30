package deamon

import (
	"github.com/soyoslab/soy_log_collector/pkg/container/ring"
	"github.com/soyoslab/soy_log_explorer/internal/global"
	"github.com/soyoslab/soy_log_explorer/internal/rest"
	"time"
)

func docsCast(obj interface{}) global.ESdocs {
	docs, _ := obj.(global.ESdocs)
	return docs
}

// Listen does background jobs
func Listen(ring *ring.Ring, duration time.Duration) {
	for true {
		next, err := ring.Pop()
		if err == nil {
			docs := docsCast(next)
			rest.ESPush(docs.Docs, docs.Index, "http://localhost:9200")
		}
		time.Sleep(time.Second * duration)
	}
}
