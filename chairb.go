package main

import (
	"fmt"
	"os"
	
	"github.com/ugorji/go/codec"
	
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

func parseIncomingMessage(msg *bw.SimpleMessage, h codec.Handle) {
	var contents []byte = msg.POs[0].GetContents()
	var dec *codec.Decoder = codec.NewDecoderBytes(contents, h)
	var cm ChairMessage = NewChairMessage()
	var err error = dec.Decode(&cm)
	fmt.Printf("Decoded message: %v\n", cm)
	exitOnError(err, "Could not decode message")
	err = cm.SanityCheck()
	exitOnError(err, "Decode message fails sanity check")
}

func main() {
	var err error
	
	var chandle codec.Handle = new(codec.MsgpackHandle)
	
	var cl *bw.BW2Client
	cl, err = bw.Connect("localhost:28589")
	exitOnError(err, "Could not connect")
	
	var vk string
	vk, err = cl.SetEntityFile("chairb.key")
	exitOnError(err, "Could not use entity")
	
	var dchain *bw.SimpleChain
	dchain, err = cl.BuildAnyChain(RAWDATAURI, "CP", vk)
	exitOnError(err, "Could not build DOT chain")
	
	var msgchan chan *bw.SimpleMessage
	msgchan, err = cl.Subscribe(&bw.SubscribeParams{
		URI: RAWDATAURI,
		PrimaryAccessChain: dchain.Hash,
		ElaboratePAC: "full",
	})
	
	var msg *bw.SimpleMessage
	for msg = range msgchan {
		parseIncomingMessage(msg, chandle)
		msg.Dump()
	}
}
