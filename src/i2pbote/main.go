package i2pbote

import (
	"i2pbote/bote"
	"i2pbote/config"
	"i2pbote/i2p"
	"i2pbote/log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

var Version = "0.0.0"

func Main() {
	log.Infof("starting i2pboted-%s", Version)
	fname := "bote.ini"
	if len(os.Args) == 2 {
		fname = os.Args[1]
	}
	cfg, err := config.Load(fname)
	if err != nil {
		log.Fatal(err.Error())
	}
	r := bote.NewRouter(cfg.Router)
	log.Info("starting up i2p network connection")
	session, err := i2p.NewPacketSession(cfg.I2P)
	if err != nil {
		log.Fatalf("failed to create i2p network context: %s", err)
	}
	r.InjectNetwork(session)
	if cfg.RPC.Enabled {
		log.Info("RPC enabled")
		l, err := cfg.RPC.CreateListener()
		if err != nil {
			log.Fatal(err.Error())
		}
		serv := rpc.NewServer()
		serv.RegisterName("i2pbote", r.RPC())
		go func() {
			defer func() {
				l.Close()
				if cfg.RPC.BindUnix != "" {
					// remove unix socket
					os.Remove(cfg.RPC.BindUnix)
				}
			}()
			for {
				c, e := l.Accept()
				if e == nil {
					log.Infof("New RPC connection from %s", c.RemoteAddr())
					go serv.ServeCodec(jsonrpc.NewServerCodec(c))
				} else {
					log.Warnf("RPC Accept() failed: %s", err)
				}
			}
		}()
	}
	go r.Run()
	err = r.Wait()
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("i2pbote exited")
}
