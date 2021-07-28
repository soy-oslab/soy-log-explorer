package server

import (
	"flag"
	"fmt"
	"os"

	"github.com/smallnest/rpcx/server"
	"github.com/soyoslab/soy_log_explorer/internal/port"
)

// Execute runs the server
func Execute() {
	var addr = flag.String("addr",
		fmt.Sprintf("%s:%s", os.Getenv("EXPLORER_HOST"), os.Getenv("EXPLORER_PORT")),
		"server address")

	s := server.NewServer()
	s.Register(new(port.Rpush), "")
	s.Serve("tcp", *addr)
}
