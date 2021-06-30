package main

import (
	"github.com/soyoslab/soy_log_explorer/internal/global"
	"github.com/soyoslab/soy_log_explorer/pkg/cmd/server"
	"github.com/soyoslab/soy_log_explorer/pkg/deamon"
)

func main() {
	go deamon.Listen(global.HotRing, 10)
	go deamon.Listen(global.ColdRing, 20)
	server.Execute()
}
