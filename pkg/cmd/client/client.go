package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"flag"
	"fmt"
	"os"

	"github.com/smallnest/rpcx/client"
	"github.com/soyoslab/soy_log_generator/pkg/compressor"
)

var addr = flag.String("addr", "localhost:"+os.Getenv("EXPLORER_PORT"), "server address")

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

	docs := esdocs{Index: "my_index", Docs: `{"name":"steels"}`}
	var buf bytes.Buffer

	err := xclient.Call(context.Background(), "HotPush", &docs, &reply)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("[Hot] reply: %s\n", reply)
	}

	enc := gob.NewEncoder(&buf)
	err = enc.Encode(docs)
	if err != nil {
		fmt.Printf("error1: %v\n", err)
		return
	}

	c := &compressor.GzipComp{}
	data, err := c.Compress(buf.Bytes())
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	err = xclient.Call(context.Background(), "ColdPush", &data, &reply)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("[Cold] reply: %s\n", reply)
	}
}
