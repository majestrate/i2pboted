package bote

import (
	"i2pbote/config"
	"i2pbote/i2p"
	"i2pbote/log"
)

type Router struct {
	session i2p.PacketSession
}

func NewRouter(cfg config.RouterConfig) *Router {
	return &Router{}
}

func (r *Router) InjectNetwork(s i2p.PacketSession) {
	log.Infof("Aquired net context with address %s", s.I2PAddr().Base32())
	r.session = s
}

// close router and network context
func (r *Router) Close() {
	log.Info("closing i2pbote context...")
	if r.session != nil {
		log.Debug("closing network session")
		r.session.Close()
	}
}
