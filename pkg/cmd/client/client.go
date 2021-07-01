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

	docs := esdocs{Index: "my_index", Docs: `{"name":"kai"}`}
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err1 := enc.Encode(docs)
	if err1 != nil {
		fmt.Printf("error1: %v\n", err1)
		return
	}

	c := &compressor.GzipComp{}
	data := c.Compress(buf.Bytes())

	err := xclient.Call(context.Background(), "HotPush", &data, &reply)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("reply: %s\n", reply)
	}
}
