package main

import (
	"fmt"
	"os"
	
	bw "github.com/immesys/bw2bind"
)

const (
	RAWDATAURI = "castle.bw2.io/sam/test/rawdata"
)

func exitOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %v\n", msg, err)
		os.Exit(1)
	}
}

func main() {
	var err error
	var cl *bw.BW2Client
	
	cl, err = bw.Connect("localhost:28589")
	exitOnError(err, "Could not connect")
	
	var vk string
	vk, err = cl.SetEntityFile("chairb.key")
	exitOnError(err, "Could not use entity")
	
	var dchain *bw.SimpleChain
	dchain, err = cl.BuildAnyChain(RAWDATAURI, "P", vk)
	exitOnError(err, "Could not build DOT chain")
	
	var msgchan chan *bw.SimpleMessage
	msgchan, err = cl.Subscribe(&bw.SubscribeParams{
		URI: RAWDATAURI,
		PrimaryAccessChain: dchain.Hash,
	})
	
	var msg *bw.SimpleMessage
	for msg = range msgchan {
		msg.Dump()
	}
}
