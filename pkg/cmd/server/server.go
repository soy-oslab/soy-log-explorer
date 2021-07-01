package server

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"github.com/soyoslab/soy_log_explorer/internal/port"
	"os"
)

// Execute runs the server
func Execute() {
	var addr = flag.String("addr", "0.0.0.0:"+os.Getenv("EXPLORER_PORT"), "server address")

	s := server.NewServer()
	s.Register(new(port.Rpush), "")
	s.Serve("tcp", *addr)
}
