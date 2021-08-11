package main

import (
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/soyoslab/soy_log_explorer/internal/global"
	"github.com/soyoslab/soy_log_explorer/internal/rest"
	"github.com/soyoslab/soy_log_explorer/pkg/cmd/server"
	"github.com/soyoslab/soy_log_explorer/pkg/daemon"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	coldTimeout, _ := strconv.ParseInt(os.Getenv("EXPLORER_COLD_TIMEOUT"), 10, 32)
	go daemon.Listen(global.HotRing, rest.ESPushHot, 0)
	go daemon.Listen(global.ColdRing, rest.ESPushCold, time.Duration(coldTimeout))
	server.Execute()
}
