package main

import (
	"github.com/soyoslab/soy_log_explorer/internal/global"
	"github.com/soyoslab/soy_log_explorer/internal/rest"
	"github.com/soyoslab/soy_log_explorer/pkg/cmd/server"
	"github.com/soyoslab/soy_log_explorer/pkg/daemon"
)

func main() {
	go daemon.Listen(global.HotRing, rest.ESPushHot, 0)
	go daemon.Listen(global.ColdRing, rest.ESPushCold, 1000)
	server.Execute()
}
