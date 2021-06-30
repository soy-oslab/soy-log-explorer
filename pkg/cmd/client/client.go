package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"time"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	var args string
	var reply string

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	xclient := client.NewXClient("Rpush", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args = fmt.Sprintf(`{ "index": { "hour": %d, "minute": %d, "second": %d } }`, time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	err := xclient.Call(context.Background(), "HotPush", &args, &reply)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("reply: %s\n", reply)
	}
}
