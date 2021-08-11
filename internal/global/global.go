package global

import (
	"os"
	"strconv"

	"github.com/soyoslab/soy_log_collector/pkg/container/ring"
)

// ColdRing is or cold port
var ColdRing *ring.Ring

func init() {
	RINGSZ, _ := strconv.ParseInt(os.Getenv("EXPLORER_RINGSZ"), 10, 32)
	ColdRing = ring.New(int(RINGSZ))
}
