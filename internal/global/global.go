package global

import (
	"github.com/soyoslab/soy_log_collector/pkg/container/ring"
)

// RINGSZ is maximum ring size
const RINGSZ = 1000

// HotRing is for hot port
var HotRing *ring.Ring

// ColdRing is or cold port
var ColdRing *ring.Ring

func init() {
	HotRing = ring.New(RINGSZ)
	ColdRing = ring.New(RINGSZ)
}
