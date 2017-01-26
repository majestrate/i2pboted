package main

import (
	"github.com/majestrate/i2pboted/bote"
	"github.com/majestrate/i2pboted/config"
	"github.com/majestrate/i2pboted/i2p"
	"github.com/majestrate/i2pboted/log"
	"os"
)

func main() {
	fname := "bote.ini"
	if len(os.Args) == 2 {
		fname = os.Args[1]
	}
	cfg, err := config.Load(fname)
	if err != nil {
		log.Fatal(err.Error())
	}
	r := bote.NewRouter(cfg.Router)
	session, err := i2p.NewPacketSession(cfg.I2P)
	if err != nil {
		log.Fatalf("failed to create i2p network context: %s", err)
	}
	r.InjectNetwork(session)
}
