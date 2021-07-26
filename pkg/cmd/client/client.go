package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"flag"
	"fmt"
	"os"

	"github.com/smallnest/rpcx/client"
	"github.com/soyoslab/soy_log_explorer/pkg/esdocs"
	"github.com/soyoslab/soy_log_generator/pkg/compressor"
)

var addr = flag.String("addr", "localhost:"+os.Getenv("EXPLORER_PORT"), "server address")

func main() {
	flag.Parse()

	var reply string

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	xclient := client.NewXClient("Rpush", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	docs := esdocs.ESdocs{Index: "application1", Docs: `{"File2_2021-07-26_06_35_33_850136975":"hello"}`}
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
