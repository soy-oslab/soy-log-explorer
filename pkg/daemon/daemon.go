package daemon

import (
	"time"

	"github.com/soyoslab/soy_log_collector/pkg/container/ring"
	"github.com/soyoslab/soy_log_explorer/pkg/signal"
)

// Listen does background jobs
// ring: contains jobs
// fn: function
// duration: sleep every x period(milli)
func Listen(ring *ring.Ring, fn func(...interface{}), duration time.Duration) {
	for true {
		signal.Wait()

		for true {
			next, err := ring.Pop()
			if err == nil {
				fn(next)
			} else {
				break
			}
		}
	}
}
