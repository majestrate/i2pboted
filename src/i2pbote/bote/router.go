package bote

import (
	"i2pbote/bote/protocol"
	"i2pbote/bote/protocol/comm"
	"i2pbote/config"
	"i2pbote/i2p"
	"i2pbote/log"
	"net"
)

type Router struct {
	session i2p.PacketSession
	done    chan error
	handler protocol.Handler
}

func NewRouter(cfg config.RouterConfig) *Router {
	handler := protocol.NewHandler()
	return &Router{
		done:    make(chan error),
		handler: handler,
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
	r.done <- nil
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
	err := <-r.done
	close(r.done)
	return err
}

func (r *Router) gotPacketFrom(data []byte, from net.Addr) {
	log.Debugf("got %d bytes from remote peer", len(data))
	pkt, err := comm.Parse(data)
	if err != nil {
		log.Warnf("%s : %s", from, err)
		return
	}
	err = r.handler.CommPacket(pkt, from, r.session)
	if err != nil {
		log.Warnf("packet error: %s", err)
		return
	}
}

// blocking run mainloop
func (r *Router) Run() {
	log.Debug("i2pbote run mainloop")
	var b [i2p.DatagramMTU]byte
	for {
		n, from, err := r.session.ReadFrom(b[:])
		if err != nil {
			r.done <- err
			return
		}
		msg := make([]byte, n)
		copy(msg, b[:n])
		r.gotPacketFrom(msg, from)
	}
}
