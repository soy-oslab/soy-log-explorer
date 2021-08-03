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
	expHost := os.Getenv("EXPLORER_HOST")
	if expHost == "" {
		expHost = "0.0.0.0"
	}
	expPort := os.Getenv("EXPLORER_PORT")
	if expPort == "" {
		expPort = "8972"
	}
	var addr = flag.String("addr",
		fmt.Sprintf("%s:%s", expHost, expPort),
		"server address")

	s := server.NewServer()
	s.Register(new(port.Rpush), "")
	s.Serve("tcp", *addr)
}
