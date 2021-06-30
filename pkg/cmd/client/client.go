package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/client"
)

var addr = flag.String("addr", "localhost:8972", "server address")

type esdocs struct {
	Index string
	Docs  string
}

func main() {
	flag.Parse()

	var reply string

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	xclient := client.NewXClient("Rpush", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	a := esdocs{Index: "my_index", Docs: `{ "name" : "drop" }`}

	err := xclient.Call(context.Background(), "HotPush", &a, &reply)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("reply: %s\n", reply)
	}
}
