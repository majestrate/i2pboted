package bote

import (
	"i2pbote/config"
	"i2pbote/i2p"
	"i2pbote/log"
)

type Router struct {
	session i2p.PacketSession
	done    chan error
}

func NewRouter(cfg config.RouterConfig) *Router {
	return &Router{
		done: make(chan error),
	}
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

func (r *Router) GracefulClose() {
	log.Info("graceful close started")
	// TODO: begin graceful shutdown
}

// get bote rpc instance
func (r *Router) RPC() *RPC {
	return &RPC{
		router: r,
	}
}

// wait until done
func (r *Router) Wait() error {
	return <-r.done
}

// blocking run mainloop
func (r *Router) Run() {
	log.Debug("i2pbote run mainloop")
	var err error
	for err == nil {

	}
	r.done <- err
}
