package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	
	"github.com/ugorji/go/codec"
	
	bw "github.com/immesys/bw2bind"
)

const (
	RAWDATAURI = "castle.bw2.io/sam/+/+/rawlog"
	CHAIREND = "/rawlog"
	CHAIRSTART = "/sam"
)

func exitOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %v\n", msg, err)
		os.Exit(1)
	}
}

func parseIncomingMessage(msg *bw.SimpleMessage, h codec.Handle) (string, ChairMessage, error) {
	var end int = strings.LastIndex(msg.URI, CHAIREND)
	var start int = strings.Index(msg.URI, CHAIRSTART) + len(CHAIRSTART)
	if end == -1 || start == -1 {
		return "", nil, errors.New(fmt.Sprintf("invalid path %v", msg.URI))
	}
	var chairid string = msg.URI[start:end]
	
	var contents []byte = msg.POs[0].GetContents()
	var dec *codec.Decoder = codec.NewDecoderBytes(contents, h)

	var cm ChairMessage
	var err error
	cm, err = NewChairMessageFrom(dec)
	if err != nil {
		return chairid, cm, errors.New(fmt.Sprintf("could not parse message: %v", err))
	}
	
	err = cm.SanityCheck()
	if err != nil {
		return chairid, cm, errors.New(fmt.Sprintf("decoded message fails sanity check: %v", err))
	} else {
		return chairid, cm, nil
	}
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
	var cid string
	var cm ChairMessage
	for msg = range msgchan {
		cid, cm, err = parseIncomingMessage(msg, chandle)
		if err == nil {
			fmt.Printf("Received message for chair %v: %v\n", cid, cm)
		} else {
			fmt.Printf("Could not handle incoming message: %v\n", err)
		}
	}
}
