package i2pbote

import (
	"i2pbote/bote"
	"i2pbote/bote/network"
	"i2pbote/config"
	"i2pbote/i2p"
	"i2pbote/log"
	"i2pbote/version"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/signal"
)

func Main() {
	log.Infof("starting %s", version.Version())
	fname := "bote.ini"
	if len(os.Args) == 2 {
		fname = os.Args[1]
	}
	cfg, err := config.Load(fname)
	if err != nil {
		log.Fatal(err.Error())
	}
	sigch := make(chan os.Signal)
	r := bote.NewRouter(&cfg.Router)
	log.Info("starting up i2p network connection")
	session, err := i2p.NewPacketSession(cfg.I2P)
	if err != nil {
		log.Errorf("failed to create i2p network context: %s", err)
		return
	}
	signal.Notify(sigch, os.Interrupt)
	// signal handler
	go func() {
		for {
			sig, ok := <-sigch
			if !ok {
				break
			}
			if sig == os.Interrupt {
				log.Info("interrupt received, closing router")
				r.Close()
			}
		}
	}()
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
			defer l.Close()
			if cfg.RPC.BindUnix != "" {
				defer os.Remove(cfg.RPC.BindUnix)
			}
			for r.IsRunning() {
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

	// bootstrap from seed nodes
	strapper := network.NewFileBootstrap(cfg.Bootstrap.NodeFile, session)
	err = r.TryBootstrap(strapper)
	if err != nil {
		log.Warnf("bootstrap failed: %s", err.Error())
	}

	err = r.Wait()
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("i2pbote exited")
}
