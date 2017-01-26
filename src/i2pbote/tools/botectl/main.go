package main

import (
	"fmt"
	"i2pbote/bote"
	"i2pbote/log"
	"net"
	"net/rpc/jsonrpc"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Infof("Usage: %s url method", os.Args[0])
		return
	}
	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Error(err.Error())
		return
	}
	conn, err := net.Dial(u.Scheme, u.Host)
	if err != nil {
		log.Errorf("Dial(): %s", err.Error())
		return
	}
	cl := jsonrpc.NewClient(conn)
	rpl := new(bote.RPCResult)
	args := new(bote.RPCArgs)
	meth := fmt.Sprintf("i2pbote.%s", os.Args[2])
	err = cl.Call(meth, args, rpl)
	if err == nil {
		log.Infof("%s result: %d", meth, rpl.Code)
	} else {
		log.Errorf("%s failed: %s", meth, err)
	}
}
