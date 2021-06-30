package global

import (
	"github.com/soyoslab/soy_log_collector/pkg/container/ring"
)

// RINGSZ is maximum ring size
var RINGSZ = 10

// HotRing is for hot port
var HotRing *ring.Ring

// ColdRing is or cold port
var ColdRing *ring.Ring

func init() {
	HotRing = ring.New(RINGSZ)
	ColdRing = ring.New(RINGSZ)
}
