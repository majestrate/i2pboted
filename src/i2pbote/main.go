package i2pbote

import (
	"i2pbote/bote"
	"i2pbote/config"
	"i2pbote/i2p"
	"i2pbote/log"
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
}
