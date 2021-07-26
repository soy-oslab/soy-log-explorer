package main

import (
	"github.com/soyoslab/soy_log_explorer/internal/global"
	"github.com/soyoslab/soy_log_explorer/internal/rest"
	"github.com/soyoslab/soy_log_explorer/pkg/cmd/server"
	deamon "github.com/soyoslab/soy_log_explorer/pkg/daemon"
)

func main() {
	go deamon.Listen(global.HotRing, rest.ESPushHot, 1)
	go deamon.Listen(global.ColdRing, rest.ESPushCold, 2)
	server.Execute()
}
