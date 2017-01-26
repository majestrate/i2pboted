package bote

import (
	"github.com/majestrate/i2pboted/config"
	"github.com/majestrate/i2pboted/i2p"
	"github.com/majestrate/i2pboted/log"
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
