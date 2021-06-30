package server

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"github.com/soyoslab/soy_log_explorer/internal/port"
)

var addr = flag.String("addr", "localhost:8972", "server address")

// Execute runs the server
func Execute() {
	s := server.NewServer()
	s.Register(new(port.Rpush), "")
	s.Serve("tcp", *addr)
}
