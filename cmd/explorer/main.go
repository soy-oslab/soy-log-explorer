package main

import (
	"github.com/soyoslab/soy_log_explorer/internal/global"
	"github.com/soyoslab/soy_log_explorer/internal/rest"
	"github.com/soyoslab/soy_log_explorer/pkg/cmd/server"
	"github.com/soyoslab/soy_log_explorer/pkg/deamon"
)

func main() {
	go deamon.Listen(global.HotRing, rest.ESPush, 1)
	go deamon.Listen(global.ColdRing, rest.ESPush, 2)
	server.Execute()
}
